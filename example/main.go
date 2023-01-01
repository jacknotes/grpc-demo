package main

import (
	"fmt"

	"github.com/jacknotes/grpc-demo/protobuf"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
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

	// Oneof 使用
	sm := &protobuf.SampleMessage{
		TestOneOf: &protobuf.SampleMessage_Sub2{
			Sub2: &protobuf.Sub2{Name: "test"},
		},
	}
	fmt.Println(sm.GetSub1()) // <nil>
	fmt.Println(sm.GetSub2()) // name:"test"

	// any 使用
	// 把sub2转化为any类型
	sub2, _ := anypb.New(&protobuf.Sub2{Name: "test"})
	resp := &protobuf.Response{
		Code: 0,
		Data: sub2,
	}
	fmt.Println(resp)

	sub3 := &protobuf.Sub2{}
	sub2.UnmarshalTo(sub3)
	fmt.Println(sub3)
}
