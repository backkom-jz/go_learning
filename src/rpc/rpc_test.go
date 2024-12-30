package rpc

import (
	"fmt"
	"testing"
)

func Add(a, b int) int {
	total := a + b
	return total
}

// json 是一种数据格式的协议 xml protobuf msgpack
// 现在网络调用的有两个端 --客户端  将数据传输到gin-服务端
// 服务端负责解析数据
// json不是一个高性能的数据协议

type Company struct {
	Name    string
	Address string
}

type Employee struct {
	Name    string
	company Company
}
type PrintResult struct {
	Info string
	Err  error
}

func RpcPrintLn(employee Employee) PrintResult {
	/*
			客户端
				1、建立链接 tcp/http
				2、将employee对象序列化成json字符串--序列化
				3、发送json字符串-- 调用成功后接受到的是一个二进制的数据
				4、等待服务器发送结果
				5、将服务器返回的数据解析成PrintResult对象 --反序列化
			服务端
				1、监听网路端口 80
				2、读取数据 - 二进制的json数据
				3、将数据进行反序列化为Employee对象
				4、开始处理业务逻辑
				5、将处理的接口PrintResult序列化成Json二进制数据
				6、将数据返回
		序列化和反序列化是可以选择的，不一定要采用json、xml、protobuf[压缩比]、msgpack
	*/

	//RPC中的第二个点 传输协议【http 1.0/2.0、tcp、udp】，数据编码协议
	/*
		http现在流行的是1.*，这种协议有性能问题，一次性，一旦结果返回，连接就断开
		1、直接自己基于tcp/udp协议封装一层协议 myhttp 没有通用性 http2.0 既有http的特性 也有长链接的特性

	*/
	return PrintResult{
		"123",
		nil,
	}
}

/*
	http协议来说，有一个问题：一次性，一旦对方返回了接口，链接断开 http2.0支持长链接
	如不使用http2.0 可使用tcp
*/

func TestAdd(t *testing.T) {
	fmt.Println(Add(1, 2))

	// 打印的工作放到另外一台服务器上，我需要将本地的内存对象struct 这样不行
	// 可行的方式就是将struct序列化成json
	// 远程服务器需要将二进制对象反解析成struct
	// 这么麻烦 直接全部使用json格式化
	// 这种方法再浏览器和gin之间的服务之间可行
	// 但是一个大型的服务
	fmt.Println(
		Employee{
			Name: "bobby",
			company: Company{
				Name:    "慕课网",
				Address: "上海市",
			},
		})
}
