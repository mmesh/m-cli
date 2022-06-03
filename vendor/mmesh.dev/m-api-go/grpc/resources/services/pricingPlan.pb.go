// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: mmesh/protobuf/resources/v1/services/pricingPlan.proto

package services

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	resource "mmesh.dev/m-api-go/grpc/resources/resource"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PlanType int32

const (
	PlanType_NETWORK PlanType = 0
)

// Enum value maps for PlanType.
var (
	PlanType_name = map[int32]string{
		0: "NETWORK",
	}
	PlanType_value = map[string]int32{
		"NETWORK": 0,
	}
)

func (x PlanType) Enum() *PlanType {
	p := new(PlanType)
	*p = x
	return p
}

func (x PlanType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlanType) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_enumTypes[0].Descriptor()
}

func (PlanType) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_enumTypes[0]
}

func (x PlanType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlanType.Descriptor instead.
func (PlanType) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_rawDescGZIP(), []int{0}
}

type PricingPlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceID                string            `protobuf:"bytes,1,opt,name=serviceID,proto3" json:"serviceID,omitempty"`
	PricingPlanID            string            `protobuf:"bytes,2,opt,name=pricingPlanID,proto3" json:"pricingPlanID,omitempty"`
	Description              string            `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	UnitLabel                string            `protobuf:"bytes,6,opt,name=unitLabel,proto3" json:"unitLabel,omitempty"`
	StripeProductID          string            `protobuf:"bytes,11,opt,name=stripeProductID,proto3" json:"stripeProductID,omitempty"`
	Prices                   map[string]*Price `protobuf:"bytes,41,rep,name=prices,proto3" json:"prices,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map[stripePriceRecurringInterval]*Price
	Resources                *Resources        `protobuf:"bytes,51,opt,name=resources,proto3" json:"resources,omitempty"`
	Type                     PlanType          `protobuf:"varint,61,opt,name=type,proto3,enum=services.PlanType" json:"type,omitempty"`
	NetworkLimits            *NetworkLimits    `protobuf:"bytes,71,opt,name=networkLimits,proto3" json:"networkLimits,omitempty"`
	Included3RdPartyServices bool              `protobuf:"varint,81,opt,name=included3rdPartyServices,proto3" json:"included3rdPartyServices,omitempty"`
	CreationDate             int64             `protobuf:"varint,101,opt,name=creationDate,proto3" json:"creationDate,omitempty"`
	LastModified             int64             `protobuf:"varint,102,opt,name=lastModified,proto3" json:"lastModified,omitempty"`
}

func (x *PricingPlan) Reset() {
	*x = PricingPlan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PricingPlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PricingPlan) ProtoMessage() {}

func (x *PricingPlan) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PricingPlan.ProtoReflect.Descriptor instead.
func (*PricingPlan) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_rawDescGZIP(), []int{0}
}

func (x *PricingPlan) GetServiceID() string {
	if x != nil {
		return x.ServiceID
	}
	return ""
}

func (x *PricingPlan) GetPricingPlanID() string {
	if x != nil {
		return x.PricingPlanID
	}
	return ""
}

func (x *PricingPlan) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PricingPlan) GetUnitLabel() string {
	if x != nil {
		return x.UnitLabel
	}
	return ""
}

func (x *PricingPlan) GetStripeProductID() string {
	if x != nil {
		return x.StripeProductID
	}
	return ""
}

func (x *PricingPlan) GetPrices() map[string]*Price {
	if x != nil {
		return x.Prices
	}
	return nil
}

func (x *PricingPlan) GetResources() *Resources {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *PricingPlan) GetType() PlanType {
	if x != nil {
		return x.Type
	}
	return PlanType_NETWORK
}

func (x *PricingPlan) GetNetworkLimits() *NetworkLimits {
	if x != nil {
		return x.NetworkLimits
	}
	return nil
}

func (x *PricingPlan) GetIncluded3RdPartyServices() bool {
	if x != nil {
		return x.Included3RdPartyServices
	}
	return false
}

func (x *PricingPlan) GetCreationDate() int64 {
	if x != nil {
		return x.CreationDate
	}
	return 0
}

func (x *PricingPlan) GetLastModified() int64 {
	if x != nil {
		return x.LastModified
	}
	return 0
}

type PricingPlans struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta         *resource.ListResponse `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	PricingPlans []*PricingPlan         `protobuf:"bytes,2,rep,name=pricingPlans,proto3" json:"pricingPlans,omitempty"`
}

func (x *PricingPlans) Reset() {
	*x = PricingPlans{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PricingPlans) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PricingPlans) ProtoMessage() {}

func (x *PricingPlans) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PricingPlans.ProtoReflect.Descriptor instead.
func (*PricingPlans) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_rawDescGZIP(), []int{1}
}

func (x *PricingPlans) GetMeta() *resource.ListResponse {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *PricingPlans) GetPricingPlans() []*PricingPlan {
	if x != nil {
		return x.PricingPlans
	}
	return nil
}

type ListPricingPlansRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta      *resource.ListRequest `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	ServiceID string                `protobuf:"bytes,2,opt,name=serviceID,proto3" json:"serviceID,omitempty"`
}

