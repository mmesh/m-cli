// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.26.0--rc2
// source: mmesh/protobuf/network/v1/mmsp/nodemgmt.proto

package mmsp

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	hsecdb "mmesh.dev/m-api-go/grpc/resources/nstore/hsecdb"
	metricsdb "mmesh.dev/m-api-go/grpc/resources/nstore/metricsdb"
	netdb "mmesh.dev/m-api-go/grpc/resources/nstore/netdb"
	topology "mmesh.dev/m-api-go/grpc/resources/topology"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NodeMgmtMsgType int32

const (
	NodeMgmtMsgType_UNDEFINED_NODEMGMT_MSG            NodeMgmtMsgType = 0
	NodeMgmtMsgType_NODE_INIT                         NodeMgmtMsgType = 11
	NodeMgmtMsgType_NODE_CONFIG                       NodeMgmtMsgType = 21
	NodeMgmtMsgType_NODE_HOST_METRICS_REQUEST         NodeMgmtMsgType = 101
	NodeMgmtMsgType_NODE_HOST_METRICS_RESPONSE        NodeMgmtMsgType = 102
	NodeMgmtMsgType_NODE_NET_CONNTRACK_STATE_REQUEST  NodeMgmtMsgType = 201
	NodeMgmtMsgType_NODE_NET_CONNTRACK_STATE_RESPONSE NodeMgmtMsgType = 202
	NodeMgmtMsgType_NODE_NET_CONNTRACK_LOG_REQUEST    NodeMgmtMsgType = 211
	NodeMgmtMsgType_NODE_NET_CONNTRACK_LOG_RESPONSE   NodeMgmtMsgType = 212
	NodeMgmtMsgType_NODE_NET_TRAFFIC_METRICS_REQUEST  NodeMgmtMsgType = 221
	NodeMgmtMsgType_NODE_NET_TRAFFIC_METRICS_RESPONSE NodeMgmtMsgType = 222
	NodeMgmtMsgType_NODE_HOST_SECURITY_REQUEST        NodeMgmtMsgType = 301
	NodeMgmtMsgType_NODE_HOST_SECURITY_RESPONSE       NodeMgmtMsgType = 302
)

// Enum value maps for NodeMgmtMsgType.
var (
	NodeMgmtMsgType_name = map[int32]string{
		0:   "UNDEFINED_NODEMGMT_MSG",
		11:  "NODE_INIT",
		21:  "NODE_CONFIG",
		101: "NODE_HOST_METRICS_REQUEST",
		102: "NODE_HOST_METRICS_RESPONSE",
		201: "NODE_NET_CONNTRACK_STATE_REQUEST",
		202: "NODE_NET_CONNTRACK_STATE_RESPONSE",
		211: "NODE_NET_CONNTRACK_LOG_REQUEST",
		212: "NODE_NET_CONNTRACK_LOG_RESPONSE",
		221: "NODE_NET_TRAFFIC_METRICS_REQUEST",
		222: "NODE_NET_TRAFFIC_METRICS_RESPONSE",
		301: "NODE_HOST_SECURITY_REQUEST",
		302: "NODE_HOST_SECURITY_RESPONSE",
	}
	NodeMgmtMsgType_value = map[string]int32{
		"UNDEFINED_NODEMGMT_MSG":            0,
		"NODE_INIT":                         11,
		"NODE_CONFIG":                       21,
		"NODE_HOST_METRICS_REQUEST":         101,
		"NODE_HOST_METRICS_RESPONSE":        102,
		"NODE_NET_CONNTRACK_STATE_REQUEST":  201,
		"NODE_NET_CONNTRACK_STATE_RESPONSE": 202,
		"NODE_NET_CONNTRACK_LOG_REQUEST":    211,
		"NODE_NET_CONNTRACK_LOG_RESPONSE":   212,
		"NODE_NET_TRAFFIC_METRICS_REQUEST":  221,
		"NODE_NET_TRAFFIC_METRICS_RESPONSE": 222,
		"NODE_HOST_SECURITY_REQUEST":        301,
		"NODE_HOST_SECURITY_RESPONSE":       302,
	}
)

