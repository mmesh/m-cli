// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.26.0--rc2
// source: mmesh/protobuf/resources/v1/events/event.proto

package events

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

type EventType int32

const (
	EventType_UNDEFINED_TYPE EventType = 0
	EventType_ALERT          EventType = 10
	EventType_CHANGE         EventType = 20
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0:  "UNDEFINED_TYPE",
		10: "ALERT",
		20: "CHANGE",
	}
	EventType_value = map[string]int32{
		"UNDEFINED_TYPE": 0,
		"ALERT":          10,
		"CHANGE":         20,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_events_event_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_events_event_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_event_proto_rawDescGZIP(), []int{0}
}

type Class int32

const (
	Class_UNDEFINED_CLASS Class = 0
	Class_BILLING         Class = 10
	Class_DATABASE        Class = 20
	Class_NETWORK         Class = 30
	Class_SECURITY        Class = 40
	Class_CONFIGURATION   Class = 50
	Class_HOST            Class = 60
	Class_WORKFLOW        Class = 70
	Class_KUBERNETES      Class = 80
	Class_USER_ACTIVITY   Class = 100
	Class_EXTERNAL        Class = 1000
)

// Enum value maps for Class.
var (
	Class_name = map[int32]string{
		0:    "UNDEFINED_CLASS",
		10:   "BILLING",
		20:   "DATABASE",
		30:   "NETWORK",
		40:   "SECURITY",
		50:   "CONFIGURATION",
		60:   "HOST",
		70:   "WORKFLOW",
		80:   "KUBERNETES",
		100:  "USER_ACTIVITY",
		1000: "EXTERNAL",
	}
	Class_value = map[string]int32{
		"UNDEFINED_CLASS": 0,
		"BILLING":         10,
		"DATABASE":        20,
		"NETWORK":         30,
		"SECURITY":        40,
		"CONFIGURATION":   50,
		"HOST":            60,
		"WORKFLOW":        70,
		"KUBERNETES":      80,
		"USER_ACTIVITY":   100,
		"EXTERNAL":        1000,
	}
)

func (x Class) Enum() *Class {
	p := new(Class)
	*p = x
	return p
}

func (x Class) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Class) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_events_event_proto_enumTypes[1].Descriptor()
}

func (Class) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_events_event_proto_enumTypes[1]
}

func (x Class) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Class.Descriptor instead.
func (Class) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_event_proto_rawDescGZIP(), []int{1}
}

type Group int32

const (
	Group_UNDEFINED_GROUP         Group = 0
	Group_ROUTING                 Group = 10
	Group_IAM                     Group = 20 // user auth, rbac...
	Group_HOST_METRICS            Group = 30
	Group_WORKFLOW_TASK           Group = 40
	Group_WEBHOOK                 Group = 50
	Group_THIRD_PARTY_INTEGRATION Group = 60
)

// Enum value maps for Group.
var (
	Group_name = map[int32]string{
		0:  "UNDEFINED_GROUP",
		10: "ROUTING",
		20: "IAM",
		30: "HOST_METRICS",
		40: "WORKFLOW_TASK",
		50: "WEBHOOK",
		60: "THIRD_PARTY_INTEGRATION",
	}
	Group_value = map[string]int32{
		"UNDEFINED_GROUP":         0,
		"ROUTING":                 10,
		"IAM":                     20,
		"HOST_METRICS":            30,
		"WORKFLOW_TASK":           40,
		"WEBHOOK":                 50,
		"THIRD_PARTY_INTEGRATION": 60,
	}
)

func (x Group) Enum() *Group {
	p := new(Group)
	*p = x
	return p
}

func (x Group) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Group) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_events_event_proto_enumTypes[2].Descriptor()
}

func (Group) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_events_event_proto_enumTypes[2]
}

func (x Group) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Group.Descriptor instead.
func (Group) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_event_proto_rawDescGZIP(), []int{2}
}

type Severity int32

const (
	Severity_UNDEFINED_SEVERITY Severity = 0
	Severity_INFO               Severity = 10
	Severity_WARNING            Severity = 20
	Severity_ERROR              Severity = 30
	Severity_CRITICAL           Severity = 40
)

