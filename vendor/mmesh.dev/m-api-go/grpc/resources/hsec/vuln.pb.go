// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.26.0--rc2
// source: mmesh/protobuf/resources/v1/hsec/vuln.proto

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

type VulnerabilityStatus int32

const (
	VulnerabilityStatus_STATUS_UNKNOWN             VulnerabilityStatus = 0
	VulnerabilityStatus_STATUS_NOT_AFFECTED        VulnerabilityStatus = 11
	VulnerabilityStatus_STATUS_AFFECTED            VulnerabilityStatus = 21
	VulnerabilityStatus_STATUS_FIXED               VulnerabilityStatus = 31
	VulnerabilityStatus_STATUS_UNDER_INVESTIGATION VulnerabilityStatus = 41
	VulnerabilityStatus_STATUS_WILL_NOT_FIX        VulnerabilityStatus = 51
	VulnerabilityStatus_STATUS_FIX_DEFERRED        VulnerabilityStatus = 61
	VulnerabilityStatus_STATUS_END_OF_LIFE         VulnerabilityStatus = 71
)

// Enum value maps for VulnerabilityStatus.
var (
	VulnerabilityStatus_name = map[int32]string{
		0:  "STATUS_UNKNOWN",
		11: "STATUS_NOT_AFFECTED",
		21: "STATUS_AFFECTED",
		31: "STATUS_FIXED",
		41: "STATUS_UNDER_INVESTIGATION",
		51: "STATUS_WILL_NOT_FIX",
		61: "STATUS_FIX_DEFERRED",
		71: "STATUS_END_OF_LIFE",
	}
	VulnerabilityStatus_value = map[string]int32{
		"STATUS_UNKNOWN":             0,
		"STATUS_NOT_AFFECTED":        11,
		"STATUS_AFFECTED":            21,
		"STATUS_FIXED":               31,
		"STATUS_UNDER_INVESTIGATION": 41,
		"STATUS_WILL_NOT_FIX":        51,
		"STATUS_FIX_DEFERRED":        61,
		"STATUS_END_OF_LIFE":         71,
	}
)

func (x VulnerabilityStatus) Enum() *VulnerabilityStatus {
	p := new(VulnerabilityStatus)
	*p = x
	return p
}

func (x VulnerabilityStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VulnerabilityStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_hsec_vuln_proto_enumTypes[0].Descriptor()
}

func (VulnerabilityStatus) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_hsec_vuln_proto_enumTypes[0]
}

func (x VulnerabilityStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VulnerabilityStatus.Descriptor instead.
func (VulnerabilityStatus) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_hsec_vuln_proto_rawDescGZIP(), []int{0}
}

type Severity int32

const (
	Severity_SEVERITY_UNKNOWN  Severity = 0
	Severity_SEVERITY_LOW      Severity = 11
	Severity_SEVERITY_MEDIUM   Severity = 21
	Severity_SEVERITY_HIGH     Severity = 31
	Severity_SEVERITY_CRITICAL Severity = 41
)

// Enum value maps for Severity.
var (
	Severity_name = map[int32]string{
		0:  "SEVERITY_UNKNOWN",
		11: "SEVERITY_LOW",
		21: "SEVERITY_MEDIUM",
		31: "SEVERITY_HIGH",
		41: "SEVERITY_CRITICAL",
	}
	Severity_value = map[string]int32{
		"SEVERITY_UNKNOWN":  0,
		"SEVERITY_LOW":      11,
		"SEVERITY_MEDIUM":   21,
		"SEVERITY_HIGH":     31,
		"SEVERITY_CRITICAL": 41,
	}
)

func (x Severity) Enum() *Severity {
	p := new(Severity)
	*p = x
	return p
}

func (x Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_hsec_vuln_proto_enumTypes[1].Descriptor()
}

func (Severity) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_hsec_vuln_proto_enumTypes[1]
}

func (x Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Severity.Descriptor instead.
func (Severity) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_hsec_vuln_proto_rawDescGZIP(), []int{1}
}