func (x NodeMgmtMsgType) Enum() *NodeMgmtMsgType {
	p := new(NodeMgmtMsgType)
	*p = x
	return p
}

func (x NodeMgmtMsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeMgmtMsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_enumTypes[0].Descriptor()
}

func (NodeMgmtMsgType) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_enumTypes[0]
}

func (x NodeMgmtMsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeMgmtMsgType.Descriptor instead.
func (NodeMgmtMsgType) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_rawDescGZIP(), []int{0}
}

type NodeMgmtConfigActionType int32

const (
	NodeMgmtConfigActionType_CFG_UNDEFINED   NodeMgmtConfigActionType = 0
	NodeMgmtConfigActionType_CFG_METADATA    NodeMgmtConfigActionType = 11
	NodeMgmtConfigActionType_CFG_NETWORKING  NodeMgmtConfigActionType = 21
	NodeMgmtConfigActionType_CFG_MANAGEMENT  NodeMgmtConfigActionType = 31
	NodeMgmtConfigActionType_CFG_MAINTENANCE NodeMgmtConfigActionType = 41
)

// Enum value maps for NodeMgmtConfigActionType.
var (
	NodeMgmtConfigActionType_name = map[int32]string{
		0:  "CFG_UNDEFINED",
		11: "CFG_METADATA",
		21: "CFG_NETWORKING",
		31: "CFG_MANAGEMENT",
		41: "CFG_MAINTENANCE",
	}
	NodeMgmtConfigActionType_value = map[string]int32{
		"CFG_UNDEFINED":   0,
		"CFG_METADATA":    11,
		"CFG_NETWORKING":  21,
		"CFG_MANAGEMENT":  31,
		"CFG_MAINTENANCE": 41,
	}
)

func (x NodeMgmtConfigActionType) Enum() *NodeMgmtConfigActionType {
	p := new(NodeMgmtConfigActionType)
	*p = x
	return p
}

func (x NodeMgmtConfigActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeMgmtConfigActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_enumTypes[1].Descriptor()
}

func (NodeMgmtConfigActionType) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_enumTypes[1]
}

func (x NodeMgmtConfigActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeMgmtConfigActionType.Descriptor instead.
func (NodeMgmtConfigActionType) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_rawDescGZIP(), []int{1}
}

