// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.26.0--rc2
// source: mmesh/protobuf/common/v1/status/status.proto

package status

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

type StatusCode int32

const (
	StatusCode_UNDEFINED   StatusCode = 0
	StatusCode_OK          StatusCode = 11
	StatusCode_CANCELED    StatusCode = 21
	StatusCode_INTERRUPTED StatusCode = 31
	StatusCode_INCOMPLETE  StatusCode = 41
	StatusCode_SUSPENDED   StatusCode = 51
	StatusCode_FAILED      StatusCode = 101
)

// Enum value maps for StatusCode.
var (
	StatusCode_name = map[int32]string{
		0:   "UNDEFINED",
		11:  "OK",
		21:  "CANCELED",
		31:  "INTERRUPTED",
		41:  "INCOMPLETE",
		51:  "SUSPENDED",
		101: "FAILED",
	}
	StatusCode_value = map[string]int32{
		"UNDEFINED":   0,
		"OK":          11,
		"CANCELED":    21,
		"INTERRUPTED": 31,
		"INCOMPLETE":  41,
		"SUSPENDED":   51,
		"FAILED":      101,
	}
)

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}

func (x StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_common_v1_status_status_proto_enumTypes[0].Descriptor()
}

func (StatusCode) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_common_v1_status_status_proto_enumTypes[0]
}

func (x StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCode.Descriptor instead.
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_common_v1_status_status_proto_rawDescGZIP(), []int{0}
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceID  string     `protobuf:"bytes,1,opt,name=sourceID,proto3" json:"sourceID,omitempty"`
	Code      StatusCode `protobuf:"varint,2,opt,name=code,proto3,enum=status.StatusCode" json:"code,omitempty"`
	Message   string     `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp int64      `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_common_v1_status_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_common_v1_status_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_common_v1_status_status_proto_rawDescGZIP(), []int{0}
}

func (x *StatusResponse) GetSourceID() string {
	if x != nil {
		return x.SourceID
	}
	return ""
}

func (x *StatusResponse) GetCode() StatusCode {
	if x != nil {
		return x.Code
	}
	return StatusCode_UNDEFINED
}

func (x *StatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *StatusResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_mmesh_protobuf_common_v1_status_status_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_common_v1_status_status_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2a, 0x6d, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x0b, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41,
	0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x15, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x54, 0x45,
	0x52, 0x52, 0x55, 0x50, 0x54, 0x45, 0x44, 0x10, 0x1f, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x43,
	0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x29, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x53,
	0x50, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x33, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65,
	0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_common_v1_status_status_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_common_v1_status_status_proto_rawDescData = file_mmesh_protobuf_common_v1_status_status_proto_rawDesc
)

func file_mmesh_protobuf_common_v1_status_status_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_common_v1_status_status_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_common_v1_status_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_common_v1_status_status_proto_rawDescData)
	})
	return file_mmesh_protobuf_common_v1_status_status_proto_rawDescData
}

var file_mmesh_protobuf_common_v1_status_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mmesh_protobuf_common_v1_status_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mmesh_protobuf_common_v1_status_status_proto_goTypes = []interface{}{
	(StatusCode)(0),        // 0: status.StatusCode
	(*StatusResponse)(nil), // 1: status.StatusResponse
}
var file_mmesh_protobuf_common_v1_status_status_proto_depIdxs = []int32{
	0, // 0: status.StatusResponse.code:type_name -> status.StatusCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_common_v1_status_status_proto_init() }
func file_mmesh_protobuf_common_v1_status_status_proto_init() {
	if File_mmesh_protobuf_common_v1_status_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_common_v1_status_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
			RawDescriptor: file_mmesh_protobuf_common_v1_status_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_common_v1_status_status_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_common_v1_status_status_proto_depIdxs,
		EnumInfos:         file_mmesh_protobuf_common_v1_status_status_proto_enumTypes,
		MessageInfos:      file_mmesh_protobuf_common_v1_status_status_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_common_v1_status_status_proto = out.File
	file_mmesh_protobuf_common_v1_status_status_proto_rawDesc = nil
	file_mmesh_protobuf_common_v1_status_status_proto_goTypes = nil
	file_mmesh_protobuf_common_v1_status_status_proto_depIdxs = nil
}
