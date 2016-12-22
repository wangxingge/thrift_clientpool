package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/wangxingge/thrift_clientpool/examples/bookservice"
	"github.com/wangxingge/thrift_clientpool/examples/entity"
	impl "github.com/wangxingge/thrift_clientpool/examples/server/serviceImplement"
	"log"
)

var (
	DefaultTransportBufferSize = 1024 * 10
	serviceEndpoint            = flag.String("endpoint", "0.0.0.0", "[endpoint] Default endpoint is 0.0.0.0.")
	servicePort                = flag.String("port", "19998", "[port] Default service port is 199998.")
)

func main() {

	binaryProtocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTBufferedTransportFactory(DefaultTransportBufferSize)
	network := fmt.Sprintf("%v:%v", *serviceEndpoint, *servicePort)
	serverTransport, err := thrift.NewTServerSocket(network)
	if err != nil {
		log.Println("Start RPC Service Failed: ", err)
		return
	}

	multi_processor := thrift.NewTMultiplexedProcessor()
	multi_processor.RegisterProcessor(entity.ServiceTag_BookService, bookservice.NewBookServiceProcessor(&impl.BookServiceImpl{}))
	multi_processor.RegisterProcessor(entity.ServiceTag_UserService, bookservice.NewUserServiceProcessor(&impl.UserServiceImpl{}))
	server := thrift.NewTSimpleServer4(multi_processor, serverTransport, transportFactory, binaryProtocolFactory)
	log.Println("Service started on:", network)
	server.Serve()
}