type NodeMgmtPDU struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type                      NodeMgmtMsgType                    `protobuf:"varint,11,opt,name=type,proto3,enum=mmsp.NodeMgmtMsgType" json:"type,omitempty"`
	NodeReq                   *topology.NodeReq                  `protobuf:"bytes,21,opt,name=nodeReq,proto3" json:"nodeReq,omitempty"`                                      // NODE_INIT (required for mmpInit)
	NodeConfig                *NodeMgmtConfig                    `protobuf:"bytes,31,opt,name=nodeConfig,proto3" json:"nodeConfig,omitempty"`                                // NODE_CONFIG (required for nodeConfig)
	HostMetricsRequest        *metricsdb.HostMetricsRequest      `protobuf:"bytes,101,opt,name=hostMetricsRequest,proto3" json:"hostMetricsRequest,omitempty"`               // NODE_HOST_METRICS_REQUEST
	HostMetricsResponse       *metricsdb.HostMetricsResponse     `protobuf:"bytes,102,opt,name=hostMetricsResponse,proto3" json:"hostMetricsResponse,omitempty"`             // NODE_HOST_METRICS_RESPONSE
	NetCtStateRequest         *netdb.ConntrackTableRequest       `protobuf:"bytes,201,opt,name=netCtStateRequest,proto3" json:"netCtStateRequest,omitempty"`                 // NODE_NET_CONNTRACK_STATE_REQUEST
	NetCtStateResponse        *netdb.ConntrackTableResponse      `protobuf:"bytes,202,opt,name=netCtStateResponse,proto3" json:"netCtStateResponse,omitempty"`               // NODE_NET_CONNTRACK_STATE_RESPONSE
	NetCtLogRequest           *netdb.ConntrackLogRequest         `protobuf:"bytes,211,opt,name=netCtLogRequest,proto3" json:"netCtLogRequest,omitempty"`                     // NODE_NET_CONNTRACK_LOG_REQUEST
	NetCtLogResponse          *netdb.ConntrackLogResponse        `protobuf:"bytes,212,opt,name=netCtLogResponse,proto3" json:"netCtLogResponse,omitempty"`                   // NODE_NET_CONNTRACK_LOG_RESPONSE
	NetTrafficMetricsRequest  *netdb.TrafficMetricsRequest       `protobuf:"bytes,221,opt,name=netTrafficMetricsRequest,proto3" json:"netTrafficMetricsRequest,omitempty"`   // NODE_NET_TRAFFIC_METRICS_REQUEST
	NetTrafficMetricsResponse *netdb.TrafficMetricsResponse      `protobuf:"bytes,222,opt,name=netTrafficMetricsResponse,proto3" json:"netTrafficMetricsResponse,omitempty"` // NODE_NET_TRAFFIC_METRICS_RESPONSE
	HsecReportRequest         *hsecdb.HostSecurityReportRequest  `protobuf:"bytes,301,opt,name=hsecReportRequest,proto3" json:"hsecReportRequest,omitempty"`                 // NODE_HOST_SECURITY_REQUEST
	HsecReportResponse        *hsecdb.HostSecurityReportResponse `protobuf:"bytes,302,opt,name=hsecReportResponse,proto3" json:"hsecReportResponse,omitempty"`               // NODE_HOST_SECURITY_RESPONSE
}

func (x *NodeMgmtPDU) Reset() {
	*x = NodeMgmtPDU{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeMgmtPDU) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMgmtPDU) ProtoMessage() {}

func (x *NodeMgmtPDU) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMgmtPDU.ProtoReflect.Descriptor instead.
func (*NodeMgmtPDU) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_rawDescGZIP(), []int{0}
}

func (x *NodeMgmtPDU) GetType() NodeMgmtMsgType {
	if x != nil {
		return x.Type
	}
	return NodeMgmtMsgType_UNDEFINED_NODEMGMT_MSG
}

func (x *NodeMgmtPDU) GetNodeReq() *topology.NodeReq {
	if x != nil {
		return x.NodeReq
	}
	return nil
}

func (x *NodeMgmtPDU) GetNodeConfig() *NodeMgmtConfig {
	if x != nil {
		return x.NodeConfig
	}
	return nil
}

func (x *NodeMgmtPDU) GetHostMetricsRequest() *metricsdb.HostMetricsRequest {
	if x != nil {
		return x.HostMetricsRequest
	}
	return nil
}

func (x *NodeMgmtPDU) GetHostMetricsResponse() *metricsdb.HostMetricsResponse {
	if x != nil {
		return x.HostMetricsResponse
	}
	return nil
}

func (x *NodeMgmtPDU) GetNetCtStateRequest() *netdb.ConntrackTableRequest {
	if x != nil {
		return x.NetCtStateRequest
	}
	return nil
}

func (x *NodeMgmtPDU) GetNetCtStateResponse() *netdb.ConntrackTableResponse {
	if x != nil {
		return x.NetCtStateResponse
	}
	return nil
}

func (x *NodeMgmtPDU) GetNetCtLogRequest() *netdb.ConntrackLogRequest {
	if x != nil {
		return x.NetCtLogRequest
	}
	return nil
}

func (x *NodeMgmtPDU) GetNetCtLogResponse() *netdb.ConntrackLogResponse {
	if x != nil {
		return x.NetCtLogResponse
	}
	return nil
}

func (x *NodeMgmtPDU) GetNetTrafficMetricsRequest() *netdb.TrafficMetricsRequest {
	if x != nil {
		return x.NetTrafficMetricsRequest
	}
	return nil
}

