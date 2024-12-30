package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 1. 建立 TCP 连接到服务端
	conn, err := net.Dial("tcp", "localhost:1234") // 使用 net.Dial 创建 TCP 连接
	if err != nil {
		fmt.Printf("Failed to connect to RPC server: %v\n", err)
		return
	}
	defer func() {
		if cerr := conn.Close(); cerr != nil {
			fmt.Printf("Failed to close connection: %v\n", cerr)
		}
	}()

	// 2. 创建一个支持 JSON-RPC 协议的客户端
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	// 3. 准备存放服务端返回结果的变量
	var reply string

	// 4. 调用远程服务方法并处理结果
	const methodName = "HelloService.Hello"
	arg := "bobby" // 服务方法的参数
	err = client.Call(methodName, arg, &reply)
	if err != nil {
		fmt.Printf("Failed to call method '%s': %v\n", methodName, err)
		return
	}

	// 5. 输出服务端返回的结果
	fmt.Printf("Server response: %s\n", reply)
}
