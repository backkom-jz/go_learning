package main

import (
	"context"
	"log"
	"net"

	"go_learning/src/grpc_test/proto"
	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedGreeterServer // 提供默认实现，便于扩展
}

// SayHello 实现了 Greeter 服务的 SayHello 方法
func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	log.Printf("Received request from: %s", request.Name) // 添加日志
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	// 创建 gRPC 服务器实例
	server := grpc.NewServer()

	// 注册 Greeter 服务
	proto.RegisterGreeterServer(server, &Server{})

	// 监听指定的网络地址和端口
	address := "0.0.0.0:8088"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err)
	}
	log.Printf("gRPC server is listening on %s", address)

	// 启动 gRPC 服务
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
