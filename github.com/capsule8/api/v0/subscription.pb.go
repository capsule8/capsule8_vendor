// Code generated by protoc-gen-go. DO NOT EDIT.
// source: capsule8/api/v0/subscription.proto

package v0

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ContainerEventView int32

const (
	// Default view of a ContainerEvent includes just basic information
	ContainerEventView_BASIC ContainerEventView = 0
	// Full view of a ContainerEvent includes raw Docker and OCI config JSON
	ContainerEventView_FULL ContainerEventView = 1
)

var ContainerEventView_name = map[int32]string{
	0: "BASIC",
	1: "FULL",
}
var ContainerEventView_value = map[string]int32{
	"BASIC": 0,
	"FULL":  1,
}

func (x ContainerEventView) String() string {
	return proto.EnumName(ContainerEventView_name, int32(x))
}
func (ContainerEventView) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

type ThrottleModifier_IntervalType int32

const (
	ThrottleModifier_MILLISECOND ThrottleModifier_IntervalType = 0
	ThrottleModifier_SECOND      ThrottleModifier_IntervalType = 1
	ThrottleModifier_MINUTE      ThrottleModifier_IntervalType = 2
	ThrottleModifier_HOUR        ThrottleModifier_IntervalType = 3
)

var ThrottleModifier_IntervalType_name = map[int32]string{
	0: "MILLISECOND",
	1: "SECOND",
	2: "MINUTE",
	3: "HOUR",
}
var ThrottleModifier_IntervalType_value = map[string]int32{
	"MILLISECOND": 0,
	"SECOND":      1,
	"MINUTE":      2,
	"HOUR":        3,
}

func (x ThrottleModifier_IntervalType) String() string {
	return proto.EnumName(ThrottleModifier_IntervalType_name, int32(x))
}
func (ThrottleModifier_IntervalType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor6, []int{12, 0}
}

//
// SignedSubscription wraps a subscription and signature.
//
type SignedSubscription struct {
	Subscription   *Subscription `protobuf:"bytes,1,opt,name=subscription" json:"subscription,omitempty"`
	SubscriptionId string        `protobuf:"bytes,2,opt,name=subscription_id,json=subscriptionId" json:"subscription_id,omitempty"`
	// (Soon to be) Required, will used to sign the subscription
	Signature string `protobuf:"bytes,3,opt,name=signature" json:"signature,omitempty"`
}

func (m *SignedSubscription) Reset()                    { *m = SignedSubscription{} }
func (m *SignedSubscription) String() string            { return proto.CompactTextString(m) }
func (*SignedSubscription) ProtoMessage()               {}
func (*SignedSubscription) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *SignedSubscription) GetSubscription() *Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

func (m *SignedSubscription) GetSubscriptionId() string {
	if m != nil {
		return m.SubscriptionId
	}
	return ""
}

func (m *SignedSubscription) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

//
// Subscription identifies a subscriber's interest in events. The optional
// Filter field restricts the subscription to the nodes and processes
// specified within. The optional Selector field specifies the types and
// properties of the events desired by the subscriber.
//
type Subscription struct {
	// Return events matching one or more of the specified event
	// filters. If no event filters are specified, then no events
	// will be returned.
	EventFilter *EventFilter `protobuf:"bytes,1,opt,name=event_filter,json=eventFilter" json:"event_filter,omitempty"`
	// If not empty, then only return events from containers matched
	// by one or more of the specified container filters.
	ContainerFilter *ContainerFilter `protobuf:"bytes,2,opt,name=container_filter,json=containerFilter" json:"container_filter,omitempty"`
	// If not empty, then only return events from Sensors matched by
	// one or more of the specified sensor filters.
	SensorFilter *SensorFilter `protobuf:"bytes,3,opt,name=sensor_filter,json=sensorFilter" json:"sensor_filter,omitempty"`
	// If not empty, then only return events that occurred after
	// the specified relative duration subtracted from the current
	// time (recorder time). If the resulting time is in the past, then the
	// subscription will search for historic events before streaming
	// live ones.
	SinceDuration *google_protobuf1.Int64Value `protobuf:"bytes,10,opt,name=since_duration,json=sinceDuration" json:"since_duration,omitempty"`
	// If not empty, then only return events that occurred before
	// the specified relative duration added to `since_duration`.
	// If `since_duration` is not supplied, return events from now and until
	// the specified relative duration is hit.
	ForDuration *google_protobuf1.Int64Value `protobuf:"bytes,11,opt,name=for_duration,json=forDuration" json:"for_duration,omitempty"`
	// If not empty, apply the specified modifier to the subscription.
	Modifier *Modifier `protobuf:"bytes,20,opt,name=modifier" json:"modifier,omitempty"`
}

