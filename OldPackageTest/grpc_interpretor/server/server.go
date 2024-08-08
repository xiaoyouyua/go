package main

import (
	"OldPackageTest/metadata/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println("get metadata error")
	}
	for k, v := range md {
		fmt.Println(k, v)
	}
	//for nameSlice, ok := md["name"]; ok{
	//	fmt.Println(nameSlice)
	//	for i, v := range nameSlice {
	//		fmt.Println(i, v)
	//	}
	//}
	return &proto.HelloReply{
		Message: "hello," + request.Name,
	}, nil
}

func main() {
	interpreto := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		fmt.Println("接收到一个新的请求")
		return handler(ctx, req)
	}
	opt := grpc.UnaryInterceptor(interpreto)
	g := grpc.NewServer(opt)
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	err = g.Serve(lis)
	if err != nil {
		panic(err)
	}

}