func (x *NodeMgmtPDU) GetNetTrafficMetricsResponse() *netdb.TrafficMetricsResponse {
	if x != nil {
		return x.NetTrafficMetricsResponse
	}
	return nil
}

func (x *NodeMgmtPDU) GetHsecReportRequest() *hsecdb.HostSecurityReportRequest {
	if x != nil {
		return x.HsecReportRequest
	}
	return nil
}

func (x *NodeMgmtPDU) GetHsecReportResponse() *hsecdb.HostSecurityReportResponse {
	if x != nil {
		return x.HsecReportResponse
	}
	return nil
}

type NodeMgmtConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type NodeMgmtConfigActionType `protobuf:"varint,1,opt,name=type,proto3,enum=mmsp.NodeMgmtConfigActionType" json:"type,omitempty"`
	Cfg  *topology.NodeCfg        `protobuf:"bytes,11,opt,name=cfg,proto3" json:"cfg,omitempty"`
}

func (x *NodeMgmtConfig) Reset() {
	*x = NodeMgmtConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeMgmtConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMgmtConfig) ProtoMessage() {}

func (x *NodeMgmtConfig) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMgmtConfig.ProtoReflect.Descriptor instead.
func (*NodeMgmtConfig) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_rawDescGZIP(), []int{1}
}

func (x *NodeMgmtConfig) GetType() NodeMgmtConfigActionType {
	if x != nil {
		return x.Type
	}
	return NodeMgmtConfigActionType_CFG_UNDEFINED
}

func (x *NodeMgmtConfig) GetCfg() *topology.NodeCfg {
	if x != nil {
		return x.Cfg
	}
	return nil
}

