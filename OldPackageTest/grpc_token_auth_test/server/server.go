package main

import (
	"OldPackageTest/metadata/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {

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
		md, ok := metadata.FromIncomingContext(ctx)
		fmt.Println(md)
		if !ok {
			return resp, status.Error(codes.Unauthenticated, "无token认证信息")
			fmt.Println("get metadata error")
		}

		var (
			appid  string
			appkey string
		)
		if val, ok := md["appid"]; ok {
			appid = val[0]
		}
		if val, ok := md["appkey"]; ok {
			appkey = val[0]
		}
		if appid != "101010" || appkey != "i am key" {
			return resp, status.Error(codes.Unauthenticated, "token认证信息错误")
		}

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
