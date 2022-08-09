// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: mmesh/protobuf/resources/v1/services/catalog/catalog.proto

package catalog

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	compute "mmesh.dev/m-api-go/grpc/resources/services/catalog/cloud/compute"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Catalog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cloud *CloudCatalog `protobuf:"bytes,1,opt,name=cloud,proto3" json:"cloud,omitempty"` // ProCatalog pro = 2;
}

func (x *Catalog) Reset() {
	*x = Catalog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Catalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Catalog) ProtoMessage() {}

func (x *Catalog) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Catalog.ProtoReflect.Descriptor instead.
func (*Catalog) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *Catalog) GetCloud() *CloudCatalog {
	if x != nil {
		return x.Cloud
	}
	return nil
}

type CloudCatalog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locations         map[string]*compute.Location     `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	InstanceTypes     map[string]*compute.InstanceType `protobuf:"bytes,2,rep,name=instanceTypes,proto3" json:"instanceTypes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OSImages          map[string]*compute.Image        `protobuf:"bytes,3,rep,name=OSImages,proto3" json:"OSImages,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AppImages         map[string]*compute.Image        `protobuf:"bytes,4,rep,name=appImages,proto3" json:"appImages,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	KubernetesOptions *compute.KubernetesOptions       `protobuf:"bytes,11,opt,name=kubernetesOptions,proto3" json:"kubernetesOptions,omitempty"`
}

func (x *CloudCatalog) Reset() {
	*x = CloudCatalog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudCatalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudCatalog) ProtoMessage() {}

func (x *CloudCatalog) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudCatalog.ProtoReflect.Descriptor instead.
func (*CloudCatalog) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_rawDescGZIP(), []int{1}
}

func (x *CloudCatalog) GetLocations() map[string]*compute.Location {
	if x != nil {
		return x.Locations
	}
	return nil
}

func (x *CloudCatalog) GetInstanceTypes() map[string]*compute.InstanceType {
	if x != nil {
		return x.InstanceTypes
	}
	return nil
}

func (x *CloudCatalog) GetOSImages() map[string]*compute.Image {
	if x != nil {
		return x.OSImages
	}
	return nil
}

func (x *CloudCatalog) GetAppImages() map[string]*compute.Image {
	if x != nil {
		return x.AppImages
	}
	return nil
}

func (x *CloudCatalog) GetKubernetesOptions() *compute.KubernetesOptions {
	if x != nil {
		return x.KubernetesOptions
	}
	return nil
}

var File_mmesh_protobuf_resources_v1_services_catalog_catalog_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x1a, 0x49, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x4d, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x46, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4b, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x07, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12,
	0x2b, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x22, 0xb6, 0x05, 0x0a,
	0x0c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x42, 0x0a,
	0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x4e, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x12, 0x3f, 0x0a, 0x08, 0x4f, 0x53, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x4f, 0x53, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x4f, 0x53, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x42, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x41, 0x70, 0x70,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x61, 0x70, 0x70,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x11, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x11, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x1a, 0x4f, 0x0a, 0x0e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x57, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4b, 0x0a, 0x0d, 0x4f, 0x53,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x24, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4c, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x34, 0x5a, 0x32, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64,
	0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_rawDescData = file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_goTypes = []interface{}{
	(*Catalog)(nil),                   // 0: catalog.Catalog
	(*CloudCatalog)(nil),              // 1: catalog.CloudCatalog
	nil,                               // 2: catalog.CloudCatalog.LocationsEntry
	nil,                               // 3: catalog.CloudCatalog.InstanceTypesEntry
	nil,                               // 4: catalog.CloudCatalog.OSImagesEntry
	nil,                               // 5: catalog.CloudCatalog.AppImagesEntry
	(*compute.KubernetesOptions)(nil), // 6: compute.KubernetesOptions
	(*compute.Location)(nil),          // 7: compute.Location
	(*compute.InstanceType)(nil),      // 8: compute.InstanceType
	(*compute.Image)(nil),             // 9: compute.Image
}
var file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_depIdxs = []int32{
	1,  // 0: catalog.Catalog.cloud:type_name -> catalog.CloudCatalog
	2,  // 1: catalog.CloudCatalog.locations:type_name -> catalog.CloudCatalog.LocationsEntry
	3,  // 2: catalog.CloudCatalog.instanceTypes:type_name -> catalog.CloudCatalog.InstanceTypesEntry
	4,  // 3: catalog.CloudCatalog.OSImages:type_name -> catalog.CloudCatalog.OSImagesEntry
	5,  // 4: catalog.CloudCatalog.appImages:type_name -> catalog.CloudCatalog.AppImagesEntry
	6,  // 5: catalog.CloudCatalog.kubernetesOptions:type_name -> compute.KubernetesOptions
	7,  // 6: catalog.CloudCatalog.LocationsEntry.value:type_name -> compute.Location
	8,  // 7: catalog.CloudCatalog.InstanceTypesEntry.value:type_name -> compute.InstanceType
	9,  // 8: catalog.CloudCatalog.OSImagesEntry.value:type_name -> compute.Image
	9,  // 9: catalog.CloudCatalog.AppImagesEntry.value:type_name -> compute.Image
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_init() }
func file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_init() {
	if File_mmesh_protobuf_resources_v1_services_catalog_catalog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Catalog); i {
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
		file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudCatalog); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_services_catalog_catalog_proto = out.File
	file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_services_catalog_catalog_proto_depIdxs = nil
}
