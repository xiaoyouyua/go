package main

import (
	"OldPackageTest/calc/new_helloworld/hanlder"
	"OldPackageTest/calc/new_helloworld/server_proxy"
	"net"
	"net/rpc"
)

func main() {
	//1。实例化一个server
	listener, _ := net.Listen("tcp", ":1234")

	//2.注册处理逻辑handler
	server_proxy.RegisterHelloServic(hanlder.NewHelloService{})
	//3．启动服务
	for {
		conn, _ := listener.Accept() //当一个新的连接进来的时候
		go rpc.ServeConn(conn)
	}

}
