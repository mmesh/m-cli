// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: mmesh/protobuf/resources/v1/iam/sshKey.proto

package iam

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

type SSHKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// string ID = 1;
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PublicKey string `protobuf:"bytes,11,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
}

func (x *SSHKey) Reset() {
	*x = SSHKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_iam_sshKey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSHKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSHKey) ProtoMessage() {}

func (x *SSHKey) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_iam_sshKey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSHKey.ProtoReflect.Descriptor instead.
func (*SSHKey) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_iam_sshKey_proto_rawDescGZIP(), []int{0}
}

func (x *SSHKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SSHKey) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

var File_mmesh_protobuf_resources_v1_iam_sshKey_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_iam_sshKey_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x73, 0x73, 0x68, 0x4b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x69, 0x61, 0x6d, 0x22, 0x3a, 0x0a, 0x06, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x42,
	0x27, 0x5a, 0x25, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61,
	0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_iam_sshKey_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_iam_sshKey_proto_rawDescData = file_mmesh_protobuf_resources_v1_iam_sshKey_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_iam_sshKey_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_iam_sshKey_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_iam_sshKey_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_iam_sshKey_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_iam_sshKey_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_iam_sshKey_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mmesh_protobuf_resources_v1_iam_sshKey_proto_goTypes = []interface{}{
	(*SSHKey)(nil), // 0: iam.SSHKey
}
var file_mmesh_protobuf_resources_v1_iam_sshKey_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_iam_sshKey_proto_init() }
func file_mmesh_protobuf_resources_v1_iam_sshKey_proto_init() {
	if File_mmesh_protobuf_resources_v1_iam_sshKey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_iam_sshKey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSHKey); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_iam_sshKey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_iam_sshKey_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_iam_sshKey_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_iam_sshKey_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_iam_sshKey_proto = out.File
	file_mmesh_protobuf_resources_v1_iam_sshKey_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_iam_sshKey_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_iam_sshKey_proto_depIdxs = nil
}