func (x *ListPricingPlansRequest) Reset() {
	*x = ListPricingPlansRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPricingPlansRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPricingPlansRequest) ProtoMessage() {}

func (x *ListPricingPlansRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPricingPlansRequest.ProtoReflect.Descriptor instead.
func (*ListPricingPlansRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_rawDescGZIP(), []int{2}
}

func (x *ListPricingPlansRequest) GetMeta() *resource.ListRequest {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ListPricingPlansRequest) GetServiceID() string {
	if x != nil {
		return x.ServiceID
	}
	return ""
}

type Resources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Discount int64 `protobuf:"varint,1,opt,name=discount,proto3" json:"discount,omitempty"`
	User     *Unit `protobuf:"bytes,11,opt,name=user,proto3" json:"user,omitempty"`
	Tenant   *Unit `protobuf:"bytes,21,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Network  *Unit `protobuf:"bytes,22,opt,name=network,proto3" json:"network,omitempty"`
	Subnet   *Unit `protobuf:"bytes,23,opt,name=subnet,proto3" json:"subnet,omitempty"`
	Node     *Unit `protobuf:"bytes,24,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *Resources) Reset() {
	*x = Resources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resources) ProtoMessage() {}

func (x *Resources) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resources.ProtoReflect.Descriptor instead.
func (*Resources) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_rawDescGZIP(), []int{3}
}

func (x *Resources) GetDiscount() int64 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *Resources) GetUser() *Unit {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Resources) GetTenant() *Unit {
	if x != nil {
		return x.Tenant
	}
	return nil
}

func (x *Resources) GetNetwork() *Unit {
	if x != nil {
		return x.Network
	}
	return nil
}

func (x *Resources) GetSubnet() *Unit {
	if x != nil {
		return x.Subnet
	}
	return nil
}

func (x *Resources) GetNode() *Unit {
	if x != nil {
		return x.Node
	}
	return nil
}

type Unit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncludedUnits int64 `protobuf:"varint,1,opt,name=includedUnits,proto3" json:"includedUnits,omitempty"` // int64 listPrice = 11;
}

func (x *Unit) Reset() {
	*x = Unit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unit) ProtoMessage() {}

func (x *Unit) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unit.ProtoReflect.Descriptor instead.
func (*Unit) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_rawDescGZIP(), []int{4}
}

func (x *Unit) GetIncludedUnits() int64 {
	if x != nil {
		return x.IncludedUnits
	}
	return 0
}

type NetworkLimits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxGiB      *ItemLimit `protobuf:"bytes,1,opt,name=maxGiB,proto3" json:"maxGiB,omitempty"`
	MaxTenants  *ItemLimit `protobuf:"bytes,11,opt,name=maxTenants,proto3" json:"maxTenants,omitempty"`
	MaxNetworks *ItemLimit `protobuf:"bytes,12,opt,name=maxNetworks,proto3" json:"maxNetworks,omitempty"`
	MaxSubnets  *ItemLimit `protobuf:"bytes,13,opt,name=maxSubnets,proto3" json:"maxSubnets,omitempty"`
	MaxNodes    *ItemLimit `protobuf:"bytes,14,opt,name=maxNodes,proto3" json:"maxNodes,omitempty"`
}

func (x *NetworkLimits) Reset() {
	*x = NetworkLimits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkLimits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkLimits) ProtoMessage() {}

func (x *NetworkLimits) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkLimits.ProtoReflect.Descriptor instead.
func (*NetworkLimits) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_rawDescGZIP(), []int{5}
}

func (x *NetworkLimits) GetMaxGiB() *ItemLimit {
	if x != nil {
		return x.MaxGiB
	}
	return nil
}

func (x *NetworkLimits) GetMaxTenants() *ItemLimit {
	if x != nil {
		return x.MaxTenants
	}
	return nil
}

func (x *NetworkLimits) GetMaxNetworks() *ItemLimit {
	if x != nil {
		return x.MaxNetworks
	}
	return nil
}

func (x *NetworkLimits) GetMaxSubnets() *ItemLimit {
	if x != nil {
		return x.MaxSubnets
	}
	return nil
}

func (x *NetworkLimits) GetMaxNodes() *ItemLimit {
	if x != nil {
		return x.MaxNodes
	}
	return nil
}

type ItemLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit              int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	EnforceRestriction bool  `protobuf:"varint,2,opt,name=enforceRestriction,proto3" json:"enforceRestriction,omitempty"`
}

func (x *ItemLimit) Reset() {
	*x = ItemLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemLimit) ProtoMessage() {}

func (x *ItemLimit) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemLimit.ProtoReflect.Descriptor instead.
func (*ItemLimit) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_rawDescGZIP(), []int{6}
}

