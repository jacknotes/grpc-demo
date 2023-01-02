package main

import (
	"context"
	"fmt"
	"io"
	"net"

	"github.com/jacknotes/grpc-demo/grpc/auth"
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

// server --> stream --> client
// 重写，HelloService_ChatServer 对应 HelloService_ChatClient
func (s *Service) Chat(stream protocol.HelloService_ChatServer) error {
	// 可以把stream当成一个channel io pipe对象
	for {
		// 接收客户端请求
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("chat exit")
				return nil
			}
			fmt.Println(err)
			return err
		}
		// 处理请求
		err = stream.Send(&protocol.Response{Value: "hello: " + req.Value})
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
}

func main() {
	// 如何把Service 作为一个rpc暴露出去，提供服务
	// 需要将auth.GrpcAuthUnaryServerInterceptor()传给grpc.UnaryInterceptor()，才能成为一个拦截器
	server := grpc.NewServer(
		// 请求响应模式的认证中间件，会走此方法去认证
		grpc.UnaryInterceptor(auth.GrpcAuthUnaryServerInterceptor()),
		// stream模式的认证中间件，会走此方法去认证
		grpc.StreamInterceptor(auth.GrpcAuthStreamServerInterceptor()),
	)

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
