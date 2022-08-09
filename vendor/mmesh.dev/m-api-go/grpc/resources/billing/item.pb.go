// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: mmesh/protobuf/resources/v1/billing/item.proto

package billing

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

type ItemClass int32

const (
	ItemClass_RESOURCE_UNIT ItemClass = 0
	ItemClass_IAAS          ItemClass = 1
)

// Enum value maps for ItemClass.
var (
	ItemClass_name = map[int32]string{
		0: "RESOURCE_UNIT",
		1: "IAAS",
	}
	ItemClass_value = map[string]int32{
		"RESOURCE_UNIT": 0,
		"IAAS":          1,
	}
)

func (x ItemClass) Enum() *ItemClass {
	p := new(ItemClass)
	*p = x
	return p
}

func (x ItemClass) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ItemClass) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_billing_item_proto_enumTypes[0].Descriptor()
}

func (ItemClass) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_billing_item_proto_enumTypes[0]
}

func (x ItemClass) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ItemClass.Descriptor instead.
func (ItemClass) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_billing_item_proto_rawDescGZIP(), []int{0}
}

type InvoiceItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID            string    `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	StripeCustomerID     string    `protobuf:"bytes,2,opt,name=stripeCustomerID,proto3" json:"stripeCustomerID,omitempty"`
	Period               string    `protobuf:"bytes,6,opt,name=period,proto3" json:"period,omitempty"` // yyyyMMhhmmss-yyyyMMhhmmss
	ServiceID            string    `protobuf:"bytes,11,opt,name=serviceID,proto3" json:"serviceID,omitempty"`
	ProviderID           string    `protobuf:"bytes,12,opt,name=providerID,proto3" json:"providerID,omitempty"`
	StripeSubscriptionID string    `protobuf:"bytes,14,opt,name=stripeSubscriptionID,proto3" json:"stripeSubscriptionID,omitempty"`
	StripeInvoiceID      string    `protobuf:"bytes,15,opt,name=stripeInvoiceID,proto3" json:"stripeInvoiceID,omitempty"`
	StripeProductID      string    `protobuf:"bytes,17,opt,name=stripeProductID,proto3" json:"stripeProductID,omitempty"`
	ItemID               string    `protobuf:"bytes,21,opt,name=itemID,proto3" json:"itemID,omitempty"`                     // resource UUID
	ProductID            string    `protobuf:"bytes,22,opt,name=productID,proto3" json:"productID,omitempty"`               // strings.ToLower(resource.Kind)
	ProductName          string    `protobuf:"bytes,23,opt,name=productName,proto3" json:"productName,omitempty"`           // "Resource Unit: Kind"
	ShortDescription     string    `protobuf:"bytes,24,opt,name=shortDescription,proto3" json:"shortDescription,omitempty"` // "Kind: Metadata.Name"
	LongDescription      string    `protobuf:"bytes,25,opt,name=longDescription,proto3" json:"longDescription,omitempty"`   // "Kind: ObjKey"
	Category             string    `protobuf:"bytes,31,opt,name=category,proto3" json:"category,omitempty"`
	Amount               int64     `protobuf:"varint,41,opt,name=amount,proto3" json:"amount,omitempty"`                      // US dollar cents (amount charged)
	Currency             string    `protobuf:"bytes,42,opt,name=currency,proto3" json:"currency,omitempty"`                   // it takes the currency from parent invoice
	UnitPrice            string    `protobuf:"bytes,45,opt,name=unitPrice,proto3" json:"unitPrice,omitempty"`                 // [metadata] US dollar cents (discount NOT included)
	Discount             string    `protobuf:"bytes,46,opt,name=discount,proto3" json:"discount,omitempty"`                   // [metadata] pricing-plan discount
	PriceWithDiscount    string    `protobuf:"bytes,47,opt,name=priceWithDiscount,proto3" json:"priceWithDiscount,omitempty"` // [metadata] US dollar cents (discount included)
	CreationDate         int64     `protobuf:"varint,51,opt,name=creationDate,proto3" json:"creationDate,omitempty"`
	CancelationDate      int64     `protobuf:"varint,52,opt,name=cancelationDate,proto3" json:"cancelationDate,omitempty"` // bool cancelAtPeriodEnd = 53;
	StartTime            int64     `protobuf:"varint,61,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              int64     `protobuf:"varint,62,opt,name=endTime,proto3" json:"endTime,omitempty"`
	PeriodStart          int64     `protobuf:"varint,71,opt,name=periodStart,proto3" json:"periodStart,omitempty"`
	PeriodEnd            int64     `protobuf:"varint,72,opt,name=periodEnd,proto3" json:"periodEnd,omitempty"`
	LastPeriodStart      int64     `protobuf:"varint,75,opt,name=lastPeriodStart,proto3" json:"lastPeriodStart,omitempty"`
	LastPeriodEnd        int64     `protobuf:"varint,76,opt,name=lastPeriodEnd,proto3" json:"lastPeriodEnd,omitempty"`
	LastModified         int64     `protobuf:"varint,101,opt,name=lastModified,proto3" json:"lastModified,omitempty"`
	Class                ItemClass `protobuf:"varint,110,opt,name=class,proto3,enum=billing.ItemClass" json:"class,omitempty"`
	Preview              bool      `protobuf:"varint,111,opt,name=preview,proto3" json:"preview,omitempty"`
	Billed               bool      `protobuf:"varint,121,opt,name=billed,proto3" json:"billed,omitempty"`
}

