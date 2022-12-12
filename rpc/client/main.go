package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/jacknotes/grpc-demo/rpc/service"
)

// 我们声明了一个空指针，强制把这个指针转换成了一个*HelloServiceClient类型
var _ service.Service = (*HelloServiceClient)(nil)

type HelloServiceClient struct {
	client *rpc.Client
}

// 客户端构建函数
func NewHelloServiceClient(network, address string) (service.Service, error) {
	// 首先是通过rpc.Dial拨号RPC服务, 建立连接
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	// 客户端实现了基于JSON的编解码
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	return &HelloServiceClient{
		client: client,
	}, nil
}

// 对于RPC客户端来说，我们就需要包装客户端的调用
func (c *HelloServiceClient) Hello(name string, response *string) error {
	// 然后通过client.Call调用具体的RPC方法
	// 在调用client.Call时:
	// 		第一个参数是用点号链接的RPC服务名字和对象的方法名字，
	// 		第二个参数是 请求参数
	//      第三个是请求响应, 必须是一个指针, 有底层rpc服务帮你赋值
	return c.client.Call(service.Name+".Hello", name, response)
}

func main() {
	client, err := NewHelloServiceClient("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	var response string
	client.Hello("World", &response)
	fmt.Println(response)
}
