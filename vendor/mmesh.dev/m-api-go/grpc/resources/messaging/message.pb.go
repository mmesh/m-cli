// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: mmesh/protobuf/resources/v1/messaging/message.proto

package messaging

import (
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

type MessageStatus int32

const (
	MessageStatus_SENT   MessageStatus = 0
	MessageStatus_FAILED MessageStatus = 1
)

// Enum value maps for MessageStatus.
var (
	MessageStatus_name = map[int32]string{
		0: "SENT",
		1: "FAILED",
	}
	MessageStatus_value = map[string]int32{
		"SENT":   0,
		"FAILED": 1,
	}
)

func (x MessageStatus) Enum() *MessageStatus {
	p := new(MessageStatus)
	*p = x
	return p
}

func (x MessageStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_messaging_message_proto_enumTypes[0].Descriptor()
}

func (MessageStatus) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_messaging_message_proto_enumTypes[0]
}

func (x MessageStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageStatus.Descriptor instead.
func (MessageStatus) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_messaging_message_proto_rawDescGZIP(), []int{0}
}

type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_messaging_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_messaging_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_messaging_message_proto_rawDescGZIP(), []int{0}
}

func (x *Contact) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Contact) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_mmesh_protobuf_resources_v1_messaging_message_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_messaging_message_proto_rawDesc = []byte{
	0x0a, 0x33, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x22, 0x37, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2a, 0x25, 0x0a, 0x0d, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45,
	0x4e, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01,
	0x42, 0x2d, 0x5a, 0x2b, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_messaging_message_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_messaging_message_proto_rawDescData = file_mmesh_protobuf_resources_v1_messaging_message_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_messaging_message_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_messaging_message_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_messaging_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_messaging_message_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_messaging_message_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_messaging_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mmesh_protobuf_resources_v1_messaging_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mmesh_protobuf_resources_v1_messaging_message_proto_goTypes = []interface{}{
	(MessageStatus)(0), // 0: messaging.MessageStatus
	(*Contact)(nil),    // 1: messaging.Contact
}
var file_mmesh_protobuf_resources_v1_messaging_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_messaging_message_proto_init() }
func file_mmesh_protobuf_resources_v1_messaging_message_proto_init() {
	if File_mmesh_protobuf_resources_v1_messaging_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_messaging_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_messaging_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_messaging_message_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_messaging_message_proto_depIdxs,
		EnumInfos:         file_mmesh_protobuf_resources_v1_messaging_message_proto_enumTypes,
		MessageInfos:      file_mmesh_protobuf_resources_v1_messaging_message_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_messaging_message_proto = out.File
	file_mmesh_protobuf_resources_v1_messaging_message_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_messaging_message_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_messaging_message_proto_depIdxs = nil
}
