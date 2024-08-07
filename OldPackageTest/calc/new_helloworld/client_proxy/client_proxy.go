package client_proxy

import (
	"OldPackageTest/calc/new_helloworld/hanlder"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

// 在go语言中没有类、对象就意味着没有初始化方法
func NewHelloServerClinet(protol, address string) *HelloServiceStub {
	conn, err := rpc.Dial(protol, address)
	if err != nil {
		panic("connect error!")
	}
	return &HelloServiceStub{conn}
}

func (c *HelloServiceStub) Hello(requsest string, reply *string) error {
	err := c.Call(hanlder.HelloServiceName+".Hello", requsest, reply)
	if err != nil {
		return err
	}
	return nil
}
