// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: protobuf/hello.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 定义枚举类型
type Option int32

const (
	// 枚举选项, 必须从0编号开始，0编号标识的是枚举的默认值
	Option_A Option = 0 // A后面版本会弃用
	Option_B Option = 1
	Option_C Option = 2
	Option_D Option = 0 // D默认不能定义0编号，但是需要跟A编号一样，则需要开启允许别名
)

// Enum value maps for Option.
var (
	Option_name = map[int32]string{
		0: "A",
		1: "B",
		2: "C",
		// Duplicate value: 0: "D",
	}
	Option_value = map[string]int32{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 0,
	}
)

func (x Option) Enum() *Option {
	p := new(Option)
	*p = x
	return p
}

func (x Option) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Option) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_hello_proto_enumTypes[0].Descriptor()
}

func (Option) Type() protoreflect.EnumType {
	return &file_protobuf_hello_proto_enumTypes[0]
}

func (x Option) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Option.Descriptor instead.
func (Option) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_hello_proto_rawDescGZIP(), []int{0}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data *anypb.Any `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_protobuf_hello_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type String struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value       string           `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`                                                                                            // 等价于Value string中go变量的写法，probobuf字段默认大写，为全部导出，不可配置为小写
	Key         *string          `protobuf:"bytes,2,opt,name=key,proto3,oneof" json:"key,omitempty"`                                                                                          // Key string , 为入optional后运行为nil "" null
	Enabled     *bool            `protobuf:"varint,3,opt,name=enabled,proto3,oneof" json:"enabled,omitempty"`                                                                                 // 默认enabled=false,  添加了optional则默认是nil了，enabled != nil 时才添加过滤条件
	MetricValue float64          `protobuf:"fixed64,4,opt,name=metric_value,json=metricValue,proto3" json:"metric_value,omitempty"`                                                           // float64，会默认把蛇形转换为大驼峰形
	Values      map[string]int32 `protobuf:"bytes,5,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` //map<string,string> 为key和value的类型
	String1     *String1         `protobuf:"bytes,6,opt,name=string1,proto3" json:"string1,omitempty"`                                                                                        // 结构体嵌套
}

func (x *String) Reset() {
	*x = String{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *String) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*String) ProtoMessage() {}

func (x *String) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use String.ProtoReflect.Descriptor instead.
func (*String) Descriptor() ([]byte, []int) {
	return file_protobuf_hello_proto_rawDescGZIP(), []int{1}
}

func (x *String) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *String) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *String) GetEnabled() bool {
	if x != nil && x.Enabled != nil {
		return *x.Enabled
	}
	return false
}

func (x *String) GetMetricValue() float64 {
	if x != nil {
		return x.MetricValue
	}
	return 0
}

func (x *String) GetValues() map[string]int32 {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *String) GetString1() *String1 {
	if x != nil {
		return x.String1
	}
	return nil
}

