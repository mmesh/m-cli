// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.26.0--rc2
// source: mmesh/protobuf/resources/v1/nstore/netdb/ctlog.proto

package netdb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	nstore "mmesh.dev/m-api-go/grpc/resources/nstore"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConntrackLogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp  int64            `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Connection *Connection      `protobuf:"bytes,11,opt,name=connection,proto3" json:"connection,omitempty"`                      // key
	Status     ConnectionStatus `protobuf:"varint,21,opt,name=status,proto3,enum=netdb.ConnectionStatus" json:"status,omitempty"` // value
}

func (x *ConntrackLogEntry) Reset() {
	*x = ConntrackLogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConntrackLogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConntrackLogEntry) ProtoMessage() {}

func (x *ConntrackLogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConntrackLogEntry.ProtoReflect.Descriptor instead.
func (*ConntrackLogEntry) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_rawDescGZIP(), []int{0}
}

func (x *ConntrackLogEntry) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ConntrackLogEntry) GetConnection() *Connection {
	if x != nil {
		return x.Connection
	}
	return nil
}

func (x *ConntrackLogEntry) GetStatus() ConnectionStatus {
	if x != nil {
		return x.Status
	}
	return ConnectionStatus_UNKNOWN_STATUS
}

type ConntrackLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *nstore.DataRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *ConntrackLogRequest) Reset() {
	*x = ConntrackLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConntrackLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConntrackLogRequest) ProtoMessage() {}

func (x *ConntrackLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConntrackLogRequest.ProtoReflect.Descriptor instead.
func (*ConntrackLogRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_rawDescGZIP(), []int{1}
}

func (x *ConntrackLogRequest) GetRequest() *nstore.DataRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type ConntrackLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string               `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	TenantID  string               `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	NodeID    string               `protobuf:"bytes,5,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	QueryID   string               `protobuf:"bytes,11,opt,name=queryID,proto3" json:"queryID,omitempty"`
	CtLog     []*ConntrackLogEntry `protobuf:"bytes,101,rep,name=ctLog,proto3" json:"ctLog,omitempty"`
	Timestamp int64                `protobuf:"varint,1001,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ConntrackLogResponse) Reset() {
	*x = ConntrackLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConntrackLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConntrackLogResponse) ProtoMessage() {}

func (x *ConntrackLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConntrackLogResponse.ProtoReflect.Descriptor instead.
func (*ConntrackLogResponse) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_rawDescGZIP(), []int{2}
}

func (x *ConntrackLogResponse) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *ConntrackLogResponse) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *ConntrackLogResponse) GetNodeID() string {
	if x != nil {
		return x.NodeID
	}
	return ""
}

func (x *ConntrackLogResponse) GetQueryID() string {
	if x != nil {
		return x.QueryID
	}
	return ""
}

func (x *ConntrackLogResponse) GetCtLog() []*ConntrackLogEntry {
	if x != nil {
		return x.CtLog
	}
	return nil
}

func (x *ConntrackLogResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_rawDesc = []byte{
	0x0a, 0x34, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x6e, 0x65, 0x74, 0x64, 0x62, 0x2f, 0x63, 0x74, 0x6c, 0x6f, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6e, 0x65, 0x74, 0x64, 0x62, 0x1a, 0x39, 0x6d,
	0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x6e, 0x65, 0x74, 0x64, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x6e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x11, 0x43, 0x6f,
	0x6e, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x31, 0x0a,
	0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x6e, 0x65, 0x74, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x6e, 0x65, 0x74, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x44, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xd1, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f,
	0x64, 0x65, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x44, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x05,
	0x63, 0x74, 0x4c, 0x6f, 0x67, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65,
	0x74, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x4c, 0x6f, 0x67,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x63, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x1d, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x30, 0x5a, 0x2e, 0x6d,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67,
	0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x6e, 0x65, 0x74, 0x64, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_rawDescData = file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_goTypes = []interface{}{
	(*ConntrackLogEntry)(nil),    // 0: netdb.ConntrackLogEntry
	(*ConntrackLogRequest)(nil),  // 1: netdb.ConntrackLogRequest
	(*ConntrackLogResponse)(nil), // 2: netdb.ConntrackLogResponse
	(*Connection)(nil),           // 3: netdb.Connection
	(ConnectionStatus)(0),        // 4: netdb.ConnectionStatus
	(*nstore.DataRequest)(nil),   // 5: nstore.DataRequest
}
var file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_depIdxs = []int32{
	3, // 0: netdb.ConntrackLogEntry.connection:type_name -> netdb.Connection
	4, // 1: netdb.ConntrackLogEntry.status:type_name -> netdb.ConnectionStatus
	5, // 2: netdb.ConntrackLogRequest.request:type_name -> nstore.DataRequest
	0, // 3: netdb.ConntrackLogResponse.ctLog:type_name -> netdb.ConntrackLogEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_init() }
func file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_init() {
	if File_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto != nil {
		return
	}
	file_mmesh_protobuf_resources_v1_nstore_netdb_connection_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConntrackLogEntry); i {
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
		file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConntrackLogRequest); i {
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
		file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConntrackLogResponse); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto = out.File
	file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_nstore_netdb_ctlog_proto_depIdxs = nil
}
