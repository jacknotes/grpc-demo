package main

import (
	"context"
	"fmt"
	"net"

	"github.com/jacknotes/grpc-demo/grpc/protocol"
	"google.golang.org/grpc"
)

type Service struct {
	// 结构体嵌套(必须是匿名)，等价于继承UnimplementedHelloServiceServer，而UnimplementedHelloServiceServer实现了
	// HelloServiceServer接口，所以Service间接实现了HelloServiceServer接口
	protocol.UnimplementedHelloServiceServer
}

// 重写protocol.UnimplementedHelloServiceServer的方法Hello,简单加个hello前缀
func (s *Service) Hello(ctx context.Context, req *protocol.Request) (*protocol.Response, error) {
	return &protocol.Response{
		Value: "hello: " + req.Value,
	}, nil
}

func main() {
	// 如何把Service 作为一个rpc暴露出去，提供服务
	server := grpc.NewServer()

	// Service生成的代码里面，提供了注册函数，把自己注册到grpc的servr内
	protocol.RegisterHelloServiceServer(server, new(Service))

	// 开启grpc服务
	ls, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	// 监听grpc服务
	if err := server.Serve(ls); err != nil {
		fmt.Println(err)
	}

}
