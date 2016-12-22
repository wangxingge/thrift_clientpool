package main

import (
	pool "github.com/wangxingge/thrift_clientpool"
	entity "github.com/wangxingge/thrift_clientpool/examples/entity"
	bookservice "github.com/wangxingge/thrift_clientpool/examples/bookservice"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
	"net"
)

var (
	boolPool pool.ThriftClientPool
address = "127.0.0.1"
port = "19998"
	defaultBufferSize int = 1024 * 500

)

func main() {
	pool = pool.NewThriftClientPool(entity.ServiceTag_BookService, address, port, nil, nil, nil, 50, 10)

}

func dialBook() (connection *bookservice.BookServiceClient, err error) {
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

func keepalvieBook(connection interface{}) (err error){

}