type DetectedVulnerability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VulnerabilityID  string              `protobuf:"bytes,1,opt,name=vulnerabilityID,proto3" json:"vulnerabilityID,omitempty"`
	VendorIDs        []string            `protobuf:"bytes,11,rep,name=vendorIDs,proto3" json:"vendorIDs,omitempty"`
	PkgID            string              `protobuf:"bytes,21,opt,name=pkgID,proto3" json:"pkgID,omitempty"`
	PkgName          string              `protobuf:"bytes,31,opt,name=pkgName,proto3" json:"pkgName,omitempty"`
	PkgPath          string              `protobuf:"bytes,41,opt,name=pkgPath,proto3" json:"pkgPath,omitempty"`
	PkgIdentifier    *PkgIdentifier      `protobuf:"bytes,45,opt,name=pkgIdentifier,proto3" json:"pkgIdentifier,omitempty"`
	InstalledVersion string              `protobuf:"bytes,51,opt,name=installedVersion,proto3" json:"installedVersion,omitempty"`
	FixedVersion     string              `protobuf:"bytes,61,opt,name=fixedVersion,proto3" json:"fixedVersion,omitempty"`
	Status           VulnerabilityStatus `protobuf:"varint,71,opt,name=status,proto3,enum=hsec.VulnerabilityStatus" json:"status,omitempty"`
	Layer            *Layer              `protobuf:"bytes,81,opt,name=layer,proto3" json:"layer,omitempty"`
	SeveritySource   string              `protobuf:"bytes,91,opt,name=severitySource,proto3" json:"severitySource,omitempty"`
	PrimaryURL       string              `protobuf:"bytes,101,opt,name=primaryURL,proto3" json:"primaryURL,omitempty"`
	DataSource       *DataSource         `protobuf:"bytes,121,opt,name=dataSource,proto3" json:"dataSource,omitempty"`
	Vulnerability    *Vulnerability      `protobuf:"bytes,200,opt,name=vulnerability,proto3" json:"vulnerability,omitempty"`
}

func (x *DetectedVulnerability) Reset() {
	*x = DetectedVulnerability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_hsec_vuln_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectedVulnerability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectedVulnerability) ProtoMessage() {}

func (x *DetectedVulnerability) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_hsec_vuln_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectedVulnerability.ProtoReflect.Descriptor instead.
func (*DetectedVulnerability) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_hsec_vuln_proto_rawDescGZIP(), []int{0}
}

func (x *DetectedVulnerability) GetVulnerabilityID() string {
	if x != nil {
		return x.VulnerabilityID
	}
	return ""
}

func (x *DetectedVulnerability) GetVendorIDs() []string {
	if x != nil {
		return x.VendorIDs
	}
	return nil
}

func (x *DetectedVulnerability) GetPkgID() string {
	if x != nil {
		return x.PkgID
	}
	return ""
}

func (x *DetectedVulnerability) GetPkgName() string {
	if x != nil {
		return x.PkgName
	}
	return ""
}

func (x *DetectedVulnerability) GetPkgPath() string {
	if x != nil {
		return x.PkgPath
	}
	return ""
}

func (x *DetectedVulnerability) GetPkgIdentifier() *PkgIdentifier {
	if x != nil {
		return x.PkgIdentifier
	}
	return nil
}

func (x *DetectedVulnerability) GetInstalledVersion() string {
	if x != nil {
		return x.InstalledVersion
	}
	return ""
}

func (x *DetectedVulnerability) GetFixedVersion() string {
	if x != nil {
		return x.FixedVersion
	}
	return ""
}

func (x *DetectedVulnerability) GetStatus() VulnerabilityStatus {
	if x != nil {
		return x.Status
	}
	return VulnerabilityStatus_STATUS_UNKNOWN
}

func (x *DetectedVulnerability) GetLayer() *Layer {
	if x != nil {
		return x.Layer
	}
	return nil
}

func (x *DetectedVulnerability) GetSeveritySource() string {
	if x != nil {
		return x.SeveritySource
	}
	return ""
}

func (x *DetectedVulnerability) GetPrimaryURL() string {
	if x != nil {
		return x.PrimaryURL
	}
	return ""
}