type String1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value   string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`                             // 等价于Value string中go变量的写法，probobuf字段默认大写，为全部导出，不可配置为小写
	Key     *string  `protobuf:"bytes,2,opt,name=key,proto3,oneof" json:"key,omitempty"`                           // Key string , 添加了optional则默认是nil了，为入optional后运行为nil "" null
	Enabled *bool    `protobuf:"varint,3,opt,name=enabled,proto3,oneof" json:"enabled,omitempty"`                  // 默认enabled=false,  enabled != nil 时才添加过滤条件
	Option  []Option `protobuf:"varint,5,rep,packed,name=option,proto3,enum=hello.Option" json:"option,omitempty"` // Option的数组
}

func (x *String1) Reset() {
	*x = String1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *String1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*String1) ProtoMessage() {}

func (x *String1) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use String1.ProtoReflect.Descriptor instead.
func (*String1) Descriptor() ([]byte, []int) {
	return file_protobuf_hello_proto_rawDescGZIP(), []int{2}
}

func (x *String1) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *String1) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *String1) GetEnabled() bool {
	if x != nil && x.Enabled != nil {
		return *x.Enabled
	}
	return false
}

func (x *String1) GetOption() []Option {
	if x != nil {
		return x.Option
	}
	return nil
}

type Sub1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Sub1) Reset() {
	*x = Sub1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_hello_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sub1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sub1) ProtoMessage() {}

func (x *Sub1) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_hello_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sub1.ProtoReflect.Descriptor instead.
func (*Sub1) Descriptor() ([]byte, []int) {
	return file_protobuf_hello_proto_rawDescGZIP(), []int{3}
}

func (x *Sub1) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Sub2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Sub2) Reset() {
	*x = Sub2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_hello_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sub2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sub2) ProtoMessage() {}

func (x *Sub2) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_hello_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sub2.ProtoReflect.Descriptor instead.
func (*Sub2) Descriptor() ([]byte, []int) {
	return file_protobuf_hello_proto_rawDescGZIP(), []int{4}
}

func (x *Sub2) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SampleMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to TestOneOf:
	//	*SampleMessage_Sub1
	//	*SampleMessage_Sub2
	TestOneOf isSampleMessage_TestOneOf `protobuf_oneof:"test_one_of"`
}

func (x *SampleMessage) Reset() {
	*x = SampleMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_hello_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleMessage) ProtoMessage() {}

func (x *SampleMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_hello_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleMessage.ProtoReflect.Descriptor instead.
func (*SampleMessage) Descriptor() ([]byte, []int) {
	return file_protobuf_hello_proto_rawDescGZIP(), []int{5}
}

func (m *SampleMessage) GetTestOneOf() isSampleMessage_TestOneOf {
	if m != nil {
		return m.TestOneOf
	}
	return nil
}

func (x *SampleMessage) GetSub1() *Sub1 {
	if x, ok := x.GetTestOneOf().(*SampleMessage_Sub1); ok {
		return x.Sub1
	}
	return nil
}

func (x *SampleMessage) GetSub2() *Sub2 {
	if x, ok := x.GetTestOneOf().(*SampleMessage_Sub2); ok {
		return x.Sub2
	}
	return nil
}

type isSampleMessage_TestOneOf interface {
	isSampleMessage_TestOneOf()
}

type SampleMessage_Sub1 struct {
	Sub1 *Sub1 `protobuf:"bytes,1,opt,name=sub1,proto3,oneof"`
}

type SampleMessage_Sub2 struct {
	Sub2 *Sub2 `protobuf:"bytes,2,opt,name=sub2,proto3,oneof"`
}

func (*SampleMessage_Sub1) isSampleMessage_TestOneOf() {}

func (*SampleMessage_Sub2) isSampleMessage_TestOneOf() {}

var File_protobuf_hello_proto protoreflect.FileDescriptor

var file_protobuf_hello_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0xa3, 0x02, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x31, 0x0a, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12,
	0x28, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31,
	0x52, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x1a, 0x39, 0x0a, 0x0b, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6b, 0x65, 0x79, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x9c, 0x01, 0x0a, 0x07, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x88, 0x01,
	0x01, 0x12, 0x1d, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x48, 0x01, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x25, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x0d, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6b, 0x65, 0x79, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4a, 0x04, 0x08, 0x04, 0x10,
	0x05, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x0b, 0x22, 0x1a, 0x0a, 0x04, 0x53, 0x75, 0x62, 0x31, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x1a, 0x0a, 0x04, 0x53, 0x75, 0x62, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x64, 0x0a, 0x0d, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x21, 0x0a, 0x04, 0x73, 0x75, 0x62, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x31, 0x48, 0x00, 0x52, 0x04, 0x73,
	0x75, 0x62, 0x31, 0x12, 0x21, 0x0a, 0x04, 0x73, 0x75, 0x62, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x32, 0x48, 0x00,
	0x52, 0x04, 0x73, 0x75, 0x62, 0x32, 0x42, 0x0d, 0x0a, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6f,
	0x6e, 0x65, 0x5f, 0x6f, 0x66, 0x2a, 0x28, 0x0a, 0x06, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x05, 0x0a, 0x01, 0x41, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x42, 0x10, 0x01, 0x12, 0x05, 0x0a,
	0x01, 0x43, 0x10, 0x02, 0x12, 0x05, 0x0a, 0x01, 0x44, 0x10, 0x00, 0x1a, 0x02, 0x10, 0x01, 0x42,
	0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x61,
	0x63, 0x6b, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x64, 0x65, 0x6d,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protobuf_hello_proto_rawDescOnce sync.Once
	file_protobuf_hello_proto_rawDescData = file_protobuf_hello_proto_rawDesc
)

func file_protobuf_hello_proto_rawDescGZIP() []byte {
	file_protobuf_hello_proto_rawDescOnce.Do(func() {
		file_protobuf_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_hello_proto_rawDescData)
	})
	return file_protobuf_hello_proto_rawDescData
}

var file_protobuf_hello_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobuf_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protobuf_hello_proto_goTypes = []interface{}{
	(Option)(0),           // 0: hello.Option
	(*Response)(nil),      // 1: hello.Response
	(*String)(nil),        // 2: hello.String
	(*String1)(nil),       // 3: hello.String1
	(*Sub1)(nil),          // 4: hello.Sub1
	(*Sub2)(nil),          // 5: hello.Sub2
	(*SampleMessage)(nil), // 6: hello.SampleMessage
	nil,                   // 7: hello.String.ValuesEntry
	(*anypb.Any)(nil),     // 8: google.protobuf.Any
}
var file_protobuf_hello_proto_depIdxs = []int32{
	8, // 0: hello.Response.data:type_name -> google.protobuf.Any
	7, // 1: hello.String.values:type_name -> hello.String.ValuesEntry
	3, // 2: hello.String.string1:type_name -> hello.String1
	0, // 3: hello.String1.option:type_name -> hello.Option
	4, // 4: hello.SampleMessage.sub1:type_name -> hello.Sub1
	5, // 5: hello.SampleMessage.sub2:type_name -> hello.Sub2
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_protobuf_hello_proto_init() }
func file_protobuf_hello_proto_init() {
	if File_protobuf_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*String); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*String1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_hello_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sub1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_hello_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sub2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_hello_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_protobuf_hello_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_protobuf_hello_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_protobuf_hello_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*SampleMessage_Sub1)(nil),
		(*SampleMessage_Sub2)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_hello_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_hello_proto_goTypes,
		DependencyIndexes: file_protobuf_hello_proto_depIdxs,
		EnumInfos:         file_protobuf_hello_proto_enumTypes,
		MessageInfos:      file_protobuf_hello_proto_msgTypes,
	}.Build()
	File_protobuf_hello_proto = out.File
	file_protobuf_hello_proto_rawDesc = nil
	file_protobuf_hello_proto_goTypes = nil
	file_protobuf_hello_proto_depIdxs = nil
}
