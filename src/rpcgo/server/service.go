package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}

func main() {
	/*
		1、实例化一个server
		2、注册处理逻辑 handler
		3、启动服务
	*/
	// 实例化一个server
	listener, _ := net.Listen("tcp", ":1234")
	// 注册服务
	_ = rpc.RegisterName("HelloService", &HelloService{})

	// 第三步 启动服务
	conn, _ := listener.Accept()
	for {
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

	// 可以跨语言调用v 1 go语言的rpc的序列化和反序列化的协议是什么（GOB）
	// 2 能否替换成常见的序列化
	//rpc.ServeConn(conn)
}
