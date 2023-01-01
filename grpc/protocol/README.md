# 如何编译proto文件

```
# 进入protocol目录
cd protocol

# 1. 编译文件, 指定protobuf文件的搜索位置为当前目录 -I=.
protoc -I=. -I='/c/Program Files/Git/usr/local/include' --go_out=. --go_opt=module="github.com/jacknotes/grpc-demo/grpc/protocol" hello.proto


# 2. 生成编译文件, 指定protobuf文件的搜索位置为当前目录 -I=.
protoc -I=. -I='/c/Program Files/Git/usr/local/include' --go_out=. --go-grpc_out=.  --go_opt=module="github.com/jacknotes/grpc-demo/grpc/protocol" --go-grpc_opt=module="github.com/jacknotes/grpc-demo/grpc/protocol" hello.proto

```