var File_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6d, 0x73, 0x70,
	0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x6d, 0x6d, 0x73, 0x70, 0x1a, 0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2f, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x64, 0x62, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x64, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x6e, 0x65, 0x74, 0x64, 0x62, 0x2f, 0x63,
	0x74, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x6d, 0x6d,
	0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2f, 0x6e, 0x65, 0x74, 0x64, 0x62, 0x2f, 0x63, 0x74, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x3d, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x6e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x6e, 0x65, 0x74, 0x64, 0x62, 0x2f, 0x6e, 0x65, 0x74,
	0x66, 0x6c, 0x6f, 0x77, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x36, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x68, 0x73, 0x65, 0x63, 0x64, 0x62, 0x2f, 0x68, 0x73, 0x65,
	0x63, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x07, 0x0a, 0x0b, 0x4e, 0x6f,
	0x64, 0x65, 0x4d, 0x67, 0x6d, 0x74, 0x50, 0x44, 0x55, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6d, 0x6d, 0x73, 0x70, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x4d, 0x67, 0x6d, 0x74, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x34, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6d, 0x73, 0x70, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x4d, 0x67, 0x6d, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0a, 0x6e, 0x6f, 0x64,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4d, 0x0a, 0x12, 0x68, 0x6f, 0x73, 0x74, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x65, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x64, 0x62, 0x2e,
	0x48, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x12, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x13, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x66, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x64, 0x62, 0x2e,
	0x48, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x13, 0x68, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x11, 0x6e, 0x65, 0x74, 0x43,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0xc9, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x74, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6e,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x11, 0x6e, 0x65, 0x74, 0x43, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x12, 0x6e, 0x65, 0x74, 0x43, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0xca, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x65, 0x74, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x12, 0x6e, 0x65, 0x74, 0x43, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0f, 0x6e, 0x65, 0x74, 0x43, 0x74, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0xd3, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x6e, 0x65, 0x74, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0f, 0x6e, 0x65, 0x74,
	0x43, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x10,
	0x6e, 0x65, 0x74, 0x43, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0xd4, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x65, 0x74, 0x64, 0x62, 0x2e,
	0x43, 0x6f, 0x6e, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x10, 0x6e, 0x65, 0x74, 0x43, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x18, 0x6e, 0x65, 0x74, 0x54, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0xdd, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x74, 0x64,
	0x62, 0x2e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x18, 0x6e, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x5c, 0x0a, 0x19, 0x6e, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0xde,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x65, 0x74, 0x64, 0x62, 0x2e, 0x54, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x19, 0x6e, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x50, 0x0a, 0x11, 0x68, 0x73, 0x65, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0xad, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x68, 0x73,
	0x65, 0x63, 0x64, 0x62, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x11,
	0x68, 0x73, 0x65, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x53, 0x0a, 0x12, 0x68, 0x73, 0x65, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0xae, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x68, 0x73, 0x65, 0x63, 0x64, 0x62, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x12, 0x68, 0x73, 0x65, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x69, 0x0a, 0x0e, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x67,
	0x6d, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6d, 0x6d, 0x73, 0x70, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x4d, 0x67, 0x6d, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x03,
	0x63, 0x66, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x6f, 0x70, 0x6f,
	0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x66, 0x67, 0x52, 0x03, 0x63, 0x66,
	0x67, 0x2a, 0xb8, 0x03, 0x0a, 0x0f, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x67, 0x6d, 0x74, 0x4d, 0x73,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e,
	0x45, 0x44, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x4d, 0x47, 0x4d, 0x54, 0x5f, 0x4d, 0x53, 0x47, 0x10,
	0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x10, 0x0b,
	0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10,
	0x15, 0x12, 0x1d, 0x0a, 0x19, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x4d,
	0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x65,
	0x12, 0x1e, 0x0a, 0x1a, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x4d, 0x45,
	0x54, 0x52, 0x49, 0x43, 0x53, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x66,
	0x12, 0x25, 0x0a, 0x20, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x4f, 0x4e,
	0x4e, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x10, 0xc9, 0x01, 0x12, 0x26, 0x0a, 0x21, 0x4e, 0x4f, 0x44, 0x45, 0x5f,
	0x4e, 0x45, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0xca, 0x01, 0x12,
	0x23, 0x0a, 0x1e, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x4e,
	0x54, 0x52, 0x41, 0x43, 0x4b, 0x5f, 0x4c, 0x4f, 0x47, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x10, 0xd3, 0x01, 0x12, 0x24, 0x0a, 0x1f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x45, 0x54,
	0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x5f, 0x4c, 0x4f, 0x47, 0x5f, 0x52,
	0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0xd4, 0x01, 0x12, 0x25, 0x0a, 0x20, 0x4e, 0x4f,
	0x44, 0x45, 0x5f, 0x4e, 0x45, 0x54, 0x5f, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x4d,
	0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0xdd,
	0x01, 0x12, 0x26, 0x0a, 0x21, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x45, 0x54, 0x5f, 0x54, 0x52,
	0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x5f, 0x52, 0x45,
	0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0xde, 0x01, 0x12, 0x1f, 0x0a, 0x1a, 0x4e, 0x4f, 0x44,
	0x45, 0x5f, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0xad, 0x02, 0x12, 0x20, 0x0a, 0x1b, 0x4e, 0x4f,
	0x44, 0x45, 0x5f, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59,
	0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0xae, 0x02, 0x2a, 0x7c, 0x0a, 0x18,
	0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x67, 0x6d, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x46, 0x47, 0x5f,
	0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x43,
	0x46, 0x47, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x10, 0x0b, 0x12, 0x12, 0x0a,
	0x0e, 0x43, 0x46, 0x47, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x49, 0x4e, 0x47, 0x10,
	0x15, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x46, 0x47, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x4d,
	0x45, 0x4e, 0x54, 0x10, 0x1f, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x46, 0x47, 0x5f, 0x4d, 0x41, 0x49,
	0x4e, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x29, 0x42, 0x26, 0x5a, 0x24, 0x6d, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6d, 0x6d,
	0x73, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_rawDescData = file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_rawDesc
)

func file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_rawDescData)
	})
	return file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_rawDescData
}

