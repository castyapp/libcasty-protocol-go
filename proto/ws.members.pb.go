// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.1
// source: ws.members.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TheaterMembers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Members []*User `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *TheaterMembers) Reset() {
	*x = TheaterMembers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ws_members_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TheaterMembers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TheaterMembers) ProtoMessage() {}

func (x *TheaterMembers) ProtoReflect() protoreflect.Message {
	mi := &file_ws_members_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TheaterMembers.ProtoReflect.Descriptor instead.
func (*TheaterMembers) Descriptor() ([]byte, []int) {
	return file_ws_members_proto_rawDescGZIP(), []int{0}
}

func (x *TheaterMembers) GetMembers() []*User {
	if x != nil {
		return x.Members
	}
	return nil
}

var File_ws_members_proto protoreflect.FileDescriptor

var file_ws_members_proto_rawDesc = []byte{
	0x0a, 0x10, 0x77, 0x73, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x0e, 0x54, 0x68,
	0x65, 0x61, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x07,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ws_members_proto_rawDescOnce sync.Once
	file_ws_members_proto_rawDescData = file_ws_members_proto_rawDesc
)

func file_ws_members_proto_rawDescGZIP() []byte {
	file_ws_members_proto_rawDescOnce.Do(func() {
		file_ws_members_proto_rawDescData = protoimpl.X.CompressGZIP(file_ws_members_proto_rawDescData)
	})
	return file_ws_members_proto_rawDescData
}

var file_ws_members_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ws_members_proto_goTypes = []interface{}{
	(*TheaterMembers)(nil), // 0: proto.TheaterMembers
	(*User)(nil),           // 1: proto.User
}
var file_ws_members_proto_depIdxs = []int32{
	1, // 0: proto.TheaterMembers.members:type_name -> proto.User
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ws_members_proto_init() }
func file_ws_members_proto_init() {
	if File_ws_members_proto != nil {
		return
	}
	file_grpc_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ws_members_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TheaterMembers); i {
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
			RawDescriptor: file_ws_members_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ws_members_proto_goTypes,
		DependencyIndexes: file_ws_members_proto_depIdxs,
		MessageInfos:      file_ws_members_proto_msgTypes,
	}.Build()
	File_ws_members_proto = out.File
	file_ws_members_proto_rawDesc = nil
	file_ws_members_proto_goTypes = nil
	file_ws_members_proto_depIdxs = nil
}
