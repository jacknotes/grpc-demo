package main

import (
	"context"
	"fmt"

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
}
