package main

import (
	helloworld "OldPackageTest/calc/helloworld/proto"
	"fmt"
	"github.com/golang/protobuf/proto"
)

type Hello struct {
	Name    string   `json:name`
	Age     int      `json:age`
	Courses []string `json:courses`
}

func main() {
	req := helloworld.HelloRequest{
		Name:    "helloworld",
		Age:     20,
		Courses: []string{"go", "gin", "微服务"},
	}
	//jsonStruct := Hello{
	//	Name:    "helloworld",
	//	Age:     18,
	//	Courses: []string{"go", "gin", "微服务"}}
	//jsonRsp, _ := json.Marshal(jsonStruct)
	//fmt.Println(len(jsonRsp))
	rsp, _ := proto.Marshal(&req)
	newReq := &helloworld.HelloRequest{}
	proto.Unmarshal(rsp, newReq)
	fmt.Println(newReq.Name, newReq.Age, newReq.Courses)
}