func (x *DetectedVulnerability) GetDataSource() *DataSource {
	if x != nil {
		return x.DataSource
	}
	return nil
}

func (x *DetectedVulnerability) GetVulnerability() *Vulnerability {
	if x != nil {
		return x.Vulnerability
	}
	return nil
}

type DataSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`
	URL  string `protobuf:"bytes,21,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *DataSource) Reset() {
	*x = DataSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_hsec_vuln_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSource) ProtoMessage() {}

func (x *DataSource) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_hsec_vuln_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSource.ProtoReflect.Descriptor instead.
func (*DataSource) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_hsec_vuln_proto_rawDescGZIP(), []int{1}
}

func (x *DataSource) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *DataSource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataSource) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

type Vulnerability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title            string              `protobuf:"bytes,11,opt,name=title,proto3" json:"title,omitempty"`
	Description      string              `protobuf:"bytes,21,opt,name=description,proto3" json:"description,omitempty"`
	Severity         string              `protobuf:"bytes,31,opt,name=severity,proto3" json:"severity,omitempty"`
	CweIDs           []string            `protobuf:"bytes,41,rep,name=cweIDs,proto3" json:"cweIDs,omitempty"`
	VendorSeverity   map[string]Severity `protobuf:"bytes,51,rep,name=vendorSeverity,proto3" json:"vendorSeverity,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=hsec.Severity"` // map[sourceID]*Severity
	CVSS             map[string]*CVSS    `protobuf:"bytes,61,rep,name=CVSS,proto3" json:"CVSS,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`                                         // map[sourceID]*CVSS
	References       []string            `protobuf:"bytes,71,rep,name=references,proto3" json:"references,omitempty"`
	PublishedDate    int64               `protobuf:"varint,81,opt,name=publishedDate,proto3" json:"publishedDate,omitempty"`
	LastModifiedDate int64               `protobuf:"varint,91,opt,name=lastModifiedDate,proto3" json:"lastModifiedDate,omitempty"`
}

func (x *Vulnerability) Reset() {
	*x = Vulnerability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_hsec_vuln_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vulnerability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vulnerability) ProtoMessage() {}

func (x *Vulnerability) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_hsec_vuln_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vulnerability.ProtoReflect.Descriptor instead.
func (*Vulnerability) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_hsec_vuln_proto_rawDescGZIP(), []int{2}
}

func (x *Vulnerability) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Vulnerability) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Vulnerability) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *Vulnerability) GetCweIDs() []string {
	if x != nil {
		return x.CweIDs
	}
	return nil
}

func (x *Vulnerability) GetVendorSeverity() map[string]Severity {
	if x != nil {
		return x.VendorSeverity
	}
	return nil
}

func (x *Vulnerability) GetCVSS() map[string]*CVSS {
	if x != nil {
		return x.CVSS
	}
	return nil
}

func (x *Vulnerability) GetReferences() []string {
	if x != nil {
		return x.References
	}
	return nil
}

func (x *Vulnerability) GetPublishedDate() int64 {
	if x != nil {
		return x.PublishedDate
	}
	return 0
}

func (x *Vulnerability) GetLastModifiedDate() int64 {
	if x != nil {
		return x.LastModifiedDate
	}
	return 0
}

type CVSS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V2Vector string  `protobuf:"bytes,1,opt,name=v2Vector,proto3" json:"v2Vector,omitempty"`
	V3Vector string  `protobuf:"bytes,11,opt,name=v3Vector,proto3" json:"v3Vector,omitempty"`
	V2Score  float64 `protobuf:"fixed64,21,opt,name=v2Score,proto3" json:"v2Score,omitempty"`
	V3Score  float64 `protobuf:"fixed64,31,opt,name=v3Score,proto3" json:"v3Score,omitempty"`
}

func (x *CVSS) Reset() {
	*x = CVSS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_hsec_vuln_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CVSS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CVSS) ProtoMessage() {}

func (x *CVSS) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_hsec_vuln_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CVSS.ProtoReflect.Descriptor instead.
func (*CVSS) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_hsec_vuln_proto_rawDescGZIP(), []int{3}
}

