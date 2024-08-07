package main

import (
	"fmt"
	// "OldPackageTest/calc"
	// "github.com/gin-gonic/gin"
)

// 一定要将代码新建到gopath目录之下的src
// 要记得设置G0111MODULE=off
func main() {
	// fmt.Println("hello go")
	// // r := gin.Default()
	// fmt.Println(calc.Add(1, 2))
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Printf("%p %d\n", &i, i)
		}()
	}
}