func (m *Subscription) Reset()                    { *m = Subscription{} }
func (m *Subscription) String() string            { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()               {}
func (*Subscription) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *Subscription) GetEventFilter() *EventFilter {
	if m != nil {
		return m.EventFilter
	}
	return nil
}

func (m *Subscription) GetContainerFilter() *ContainerFilter {
	if m != nil {
		return m.ContainerFilter
	}
	return nil
}

func (m *Subscription) GetSensorFilter() *SensorFilter {
	if m != nil {
		return m.SensorFilter
	}
	return nil
}

func (m *Subscription) GetSinceDuration() *google_protobuf1.Int64Value {
	if m != nil {
		return m.SinceDuration
	}
	return nil
}

func (m *Subscription) GetForDuration() *google_protobuf1.Int64Value {
	if m != nil {
		return m.ForDuration
	}
	return nil
}

func (m *Subscription) GetModifier() *Modifier {
	if m != nil {
		return m.Modifier
	}
	return nil
}

type SensorFilter struct {
}

func (m *SensorFilter) Reset()                    { *m = SensorFilter{} }
func (m *SensorFilter) String() string            { return proto.CompactTextString(m) }
func (*SensorFilter) ProtoMessage()               {}
func (*SensorFilter) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

type ContainerFilter struct {
	Ids      []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
	Names    []string `protobuf:"bytes,2,rep,name=names" json:"names,omitempty"`
	ImageIds []string `protobuf:"bytes,3,rep,name=image_ids,json=imageIds" json:"image_ids,omitempty"`
	// Container image name (shell-style globs are supported)
	ImageNames []string `protobuf:"bytes,4,rep,name=image_names,json=imageNames" json:"image_names,omitempty"`
}

func (m *ContainerFilter) Reset()                    { *m = ContainerFilter{} }
func (m *ContainerFilter) String() string            { return proto.CompactTextString(m) }
func (*ContainerFilter) ProtoMessage()               {}
func (*ContainerFilter) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *ContainerFilter) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *ContainerFilter) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *ContainerFilter) GetImageIds() []string {
	if m != nil {
		return m.ImageIds
	}
	return nil
}

func (m *ContainerFilter) GetImageNames() []string {
	if m != nil {
		return m.ImageNames
	}
	return nil
}

type EventFilter struct {
	//
	// Kernel-level events
	//
	SyscallEvents []*SyscallEventFilter `protobuf:"bytes,1,rep,name=syscall_events,json=syscallEvents" json:"syscall_events,omitempty"`
	ProcessEvents []*ProcessEventFilter `protobuf:"bytes,2,rep,name=process_events,json=processEvents" json:"process_events,omitempty"`
	FileEvents    []*FileEventFilter    `protobuf:"bytes,3,rep,name=file_events,json=fileEvents" json:"file_events,omitempty"`
	//
	// Operating System-level events (containers, etc)
	//
	ContainerEvents []*ContainerEventFilter `protobuf:"bytes,10,rep,name=container_events,json=containerEvents" json:"container_events,omitempty"`
	//
	// Debugging events (>= 100)
	//
	ChargenEvents []*ChargenEventFilter `protobuf:"bytes,100,rep,name=chargen_events,json=chargenEvents" json:"chargen_events,omitempty"`
	TickerEvents  []*TickerEventFilter  `protobuf:"bytes,101,rep,name=ticker_events,json=tickerEvents" json:"ticker_events,omitempty"`
}

func (m *EventFilter) Reset()                    { *m = EventFilter{} }
func (m *EventFilter) String() string            { return proto.CompactTextString(m) }
func (*EventFilter) ProtoMessage()               {}
func (*EventFilter) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

func (m *EventFilter) GetSyscallEvents() []*SyscallEventFilter {
	if m != nil {
		return m.SyscallEvents
	}
	return nil
}

func (m *EventFilter) GetProcessEvents() []*ProcessEventFilter {
	if m != nil {
		return m.ProcessEvents
	}
	return nil
}

func (m *EventFilter) GetFileEvents() []*FileEventFilter {
	if m != nil {
		return m.FileEvents
	}
	return nil
}

func (m *EventFilter) GetContainerEvents() []*ContainerEventFilter {
	if m != nil {
		return m.ContainerEvents
	}
	return nil
}

