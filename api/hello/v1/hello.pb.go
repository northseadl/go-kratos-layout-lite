// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: hello/v1/hello.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_v1_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_v1_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_hello_v1_hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_v1_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_hello_v1_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_hello_v1_hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type HiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HiRequest) Reset() {
	*x = HiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_v1_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HiRequest) ProtoMessage() {}

func (x *HiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_v1_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HiRequest.ProtoReflect.Descriptor instead.
func (*HiRequest) Descriptor() ([]byte, []int) {
	return file_hello_v1_hello_proto_rawDescGZIP(), []int{2}
}

func (x *HiRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type HiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HiReply) Reset() {
	*x = HiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_v1_hello_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HiReply) ProtoMessage() {}

func (x *HiReply) ProtoReflect() protoreflect.Message {
	mi := &file_hello_v1_hello_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HiReply.ProtoReflect.Descriptor instead.
func (*HiReply) Descriptor() ([]byte, []int) {
	return file_hello_v1_hello_proto_rawDescGZIP(), []int{3}
}

func (x *HiReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_hello_v1_hello_proto protoreflect.FileDescriptor

var file_hello_v1_hello_proto_rawDesc = []byte{
	0x0a, 0x14, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0c, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x1e, 0x0a,
	0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x26, 0x0a,
	0x09, 0x48, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x03,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x1b, 0x0a, 0x07, 0x48, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x32, 0xa5, 0x01, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x50, 0x0a, 0x05,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x4a,
	0x0a, 0x02, 0x48, 0x69, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x69, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76,
	0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x3a, 0x01, 0x2a, 0x42, 0x37, 0x0a, 0x0c, 0x61, 0x70,
	0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x25, 0x67, 0x6f,
	0x2d, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2d, 0x6c,
	0x69, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hello_v1_hello_proto_rawDescOnce sync.Once
	file_hello_v1_hello_proto_rawDescData = file_hello_v1_hello_proto_rawDesc
)

func file_hello_v1_hello_proto_rawDescGZIP() []byte {
	file_hello_v1_hello_proto_rawDescOnce.Do(func() {
		file_hello_v1_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_v1_hello_proto_rawDescData)
	})
	return file_hello_v1_hello_proto_rawDescData
}

var file_hello_v1_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_hello_v1_hello_proto_goTypes = []interface{}{
	(*HelloRequest)(nil), // 0: api.hello.v1.HelloRequest
	(*HelloReply)(nil),   // 1: api.hello.v1.HelloReply
	(*HiRequest)(nil),    // 2: api.hello.v1.HiRequest
	(*HiReply)(nil),      // 3: api.hello.v1.HiReply
}
var file_hello_v1_hello_proto_depIdxs = []int32{
	0, // 0: api.hello.v1.Hello.Hello:input_type -> api.hello.v1.HelloRequest
	2, // 1: api.hello.v1.Hello.Hi:input_type -> api.hello.v1.HiRequest
	1, // 2: api.hello.v1.Hello.Hello:output_type -> api.hello.v1.HelloReply
	3, // 3: api.hello.v1.Hello.Hi:output_type -> api.hello.v1.HiReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hello_v1_hello_proto_init() }
func file_hello_v1_hello_proto_init() {
	if File_hello_v1_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_v1_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_hello_v1_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_hello_v1_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HiRequest); i {
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
		file_hello_v1_hello_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HiReply); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hello_v1_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hello_v1_hello_proto_goTypes,
		DependencyIndexes: file_hello_v1_hello_proto_depIdxs,
		MessageInfos:      file_hello_v1_hello_proto_msgTypes,
	}.Build()
	File_hello_v1_hello_proto = out.File
	file_hello_v1_hello_proto_rawDesc = nil
	file_hello_v1_hello_proto_goTypes = nil
	file_hello_v1_hello_proto_depIdxs = nil
}