// Enum value maps for Severity.
var (
	Severity_name = map[int32]string{
		0:  "UNDEFINED_SEVERITY",
		10: "INFO",
		20: "WARNING",
		30: "ERROR",
		40: "CRITICAL",
	}
	Severity_value = map[string]int32{
		"UNDEFINED_SEVERITY": 0,
		"INFO":               10,
		"WARNING":            20,
		"ERROR":              30,
		"CRITICAL":           40,
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
	return file_mmesh_protobuf_resources_v1_events_event_proto_enumTypes[3].Descriptor()
}

func (Severity) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_events_event_proto_enumTypes[3]
}

func (x Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Severity.Descriptor instead.
func (Severity) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_event_proto_rawDescGZIP(), []int{3}
}

type ActionType int32

const (
	ActionType_UNDEFINED_ACTION ActionType = 0
	ActionType_TRIGGER          ActionType = 10
	ActionType_ACKNOWLEDGE      ActionType = 20
	ActionType_RESOLVE          ActionType = 30
)

// Enum value maps for ActionType.
var (
	ActionType_name = map[int32]string{
		0:  "UNDEFINED_ACTION",
		10: "TRIGGER",
		20: "ACKNOWLEDGE",
		30: "RESOLVE",
	}
	ActionType_value = map[string]int32{
		"UNDEFINED_ACTION": 0,
		"TRIGGER":          10,
		"ACKNOWLEDGE":      20,
		"RESOLVE":          30,
	}
)

func (x ActionType) Enum() *ActionType {
	p := new(ActionType)
	*p = x
	return p
}

func (x ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_events_event_proto_enumTypes[4].Descriptor()
}

func (ActionType) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_events_event_proto_enumTypes[4]
}

func (x ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionType.Descriptor instead.
func (ActionType) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_event_proto_rawDescGZIP(), []int{4}
}

type Status int32

const (
	Status_UNDEFINED_STATUS Status = 0
	Status_TRIGGERED        Status = 10
	Status_ACKNOWLEDGED     Status = 20
	Status_RESOLVED         Status = 30
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0:  "UNDEFINED_STATUS",
		10: "TRIGGERED",
		20: "ACKNOWLEDGED",
		30: "RESOLVED",
	}
	Status_value = map[string]int32{
		"UNDEFINED_STATUS": 0,
		"TRIGGERED":        10,
		"ACKNOWLEDGED":     20,
		"RESOLVED":         30,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_events_event_proto_enumTypes[5].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_events_event_proto_enumTypes[5]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_event_proto_rawDescGZIP(), []int{5}
}

type EventResult int32

const (
	EventResult_UNKNOWN_RESULT EventResult = 0
	EventResult_SUCCESS        EventResult = 10
	EventResult_FAIL           EventResult = 20
)

// Enum value maps for EventResult.
var (
	EventResult_name = map[int32]string{
		0:  "UNKNOWN_RESULT",
		10: "SUCCESS",
		20: "FAIL",
	}
	EventResult_value = map[string]int32{
		"UNKNOWN_RESULT": 0,
		"SUCCESS":        10,
		"FAIL":           20,
	}
)

func (x EventResult) Enum() *EventResult {
	p := new(EventResult)
	*p = x
	return p
}

func (x EventResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventResult) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_events_event_proto_enumTypes[6].Descriptor()
}

func (EventResult) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_events_event_proto_enumTypes[6]
}

