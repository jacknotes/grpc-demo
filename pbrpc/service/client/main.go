package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/jacknotes/grpc-demo/pbrpc/codec/client"
	"github.com/jacknotes/grpc-demo/pbrpc/service"
)

// 约束客户端
var _ service.HelloService = (*HelloServiceClient)(nil)

type HelloServiceClient struct {
	*rpc.Client
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	// c, err := rpc.Dial(network, address)
	// if err != nil {
	// 	return nil, err
	// }

	// 建立链接
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	// 采用Json编解码的客户端
	// c := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	c := rpc.NewClientWithCodec(client.NewClientCodec(conn))
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(req *service.Request, resp *service.Response) error {
	return p.Client.Call(service.Name+".Hello", req, resp)
}

func main() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	resp := &service.Response{}
	err = client.Hello(&service.Request{Value: "k8s"}, resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
