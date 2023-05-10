// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: mmesh/protobuf/resources/v1/messaging/mail.proto

package messaging

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

type MailTemplate int32

const (
	MailTemplate_UNDEFINED            MailTemplate = 0
	MailTemplate_ACCOUNT_CONFIRMATION MailTemplate = 11
	MailTemplate_EMAIL_CONFIRMATION   MailTemplate = 21
	MailTemplate_PASSWORD_RESET       MailTemplate = 31
	MailTemplate_MAGIC_LOGIN_LINK     MailTemplate = 41
)

// Enum value maps for MailTemplate.
var (
	MailTemplate_name = map[int32]string{
		0:  "UNDEFINED",
		11: "ACCOUNT_CONFIRMATION",
		21: "EMAIL_CONFIRMATION",
		31: "PASSWORD_RESET",
		41: "MAGIC_LOGIN_LINK",
	}
	MailTemplate_value = map[string]int32{
		"UNDEFINED":            0,
		"ACCOUNT_CONFIRMATION": 11,
		"EMAIL_CONFIRMATION":   21,
		"PASSWORD_RESET":       31,
		"MAGIC_LOGIN_LINK":     41,
	}
)

func (x MailTemplate) Enum() *MailTemplate {
	p := new(MailTemplate)
	*p = x
	return p
}

func (x MailTemplate) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MailTemplate) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_messaging_mail_proto_enumTypes[0].Descriptor()
}

func (MailTemplate) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_messaging_mail_proto_enumTypes[0]
}

func (x MailTemplate) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MailTemplate.Descriptor instead.
func (MailTemplate) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_messaging_mail_proto_rawDescGZIP(), []int{0}
}

type MailMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageID string   `protobuf:"bytes,1,opt,name=messageID,proto3" json:"messageID,omitempty"`
	Timestamp int64    `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Sender    *Contact `protobuf:"bytes,11,opt,name=sender,proto3" json:"sender,omitempty"`
	ReplyTo   *Contact `protobuf:"bytes,12,opt,name=replyTo,proto3" json:"replyTo,omitempty"`
	Recipient *Contact `protobuf:"bytes,13,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Subject   string   `protobuf:"bytes,21,opt,name=subject,proto3" json:"subject,omitempty"`
	HtmlBody  string   `protobuf:"bytes,31,opt,name=htmlBody,proto3" json:"htmlBody,omitempty"`
	TextBody  string   `protobuf:"bytes,32,opt,name=textBody,proto3" json:"textBody,omitempty"`
	// MailTemplate template = 41;
	ProviderID   int64         `protobuf:"varint,51,opt,name=providerID,proto3" json:"providerID,omitempty"`
	Status       MessageStatus `protobuf:"varint,61,opt,name=status,proto3,enum=messaging.MessageStatus" json:"status,omitempty"`
	Retries      int32         `protobuf:"varint,71,opt,name=retries,proto3" json:"retries,omitempty"`
	LastModified int64         `protobuf:"varint,101,opt,name=lastModified,proto3" json:"lastModified,omitempty"`
}

func (x *MailMessage) Reset() {
	*x = MailMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_messaging_mail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailMessage) ProtoMessage() {}

func (x *MailMessage) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_messaging_mail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailMessage.ProtoReflect.Descriptor instead.
func (*MailMessage) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_messaging_mail_proto_rawDescGZIP(), []int{0}
}

func (x *MailMessage) GetMessageID() string {
	if x != nil {
		return x.MessageID
	}
	return ""
}

func (x *MailMessage) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *MailMessage) GetSender() *Contact {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *MailMessage) GetReplyTo() *Contact {
	if x != nil {
		return x.ReplyTo
	}
	return nil
}

func (x *MailMessage) GetRecipient() *Contact {
	if x != nil {
		return x.Recipient
	}
	return nil
}

func (x *MailMessage) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *MailMessage) GetHtmlBody() string {
	if x != nil {
		return x.HtmlBody
	}
	return ""
}

func (x *MailMessage) GetTextBody() string {
	if x != nil {
		return x.TextBody
	}
	return ""
}

func (x *MailMessage) GetProviderID() int64 {
	if x != nil {
		return x.ProviderID
	}
	return 0
}

func (x *MailMessage) GetStatus() MessageStatus {
	if x != nil {
		return x.Status
	}
	return MessageStatus_SENT
}

func (x *MailMessage) GetRetries() int32 {
	if x != nil {
		return x.Retries
	}
	return 0
}

func (x *MailMessage) GetLastModified() int64 {
	if x != nil {
		return x.LastModified
	}
	return 0
}

type MailMessages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta         *resource.ListResponse `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	MailMessages []*MailMessage         `protobuf:"bytes,2,rep,name=mailMessages,proto3" json:"mailMessages,omitempty"`
}

func (x *MailMessages) Reset() {
	*x = MailMessages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_messaging_mail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailMessages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailMessages) ProtoMessage() {}

func (x *MailMessages) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_messaging_mail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailMessages.ProtoReflect.Descriptor instead.
func (*MailMessages) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_messaging_mail_proto_rawDescGZIP(), []int{1}
}

func (x *MailMessages) GetMeta() *resource.ListResponse {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *MailMessages) GetMailMessages() []*MailMessage {
	if x != nil {
		return x.MailMessages
	}
	return nil
}

type ListMailMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *resource.ListRequest `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *ListMailMessagesRequest) Reset() {
	*x = ListMailMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_messaging_mail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMailMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMailMessagesRequest) ProtoMessage() {}

