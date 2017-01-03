package thrift_clientpool

import (
	"errors"
	"fmt"
	"time"
)

var (
	ClientFactory = &ThriftPoolFactory{}
)

type ThriftPoolFactory struct {
	pools map[string]*ThriftClientPool
}

func (factory *ThriftPoolFactory) NewPoolFullParam(poolName, address, port string, dialFn func(name, address, port string) (connection interface{}, err error), closeFn func(connection interface{}) (err error), keepAliveFn func(connection interface{}) (err error), poolSize, initialPoolSize int) error {

	if factory.pools == nil {
		factory.pools = make(map[string]*ThriftClientPool)
	}

	if _, ok := factory.pools[poolName]; ok {
		return errors.New(fmt.Sprintf("duplicate pool"))
	}

	tmp, err := NewThriftClientPool(poolName, address, port, dialFn, closeFn, keepAliveFn, poolSize, initialPoolSize)
	if err != nil {
		return err
	}
	tmp.KeepAliveInterval= time.Second

	factory.pools[poolName] = tmp
	tmp.Start()
	return nil
}

func (factory *ThriftPoolFactory) GetConnection(poolName string) (connection interface{}, err error) {

	if p, ok := factory.pools[poolName]; ok {
		return p.Get()
	}

	return nil, errors.New(fmt.Sprintf("Not found pool with name %v.", poolName))
}

func (factory *ThriftPoolFactory) PutConnection(poolName string, connection interface{}) {

	if p, ok := factory.pools[poolName]; ok {
		p.Put(connection)
	}
}
