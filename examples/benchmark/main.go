package main

import (
	"errors"
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/wandoulabs/codis/pkg/utils/log"
	pool "github.com/wangxingge/thrift_clientpool"
	"github.com/wangxingge/thrift_clientpool/examples/bookservice"
	"github.com/wangxingge/thrift_clientpool/examples/entity"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

var (
	address                 = flag.String("addr", "127.0.0.1", "[addr] Server address, default 127.0.0.0.1")
	port                    = "19998"
	defaultBufferSize int   = 1024 * 500
	maxCalls                = flag.Int("maxcall", 1000, "[maxcall] Max count of api invokes, default is 1000.")
	maxConnection           = flag.Int("maxcon", 50, "[maxcon] Max connections of connection pool, default 50.")
	initConnection          = flag.Int("intcon", 1, "[intcon] Init connections of connection pool, default 1.")
	threadsCount            = flag.Int("t", 1, "[t] Concurrent thread count, default 1.")
	remainingCall     int32 = 0
	failedCall        int32 = 0
	successCall       int32 = 0
	wg                sync.WaitGroup
)

func main() {
	flag.Parse()

	if *maxConnection < 0 {
		log.Println("maxcon less than 0.")
		return
	}

	if *maxCalls < 0 {
		log.Println("maxCalls less than 0.")
		return
	}

	if *initConnection < 0 {
		log.Println("initConnection less than 0.")
		return
	}

	if *threadsCount < 0 {
		log.Println("threadsCount less than 0.")
		return
	}

	pool.ClientFactory.NewPoolFullParam(entity.ServiceTag_BookService, *address, port, dialBook, closeBook, keepAlvieBook, *maxConnection, *initConnection)

	start := time.Now()
	for i := 0; i < *threadsCount; i++ {
		go ApiTest()
	}
	time.Sleep(time.Second)
	wg.Wait()
	fmt.Printf("Success: %v, Failed: %v, Cost time: %v\n", successCall, failedCall, time.Now().Sub(start).Nanoseconds())
}

func ApiTest() {
	wg.Add(1)
	con, err := pool.ClientFactory.GetConnection(entity.ServiceTag_BookService)
	defer func() {
		wg.Done()
		pool.ClientFactory.PutConnection(entity.ServiceTag_BookService, con)
	}()

	if err != nil {
		fmt.Printf("Get Connection Failed: %v\n", err)
		return
	}

	tmpCon, ok := con.(*bookservice.BookServiceClient)
	if !ok {
		fmt.Println("Convert Connection Failed")
		return
	}

	for remainingCall < int32(*maxCalls) {
		if _, err := tmpCon.GetAllBooks(); err == nil {
			atomic.AddInt32(&successCall, 1)
		} else {
			atomic.AddInt32(&failedCall, 1)
		}

		atomic.AddInt32(&remainingCall, 1)
	}
}

func getBooks() {
	con, err := pool.ClientFactory.GetConnection(entity.ServiceTag_BookService)
	if err != nil {
		fmt.Printf("Get Connection Failed: %v\n", err)
		return
	}

	defer pool.ClientFactory.PutConnection(entity.ServiceTag_BookService, con)

	tmpCon, ok := con.(*bookservice.BookServiceClient)
	if !ok {
		fmt.Println("Convert Connection Failed")
		return
	}

	if v, err := tmpCon.GetAllBooks(); err != nil {
		fmt.Println("Get All Books Failed: ", err)
	} else {
		fmt.Println(v)
	}
}

func addBook() {
	con, err := pool.ClientFactory.GetConnection(entity.ServiceTag_BookService)
	if err != nil {
		fmt.Printf("Get Connection Failed: %v\n", err)
		return
	}

	defer pool.ClientFactory.PutConnection(entity.ServiceTag_BookService, con)

	tmpCon, ok := con.(*bookservice.BookServiceClient)
	if !ok {
		fmt.Println("Convert Connection Failed")
		return
	}
	r, err := tmpCon.AddBook(&entity.Book{
		BookId:   "1",
		BookName: "How to use thrift",
		Author:   "Wang Xingge"},
	)

	if err != nil {
		fmt.Printf("Add Book Failed: %v\n", err)
		os.Exit(0)
	}

	fmt.Printf("Add Book %v\n", r)

}

func dialBook() (connection interface{}, err error) {

	srvSocket, err := thrift.NewTSocketTimeout(net.JoinHostPort(*address, port), time.Second*5)
	if err != nil {
		return nil, err
	}

	transport := thrift.NewTBufferedTransport(srvSocket, defaultBufferSize)
	protocol := thrift.NewTBinaryProtocolTransport(transport)
	mp_group := thrift.NewTMultiplexedProtocol(protocol, entity.ServiceTag_BookService)
	if err = transport.Open(); err != nil {
		return nil, err
	}
	connection = bookservice.NewBookServiceClientProtocol(transport, mp_group, mp_group)
	return
}

func closeBook(connection interface{}) (err error) {

	return
}

func keepAlvieBook(connection interface{}) (err error) {

	if t, ok := connection.(*bookservice.BookServiceClient); ok {
		ok, err = t.DefaultKeepAlive("book service client A")
		if !ok || err != nil {
			return errors.New(fmt.Sprintf("BookService KeepAlive Failed: %v", err))
		}

	} else {
		fmt.Printf("BookService KeepAlive Failed\n")
		return errors.New(fmt.Sprintf("BookService KeepAlive Failed"))
	}

	return nil
}
