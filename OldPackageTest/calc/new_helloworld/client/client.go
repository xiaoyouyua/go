package main

import (
	"OldPackageTest/calc/new_helloworld/client_proxy"
	"fmt"
)

func main() {
	//建立连接
	client := client_proxy.NewHelloServerClinet("tcp", "localhost:1234")
	//1. 只想写业务逻辑量不 想关注每个函数的名称
	var reply string
	err := client.Hello("bobby", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)

}