func (x EventResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventResult.Descriptor instead.
func (EventResult) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_event_proto_rawDescGZIP(), []int{6}
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventID string `protobuf:"bytes,1,opt,name=eventID,proto3" json:"eventID,omitempty"`
	// AccountAlert is true for customer-specific alerts
	AccountAlert bool   `protobuf:"varint,11,opt,name=accountAlert,proto3" json:"accountAlert,omitempty"`
	AccountID    string `protobuf:"bytes,12,opt,name=accountID,proto3" json:"accountID,omitempty"`
	// timestamp
	Timestamp int64 `protobuf:"varint,21,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The type of the event
	Type EventType `protobuf:"varint,31,opt,name=type,proto3,enum=events.EventType" json:"type,omitempty"`
	// Source entity of the event or the affected system
	Source *Source `protobuf:"bytes,41,opt,name=source,proto3" json:"source,omitempty"`
	// The class of the event
	Class Class `protobuf:"varint,51,opt,name=class,proto3,enum=events.Class" json:"class,omitempty"`
	// Logical grouping of components of a service,
	// for example 'routing', 'workflow', 'IAM', 'app-stack'.
	Group Group `protobuf:"varint,61,opt,name=group,proto3,enum=events.Group" json:"group,omitempty"`
	// Component of the source machine that is responsible for the event,
	// for example 'function', 'workflowID/workflowName/taskName', 'userID',
	// 'mysql' or 'eth0'.
	Component string `protobuf:"bytes,71,opt,name=component,proto3" json:"component,omitempty"`
	// The perceived severity of the status the event is describing with respect
	// to the affected system. This can be 'critical', 'error', 'warning' or
	// 'info'.
	Severity Severity `protobuf:"varint,81,opt,name=severity,proto3,enum=events.Severity" json:"severity,omitempty"`
	// The type of event. Can be 'trigger', 'acknowledge' or 'resolve'.
	ActionType ActionType `protobuf:"varint,91,opt,name=actionType,proto3,enum=events.ActionType" json:"actionType,omitempty"`
	// A brief text summary of the event, used to generate the summaries/titles
	// of any associated alerts.
	// The maximum permitted length of this property is 1024 characters.
	Summary string `protobuf:"bytes,101,opt,name=summary,proto3" json:"summary,omitempty"`
	// Additional details about the event and affected system
	CustomDetails map[string]string `protobuf:"bytes,111,rep,name=customDetails,proto3" json:"customDetails,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map[string]string
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_events_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_events_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetEventID() string {
	if x != nil {
		return x.EventID
	}
	return ""
}

func (x *Event) GetAccountAlert() bool {
	if x != nil {
		return x.AccountAlert
	}
	return false
}

func (x *Event) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *Event) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Event) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_UNDEFINED_TYPE
}

func (x *Event) GetSource() *Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *Event) GetClass() Class {
	if x != nil {
		return x.Class
	}
	return Class_UNDEFINED_CLASS
}

func (x *Event) GetGroup() Group {
	if x != nil {
		return x.Group
	}
	return Group_UNDEFINED_GROUP
}

func (x *Event) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *Event) GetSeverity() Severity {
	if x != nil {
		return x.Severity
	}
	return Severity_UNDEFINED_SEVERITY
}

func (x *Event) GetActionType() ActionType {
	if x != nil {
		return x.ActionType
	}
	return ActionType_UNDEFINED_ACTION
}

func (x *Event) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Event) GetCustomDetails() map[string]string {
	if x != nil {
		return x.CustomDetails
	}
	return nil
}

