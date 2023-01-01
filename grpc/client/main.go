package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jacknotes/grpc-demo/grpc/protocol"
	"google.golang.org/grpc"
)

func main() {
	// grpc.WithInsecure()表示不使用tls
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := protocol.NewHelloServiceClient(conn)
	// context.Background()放在后台一直跑，相当于root的上下文
	resp, err := client.Hello(context.Background(), &protocol.Request{Value: "jack"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	// 返回一个channel io pipe 对象
	// HelloService_ChatClient 对应 HelloService_ChatServer
	stream, err := client.Chat(context.Background())
	if err != nil {
		panic(err)
	}

	// 在客户端我们将恢复原状和接收操作放到两个独立的 Goroutine
	// 首先是向服务端发送数据
	go func() {
		for {
			err := stream.Send(&protocol.Request{Value: "jack li"})
			if err != nil {
				fmt.Println(err)
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// 然后在循环中接收服务端返回的数据
	for {
		resp, err := stream.Recv()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp)
	}
}
