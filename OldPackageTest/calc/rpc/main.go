package main

import "fmt"

func Add(a, b int) int {
	total := a + b
	return total
}

func main() {
	//现在我们想把Add函数变成一个远程的函数调用，也就意味着需要把Add函数放在远程服务器上去运行
	/*
		我们原本的电商系统，这里地方有一段逻辑，这个逻辑是扣减库存 但是库存服务是一个独立的系统reduce
		如何调用 一定会牵扯到网络，做成一个web服务（gin）

		1.这个函数的调用参数如何传递-json
	*/
	fmt.Println(Add(2, 3))

}
