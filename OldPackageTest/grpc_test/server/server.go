package main

import (
	"OldPackageTest/grpc_test/proto"
	"context"
	"google.golang.org/grpc"
	"net"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "hello," + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
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
