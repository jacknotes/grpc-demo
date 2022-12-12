package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/jacknotes/grpc-demo/rpc/service"
)

// var _ Service = &HelloService{}
// 我们声明了一个空指针，强制把这个指针转换成了一个*HelloService类型
var _ service.Service = (*HelloService)(nil)

type HelloService struct{}

// 业务场景
// 该函数需要被客户端调用
// 改造成符合rpc规范的 函数签名
// 1. 第一个参数request, interface{}类型
// 2. 第二个参数是一个响应response,interface{}类型，必须是一个指针，需要回写到response
// 3. 返回一个error，不过一般我们通过response进行返回错误信息
func (s *HelloService) Hello(name string, response *string) error {
	*response = fmt.Sprintf("Hello, %s", name)
	return nil
}

// tcp rpc
// func main() {
// 	// 把要提供的服务注册给RPC框架
// 	err := rpc.RegisterName(service.Name, new(HelloService))
// 	if err != nil {
// 		panic(err)
// 	}

// 	// 监听Socket
// 	listener, err := net.Listen("tcp", ":1234")
// 	if err != nil {
// 		panic(err)
// 	}
// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			panic(err)
// 		}

// 		// 前面都是tcp的知识，到这里RPC就接管了,默认是gob类型的rpc
// 		// 因此，你可以认为RPC帮我们封装消息到函数调用的这个逻辑，
// 		// 提升了工作效率，逻辑比较简洁，可以看看相关代码
// 		// go rpc.ServeConn(conn)

// 		// json on tcp 类型的rpc
// 		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
// 	}
// }

type RPCReadWriteCloser struct {
	io.Writer
	io.ReadCloser
}

func NewRPCReadWriteCloserFromHTTP(w http.ResponseWriter, r *http.Request) *RPCReadWriteCloser {
	return &RPCReadWriteCloser{w, r.Body}
}

// http rpc
func main() {
	// 把要提供的服务注册给RPC框架
	err := rpc.RegisterName(service.Name, new(HelloService))
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		// 把HTTP请求  -->  RPC请求
		// 基于HTTP Request Body和Response构造了一个Conn
		conn := NewRPCReadWriteCloserFromHTTP(w, r)
		// rpc框架函数进行处理
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe("127.0.0.1:1234", nil)
}
