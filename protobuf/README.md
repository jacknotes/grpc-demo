cd /d/share/golang-study/day15-grpc-demo/grpc-demo
protoc -I=. --go_out=./protobuf/ --go_opt=module="github.com/jacknotes/grpc-demo/protobuf" ./protobuf/hello.proto 


$ protoc -I=. -I='/c/Program Files/Git/usr/local/include' --go_out=./protobuf/ --go_opt=module="github.com/jacknotes/grpc-demo/protobuf" ./protobuf/hello.proto


$ protoc -I=. -I='/c/Program Files/Git/usr/local/include' --go_out=./protobuf/ --go_opt=module="github.com/jacknotes/grpc-demo/protobuf" ./protobuf/pkg.proto