var File_mmesh_protobuf_resources_v1_events_event_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_events_event_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x04, 0x0a, 0x05, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x22, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x25, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x29,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x05,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x12, 0x23, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x18, 0x47, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x51, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x32, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x5b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x12, 0x46, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x18, 0x6f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x40, 0x0a, 0x12, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x36, 0x0a, 0x09, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x44, 0x45, 0x46,
	0x49, 0x4e, 0x45, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41,
	0x4c, 0x45, 0x52, 0x54, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45,
	0x10, 0x14, 0x2a, 0xaf, 0x01, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x13, 0x0a, 0x0f,
	0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x5f, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x49, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x0a, 0x12, 0x0c,
	0x0a, 0x08, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x10, 0x14, 0x12, 0x0b, 0x0a, 0x07,
	0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x1e, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x45, 0x43,
	0x55, 0x52, 0x49, 0x54, 0x59, 0x10, 0x28, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4f, 0x4e, 0x46, 0x49,
	0x47, 0x55, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x32, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f,
	0x53, 0x54, 0x10, 0x3c, 0x12, 0x0c, 0x0a, 0x08, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57,
	0x10, 0x46, 0x12, 0x0e, 0x0a, 0x0a, 0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x45, 0x53,
	0x10, 0x50, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x49, 0x54, 0x59, 0x10, 0x64, 0x12, 0x0d, 0x0a, 0x08, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x4c, 0x10, 0xe8, 0x07, 0x2a, 0x81, 0x01, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x13,
	0x0a, 0x0f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55,
	0x50, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x4f, 0x55, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x0a,
	0x12, 0x07, 0x0a, 0x03, 0x49, 0x41, 0x4d, 0x10, 0x14, 0x12, 0x10, 0x0a, 0x0c, 0x48, 0x4f, 0x53,
	0x54, 0x5f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x10, 0x1e, 0x12, 0x11, 0x0a, 0x0d, 0x57,
	0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x10, 0x28, 0x12, 0x0b,
	0x0a, 0x07, 0x57, 0x45, 0x42, 0x48, 0x4f, 0x4f, 0x4b, 0x10, 0x32, 0x12, 0x1b, 0x0a, 0x17, 0x54,
	0x48, 0x49, 0x52, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x47,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x3c, 0x2a, 0x52, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45,
	0x44, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x49, 0x4e, 0x46, 0x4f, 0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e,
	0x47, 0x10, 0x14, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x1e, 0x12, 0x0c,
	0x0a, 0x08, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x28, 0x2a, 0x4d, 0x0a, 0x0a,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x4e,
	0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x54, 0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x10, 0x0a, 0x12, 0x0f, 0x0a,
	0x0b, 0x41, 0x43, 0x4b, 0x4e, 0x4f, 0x57, 0x4c, 0x45, 0x44, 0x47, 0x45, 0x10, 0x14, 0x12, 0x0b,
	0x0a, 0x07, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x10, 0x1e, 0x2a, 0x4d, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e,
	0x45, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x54,
	0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x43,
	0x4b, 0x4e, 0x4f, 0x57, 0x4c, 0x45, 0x44, 0x47, 0x45, 0x44, 0x10, 0x14, 0x12, 0x0c, 0x0a, 0x08,
	0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x44, 0x10, 0x1e, 0x2a, 0x38, 0x0a, 0x0b, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x0a, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41,
	0x49, 0x4c, 0x10, 0x14, 0x42, 0x2a, 0x5a, 0x28, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65,
	0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_events_event_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_events_event_proto_rawDescData = file_mmesh_protobuf_resources_v1_events_event_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_events_event_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_events_event_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_events_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_events_event_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_events_event_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_events_event_proto_enumTypes = make([]protoimpl.EnumInfo, 7)
var file_mmesh_protobuf_resources_v1_events_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mmesh_protobuf_resources_v1_events_event_proto_goTypes = []interface{}{
	(EventType)(0),   // 0: events.EventType
	(Class)(0),       // 1: events.Class
	(Group)(0),       // 2: events.Group
	(Severity)(0),    // 3: events.Severity
	(ActionType)(0),  // 4: events.ActionType
	(Status)(0),      // 5: events.Status
	(EventResult)(0), // 6: events.EventResult
	(*Event)(nil),    // 7: events.Event
	nil,              // 8: events.Event.CustomDetailsEntry
	(*Source)(nil),   // 9: events.Source
}
var file_mmesh_protobuf_resources_v1_events_event_proto_depIdxs = []int32{
	0, // 0: events.Event.type:type_name -> events.EventType
	9, // 1: events.Event.source:type_name -> events.Source
	1, // 2: events.Event.class:type_name -> events.Class
	2, // 3: events.Event.group:type_name -> events.Group
	3, // 4: events.Event.severity:type_name -> events.Severity
	4, // 5: events.Event.actionType:type_name -> events.ActionType
	8, // 6: events.Event.customDetails:type_name -> events.Event.CustomDetailsEntry
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_events_event_proto_init() }
func file_mmesh_protobuf_resources_v1_events_event_proto_init() {
	if File_mmesh_protobuf_resources_v1_events_event_proto != nil {
		return
	}
	file_mmesh_protobuf_resources_v1_events_source_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_events_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_events_event_proto_rawDesc,
			NumEnums:      7,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_events_event_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_events_event_proto_depIdxs,
		EnumInfos:         file_mmesh_protobuf_resources_v1_events_event_proto_enumTypes,
		MessageInfos:      file_mmesh_protobuf_resources_v1_events_event_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_events_event_proto = out.File
	file_mmesh_protobuf_resources_v1_events_event_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_events_event_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_events_event_proto_depIdxs = nil
}