var file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_goTypes = []interface{}{
	(NodeMgmtMsgType)(0),                      // 0: mmsp.NodeMgmtMsgType
	(NodeMgmtConfigActionType)(0),             // 1: mmsp.NodeMgmtConfigActionType
	(*NodeMgmtPDU)(nil),                       // 2: mmsp.NodeMgmtPDU
	(*NodeMgmtConfig)(nil),                    // 3: mmsp.NodeMgmtConfig
	(*topology.NodeReq)(nil),                  // 4: topology.NodeReq
	(*metricsdb.HostMetricsRequest)(nil),      // 5: metricsdb.HostMetricsRequest
	(*metricsdb.HostMetricsResponse)(nil),     // 6: metricsdb.HostMetricsResponse
	(*netdb.ConntrackTableRequest)(nil),       // 7: netdb.ConntrackTableRequest
	(*netdb.ConntrackTableResponse)(nil),      // 8: netdb.ConntrackTableResponse
	(*netdb.ConntrackLogRequest)(nil),         // 9: netdb.ConntrackLogRequest
	(*netdb.ConntrackLogResponse)(nil),        // 10: netdb.ConntrackLogResponse
	(*netdb.TrafficMetricsRequest)(nil),       // 11: netdb.TrafficMetricsRequest
	(*netdb.TrafficMetricsResponse)(nil),      // 12: netdb.TrafficMetricsResponse
	(*hsecdb.HostSecurityReportRequest)(nil),  // 13: hsecdb.HostSecurityReportRequest
	(*hsecdb.HostSecurityReportResponse)(nil), // 14: hsecdb.HostSecurityReportResponse
	(*topology.NodeCfg)(nil),                  // 15: topology.NodeCfg
}
var file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_depIdxs = []int32{
	0,  // 0: mmsp.NodeMgmtPDU.type:type_name -> mmsp.NodeMgmtMsgType
	4,  // 1: mmsp.NodeMgmtPDU.nodeReq:type_name -> topology.NodeReq
	3,  // 2: mmsp.NodeMgmtPDU.nodeConfig:type_name -> mmsp.NodeMgmtConfig
	5,  // 3: mmsp.NodeMgmtPDU.hostMetricsRequest:type_name -> metricsdb.HostMetricsRequest
	6,  // 4: mmsp.NodeMgmtPDU.hostMetricsResponse:type_name -> metricsdb.HostMetricsResponse
	7,  // 5: mmsp.NodeMgmtPDU.netCtStateRequest:type_name -> netdb.ConntrackTableRequest
	8,  // 6: mmsp.NodeMgmtPDU.netCtStateResponse:type_name -> netdb.ConntrackTableResponse
	9,  // 7: mmsp.NodeMgmtPDU.netCtLogRequest:type_name -> netdb.ConntrackLogRequest
	10, // 8: mmsp.NodeMgmtPDU.netCtLogResponse:type_name -> netdb.ConntrackLogResponse
	11, // 9: mmsp.NodeMgmtPDU.netTrafficMetricsRequest:type_name -> netdb.TrafficMetricsRequest
	12, // 10: mmsp.NodeMgmtPDU.netTrafficMetricsResponse:type_name -> netdb.TrafficMetricsResponse
	13, // 11: mmsp.NodeMgmtPDU.hsecReportRequest:type_name -> hsecdb.HostSecurityReportRequest
	14, // 12: mmsp.NodeMgmtPDU.hsecReportResponse:type_name -> hsecdb.HostSecurityReportResponse
	1,  // 13: mmsp.NodeMgmtConfig.type:type_name -> mmsp.NodeMgmtConfigActionType
	15, // 14: mmsp.NodeMgmtConfig.cfg:type_name -> topology.NodeCfg
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_init() }
func file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_init() {
	if File_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeMgmtPDU); i {
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
		file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeMgmtConfig); i {
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
			RawDescriptor: file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_depIdxs,
		EnumInfos:         file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_enumTypes,
		MessageInfos:      file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto = out.File
	file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_rawDesc = nil
	file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_goTypes = nil
	file_mmesh_protobuf_network_v1_mmsp_nodemgmt_proto_depIdxs = nil
}
