// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: ThreadDump.proto

package grpc

import (
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/wrappers"
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

type PThreadState int32

const (
	PThreadState_THREAD_STATE_NEW           PThreadState = 0
	PThreadState_THREAD_STATE_RUNNABLE      PThreadState = 1
	PThreadState_THREAD_STATE_BLOCKED       PThreadState = 2
	PThreadState_THREAD_STATE_WAITING       PThreadState = 3
	PThreadState_THREAD_STATE_TIMED_WAITING PThreadState = 4
	PThreadState_THREAD_STATE_TERMINATED    PThreadState = 5
	PThreadState_THREAD_STATE_UNKNOWN       PThreadState = 6
)

// Enum value maps for PThreadState.
var (
	PThreadState_name = map[int32]string{
		0: "THREAD_STATE_NEW",
		1: "THREAD_STATE_RUNNABLE",
		2: "THREAD_STATE_BLOCKED",
		3: "THREAD_STATE_WAITING",
		4: "THREAD_STATE_TIMED_WAITING",
		5: "THREAD_STATE_TERMINATED",
		6: "THREAD_STATE_UNKNOWN",
	}
	PThreadState_value = map[string]int32{
		"THREAD_STATE_NEW":           0,
		"THREAD_STATE_RUNNABLE":      1,
		"THREAD_STATE_BLOCKED":       2,
		"THREAD_STATE_WAITING":       3,
		"THREAD_STATE_TIMED_WAITING": 4,
		"THREAD_STATE_TERMINATED":    5,
		"THREAD_STATE_UNKNOWN":       6,
	}
)

func (x PThreadState) Enum() *PThreadState {
	p := new(PThreadState)
	*p = x
	return p
}

func (x PThreadState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PThreadState) Descriptor() protoreflect.EnumDescriptor {
	return file_ThreadDump_proto_enumTypes[0].Descriptor()
}

func (PThreadState) Type() protoreflect.EnumType {
	return &file_ThreadDump_proto_enumTypes[0]
}

func (x PThreadState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PThreadState.Descriptor instead.
func (PThreadState) EnumDescriptor() ([]byte, []int) {
	return file_ThreadDump_proto_rawDescGZIP(), []int{0}
}

type PThreadDumpType int32

const (
	PThreadDumpType_TARGET  PThreadDumpType = 0
	PThreadDumpType_PENDING PThreadDumpType = 1
)

// Enum value maps for PThreadDumpType.
var (
	PThreadDumpType_name = map[int32]string{
		0: "TARGET",
		1: "PENDING",
	}
	PThreadDumpType_value = map[string]int32{
		"TARGET":  0,
		"PENDING": 1,
	}
)

func (x PThreadDumpType) Enum() *PThreadDumpType {
	p := new(PThreadDumpType)
	*p = x
	return p
}

func (x PThreadDumpType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PThreadDumpType) Descriptor() protoreflect.EnumDescriptor {
	return file_ThreadDump_proto_enumTypes[1].Descriptor()
}

func (PThreadDumpType) Type() protoreflect.EnumType {
	return &file_ThreadDump_proto_enumTypes[1]
}

func (x PThreadDumpType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PThreadDumpType.Descriptor instead.
func (PThreadDumpType) EnumDescriptor() ([]byte, []int) {
	return file_ThreadDump_proto_rawDescGZIP(), []int{1}
}

type PMonitorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StackDepth int32  `protobuf:"varint,1,opt,name=stackDepth,proto3" json:"stackDepth,omitempty"`
	StackFrame string `protobuf:"bytes,2,opt,name=stackFrame,proto3" json:"stackFrame,omitempty"`
}

func (x *PMonitorInfo) Reset() {
	*x = PMonitorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ThreadDump_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PMonitorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PMonitorInfo) ProtoMessage() {}

func (x *PMonitorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ThreadDump_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PMonitorInfo.ProtoReflect.Descriptor instead.
func (*PMonitorInfo) Descriptor() ([]byte, []int) {
	return file_ThreadDump_proto_rawDescGZIP(), []int{0}
}

func (x *PMonitorInfo) GetStackDepth() int32 {
	if x != nil {
		return x.StackDepth
	}
	return 0
}

func (x *PMonitorInfo) GetStackFrame() string {
	if x != nil {
		return x.StackFrame
	}
	return ""
}

type PThreadDump struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThreadName         string          `protobuf:"bytes,1,opt,name=threadName,proto3" json:"threadName,omitempty"`
	ThreadId           int64           `protobuf:"varint,2,opt,name=threadId,proto3" json:"threadId,omitempty"`
	BlockedTime        int64           `protobuf:"varint,3,opt,name=blockedTime,proto3" json:"blockedTime,omitempty"`
	BlockedCount       int64           `protobuf:"varint,4,opt,name=blockedCount,proto3" json:"blockedCount,omitempty"`
	WaitedTime         int64           `protobuf:"varint,5,opt,name=waitedTime,proto3" json:"waitedTime,omitempty"`
	WaitedCount        int64           `protobuf:"varint,6,opt,name=waitedCount,proto3" json:"waitedCount,omitempty"`
	LockName           string          `protobuf:"bytes,7,opt,name=lockName,proto3" json:"lockName,omitempty"`
	LockOwnerId        int64           `protobuf:"varint,8,opt,name=lockOwnerId,proto3" json:"lockOwnerId,omitempty"`
	LockOwnerName      string          `protobuf:"bytes,9,opt,name=lockOwnerName,proto3" json:"lockOwnerName,omitempty"`
	InNative           bool            `protobuf:"varint,10,opt,name=inNative,proto3" json:"inNative,omitempty"`
	Suspended          bool            `protobuf:"varint,11,opt,name=suspended,proto3" json:"suspended,omitempty"`
	ThreadState        PThreadState    `protobuf:"varint,12,opt,name=threadState,proto3,enum=v1.PThreadState" json:"threadState,omitempty"`
	StackTrace         []string        `protobuf:"bytes,13,rep,name=stackTrace,proto3" json:"stackTrace,omitempty"`
	LockedMonitor      []*PMonitorInfo `protobuf:"bytes,14,rep,name=lockedMonitor,proto3" json:"lockedMonitor,omitempty"`
	LockedSynchronizer []string        `protobuf:"bytes,15,rep,name=lockedSynchronizer,proto3" json:"lockedSynchronizer,omitempty"`
}

func (x *PThreadDump) Reset() {
	*x = PThreadDump{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ThreadDump_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PThreadDump) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PThreadDump) ProtoMessage() {}

func (x *PThreadDump) ProtoReflect() protoreflect.Message {
	mi := &file_ThreadDump_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PThreadDump.ProtoReflect.Descriptor instead.
func (*PThreadDump) Descriptor() ([]byte, []int) {
	return file_ThreadDump_proto_rawDescGZIP(), []int{1}
}

func (x *PThreadDump) GetThreadName() string {
	if x != nil {
		return x.ThreadName
	}
	return ""
}

func (x *PThreadDump) GetThreadId() int64 {
	if x != nil {
		return x.ThreadId
	}
	return 0
}

func (x *PThreadDump) GetBlockedTime() int64 {
	if x != nil {
		return x.BlockedTime
	}
	return 0
}

func (x *PThreadDump) GetBlockedCount() int64 {
	if x != nil {
		return x.BlockedCount
	}
	return 0
}

func (x *PThreadDump) GetWaitedTime() int64 {
	if x != nil {
		return x.WaitedTime
	}
	return 0
}

func (x *PThreadDump) GetWaitedCount() int64 {
	if x != nil {
		return x.WaitedCount
	}
	return 0
}

func (x *PThreadDump) GetLockName() string {
	if x != nil {
		return x.LockName
	}
	return ""
}

func (x *PThreadDump) GetLockOwnerId() int64 {
	if x != nil {
		return x.LockOwnerId
	}
	return 0
}

func (x *PThreadDump) GetLockOwnerName() string {
	if x != nil {
		return x.LockOwnerName
	}
	return ""
}

func (x *PThreadDump) GetInNative() bool {
	if x != nil {
		return x.InNative
	}
	return false
}

func (x *PThreadDump) GetSuspended() bool {
	if x != nil {
		return x.Suspended
	}
	return false
}

func (x *PThreadDump) GetThreadState() PThreadState {
	if x != nil {
		return x.ThreadState
	}
	return PThreadState_THREAD_STATE_NEW
}

func (x *PThreadDump) GetStackTrace() []string {
	if x != nil {
		return x.StackTrace
	}
	return nil
}

func (x *PThreadDump) GetLockedMonitor() []*PMonitorInfo {
	if x != nil {
		return x.LockedMonitor
	}
	return nil
}

func (x *PThreadDump) GetLockedSynchronizer() []string {
	if x != nil {
		return x.LockedSynchronizer
	}
	return nil
}

type PThreadLightDump struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThreadName  string       `protobuf:"bytes,1,opt,name=threadName,proto3" json:"threadName,omitempty"`
	ThreadId    int64        `protobuf:"varint,2,opt,name=threadId,proto3" json:"threadId,omitempty"`
	ThreadState PThreadState `protobuf:"varint,3,opt,name=threadState,proto3,enum=v1.PThreadState" json:"threadState,omitempty"`
}

func (x *PThreadLightDump) Reset() {
	*x = PThreadLightDump{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ThreadDump_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PThreadLightDump) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PThreadLightDump) ProtoMessage() {}

func (x *PThreadLightDump) ProtoReflect() protoreflect.Message {
	mi := &file_ThreadDump_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PThreadLightDump.ProtoReflect.Descriptor instead.
func (*PThreadLightDump) Descriptor() ([]byte, []int) {
	return file_ThreadDump_proto_rawDescGZIP(), []int{2}
}

func (x *PThreadLightDump) GetThreadName() string {
	if x != nil {
		return x.ThreadName
	}
	return ""
}

func (x *PThreadLightDump) GetThreadId() int64 {
	if x != nil {
		return x.ThreadId
	}
	return 0
}

func (x *PThreadLightDump) GetThreadState() PThreadState {
	if x != nil {
		return x.ThreadState
	}
	return PThreadState_THREAD_STATE_NEW
}

type PActiveThreadDump struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime     int64        `protobuf:"varint,1,opt,name=startTime,proto3" json:"startTime,omitempty"`
	LocalTraceId  int64        `protobuf:"varint,2,opt,name=localTraceId,proto3" json:"localTraceId,omitempty"`
	ThreadDump    *PThreadDump `protobuf:"bytes,3,opt,name=threadDump,proto3" json:"threadDump,omitempty"`
	Sampled       bool         `protobuf:"varint,4,opt,name=sampled,proto3" json:"sampled,omitempty"`
	TransactionId string       `protobuf:"bytes,5,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
	EntryPoint    string       `protobuf:"bytes,6,opt,name=entryPoint,proto3" json:"entryPoint,omitempty"`
}

func (x *PActiveThreadDump) Reset() {
	*x = PActiveThreadDump{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ThreadDump_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PActiveThreadDump) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PActiveThreadDump) ProtoMessage() {}

func (x *PActiveThreadDump) ProtoReflect() protoreflect.Message {
	mi := &file_ThreadDump_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PActiveThreadDump.ProtoReflect.Descriptor instead.
func (*PActiveThreadDump) Descriptor() ([]byte, []int) {
	return file_ThreadDump_proto_rawDescGZIP(), []int{3}
}

func (x *PActiveThreadDump) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *PActiveThreadDump) GetLocalTraceId() int64 {
	if x != nil {
		return x.LocalTraceId
	}
	return 0
}

func (x *PActiveThreadDump) GetThreadDump() *PThreadDump {
	if x != nil {
		return x.ThreadDump
	}
	return nil
}

func (x *PActiveThreadDump) GetSampled() bool {
	if x != nil {
		return x.Sampled
	}
	return false
}

func (x *PActiveThreadDump) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *PActiveThreadDump) GetEntryPoint() string {
	if x != nil {
		return x.EntryPoint
	}
	return ""
}

type PActiveThreadLightDump struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime     int64             `protobuf:"varint,1,opt,name=startTime,proto3" json:"startTime,omitempty"`
	LocalTraceId  int64             `protobuf:"varint,2,opt,name=localTraceId,proto3" json:"localTraceId,omitempty"`
	ThreadDump    *PThreadLightDump `protobuf:"bytes,3,opt,name=threadDump,proto3" json:"threadDump,omitempty"`
	Sampled       bool              `protobuf:"varint,4,opt,name=sampled,proto3" json:"sampled,omitempty"`
	TransactionId string            `protobuf:"bytes,5,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
	EntryPoint    string            `protobuf:"bytes,6,opt,name=entryPoint,proto3" json:"entryPoint,omitempty"`
}

func (x *PActiveThreadLightDump) Reset() {
	*x = PActiveThreadLightDump{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ThreadDump_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PActiveThreadLightDump) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PActiveThreadLightDump) ProtoMessage() {}

func (x *PActiveThreadLightDump) ProtoReflect() protoreflect.Message {
	mi := &file_ThreadDump_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PActiveThreadLightDump.ProtoReflect.Descriptor instead.
func (*PActiveThreadLightDump) Descriptor() ([]byte, []int) {
	return file_ThreadDump_proto_rawDescGZIP(), []int{4}
}

func (x *PActiveThreadLightDump) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *PActiveThreadLightDump) GetLocalTraceId() int64 {
	if x != nil {
		return x.LocalTraceId
	}
	return 0
}

func (x *PActiveThreadLightDump) GetThreadDump() *PThreadLightDump {
	if x != nil {
		return x.ThreadDump
	}
	return nil
}

func (x *PActiveThreadLightDump) GetSampled() bool {
	if x != nil {
		return x.Sampled
	}
	return false
}

func (x *PActiveThreadLightDump) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *PActiveThreadLightDump) GetEntryPoint() string {
	if x != nil {
		return x.EntryPoint
	}
	return ""
}

var File_ThreadDump_proto protoreflect.FileDescriptor

var file_ThreadDump_proto_rawDesc = []byte{
	0x0a, 0x10, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x44, 0x75, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x0c, 0x50, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x44, 0x65, 0x70, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x44, 0x65,
	0x70, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x46, 0x72, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x46, 0x72,
	0x61, 0x6d, 0x65, 0x22, 0xab, 0x04, 0x0a, 0x0b, 0x50, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x44,
	0x75, 0x6d, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x61, 0x69, 0x74, 0x65, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x77, 0x61, 0x69, 0x74, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x61, 0x69, 0x74, 0x65, 0x64, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x77, 0x61, 0x69, 0x74,
	0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x6b, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x6f, 0x63, 0x6b, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x6b, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x6f,
	0x63, 0x6b, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x6e, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x6e, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x73, 0x70, 0x65,
	0x6e, 0x64, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x75, 0x73, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x0b, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x74, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x0d, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x64, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0d, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x12, 0x2e, 0x0a, 0x12, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x53, 0x79, 0x6e, 0x63, 0x68,
	0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65,
	0x72, 0x22, 0x82, 0x01, 0x0a, 0x10, 0x50, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x4c, 0x69, 0x67,
	0x68, 0x74, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x49, 0x64, 0x12, 0x32, 0x0a, 0x0b, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x54, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x74, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0xe6, 0x01, 0x0a, 0x11, 0x50, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2f,
	0x0a, 0x0a, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x44, 0x75, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x44,
	0x75, 0x6d, 0x70, 0x52, 0x0a, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x44, 0x75, 0x6d, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x22,
	0xf0, 0x01, 0x0a, 0x16, 0x50, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x0a,
	0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x44, 0x75, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x4c, 0x69, 0x67,
	0x68, 0x74, 0x44, 0x75, 0x6d, 0x70, 0x52, 0x0a, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x44, 0x75,
	0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x2a, 0xca, 0x01, 0x0a, 0x0c, 0x50, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x48, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x4e, 0x45, 0x57, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x48, 0x52,
	0x45, 0x41, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x48, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x02, 0x12, 0x18,
	0x0a, 0x14, 0x54, 0x48, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x57,
	0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x48, 0x52, 0x45,
	0x41, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x44, 0x5f, 0x57,
	0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x48, 0x52, 0x45,
	0x41, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41,
	0x54, 0x45, 0x44, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x48, 0x52, 0x45, 0x41, 0x44, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x06, 0x2a,
	0x2a, 0x0a, 0x0f, 0x50, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x44, 0x75, 0x6d, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x42, 0x5b, 0x0a, 0x21, 0x63,
	0x6f, 0x6d, 0x2e, 0x6e, 0x61, 0x76, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x70, 0x69, 0x6e,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x42, 0x0f, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x44, 0x75, 0x6d, 0x70, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x69, 0x6e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x67, 0x6f, 0x70, 0x69, 0x6e, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ThreadDump_proto_rawDescOnce sync.Once
	file_ThreadDump_proto_rawDescData = file_ThreadDump_proto_rawDesc
)

func file_ThreadDump_proto_rawDescGZIP() []byte {
	file_ThreadDump_proto_rawDescOnce.Do(func() {
		file_ThreadDump_proto_rawDescData = protoimpl.X.CompressGZIP(file_ThreadDump_proto_rawDescData)
	})
	return file_ThreadDump_proto_rawDescData
}

var file_ThreadDump_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ThreadDump_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ThreadDump_proto_goTypes = []interface{}{
	(PThreadState)(0),              // 0: v1.PThreadState
	(PThreadDumpType)(0),           // 1: v1.PThreadDumpType
	(*PMonitorInfo)(nil),           // 2: v1.PMonitorInfo
	(*PThreadDump)(nil),            // 3: v1.PThreadDump
	(*PThreadLightDump)(nil),       // 4: v1.PThreadLightDump
	(*PActiveThreadDump)(nil),      // 5: v1.PActiveThreadDump
	(*PActiveThreadLightDump)(nil), // 6: v1.PActiveThreadLightDump
}
var file_ThreadDump_proto_depIdxs = []int32{
	0, // 0: v1.PThreadDump.threadState:type_name -> v1.PThreadState
	2, // 1: v1.PThreadDump.lockedMonitor:type_name -> v1.PMonitorInfo
	0, // 2: v1.PThreadLightDump.threadState:type_name -> v1.PThreadState
	3, // 3: v1.PActiveThreadDump.threadDump:type_name -> v1.PThreadDump
	4, // 4: v1.PActiveThreadLightDump.threadDump:type_name -> v1.PThreadLightDump
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ThreadDump_proto_init() }
func file_ThreadDump_proto_init() {
	if File_ThreadDump_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ThreadDump_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PMonitorInfo); i {
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
		file_ThreadDump_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PThreadDump); i {
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
		file_ThreadDump_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PThreadLightDump); i {
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
		file_ThreadDump_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PActiveThreadDump); i {
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
		file_ThreadDump_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PActiveThreadLightDump); i {
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
			RawDescriptor: file_ThreadDump_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ThreadDump_proto_goTypes,
		DependencyIndexes: file_ThreadDump_proto_depIdxs,
		EnumInfos:         file_ThreadDump_proto_enumTypes,
		MessageInfos:      file_ThreadDump_proto_msgTypes,
	}.Build()
	File_ThreadDump_proto = out.File
	file_ThreadDump_proto_rawDesc = nil
	file_ThreadDump_proto_goTypes = nil
	file_ThreadDump_proto_depIdxs = nil
}
