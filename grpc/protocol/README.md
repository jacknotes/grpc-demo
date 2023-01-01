# 如何编译proto文件

```
# 安装protoc-gen-go-grpc插件，从protobuf编译生成grpc相关代码
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
$ protoc-gen-go-grpc.exe --version
protoc-gen-go-grpc 1.1.0


# 进入protocol目录
cd protocol

# 生成grpc 和 protobuf编译文件, 指定protobuf文件的搜索位置为当前目录 -I=.
protoc -I=. -I='/c/Program Files/Git/usr/local/include' --go_out=. --go-grpc_out=.  --go_opt=module="github.com/jacknotes/grpc-demo/grpc/protocol" --go-grpc_opt=module="github.com/jacknotes/grpc-demo/grpc/protocol" hello.proto

```