func (x *CVSS) GetV2Vector() string {
	if x != nil {
		return x.V2Vector
	}
	return ""
}

func (x *CVSS) GetV3Vector() string {
	if x != nil {
		return x.V3Vector
	}
	return ""
}

func (x *CVSS) GetV2Score() float64 {
	if x != nil {
		return x.V2Score
	}
	return 0
}

func (x *CVSS) GetV3Score() float64 {
	if x != nil {
		return x.V3Score
	}
	return 0
}

var File_mmesh_protobuf_resources_v1_hsec_vuln_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_hsec_vuln_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x73,
	0x65, 0x63, 0x2f, 0x76, 0x75, 0x6c, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x68,
	0x73, 0x65, 0x63, 0x1a, 0x2d, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x68, 0x73, 0x65, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc0, 0x04, 0x0a, 0x15, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x56,
	0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x28, 0x0a, 0x0f,
	0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x49, 0x44, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x49, 0x44, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6b, 0x67, 0x49, 0x44, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6b, 0x67, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6b,
	0x67, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6b, 0x67,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6b, 0x67, 0x50, 0x61, 0x74, 0x68, 0x18,
	0x29, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6b, 0x67, 0x50, 0x61, 0x74, 0x68, 0x12, 0x39,
	0x0a, 0x0d, 0x70, 0x6b, 0x67, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18,
	0x2d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x73, 0x65, 0x63, 0x2e, 0x50, 0x6b, 0x67,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0d, 0x70, 0x6b, 0x67, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x33, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x69, 0x78, 0x65, 0x64, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x47, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x68, 0x73, 0x65, 0x63,
	0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x05,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x51, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x68, 0x73,
	0x65, 0x63, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x05, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12,
	0x26, 0x0a, 0x0e, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x5b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74,
	0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x55, 0x52, 0x4c, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x55, 0x52, 0x4c, 0x12, 0x30, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x79, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x68, 0x73,
	0x65, 0x63, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0a, 0x64,
	0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x76, 0x75, 0x6c,
	0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x68, 0x73, 0x65, 0x63, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0d, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x42, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x22, 0x89, 0x04, 0x0a, 0x0d, 0x56, 0x75,
	0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x77, 0x65, 0x49, 0x44, 0x73, 0x18, 0x29, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x77, 0x65, 0x49, 0x44, 0x73, 0x12, 0x4f, 0x0a, 0x0e, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x33, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x68, 0x73, 0x65, 0x63, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x31, 0x0a, 0x04, 0x43, 0x56, 0x53, 0x53,
	0x18, 0x3d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x68, 0x73, 0x65, 0x63, 0x2e, 0x56, 0x75,
	0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x56, 0x53, 0x53,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x43, 0x56, 0x53, 0x53, 0x12, 0x1e, 0x0a, 0x0a, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x47, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x51, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x2a, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x5b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6c, 0x61, 0x73,
	0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x1a, 0x51, 0x0a,
	0x13, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x68, 0x73, 0x65, 0x63, 0x2e, 0x53, 0x65, 0x76,
	0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x43, 0x0a, 0x09, 0x43, 0x56, 0x53, 0x53, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x68, 0x73, 0x65, 0x63, 0x2e, 0x43, 0x56, 0x53, 0x53, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x72, 0x0a, 0x04, 0x43, 0x56, 0x53, 0x53, 0x12, 0x1a, 0x0a,
	0x08, 0x76, 0x32, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x76, 0x32, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x33, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x33, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x32, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x76, 0x32, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x33, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x07, 0x76, 0x33, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x2a, 0xd3, 0x01, 0x0a, 0x13, 0x56, 0x75,
	0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x46, 0x46, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x13,
	0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x46, 0x46, 0x45, 0x43, 0x54, 0x45,
	0x44, 0x10, 0x15, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x49,
	0x58, 0x45, 0x44, 0x10, 0x1f, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x55, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x45, 0x53, 0x54, 0x49, 0x47, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x29, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x57, 0x49, 0x4c, 0x4c, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x49, 0x58, 0x10, 0x33, 0x12, 0x17,
	0x0a, 0x13, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x49, 0x58, 0x5f, 0x44, 0x45, 0x46,
	0x45, 0x52, 0x52, 0x45, 0x44, 0x10, 0x3d, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x4f, 0x46, 0x5f, 0x4c, 0x49, 0x46, 0x45, 0x10, 0x47, 0x2a,
	0x71, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x10, 0x53,
	0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x4f,
	0x57, 0x10, 0x0b, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f,
	0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x15, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x45, 0x56, 0x45,
	0x52, 0x49, 0x54, 0x59, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x1f, 0x12, 0x15, 0x0a, 0x11, 0x53,
	0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c,
	0x10, 0x29, 0x42, 0x28, 0x5a, 0x26, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f,
	0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x68, 0x73, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_hsec_vuln_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_hsec_vuln_proto_rawDescData = file_mmesh_protobuf_resources_v1_hsec_vuln_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_hsec_vuln_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_hsec_vuln_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_hsec_vuln_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_hsec_vuln_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_hsec_vuln_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_hsec_vuln_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_mmesh_protobuf_resources_v1_hsec_vuln_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mmesh_protobuf_resources_v1_hsec_vuln_proto_goTypes = []interface{}{
	(VulnerabilityStatus)(0),      // 0: hsec.VulnerabilityStatus
	(Severity)(0),                 // 1: hsec.Severity
	(*DetectedVulnerability)(nil), // 2: hsec.DetectedVulnerability
	(*DataSource)(nil),            // 3: hsec.DataSource
	(*Vulnerability)(nil),         // 4: hsec.Vulnerability
	(*CVSS)(nil),                  // 5: hsec.CVSS
	nil,                           // 6: hsec.Vulnerability.VendorSeverityEntry
	nil,                           // 7: hsec.Vulnerability.CVSSEntry
	(*PkgIdentifier)(nil),         // 8: hsec.PkgIdentifier
	(*Layer)(nil),                 // 9: hsec.Layer
}
var file_mmesh_protobuf_resources_v1_hsec_vuln_proto_depIdxs = []int32{
	8, // 0: hsec.DetectedVulnerability.pkgIdentifier:type_name -> hsec.PkgIdentifier
	0, // 1: hsec.DetectedVulnerability.status:type_name -> hsec.VulnerabilityStatus
	9, // 2: hsec.DetectedVulnerability.layer:type_name -> hsec.Layer
	3, // 3: hsec.DetectedVulnerability.dataSource:type_name -> hsec.DataSource
	4, // 4: hsec.DetectedVulnerability.vulnerability:type_name -> hsec.Vulnerability
	6, // 5: hsec.Vulnerability.vendorSeverity:type_name -> hsec.Vulnerability.VendorSeverityEntry
	7, // 6: hsec.Vulnerability.CVSS:type_name -> hsec.Vulnerability.CVSSEntry
	1, // 7: hsec.Vulnerability.VendorSeverityEntry.value:type_name -> hsec.Severity
	5, // 8: hsec.Vulnerability.CVSSEntry.value:type_name -> hsec.CVSS
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_hsec_vuln_proto_init() }
func file_mmesh_protobuf_resources_v1_hsec_vuln_proto_init() {
	if File_mmesh_protobuf_resources_v1_hsec_vuln_proto != nil {
		return
	}
	file_mmesh_protobuf_resources_v1_hsec_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_hsec_vuln_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectedVulnerability); i {
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
		file_mmesh_protobuf_resources_v1_hsec_vuln_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSource); i {
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
		file_mmesh_protobuf_resources_v1_hsec_vuln_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vulnerability); i {
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
		file_mmesh_protobuf_resources_v1_hsec_vuln_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CVSS); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_hsec_vuln_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_hsec_vuln_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_hsec_vuln_proto_depIdxs,
		EnumInfos:         file_mmesh_protobuf_resources_v1_hsec_vuln_proto_enumTypes,
		MessageInfos:      file_mmesh_protobuf_resources_v1_hsec_vuln_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_hsec_vuln_proto = out.File
	file_mmesh_protobuf_resources_v1_hsec_vuln_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_hsec_vuln_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_hsec_vuln_proto_depIdxs = nil
}
