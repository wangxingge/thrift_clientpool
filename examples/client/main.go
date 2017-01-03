package main

import (
	"errors"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	pool "github.com/wangxingge/thrift_clientpool"
	"github.com/wangxingge/thrift_clientpool/examples/bookservice"
	"github.com/wangxingge/thrift_clientpool/examples/entity"
	"net"
	"os"
	"time"
)

var (
	address               = "127.0.0.1"
	port                  = "19998"
	defaultBufferSize int = 1024 * 500
)

func main() {

	pool.ClientFactory.NewPoolFullParam(entity.ServiceTag_BookService, address, port, dialBook, closeBook, keepAlvieBook, 50, 1)

	addBook()

	getBooks()

	select {}
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

func dialBook(name, address, port string) (connection interface{}, err error) {

	srvSocket, err := thrift.NewTSocketTimeout(net.JoinHostPort(address, port), time.Second*5)
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