func (m *EventFilter) GetChargenEvents() []*ChargenEventFilter {
	if m != nil {
		return m.ChargenEvents
	}
	return nil
}

func (m *EventFilter) GetTickerEvents() []*TickerEventFilter {
	if m != nil {
		return m.TickerEvents
	}
	return nil
}

type SyscallEventFilter struct {
	// Type of system call event (entry w/ args or exit w/ ret)
	Type SyscallEventType `protobuf:"varint,1,opt,name=type,enum=capsule8.api.v0.SyscallEventType" json:"type,omitempty"`
	// System call number (arch/x86/entry/syscalls/syscall_64.tbl)
	Id   *google_protobuf1.Int64Value  `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Arg0 *google_protobuf1.UInt64Value `protobuf:"bytes,10,opt,name=arg0" json:"arg0,omitempty"`
	Arg1 *google_protobuf1.UInt64Value `protobuf:"bytes,11,opt,name=arg1" json:"arg1,omitempty"`
	Arg2 *google_protobuf1.UInt64Value `protobuf:"bytes,12,opt,name=arg2" json:"arg2,omitempty"`
	Arg3 *google_protobuf1.UInt64Value `protobuf:"bytes,13,opt,name=arg3" json:"arg3,omitempty"`
	Arg4 *google_protobuf1.UInt64Value `protobuf:"bytes,14,opt,name=arg4" json:"arg4,omitempty"`
	Arg5 *google_protobuf1.UInt64Value `protobuf:"bytes,15,opt,name=arg5" json:"arg5,omitempty"`
	Ret  *google_protobuf1.Int64Value  `protobuf:"bytes,20,opt,name=ret" json:"ret,omitempty"`
}

func (m *SyscallEventFilter) Reset()                    { *m = SyscallEventFilter{} }
func (m *SyscallEventFilter) String() string            { return proto.CompactTextString(m) }
func (*SyscallEventFilter) ProtoMessage()               {}
func (*SyscallEventFilter) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{5} }

func (m *SyscallEventFilter) GetType() SyscallEventType {
	if m != nil {
		return m.Type
	}
	return SyscallEventType_SYSCALL_EVENT_TYPE_UNKNOWN
}

func (m *SyscallEventFilter) GetId() *google_protobuf1.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SyscallEventFilter) GetArg0() *google_protobuf1.UInt64Value {
	if m != nil {
		return m.Arg0
	}
	return nil
}

func (m *SyscallEventFilter) GetArg1() *google_protobuf1.UInt64Value {
	if m != nil {
		return m.Arg1
	}
	return nil
}

func (m *SyscallEventFilter) GetArg2() *google_protobuf1.UInt64Value {
	if m != nil {
		return m.Arg2
	}
	return nil
}

func (m *SyscallEventFilter) GetArg3() *google_protobuf1.UInt64Value {
	if m != nil {
		return m.Arg3
	}
	return nil
}

func (m *SyscallEventFilter) GetArg4() *google_protobuf1.UInt64Value {
	if m != nil {
		return m.Arg4
	}
	return nil
}

func (m *SyscallEventFilter) GetArg5() *google_protobuf1.UInt64Value {
	if m != nil {
		return m.Arg5
	}
	return nil
}

func (m *SyscallEventFilter) GetRet() *google_protobuf1.Int64Value {
	if m != nil {
		return m.Ret
	}
	return nil
}

type ProcessEventFilter struct {
	// Required; the process event type to match
	Type ProcessEventType `protobuf:"varint,1,opt,name=type,enum=capsule8.api.v0.ProcessEventType" json:"type,omitempty"`
}

func (m *ProcessEventFilter) Reset()                    { *m = ProcessEventFilter{} }
func (m *ProcessEventFilter) String() string            { return proto.CompactTextString(m) }
func (*ProcessEventFilter) ProtoMessage()               {}
func (*ProcessEventFilter) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{6} }

func (m *ProcessEventFilter) GetType() ProcessEventType {
	if m != nil {
		return m.Type
	}
	return ProcessEventType_PROCESS_EVENT_TYPE_UNKNOWN
}

type FileEventFilter struct {
	// Required
	Type FileEventType `protobuf:"varint,1,opt,name=type,enum=capsule8.api.v0.FileEventType" json:"type,omitempty"`
	// Optional filename exact match
	Filename *google_protobuf1.StringValue `protobuf:"bytes,10,opt,name=filename" json:"filename,omitempty"`
	// Optional filename pattern
	FilenamePattern *google_protobuf1.StringValue `protobuf:"bytes,11,opt,name=filename_pattern,json=filenamePattern" json:"filename_pattern,omitempty"`
	// Optional open(2) flags mask value
	OpenFlagsMask *google_protobuf1.Int32Value `protobuf:"bytes,12,opt,name=open_flags_mask,json=openFlagsMask" json:"open_flags_mask,omitempty"`
	// Optional open(2)/creat(2) mode mask value
	CreateModeMask *google_protobuf1.Int32Value `protobuf:"bytes,13,opt,name=create_mode_mask,json=createModeMask" json:"create_mode_mask,omitempty"`
}

func (m *FileEventFilter) Reset()                    { *m = FileEventFilter{} }
func (m *FileEventFilter) String() string            { return proto.CompactTextString(m) }
func (*FileEventFilter) ProtoMessage()               {}
func (*FileEventFilter) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{7} }

func (m *FileEventFilter) GetType() FileEventType {
	if m != nil {
		return m.Type
	}
	return FileEventType_FILE_EVENT_TYPE_UNKNOWN
}

func (m *FileEventFilter) GetFilename() *google_protobuf1.StringValue {
	if m != nil {
		return m.Filename
	}
	return nil
}

func (m *FileEventFilter) GetFilenamePattern() *google_protobuf1.StringValue {
	if m != nil {
		return m.FilenamePattern
	}
	return nil
}

func (m *FileEventFilter) GetOpenFlagsMask() *google_protobuf1.Int32Value {
	if m != nil {
		return m.OpenFlagsMask
	}
	return nil
}

func (m *FileEventFilter) GetCreateModeMask() *google_protobuf1.Int32Value {
	if m != nil {
		return m.CreateModeMask
	}
	return nil
}

type ContainerEventFilter struct {
	// Required, specify the particular type of event type to match
	Type ContainerEventType `protobuf:"varint,1,opt,name=type,enum=capsule8.api.v0.ContainerEventType" json:"type,omitempty"`
	// Optional, specifies how much detail to include in container events
	View ContainerEventView `protobuf:"varint,2,opt,name=view,enum=capsule8.api.v0.ContainerEventView" json:"view,omitempty"`
}

func (m *ContainerEventFilter) Reset()                    { *m = ContainerEventFilter{} }
func (m *ContainerEventFilter) String() string            { return proto.CompactTextString(m) }
func (*ContainerEventFilter) ProtoMessage()               {}
func (*ContainerEventFilter) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{8} }

func (m *ContainerEventFilter) GetType() ContainerEventType {
	if m != nil {
		return m.Type
	}
	return ContainerEventType_CONTAINER_EVENT_TYPE_UNKNOWN
}

func (m *ContainerEventFilter) GetView() ContainerEventView {
	if m != nil {
		return m.View
	}
	return ContainerEventView_BASIC
}

type ChargenEventFilter struct {
	// Length of character sequence strings to generate
	Length uint64 `protobuf:"varint,1,opt,name=length" json:"length,omitempty"`
}

func (m *ChargenEventFilter) Reset()                    { *m = ChargenEventFilter{} }
func (m *ChargenEventFilter) String() string            { return proto.CompactTextString(m) }
func (*ChargenEventFilter) ProtoMessage()               {}
func (*ChargenEventFilter) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{9} }

func (m *ChargenEventFilter) GetLength() uint64 {
	if m != nil {
		return m.Length
	}
	return 0
}

type TickerEventFilter struct {
	// The interval at which ticker events are generated
	Interval int64 `protobuf:"varint,1,opt,name=interval" json:"interval,omitempty"`
}

func (m *TickerEventFilter) Reset()                    { *m = TickerEventFilter{} }
func (m *TickerEventFilter) String() string            { return proto.CompactTextString(m) }
func (*TickerEventFilter) ProtoMessage()               {}
func (*TickerEventFilter) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{10} }

func (m *TickerEventFilter) GetInterval() int64 {
	if m != nil {
		return m.Interval
	}
	return 0
}

//
// Modifier specifies which stream modifiers to apply if any. For a given
// stream, a modifier can apply a throttle or limit etc. Modifiers can be
// used together.
//
type Modifier struct {
	Throttle *ThrottleModifier `protobuf:"bytes,1,opt,name=throttle" json:"throttle,omitempty"`
	Limit    *LimitModifier    `protobuf:"bytes,2,opt,name=limit" json:"limit,omitempty"`
}

func (m *Modifier) Reset()                    { *m = Modifier{} }
func (m *Modifier) String() string            { return proto.CompactTextString(m) }
func (*Modifier) ProtoMessage()               {}
func (*Modifier) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{11} }

func (m *Modifier) GetThrottle() *ThrottleModifier {
	if m != nil {
		return m.Throttle
	}
	return nil
}

func (m *Modifier) GetLimit() *LimitModifier {
	if m != nil {
		return m.Limit
	}
	return nil
}

type ThrottleModifier struct {
	Interval     int64                         `protobuf:"varint,1,opt,name=interval" json:"interval,omitempty"`
	IntervalType ThrottleModifier_IntervalType `protobuf:"varint,2,opt,name=interval_type,json=intervalType,enum=capsule8.api.v0.ThrottleModifier_IntervalType" json:"interval_type,omitempty"`
}

func (m *ThrottleModifier) Reset()                    { *m = ThrottleModifier{} }
func (m *ThrottleModifier) String() string            { return proto.CompactTextString(m) }
func (*ThrottleModifier) ProtoMessage()               {}
func (*ThrottleModifier) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{12} }

func (m *ThrottleModifier) GetInterval() int64 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *ThrottleModifier) GetIntervalType() ThrottleModifier_IntervalType {
	if m != nil {
		return m.IntervalType
	}
	return ThrottleModifier_MILLISECOND
}

type LimitModifier struct {
	// Limit the number of events
	Limit int64 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
}

func (m *LimitModifier) Reset()                    { *m = LimitModifier{} }
func (m *LimitModifier) String() string            { return proto.CompactTextString(m) }
func (*LimitModifier) ProtoMessage()               {}
func (*LimitModifier) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{13} }

func (m *LimitModifier) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func init() {
	proto.RegisterType((*SignedSubscription)(nil), "capsule8.api.v0.SignedSubscription")
	proto.RegisterType((*Subscription)(nil), "capsule8.api.v0.Subscription")
	proto.RegisterType((*SensorFilter)(nil), "capsule8.api.v0.SensorFilter")
	proto.RegisterType((*ContainerFilter)(nil), "capsule8.api.v0.ContainerFilter")
	proto.RegisterType((*EventFilter)(nil), "capsule8.api.v0.EventFilter")
	proto.RegisterType((*SyscallEventFilter)(nil), "capsule8.api.v0.SyscallEventFilter")
	proto.RegisterType((*ProcessEventFilter)(nil), "capsule8.api.v0.ProcessEventFilter")
	proto.RegisterType((*FileEventFilter)(nil), "capsule8.api.v0.FileEventFilter")
	proto.RegisterType((*ContainerEventFilter)(nil), "capsule8.api.v0.ContainerEventFilter")
	proto.RegisterType((*ChargenEventFilter)(nil), "capsule8.api.v0.ChargenEventFilter")
	proto.RegisterType((*TickerEventFilter)(nil), "capsule8.api.v0.TickerEventFilter")
	proto.RegisterType((*Modifier)(nil), "capsule8.api.v0.Modifier")
	proto.RegisterType((*ThrottleModifier)(nil), "capsule8.api.v0.ThrottleModifier")
	proto.RegisterType((*LimitModifier)(nil), "capsule8.api.v0.LimitModifier")
	proto.RegisterEnum("capsule8.api.v0.ContainerEventView", ContainerEventView_name, ContainerEventView_value)
	proto.RegisterEnum("capsule8.api.v0.ThrottleModifier_IntervalType", ThrottleModifier_IntervalType_name, ThrottleModifier_IntervalType_value)
}

func init() { proto.RegisterFile("capsule8/api/v0/subscription.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 1014 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0x5d, 0x4f, 0xe3, 0x46,
	0x17, 0xde, 0xc4, 0x01, 0x85, 0x13, 0x27, 0xf1, 0x3b, 0x42, 0xaf, 0x5c, 0xa0, 0x5b, 0xea, 0xd5,
	0xaa, 0xf4, 0x2b, 0x61, 0x03, 0x74, 0xf7, 0xa6, 0xad, 0x80, 0x85, 0x6d, 0xba, 0x09, 0x8b, 0x1c,
	0xd8, 0x5b, 0xcb, 0xd8, 0x13, 0x33, 0xc2, 0xb1, 0xad, 0x99, 0x49, 0x10, 0x57, 0xbd, 0xed, 0x7f,
	0xa8, 0xd4, 0x9f, 0xd3, 0xdb, 0x5e, 0xb4, 0x3f, 0xa8, 0xf2, 0x78, 0x9c, 0x4c, 0xe2, 0x0d, 0xe1,
	0x6e, 0xce, 0x99, 0xe7, 0x79, 0x32, 0xe7, 0xd3, 0x01, 0xcb, 0x73, 0x13, 0x36, 0x0e, 0xf1, 0x9b,
	0xb6, 0x9b, 0x90, 0xf6, 0x64, 0xbf, 0xcd, 0xc6, 0x37, 0xcc, 0xa3, 0x24, 0xe1, 0x24, 0x8e, 0x5a,
	0x09, 0x8d, 0x79, 0x8c, 0x9a, 0x39, 0xa6, 0xe5, 0x26, 0xa4, 0x35, 0xd9, 0xdf, 0xda, 0x5e, 0x24,
	0xe1, 0x09, 0x8e, 0x78, 0x86, 0xde, 0x7a, 0x1e, 0xc4, 0x71, 0x10, 0xe2, 0xb6, 0xb0, 0x6e, 0xc6,
	0xc3, 0xf6, 0x3d, 0x75, 0x93, 0x04, 0x53, 0x96, 0xdd, 0x5b, 0x7f, 0x96, 0x00, 0x0d, 0x48, 0x10,
	0x61, 0x7f, 0xa0, 0xfc, 0x14, 0x3a, 0x06, 0x5d, 0xfd, 0x69, 0xb3, 0xb4, 0x5b, 0xda, 0xab, 0x75,
	0x3e, 0x6f, 0x2d, 0xfc, 0x76, 0x4b, 0x25, 0xd9, 0x73, 0x14, 0xf4, 0x15, 0x34, 0x55, 0xdb, 0x21,
	0xbe, 0x59, 0xde, 0x2d, 0xed, 0x6d, 0xd8, 0x0d, 0xd5, 0xdd, 0xf5, 0xd1, 0x0e, 0x6c, 0x30, 0x12,
	0x44, 0x2e, 0x1f, 0x53, 0x6c, 0x6a, 0x02, 0x32, 0x73, 0x58, 0x7f, 0x68, 0xa0, 0xcf, 0x3d, 0xed,
	0x67, 0xd0, 0x45, 0x80, 0xce, 0x90, 0x84, 0x1c, 0x53, 0xf9, 0xb4, 0x9d, 0xc2, 0xd3, 0xce, 0x52,
	0xd0, 0xb9, 0xc0, 0xd8, 0x35, 0x3c, 0x33, 0xd0, 0x7b, 0x30, 0xbc, 0x38, 0xe2, 0x2e, 0x89, 0x30,
	0xcd, 0x45, 0xca, 0x42, 0x64, 0xb7, 0x20, 0x72, 0x9a, 0x03, 0xa5, 0x50, 0xd3, 0x9b, 0x77, 0xa0,
	0x13, 0xa8, 0x33, 0x1c, 0xb1, 0x78, 0xaa, 0xa4, 0x2d, 0xcb, 0x94, 0x40, 0x49, 0x19, 0x9d, 0x29,
	0x16, 0x3a, 0x81, 0x06, 0x23, 0x91, 0x87, 0x1d, 0x7f, 0x4c, 0x5d, 0x91, 0x6e, 0x10, 0x22, 0xdb,
	0xad, 0xac, 0x78, 0xad, 0xbc, 0x78, 0xad, 0x6e, 0xc4, 0x7f, 0x38, 0xfc, 0xe8, 0x86, 0x63, 0x6c,
	0xd7, 0x05, 0xe5, 0xad, 0x64, 0xa0, 0x9f, 0x40, 0x1f, 0xc6, 0x74, 0xa6, 0x50, 0x5b, 0xad, 0x50,
	0x1b, 0xc6, 0x74, 0xca, 0x3f, 0x82, 0xea, 0x28, 0xf6, 0xc9, 0x90, 0x60, 0x6a, 0x6e, 0x0a, 0xee,
	0x67, 0x85, 0x10, 0xfa, 0x12, 0x60, 0x4f, 0xa1, 0x56, 0x03, 0x74, 0x35, 0x30, 0xeb, 0x1e, 0x9a,
	0x0b, 0x29, 0x43, 0x06, 0x68, 0xc4, 0x67, 0x66, 0x69, 0x57, 0xdb, 0xdb, 0xb0, 0xd3, 0x23, 0xda,
	0x84, 0xb5, 0xc8, 0x1d, 0x61, 0x66, 0x96, 0x85, 0x2f, 0x33, 0xd0, 0x36, 0x6c, 0x90, 0x91, 0x1b,
	0x60, 0x27, 0x45, 0x6b, 0xe2, 0xa6, 0x2a, 0x1c, 0x5d, 0x9f, 0xa1, 0x2f, 0xa0, 0x96, 0x5d, 0x66,
	0xc4, 0x8a, 0xb8, 0x06, 0xe1, 0xba, 0x48, 0x3d, 0xd6, 0xdf, 0x1a, 0xd4, 0x94, 0x8a, 0xa3, 0x5f,
	0xa1, 0xc1, 0x1e, 0x98, 0xe7, 0x86, 0xa1, 0x23, 0x6a, 0x9f, 0x3d, 0xa0, 0xd6, 0x79, 0x51, 0x2c,
	0x4c, 0x06, 0x53, 0xdb, 0xa5, 0xce, 0x14, 0x1f, 0x4b, 0xb5, 0x12, 0x1a, 0x7b, 0x98, 0xb1, 0x5c,
	0xab, 0xbc, 0x44, 0xeb, 0x32, 0x83, 0xcd, 0x69, 0x25, 0x8a, 0x8f, 0xa1, 0x63, 0xa8, 0x0d, 0x49,
	0x88, 0x73, 0x21, 0x4d, 0x08, 0x15, 0xfb, 0xee, 0x9c, 0x84, 0x58, 0x55, 0x81, 0x61, 0xee, 0x60,
	0xe8, 0x52, 0xed, 0x5f, 0xa9, 0x03, 0x42, 0xe7, 0xe5, 0xf2, 0xfe, 0x55, 0xc5, 0x66, 0x4d, 0x3c,
	0x0b, 0xd0, 0xbb, 0x75, 0x69, 0x80, 0xa3, 0x5c, 0xcf, 0x5f, 0x12, 0xe0, 0x69, 0x06, 0x9b, 0x0b,
	0xd0, 0x53, 0x7c, 0x0c, 0xbd, 0x83, 0x3a, 0x27, 0xde, 0xdd, 0xec, 0x69, 0x58, 0x48, 0x59, 0x05,
	0xa9, 0x2b, 0x81, 0x52, 0x95, 0x74, 0x3e, 0x73, 0x31, 0xeb, 0x1f, 0x0d, 0x50, 0xb1, 0x36, 0xe8,
	0x08, 0x2a, 0xfc, 0x21, 0xc1, 0x62, 0xec, 0x1b, 0x9d, 0x2f, 0x1f, 0x2d, 0xe7, 0xd5, 0x43, 0x82,
	0x6d, 0x01, 0x47, 0xdf, 0x42, 0x59, 0x2e, 0xa0, 0x15, 0x53, 0x51, 0x26, 0x3e, 0xda, 0x87, 0x8a,
	0x4b, 0x83, 0x7d, 0x39, 0x86, 0x3b, 0x05, 0xf8, 0xb5, 0x82, 0x17, 0x48, 0xc9, 0x78, 0x25, 0xc7,
	0x6e, 0x35, 0xe3, 0x95, 0x64, 0x74, 0x4c, 0xfd, 0x89, 0x8c, 0x8e, 0x64, 0x1c, 0x98, 0xf5, 0x27,
	0x32, 0x0e, 0x24, 0xe3, 0xd0, 0x6c, 0x3c, 0x91, 0x71, 0x28, 0x19, 0x47, 0x66, 0xf3, 0x89, 0x8c,
	0x23, 0xf4, 0x3d, 0x68, 0x14, 0x73, 0xb9, 0x33, 0x1e, 0xcd, 0x6c, 0x8a, 0xb3, 0xde, 0x03, 0x2a,
	0x0e, 0xc9, 0xca, 0xa2, 0xaa, 0x94, 0x59, 0x51, 0xad, 0x7f, 0xcb, 0xd0, 0x5c, 0x98, 0x14, 0xd4,
	0x99, 0x93, 0x7a, 0xbe, 0x7c, 0xb2, 0x94, 0xe6, 0x78, 0x03, 0xd5, 0x74, 0xbe, 0xd2, 0xdd, 0xb2,
	0xb4, 0xe6, 0x03, 0x4e, 0x49, 0x14, 0x64, 0x91, 0x4c, 0xd1, 0xe8, 0x1d, 0x18, 0xf9, 0xd9, 0x49,
	0x5c, 0xce, 0x31, 0x8d, 0x96, 0xf6, 0x80, 0xaa, 0xd0, 0xcc, 0x59, 0x97, 0x19, 0x09, 0x9d, 0x42,
	0x33, 0x4e, 0x70, 0xe4, 0x0c, 0x43, 0x37, 0x60, 0xce, 0xc8, 0x65, 0x77, 0xb2, 0x33, 0x3e, 0x99,
	0xd2, 0x83, 0x8e, 0xfc, 0x08, 0xa4, 0x9c, 0xf3, 0x94, 0xd2, 0x77, 0xd9, 0x1d, 0x3a, 0x03, 0xc3,
	0xa3, 0xd8, 0xe5, 0xd8, 0x19, 0xc5, 0x3e, 0xce, 0x54, 0xea, 0xab, 0x55, 0x1a, 0x19, 0xa9, 0x1f,
	0xfb, 0x38, 0x95, 0xb1, 0x7e, 0x2f, 0xc1, 0xe6, 0xa7, 0x16, 0x07, 0x7a, 0x3d, 0x97, 0xdb, 0x17,
	0x2b, 0xb6, 0x8d, 0x92, 0xe0, 0xd7, 0x50, 0x99, 0x10, 0x7c, 0x2f, 0xe6, 0x6f, 0x35, 0xf1, 0x23,
	0xc1, 0xf7, 0xb6, 0x20, 0x58, 0xdf, 0x01, 0x2a, 0xae, 0x1c, 0xf4, 0x7f, 0x58, 0x0f, 0x71, 0x14,
	0xf0, 0x5b, 0xf1, 0x92, 0x8a, 0x2d, 0x2d, 0xab, 0x0d, 0xff, 0x2b, 0x6c, 0x15, 0xb4, 0x05, 0x55,
	0x12, 0x71, 0x4c, 0x27, 0x6e, 0x28, 0xe0, 0x9a, 0x3d, 0xb5, 0xad, 0xdf, 0xa0, 0x9a, 0x7f, 0xd4,
	0xd0, 0x8f, 0x50, 0xe5, 0xb7, 0x34, 0xe6, 0x3c, 0xc4, 0xf2, 0x3f, 0x45, 0xb1, 0x0f, 0xaf, 0x24,
	0x60, 0xf6, 0x25, 0xcc, 0x29, 0xe8, 0x10, 0xd6, 0x42, 0x32, 0x22, 0x5c, 0xee, 0x98, 0x62, 0xe3,
	0xf5, 0xd2, 0xdb, 0x29, 0x31, 0x03, 0x5b, 0x7f, 0x95, 0xc0, 0x58, 0x14, 0x7d, 0xec, 0xc5, 0x68,
	0x00, 0xf5, 0xfc, 0xec, 0x88, 0x5a, 0x64, 0x29, 0x6d, 0xad, 0x7c, 0x6a, 0x5a, 0x70, 0x41, 0x13,
	0x65, 0xd1, 0x89, 0x62, 0x59, 0xc7, 0xa0, 0xab, 0xb7, 0xa8, 0x09, 0xb5, 0x7e, 0xb7, 0xd7, 0xeb,
	0x0e, 0xce, 0x4e, 0x3f, 0x5c, 0xbc, 0x35, 0x9e, 0x21, 0x80, 0x75, 0x79, 0x2e, 0xa5, 0xe7, 0x7e,
	0xf7, 0xe2, 0xfa, 0xea, 0xcc, 0x28, 0xa3, 0x2a, 0x54, 0x7e, 0xf9, 0x70, 0x6d, 0x1b, 0x9a, 0xf5,
	0x12, 0xea, 0x73, 0x01, 0xa6, 0x1f, 0xf9, 0x2c, 0x1f, 0x59, 0x04, 0x99, 0xf1, 0xcd, 0xd7, 0x80,
	0x8a, 0xb5, 0x46, 0x1b, 0xb0, 0x76, 0x72, 0x3c, 0xe8, 0x9e, 0x1a, 0xcf, 0x52, 0xc5, 0xf3, 0xeb,
	0x5e, 0xcf, 0x28, 0xdd, 0xac, 0x8b, 0x56, 0x3d, 0xf8, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x76, 0xd7,
	0x5b, 0x19, 0x14, 0x0b, 0x00, 0x00,
}
