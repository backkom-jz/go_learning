package main

import (
	"context"
	"log"

	"go_learning/src/grpc_test/proto"
	"google.golang.org/grpc"
)

func main() {
	// 配置 gRPC 服务地址
	serverAddress := "127.0.0.1:8088"

	// 创建 gRPC 连接
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server at %s: %v", serverAddress, err)
	}
	defer conn.Close()
	log.Printf("Connected to gRPC server at %s", serverAddress)

	// 创建 Greeter 客户端
	client := proto.NewGreeterClient(conn)

	// 调用 SayHello 方法
	request := &proto.HelloRequest{Name: "zhangliyuan"}
	response, err := client.SayHello(context.Background(), request)
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}
	log.Printf("Server response: %s", response.Message)
}
