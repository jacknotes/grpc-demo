package main

import (
	"fmt"

	"github.com/jacknotes/grpc-demo/protobuf"
	"google.golang.org/protobuf/proto"
)

func main() {
	s1 := &protobuf.String{
		Value: "number 1中文",
	}

	payload, err := proto.Marshal(s1)
	if err != nil {
		panic(err)
	}

	fmt.Println(payload)
	// fmt.Println(string(payload))

	s2 := &protobuf.String{}
	proto.Unmarshal(payload, s2)
	fmt.Println(s2)
}
