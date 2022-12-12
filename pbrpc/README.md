#

```
# 进入pbrbc目录
cd pbrbc

# 编译文件
protoc -I=. --go_out=. --go_opt=module="github.com/jacknotes/grpc-demo/pbrpc" ./service/service.proto


# 生成Codec的编译文件
protoc -I=. --go_out=. --go_opt=module="github.com/jacknotes/grpc-demo/pbrpc" ./codec/pb/header.proto
```