func (x *ItemLimit) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ItemLimit) GetEnforceRestriction() bool {
	if x != nil {
		return x.EnforceRestriction
	}
	return false
}

var File_mmesh_protobuf_resources_v1_services_pricingPlan_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_rawDesc = []byte{
	0x0a, 0x36, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x50, 0x6c,
	0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x04, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e,
	0x67, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x50, 0x6c,
	0x61, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x63,
	0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x75,
	0x6e, 0x69, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x6e, 0x69, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x74, 0x72,
	0x69, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x44, 0x12, 0x39, 0x0a, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x29, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x6e, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x31,
	0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x33, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x47, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x52, 0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x18, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x64, 0x33, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x18, 0x51, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x64, 0x33, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x1a, 0x4a, 0x0a, 0x0b,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x75, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x63,
	0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x50,
	0x6c, 0x61, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61,
	0x6e, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x22,
	0x62, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x50, 0x6c,
	0x61, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x44, 0x22, 0xe9, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x26, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69,
	0x74, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x18, 0x17, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55,
	0x6e, 0x69, 0x74, 0x52, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x6e,
	0x6f, 0x64, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22,
	0x2c, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x22, 0x8e, 0x02,
	0x0a, 0x0d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12,
	0x2b, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x47, 0x69, 0x42, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x47, 0x69, 0x42, 0x12, 0x33, 0x0a, 0x0a,
	0x6d, 0x61, 0x78, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x73, 0x12, 0x35, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x0b, 0x6d, 0x61, 0x78,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x33, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x53,
	0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x12, 0x2f, 0x0a,
	0x08, 0x6d, 0x61, 0x78, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x51,
	0x0a, 0x09, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x2e, 0x0a, 0x12, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x65,
	0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2a, 0x17, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x6d, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_rawDescData = file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_goTypes = []interface{}{
	(PlanType)(0),                   // 0: services.PlanType
	(*PricingPlan)(nil),             // 1: services.PricingPlan
	(*PricingPlans)(nil),            // 2: services.PricingPlans
	(*ListPricingPlansRequest)(nil), // 3: services.ListPricingPlansRequest
	(*Resources)(nil),               // 4: services.Resources
	(*Unit)(nil),                    // 5: services.Unit
	(*NetworkLimits)(nil),           // 6: services.NetworkLimits
	(*ItemLimit)(nil),               // 7: services.ItemLimit
	nil,                             // 8: services.PricingPlan.PricesEntry
	(*resource.ListResponse)(nil),   // 9: resource.ListResponse
	(*resource.ListRequest)(nil),    // 10: resource.ListRequest
	(*Price)(nil),                   // 11: services.Price
}
var file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_depIdxs = []int32{
	8,  // 0: services.PricingPlan.prices:type_name -> services.PricingPlan.PricesEntry
	4,  // 1: services.PricingPlan.resources:type_name -> services.Resources
	0,  // 2: services.PricingPlan.type:type_name -> services.PlanType
	6,  // 3: services.PricingPlan.networkLimits:type_name -> services.NetworkLimits
	9,  // 4: services.PricingPlans.meta:type_name -> resource.ListResponse
	1,  // 5: services.PricingPlans.pricingPlans:type_name -> services.PricingPlan
	10, // 6: services.ListPricingPlansRequest.meta:type_name -> resource.ListRequest
	5,  // 7: services.Resources.user:type_name -> services.Unit
	5,  // 8: services.Resources.tenant:type_name -> services.Unit
	5,  // 9: services.Resources.network:type_name -> services.Unit
	5,  // 10: services.Resources.subnet:type_name -> services.Unit
	5,  // 11: services.Resources.node:type_name -> services.Unit
	7,  // 12: services.NetworkLimits.maxGiB:type_name -> services.ItemLimit
	7,  // 13: services.NetworkLimits.maxTenants:type_name -> services.ItemLimit
	7,  // 14: services.NetworkLimits.maxNetworks:type_name -> services.ItemLimit
	7,  // 15: services.NetworkLimits.maxSubnets:type_name -> services.ItemLimit
	7,  // 16: services.NetworkLimits.maxNodes:type_name -> services.ItemLimit
	11, // 17: services.PricingPlan.PricesEntry.value:type_name -> services.Price
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_init() }
func file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_init() {
	if File_mmesh_protobuf_resources_v1_services_pricingPlan_proto != nil {
		return
	}
	file_mmesh_protobuf_resources_v1_services_price_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PricingPlan); i {
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
		file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PricingPlans); i {
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
		file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPricingPlansRequest); i {
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
		file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resources); i {
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
		file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unit); i {
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
		file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkLimits); i {
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
		file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemLimit); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_depIdxs,
		EnumInfos:         file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_enumTypes,
		MessageInfos:      file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_services_pricingPlan_proto = out.File
	file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_services_pricingPlan_proto_depIdxs = nil
}