func (x *ListMailMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_messaging_mail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMailMessagesRequest.ProtoReflect.Descriptor instead.
func (*ListMailMessagesRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_messaging_mail_proto_rawDescGZIP(), []int{2}
}

func (x *ListMailMessagesRequest) GetMeta() *resource.ListRequest {
	if x != nil {
		return x.Meta
	}
	return nil
}

type SendMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RcptEmail    string       `protobuf:"bytes,1,opt,name=rcptEmail,proto3" json:"rcptEmail,omitempty"`
	MailTemplate MailTemplate `protobuf:"varint,11,opt,name=mailTemplate,proto3,enum=messaging.MailTemplate" json:"mailTemplate,omitempty"`
	Link         string       `protobuf:"bytes,21,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *SendMailRequest) Reset() {
	*x = SendMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_messaging_mail_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailRequest) ProtoMessage() {}

func (x *SendMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_messaging_mail_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailRequest.ProtoReflect.Descriptor instead.
func (*SendMailRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_messaging_mail_proto_rawDescGZIP(), []int{3}
}

func (x *SendMailRequest) GetRcptEmail() string {
	if x != nil {
		return x.RcptEmail
	}
	return ""
}

func (x *SendMailRequest) GetMailTemplate() MailTemplate {
	if x != nil {
		return x.MailTemplate
	}
	return MailTemplate_UNDEFINED
}

func (x *SendMailRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

var File_mmesh_protobuf_resources_v1_messaging_mail_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_messaging_mail_proto_rawDesc = []byte{
	0x0a, 0x30, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x1a, 0x33, 0x6d,
	0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x03, 0x0a, 0x0b, 0x4d, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x44, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x2a, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x07, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x12, 0x30, 0x0a, 0x09, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x74, 0x6d, 0x6c, 0x42, 0x6f, 0x64,
	0x79, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x74, 0x6d, 0x6c, 0x42, 0x6f, 0x64,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x78, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x20, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x78, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x33, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x30, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x47, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73,
	0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0x76, 0x0a,
	0x0c, 0x4d, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x2a, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x3a, 0x0a, 0x0c, 0x6d, 0x61, 0x69,
	0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x61, 0x69, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0c, 0x6d, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x44, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x69,
	0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x29, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x80, 0x01, 0x0a, 0x0f,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x63, 0x70, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x63, 0x70, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x3b, 0x0a,
	0x0c, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e,
	0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x6d, 0x61,
	0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69,
	0x6e, 0x6b, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x2a, 0x79,
	0x0a, 0x0c, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x0d,
	0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a,
	0x14, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0b, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x4d, 0x41, 0x49, 0x4c,
	0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x15, 0x12,
	0x12, 0x0a, 0x0e, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x45,
	0x54, 0x10, 0x1f, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x41, 0x47, 0x49, 0x43, 0x5f, 0x4c, 0x4f, 0x47,
	0x49, 0x4e, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x29, 0x42, 0x2d, 0x5a, 0x2b, 0x6d, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_messaging_mail_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_messaging_mail_proto_rawDescData = file_mmesh_protobuf_resources_v1_messaging_mail_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_messaging_mail_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_messaging_mail_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_messaging_mail_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_messaging_mail_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_messaging_mail_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_messaging_mail_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mmesh_protobuf_resources_v1_messaging_mail_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mmesh_protobuf_resources_v1_messaging_mail_proto_goTypes = []interface{}{
	(MailTemplate)(0),               // 0: messaging.MailTemplate
	(*MailMessage)(nil),             // 1: messaging.MailMessage
	(*MailMessages)(nil),            // 2: messaging.MailMessages
	(*ListMailMessagesRequest)(nil), // 3: messaging.ListMailMessagesRequest
	(*SendMailRequest)(nil),         // 4: messaging.SendMailRequest
	(*Contact)(nil),                 // 5: messaging.Contact
	(MessageStatus)(0),              // 6: messaging.MessageStatus
	(*resource.ListResponse)(nil),   // 7: resource.ListResponse
	(*resource.ListRequest)(nil),    // 8: resource.ListRequest
}
var file_mmesh_protobuf_resources_v1_messaging_mail_proto_depIdxs = []int32{
	5, // 0: messaging.MailMessage.sender:type_name -> messaging.Contact
	5, // 1: messaging.MailMessage.replyTo:type_name -> messaging.Contact
	5, // 2: messaging.MailMessage.recipient:type_name -> messaging.Contact
	6, // 3: messaging.MailMessage.status:type_name -> messaging.MessageStatus
	7, // 4: messaging.MailMessages.meta:type_name -> resource.ListResponse
	1, // 5: messaging.MailMessages.mailMessages:type_name -> messaging.MailMessage
	8, // 6: messaging.ListMailMessagesRequest.meta:type_name -> resource.ListRequest
	0, // 7: messaging.SendMailRequest.mailTemplate:type_name -> messaging.MailTemplate
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_messaging_mail_proto_init() }
func file_mmesh_protobuf_resources_v1_messaging_mail_proto_init() {
	if File_mmesh_protobuf_resources_v1_messaging_mail_proto != nil {
		return
	}
	file_mmesh_protobuf_resources_v1_messaging_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_messaging_mail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailMessage); i {
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
		file_mmesh_protobuf_resources_v1_messaging_mail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailMessages); i {
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
		file_mmesh_protobuf_resources_v1_messaging_mail_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMailMessagesRequest); i {
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
		file_mmesh_protobuf_resources_v1_messaging_mail_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMailRequest); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_messaging_mail_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_messaging_mail_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_messaging_mail_proto_depIdxs,
		EnumInfos:         file_mmesh_protobuf_resources_v1_messaging_mail_proto_enumTypes,
		MessageInfos:      file_mmesh_protobuf_resources_v1_messaging_mail_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_messaging_mail_proto = out.File
	file_mmesh_protobuf_resources_v1_messaging_mail_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_messaging_mail_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_messaging_mail_proto_depIdxs = nil
}