func (x *InvoiceItem) Reset() {
	*x = InvoiceItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoiceItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceItem) ProtoMessage() {}

func (x *InvoiceItem) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvoiceItem.ProtoReflect.Descriptor instead.
func (*InvoiceItem) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_billing_item_proto_rawDescGZIP(), []int{0}
}

func (x *InvoiceItem) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *InvoiceItem) GetStripeCustomerID() string {
	if x != nil {
		return x.StripeCustomerID
	}
	return ""
}

func (x *InvoiceItem) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *InvoiceItem) GetServiceID() string {
	if x != nil {
		return x.ServiceID
	}
	return ""
}

func (x *InvoiceItem) GetProviderID() string {
	if x != nil {
		return x.ProviderID
	}
	return ""
}

func (x *InvoiceItem) GetStripeSubscriptionID() string {
	if x != nil {
		return x.StripeSubscriptionID
	}
	return ""
}

func (x *InvoiceItem) GetStripeInvoiceID() string {
	if x != nil {
		return x.StripeInvoiceID
	}
	return ""
}

func (x *InvoiceItem) GetStripeProductID() string {
	if x != nil {
		return x.StripeProductID
	}
	return ""
}

func (x *InvoiceItem) GetItemID() string {
	if x != nil {
		return x.ItemID
	}
	return ""
}

func (x *InvoiceItem) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *InvoiceItem) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *InvoiceItem) GetShortDescription() string {
	if x != nil {
		return x.ShortDescription
	}
	return ""
}

func (x *InvoiceItem) GetLongDescription() string {
	if x != nil {
		return x.LongDescription
	}
	return ""
}

func (x *InvoiceItem) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *InvoiceItem) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *InvoiceItem) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *InvoiceItem) GetUnitPrice() string {
	if x != nil {
		return x.UnitPrice
	}
	return ""
}

func (x *InvoiceItem) GetDiscount() string {
	if x != nil {
		return x.Discount
	}
	return ""
}

func (x *InvoiceItem) GetPriceWithDiscount() string {
	if x != nil {
		return x.PriceWithDiscount
	}
	return ""
}

func (x *InvoiceItem) GetCreationDate() int64 {
	if x != nil {
		return x.CreationDate
	}
	return 0
}

func (x *InvoiceItem) GetCancelationDate() int64 {
	if x != nil {
		return x.CancelationDate
	}
	return 0
}

func (x *InvoiceItem) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *InvoiceItem) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *InvoiceItem) GetPeriodStart() int64 {
	if x != nil {
		return x.PeriodStart
	}
	return 0
}

func (x *InvoiceItem) GetPeriodEnd() int64 {
	if x != nil {
		return x.PeriodEnd
	}
	return 0
}

func (x *InvoiceItem) GetLastPeriodStart() int64 {
	if x != nil {
		return x.LastPeriodStart
	}
	return 0
}

func (x *InvoiceItem) GetLastPeriodEnd() int64 {
	if x != nil {
		return x.LastPeriodEnd
	}
	return 0
}

func (x *InvoiceItem) GetLastModified() int64 {
	if x != nil {
		return x.LastModified
	}
	return 0
}

func (x *InvoiceItem) GetClass() ItemClass {
	if x != nil {
		return x.Class
	}
	return ItemClass_RESOURCE_UNIT
}

func (x *InvoiceItem) GetPreview() bool {
	if x != nil {
		return x.Preview
	}
	return false
}

func (x *InvoiceItem) GetBilled() bool {
	if x != nil {
		return x.Billed
	}
	return false
}

type InvoiceItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta         *resource.ListResponse `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	InvoiceItems []*InvoiceItem         `protobuf:"bytes,2,rep,name=invoiceItems,proto3" json:"invoiceItems,omitempty"`
}

func (x *InvoiceItems) Reset() {
	*x = InvoiceItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoiceItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceItems) ProtoMessage() {}

func (x *InvoiceItems) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvoiceItems.ProtoReflect.Descriptor instead.
func (*InvoiceItems) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_billing_item_proto_rawDescGZIP(), []int{1}
}

func (x *InvoiceItems) GetMeta() *resource.ListResponse {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *InvoiceItems) GetInvoiceItems() []*InvoiceItem {
	if x != nil {
		return x.InvoiceItems
	}
	return nil
}

type ListInvoiceItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta      *resource.ListRequest `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	AccountID string                `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
}

func (x *ListInvoiceItemsRequest) Reset() {
	*x = ListInvoiceItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInvoiceItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInvoiceItemsRequest) ProtoMessage() {}

func (x *ListInvoiceItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInvoiceItemsRequest.ProtoReflect.Descriptor instead.
func (*ListInvoiceItemsRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_billing_item_proto_rawDescGZIP(), []int{2}
}

func (x *ListInvoiceItemsRequest) GetMeta() *resource.ListRequest {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ListInvoiceItemsRequest) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

var File_mmesh_protobuf_resources_v1_billing_item_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_billing_item_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x1a, 0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x08, 0x0a, 0x0b, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x74, 0x72, 0x69,
	0x70, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x32, 0x0a, 0x14, 0x73, 0x74,
	0x72, 0x69, 0x70, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x28,
	0x0a, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49,
	0x44, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x74, 0x72, 0x69,
	0x70, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x18,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x6c, 0x6f, 0x6e, 0x67, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x6c, 0x6f, 0x6e, 0x67, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x1f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x29, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x2a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x1c, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x2d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x2e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x57, 0x69, 0x74, 0x68, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x2f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x69, 0x63, 0x65, 0x57, 0x69, 0x74, 0x68,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x33, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x34, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x3e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x47, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x45, 0x6e, 0x64, 0x18, 0x48, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x45, 0x6e, 0x64, 0x12, 0x28,
	0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x4b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x45, 0x6e, 0x64, 0x18, 0x4c, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x45, 0x6e, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x65,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x12, 0x28, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x6e, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x69, 0x6c, 0x6c, 0x65, 0x64,
	0x18, 0x79, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x62, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x22, 0x74,
	0x0a, 0x0c, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2a,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x0c, 0x69, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x62, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x29, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x2a, 0x28, 0x0a, 0x09, 0x49, 0x74, 0x65, 0x6d,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x41, 0x41, 0x53,
	0x10, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f,
	0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_billing_item_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_billing_item_proto_rawDescData = file_mmesh_protobuf_resources_v1_billing_item_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_billing_item_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_billing_item_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_billing_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_billing_item_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_billing_item_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_billing_item_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mmesh_protobuf_resources_v1_billing_item_proto_goTypes = []interface{}{
	(ItemClass)(0),                  // 0: billing.ItemClass
	(*InvoiceItem)(nil),             // 1: billing.InvoiceItem
	(*InvoiceItems)(nil),            // 2: billing.InvoiceItems
	(*ListInvoiceItemsRequest)(nil), // 3: billing.ListInvoiceItemsRequest
	(*resource.ListResponse)(nil),   // 4: resource.ListResponse
	(*resource.ListRequest)(nil),    // 5: resource.ListRequest
}
var file_mmesh_protobuf_resources_v1_billing_item_proto_depIdxs = []int32{
	0, // 0: billing.InvoiceItem.class:type_name -> billing.ItemClass
	4, // 1: billing.InvoiceItems.meta:type_name -> resource.ListResponse
	1, // 2: billing.InvoiceItems.invoiceItems:type_name -> billing.InvoiceItem
	5, // 3: billing.ListInvoiceItemsRequest.meta:type_name -> resource.ListRequest
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_billing_item_proto_init() }
func file_mmesh_protobuf_resources_v1_billing_item_proto_init() {
	if File_mmesh_protobuf_resources_v1_billing_item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvoiceItem); i {
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
		file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvoiceItems); i {
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
		file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInvoiceItemsRequest); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_billing_item_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_billing_item_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_billing_item_proto_depIdxs,
		EnumInfos:         file_mmesh_protobuf_resources_v1_billing_item_proto_enumTypes,
		MessageInfos:      file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_billing_item_proto = out.File
	file_mmesh_protobuf_resources_v1_billing_item_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_billing_item_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_billing_item_proto_depIdxs = nil
}
