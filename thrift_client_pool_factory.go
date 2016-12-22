package thrift_clientpool

import (
	"errors"
	"fmt"
	//"git.apache.org/thrift.git/lib/go/thrift"
	//"net"
	//"time"
)
var(
	defaultBufferSize int = 1024 * 500
)

type ThriftPoolFactory struct {
	pools map[string]*ThriftClientPool
}

func (factory *ThriftPoolFactory) NewPool(poolName, address, port string) error {

	if _, ok := factory.pools; ok {
		return errors.New(fmt.Sprintf("duplicate pool"))
	}

	return nil
}

//func (factory *ThriftPoolFactory) defaultDial(name string) (connection interface{}, err error) {
//
//	pool, ok := factory.pools[name]
//	if !ok {
//		return nil, errors.New(fmt.Sprintf("Pool not exists %v,Should create pool first.", name))
//	}
//
//	srvSocket, err := thrift.NewTSocketTimeout(net.JoinHostPort(pool.Address, pool.Port), time.Second*5)
//	if err != nil {
//		return nil, err
//	}
//
//	transport := thrift.NewTBufferedTransport(srvSocket, defaultBufferSize)
//	protocol := thrift.NewTBinaryProtocolTransport(transport)
//	mp_group := thrift.NewTMultiplexedProtocol(protocol, name)
//	if err = transport.Open(); err != nil {
//		return nil, err
//	}
//	//connection = emRpcSrv.NewGroupServiceClientProtocol(transport, mp_group, mp_group)
//	return
//}
