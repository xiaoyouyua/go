package server_proxy

import (
	"OldPackageTest/calc/new_helloworld/hanlder"
	"net/rpc"
)

type HelloService interface {
	Hello(request string, reply *string) error
}

// 如果做到解耦－我们关心的是函数
func RegisterHelloServic(srv hanlder.NewHelloService) error {
	return rpc.RegisterName(hanlder.HelloServiceName, srv)
}
