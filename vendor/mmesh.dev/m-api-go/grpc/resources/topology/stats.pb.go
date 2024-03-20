// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0--rc3
// source: mmesh/protobuf/resources/v1/topology/stats.proto

package topology

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

type StatsTopology struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalNetworks  int32        `protobuf:"varint,11,opt,name=totalNetworks,proto3" json:"totalNetworks,omitempty"`
	TotalSubnets   int32        `protobuf:"varint,21,opt,name=totalSubnets,proto3" json:"totalSubnets,omitempty"`
	TotalNodes     int32        `protobuf:"varint,31,opt,name=totalNodes,proto3" json:"totalNodes,omitempty"`
	TotalEndpoints int32        `protobuf:"varint,37,opt,name=totalEndpoints,proto3" json:"totalEndpoints,omitempty"`
	TotalVSs       int32        `protobuf:"varint,51,opt,name=totalVSs,proto3" json:"totalVSs,omitempty"`
	Nodes          *NodeMetrics `protobuf:"bytes,101,opt,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *StatsTopology) Reset() {
	*x = StatsTopology{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsTopology) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsTopology) ProtoMessage() {}

func (x *StatsTopology) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsTopology.ProtoReflect.Descriptor instead.
func (*StatsTopology) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_topology_stats_proto_rawDescGZIP(), []int{0}
}

func (x *StatsTopology) GetTotalNetworks() int32 {
	if x != nil {
		return x.TotalNetworks
	}
	return 0
}

func (x *StatsTopology) GetTotalSubnets() int32 {
	if x != nil {
		return x.TotalSubnets
	}
	return 0
}

func (x *StatsTopology) GetTotalNodes() int32 {
	if x != nil {
		return x.TotalNodes
	}
	return 0
}

func (x *StatsTopology) GetTotalEndpoints() int32 {
	if x != nil {
		return x.TotalEndpoints
	}
	return 0
}

func (x *StatsTopology) GetTotalVSs() int32 {
	if x != nil {
		return x.TotalVSs
	}
	return 0
}

func (x *StatsTopology) GetNodes() *NodeMetrics {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type NodeMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ByStatus    *ByStatus    `protobuf:"bytes,11,opt,name=byStatus,proto3" json:"byStatus,omitempty"`
	ByNodeClass *ByNodeClass `protobuf:"bytes,21,opt,name=byNodeClass,proto3" json:"byNodeClass,omitempty"`
	ByNodeState *ByNodeState `protobuf:"bytes,31,opt,name=byNodeState,proto3" json:"byNodeState,omitempty"`
	ByOS        *ByOS        `protobuf:"bytes,41,opt,name=byOS,proto3" json:"byOS,omitempty"`
}

func (x *NodeMetrics) Reset() {
	*x = NodeMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMetrics) ProtoMessage() {}

func (x *NodeMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMetrics.ProtoReflect.Descriptor instead.
func (*NodeMetrics) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_topology_stats_proto_rawDescGZIP(), []int{1}
}

func (x *NodeMetrics) GetByStatus() *ByStatus {
	if x != nil {
		return x.ByStatus
	}
	return nil
}

func (x *NodeMetrics) GetByNodeClass() *ByNodeClass {
	if x != nil {
		return x.ByNodeClass
	}
	return nil
}

func (x *NodeMetrics) GetByNodeState() *ByNodeState {
	if x != nil {
		return x.ByNodeState
	}
	return nil
}

func (x *NodeMetrics) GetByOS() *ByOS {
	if x != nil {
		return x.ByOS
	}
	return nil
}

type ByStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Online  int32 `protobuf:"varint,11,opt,name=online,proto3" json:"online,omitempty"`
	Offline int32 `protobuf:"varint,21,opt,name=offline,proto3" json:"offline,omitempty"`
}

func (x *ByStatus) Reset() {
	*x = ByStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByStatus) ProtoMessage() {}

func (x *ByStatus) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByStatus.ProtoReflect.Descriptor instead.
func (*ByStatus) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_topology_stats_proto_rawDescGZIP(), []int{2}
}

func (x *ByStatus) GetOnline() int32 {
	if x != nil {
		return x.Online
	}
	return 0
}

func (x *ByStatus) GetOffline() int32 {
	if x != nil {
		return x.Offline
	}
	return 0
}

type ByNodeClass struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Compute    int32 `protobuf:"varint,11,opt,name=compute,proto3" json:"compute,omitempty"`
	Serverless int32 `protobuf:"varint,21,opt,name=serverless,proto3" json:"serverless,omitempty"`
}

func (x *ByNodeClass) Reset() {
	*x = ByNodeClass{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByNodeClass) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByNodeClass) ProtoMessage() {}

func (x *ByNodeClass) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByNodeClass.ProtoReflect.Descriptor instead.
func (*ByNodeClass) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_topology_stats_proto_rawDescGZIP(), []int{3}
}

func (x *ByNodeClass) GetCompute() int32 {
	if x != nil {
		return x.Compute
	}
	return 0
}

func (x *ByNodeClass) GetServerless() int32 {
	if x != nil {
		return x.Serverless
	}
	return 0
}

type ByNodeState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connected    int32 `protobuf:"varint,11,opt,name=connected,proto3" json:"connected,omitempty"`
	Stubby       int32 `protobuf:"varint,21,opt,name=stubby,proto3" json:"stubby,omitempty"`
	PendingSetup int32 `protobuf:"varint,31,opt,name=pendingSetup,proto3" json:"pendingSetup,omitempty"`
}

func (x *ByNodeState) Reset() {
	*x = ByNodeState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByNodeState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByNodeState) ProtoMessage() {}

func (x *ByNodeState) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByNodeState.ProtoReflect.Descriptor instead.
func (*ByNodeState) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_topology_stats_proto_rawDescGZIP(), []int{4}
}

func (x *ByNodeState) GetConnected() int32 {
	if x != nil {
		return x.Connected
	}
	return 0
}

func (x *ByNodeState) GetStubby() int32 {
	if x != nil {
		return x.Stubby
	}
	return 0
}

func (x *ByNodeState) GetPendingSetup() int32 {
	if x != nil {
		return x.PendingSetup
	}
	return 0
}

type ByOS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unknown int32 `protobuf:"varint,1,opt,name=unknown,proto3" json:"unknown,omitempty"`
	Linux   int32 `protobuf:"varint,11,opt,name=linux,proto3" json:"linux,omitempty"`
	MacOS   int32 `protobuf:"varint,21,opt,name=macOS,proto3" json:"macOS,omitempty"`
	Windows int32 `protobuf:"varint,31,opt,name=windows,proto3" json:"windows,omitempty"`
}

func (x *ByOS) Reset() {
	*x = ByOS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByOS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByOS) ProtoMessage() {}

func (x *ByOS) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByOS.ProtoReflect.Descriptor instead.
func (*ByOS) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_topology_stats_proto_rawDescGZIP(), []int{5}
}

func (x *ByOS) GetUnknown() int32 {
	if x != nil {
		return x.Unknown
	}
	return 0
}

func (x *ByOS) GetLinux() int32 {
	if x != nil {
		return x.Linux
	}
	return 0
}

func (x *ByOS) GetMacOS() int32 {
	if x != nil {
		return x.MacOS
	}
	return 0
}

func (x *ByOS) GetWindows() int32 {
	if x != nil {
		return x.Windows
	}
	return 0
}

type StatsAlerts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalAlerts    int32 `protobuf:"varint,1,opt,name=totalAlerts,proto3" json:"totalAlerts,omitempty"`
	TotalTriggered int32 `protobuf:"varint,11,opt,name=totalTriggered,proto3" json:"totalTriggered,omitempty"`
	TotalResolved  int32 `protobuf:"varint,12,opt,name=totalResolved,proto3" json:"totalResolved,omitempty"`
	TotalInfo      int32 `protobuf:"varint,21,opt,name=totalInfo,proto3" json:"totalInfo,omitempty"`
	TotalWarning   int32 `protobuf:"varint,22,opt,name=totalWarning,proto3" json:"totalWarning,omitempty"`
	TotalError     int32 `protobuf:"varint,23,opt,name=totalError,proto3" json:"totalError,omitempty"`
	TotalCritical  int32 `protobuf:"varint,24,opt,name=totalCritical,proto3" json:"totalCritical,omitempty"`
}

func (x *StatsAlerts) Reset() {
	*x = StatsAlerts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsAlerts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsAlerts) ProtoMessage() {}

func (x *StatsAlerts) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsAlerts.ProtoReflect.Descriptor instead.
func (*StatsAlerts) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_topology_stats_proto_rawDescGZIP(), []int{6}
}

func (x *StatsAlerts) GetTotalAlerts() int32 {
	if x != nil {
		return x.TotalAlerts
	}
	return 0
}

func (x *StatsAlerts) GetTotalTriggered() int32 {
	if x != nil {
		return x.TotalTriggered
	}
	return 0
}

func (x *StatsAlerts) GetTotalResolved() int32 {
	if x != nil {
		return x.TotalResolved
	}
	return 0
}

func (x *StatsAlerts) GetTotalInfo() int32 {
	if x != nil {
		return x.TotalInfo
	}
	return 0
}

func (x *StatsAlerts) GetTotalWarning() int32 {
	if x != nil {
		return x.TotalWarning
	}
	return 0
}

func (x *StatsAlerts) GetTotalError() int32 {
	if x != nil {
		return x.TotalError
	}
	return 0
}

func (x *StatsAlerts) GetTotalCritical() int32 {
	if x != nil {
		return x.TotalCritical
	}
	return 0
}

type StatsWorkflows struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalProjects  int32            `protobuf:"varint,11,opt,name=totalProjects,proto3" json:"totalProjects,omitempty"`
	TotalWorkflows int32            `protobuf:"varint,21,opt,name=totalWorkflows,proto3" json:"totalWorkflows,omitempty"`
	TotalTasks     int32            `protobuf:"varint,31,opt,name=totalTasks,proto3" json:"totalTasks,omitempty"`
	LastResult     *WorkflowMetrics `protobuf:"bytes,101,opt,name=lastResult,proto3" json:"lastResult,omitempty"`
}

func (x *StatsWorkflows) Reset() {
	*x = StatsWorkflows{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsWorkflows) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsWorkflows) ProtoMessage() {}

func (x *StatsWorkflows) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsWorkflows.ProtoReflect.Descriptor instead.
func (*StatsWorkflows) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_topology_stats_proto_rawDescGZIP(), []int{7}
}

func (x *StatsWorkflows) GetTotalProjects() int32 {
	if x != nil {
		return x.TotalProjects
	}
	return 0
}

func (x *StatsWorkflows) GetTotalWorkflows() int32 {
	if x != nil {
		return x.TotalWorkflows
	}
	return 0
}

func (x *StatsWorkflows) GetTotalTasks() int32 {
	if x != nil {
		return x.TotalTasks
	}
	return 0
}

func (x *StatsWorkflows) GetLastResult() *WorkflowMetrics {
	if x != nil {
		return x.LastResult
	}
	return nil
}

type WorkflowMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnknownResultWorkflows int32 `protobuf:"varint,1,opt,name=unknownResultWorkflows,proto3" json:"unknownResultWorkflows,omitempty"`
	SuccessfulWorkflows    int32 `protobuf:"varint,11,opt,name=successfulWorkflows,proto3" json:"successfulWorkflows,omitempty"`
	FailedWorkflows        int32 `protobuf:"varint,21,opt,name=failedWorkflows,proto3" json:"failedWorkflows,omitempty"`
}

func (x *WorkflowMetrics) Reset() {
	*x = WorkflowMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowMetrics) ProtoMessage() {}

func (x *WorkflowMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowMetrics.ProtoReflect.Descriptor instead.
func (*WorkflowMetrics) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_topology_stats_proto_rawDescGZIP(), []int{8}
}

func (x *WorkflowMetrics) GetUnknownResultWorkflows() int32 {
	if x != nil {
		return x.UnknownResultWorkflows
	}
	return 0
}

func (x *WorkflowMetrics) GetSuccessfulWorkflows() int32 {
	if x != nil {
		return x.SuccessfulWorkflows
	}
	return 0
}

func (x *WorkflowMetrics) GetFailedWorkflows() int32 {
	if x != nil {
		return x.FailedWorkflows
	}
	return 0
}

var File_mmesh_protobuf_resources_v1_topology_stats_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_topology_stats_proto_rawDesc = []byte{
	0x0a, 0x30, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f,
	0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x22, 0xea, 0x01, 0x0a,
	0x0d, 0x53, 0x74, 0x61, 0x74, 0x73, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12, 0x24,
	0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x25, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x53, 0x73, 0x18, 0x33, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x53, 0x73, 0x12, 0x2b, 0x0a, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x6f,
	0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0xd3, 0x01, 0x0a, 0x0b, 0x4e, 0x6f,
	0x64, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x62, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x6f,
	0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x42, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x08, 0x62, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x0b, 0x62, 0x79, 0x4e,
	0x6f, 0x64, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x42, 0x79, 0x4e, 0x6f, 0x64, 0x65,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x0b, 0x62, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x12, 0x37, 0x0a, 0x0b, 0x62, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f,
	0x67, 0x79, 0x2e, 0x42, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0b,
	0x62, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x62,
	0x79, 0x4f, 0x53, 0x18, 0x29, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x6f, 0x70, 0x6f,
	0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x42, 0x79, 0x4f, 0x53, 0x52, 0x04, 0x62, 0x79, 0x4f, 0x53, 0x22,
	0x3c, 0x0a, 0x08, 0x42, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x47, 0x0a,
	0x0b, 0x42, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x6c, 0x65, 0x73, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x22, 0x67, 0x0a, 0x0b, 0x42, 0x79, 0x4e, 0x6f, 0x64, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x75, 0x62, 0x62, 0x79, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x75, 0x62, 0x62, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x70,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75, 0x70, 0x18, 0x1f, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75, 0x70, 0x22,
	0x66, 0x0a, 0x04, 0x42, 0x79, 0x4f, 0x53, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x63, 0x4f, 0x53,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x61, 0x63, 0x4f, 0x53, 0x12, 0x18, 0x0a,
	0x07, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x22, 0x85, 0x02, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x65,
	0x64, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x57, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x16, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x17, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x18, 0x18, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x22,
	0xb9, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x1f,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x12, 0x39, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x65,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52,
	0x0a, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xa5, 0x01, 0x0a, 0x0f,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12,
	0x36, 0x0a, 0x16, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x16, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x66, 0x75, 0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x66, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x73, 0x42, 0x2c, 0x5a, 0x2a, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76,
	0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_topology_stats_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_topology_stats_proto_rawDescData = file_mmesh_protobuf_resources_v1_topology_stats_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_topology_stats_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_topology_stats_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_topology_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_topology_stats_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_topology_stats_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_mmesh_protobuf_resources_v1_topology_stats_proto_goTypes = []interface{}{
	(*StatsTopology)(nil),   // 0: topology.StatsTopology
	(*NodeMetrics)(nil),     // 1: topology.NodeMetrics
	(*ByStatus)(nil),        // 2: topology.ByStatus
	(*ByNodeClass)(nil),     // 3: topology.ByNodeClass
	(*ByNodeState)(nil),     // 4: topology.ByNodeState
	(*ByOS)(nil),            // 5: topology.ByOS
	(*StatsAlerts)(nil),     // 6: topology.StatsAlerts
	(*StatsWorkflows)(nil),  // 7: topology.StatsWorkflows
	(*WorkflowMetrics)(nil), // 8: topology.WorkflowMetrics
}
var file_mmesh_protobuf_resources_v1_topology_stats_proto_depIdxs = []int32{
	1, // 0: topology.StatsTopology.nodes:type_name -> topology.NodeMetrics
	2, // 1: topology.NodeMetrics.byStatus:type_name -> topology.ByStatus
	3, // 2: topology.NodeMetrics.byNodeClass:type_name -> topology.ByNodeClass
	4, // 3: topology.NodeMetrics.byNodeState:type_name -> topology.ByNodeState
	5, // 4: topology.NodeMetrics.byOS:type_name -> topology.ByOS
	8, // 5: topology.StatsWorkflows.lastResult:type_name -> topology.WorkflowMetrics
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_topology_stats_proto_init() }
func file_mmesh_protobuf_resources_v1_topology_stats_proto_init() {
	if File_mmesh_protobuf_resources_v1_topology_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsTopology); i {
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
		file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeMetrics); i {
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
		file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByStatus); i {
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
		file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByNodeClass); i {
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
		file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByNodeState); i {
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
		file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByOS); i {
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
		file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsAlerts); i {
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
		file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsWorkflows); i {
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
		file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowMetrics); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_topology_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_topology_stats_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_topology_stats_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_topology_stats_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_topology_stats_proto = out.File
	file_mmesh_protobuf_resources_v1_topology_stats_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_topology_stats_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_topology_stats_proto_depIdxs = nil
}
