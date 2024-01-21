// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: mmesh/protobuf/resources/v1/hsec/common.proto

package hsec

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

type Layer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Digest    string `protobuf:"bytes,11,opt,name=digest,proto3" json:"digest,omitempty"`
	DiffID    string `protobuf:"bytes,21,opt,name=diffID,proto3" json:"diffID,omitempty"`
	CreatedBy string `protobuf:"bytes,31,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
}

func (x *Layer) Reset() {
	*x = Layer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_hsec_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Layer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Layer) ProtoMessage() {}

func (x *Layer) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_hsec_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Layer.ProtoReflect.Descriptor instead.
func (*Layer) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_hsec_common_proto_rawDescGZIP(), []int{0}
}

func (x *Layer) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *Layer) GetDiffID() string {
	if x != nil {
		return x.DiffID
	}
	return ""
}

func (x *Layer) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartLine int32 `protobuf:"varint,11,opt,name=startLine,proto3" json:"startLine,omitempty"`
	EndLine   int32 `protobuf:"varint,21,opt,name=endLine,proto3" json:"endLine,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_hsec_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_hsec_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_hsec_common_proto_rawDescGZIP(), []int{1}
}

func (x *Location) GetStartLine() int32 {
	if x != nil {
		return x.StartLine
	}
	return 0
}

func (x *Location) GetEndLine() int32 {
	if x != nil {
		return x.EndLine
	}
	return 0
}

type Code struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lines []*Line `protobuf:"bytes,11,rep,name=lines,proto3" json:"lines,omitempty"`
}

func (x *Code) Reset() {
	*x = Code{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_hsec_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code) ProtoMessage() {}

func (x *Code) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_hsec_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Code.ProtoReflect.Descriptor instead.
func (*Code) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_hsec_common_proto_rawDescGZIP(), []int{2}
}

func (x *Code) GetLines() []*Line {
	if x != nil {
		return x.Lines
	}
	return nil
}

type Line struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number      int32  `protobuf:"varint,11,opt,name=number,proto3" json:"number,omitempty"`
	Content     string `protobuf:"bytes,21,opt,name=content,proto3" json:"content,omitempty"`
	IsCause     bool   `protobuf:"varint,31,opt,name=isCause,proto3" json:"isCause,omitempty"`
	Annotation  string `protobuf:"bytes,41,opt,name=annotation,proto3" json:"annotation,omitempty"`
	Truncated   bool   `protobuf:"varint,51,opt,name=truncated,proto3" json:"truncated,omitempty"`
	Highlighted string `protobuf:"bytes,61,opt,name=highlighted,proto3" json:"highlighted,omitempty"`
	FirstCause  bool   `protobuf:"varint,71,opt,name=firstCause,proto3" json:"firstCause,omitempty"`
	LastCause   bool   `protobuf:"varint,81,opt,name=lastCause,proto3" json:"lastCause,omitempty"`
}

func (x *Line) Reset() {
	*x = Line{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_hsec_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Line) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Line) ProtoMessage() {}

func (x *Line) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_hsec_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Line.ProtoReflect.Descriptor instead.
func (*Line) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_hsec_common_proto_rawDescGZIP(), []int{3}
}

func (x *Line) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Line) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Line) GetIsCause() bool {
	if x != nil {
		return x.IsCause
	}
	return false
}

func (x *Line) GetAnnotation() string {
	if x != nil {
		return x.Annotation
	}
	return ""
}

func (x *Line) GetTruncated() bool {
	if x != nil {
		return x.Truncated
	}
	return false
}

func (x *Line) GetHighlighted() string {
	if x != nil {
		return x.Highlighted
	}
	return ""
}

func (x *Line) GetFirstCause() bool {
	if x != nil {
		return x.FirstCause
	}
	return false
}

func (x *Line) GetLastCause() bool {
	if x != nil {
		return x.LastCause
	}
	return false
}

var File_mmesh_protobuf_resources_v1_hsec_common_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_hsec_common_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x73,
	0x65, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x68, 0x73, 0x65, 0x63, 0x22, 0x55, 0x0a, 0x05, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x66, 0x66, 0x49, 0x44,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x66, 0x66, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x1f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x42, 0x0a, 0x08,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x4c, 0x69, 0x6e, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e,
	0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65,
	0x22, 0x28, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x65,
	0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x68, 0x73, 0x65, 0x63, 0x2e, 0x4c,
	0x69, 0x6e, 0x65, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x22, 0xf0, 0x01, 0x0a, 0x04, 0x4c,
	0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x43, 0x61, 0x75, 0x73, 0x65,
	0x18, 0x1f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x43, 0x61, 0x75, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x29, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x33, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x74, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x68, 0x69, 0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x18, 0x3d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x68, 0x69, 0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x43, 0x61, 0x75, 0x73, 0x65, 0x18, 0x47, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x43, 0x61, 0x75, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x61, 0x75, 0x73, 0x65, 0x18, 0x51, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x61, 0x75, 0x73, 0x65, 0x42, 0x28, 0x5a,
	0x26, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69,
	0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x68, 0x73, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_hsec_common_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_hsec_common_proto_rawDescData = file_mmesh_protobuf_resources_v1_hsec_common_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_hsec_common_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_hsec_common_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_hsec_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_hsec_common_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_hsec_common_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_hsec_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mmesh_protobuf_resources_v1_hsec_common_proto_goTypes = []interface{}{
	(*Layer)(nil),    // 0: hsec.Layer
	(*Location)(nil), // 1: hsec.Location
	(*Code)(nil),     // 2: hsec.Code
	(*Line)(nil),     // 3: hsec.Line
}
var file_mmesh_protobuf_resources_v1_hsec_common_proto_depIdxs = []int32{
	3, // 0: hsec.Code.lines:type_name -> hsec.Line
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_hsec_common_proto_init() }
func file_mmesh_protobuf_resources_v1_hsec_common_proto_init() {
	if File_mmesh_protobuf_resources_v1_hsec_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_hsec_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Layer); i {
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
		file_mmesh_protobuf_resources_v1_hsec_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_mmesh_protobuf_resources_v1_hsec_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Code); i {
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
		file_mmesh_protobuf_resources_v1_hsec_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Line); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_hsec_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_hsec_common_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_hsec_common_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_hsec_common_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_hsec_common_proto = out.File
	file_mmesh_protobuf_resources_v1_hsec_common_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_hsec_common_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_hsec_common_proto_depIdxs = nil
}
