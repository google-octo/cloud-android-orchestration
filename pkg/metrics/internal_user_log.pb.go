// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal_user_log.proto

package metrics

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UserType int32

const (
	UserType_GOOGLE   UserType = 0
	UserType_EXTERNAL UserType = 1
)

var UserType_name = map[int32]string{
	0: "GOOGLE",
	1: "EXTERNAL",
}

var UserType_value = map[string]int32{
	"GOOGLE":   0,
	"EXTERNAL": 1,
}

func (x UserType) Enum() *UserType {
	p := new(UserType)
	*p = x
	return p
}

func (x UserType) String() string {
	return proto.EnumName(UserType_name, int32(x))
}

func (x *UserType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(UserType_value, data, "UserType")
	if err != nil {
		return err
	}
	*x = UserType(value)
	return nil
}

func (UserType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36532fe67711fe38, []int{0}
}

type Duration struct {
	Seconds              *int64   `protobuf:"varint,1,req,name=seconds" json:"seconds,omitempty"`
	Nanos                *int32   `protobuf:"varint,2,req,name=nanos" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Duration) Reset()         { *m = Duration{} }
func (m *Duration) String() string { return proto.CompactTextString(m) }
func (*Duration) ProtoMessage()    {}
func (*Duration) Descriptor() ([]byte, []int) {
	return fileDescriptor_36532fe67711fe38, []int{0}
}

func (m *Duration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Duration.Unmarshal(m, b)
}
func (m *Duration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Duration.Marshal(b, m, deterministic)
}
func (m *Duration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Duration.Merge(m, src)
}
func (m *Duration) XXX_Size() int {
	return xxx_messageInfo_Duration.Size(m)
}
func (m *Duration) XXX_DiscardUnknown() {
	xxx_messageInfo_Duration.DiscardUnknown(m)
}

var xxx_messageInfo_Duration proto.InternalMessageInfo

func (m *Duration) GetSeconds() int64 {
	if m != nil && m.Seconds != nil {
		return *m.Seconds
	}
	return 0
}

func (m *Duration) GetNanos() int32 {
	if m != nil && m.Nanos != nil {
		return *m.Nanos
	}
	return 0
}

// Proto used by Atest CLI Tool for Internal PII Users
type AtestLogEventInternal struct {
	// ------------------------
	// FIELDS FOR ATESTLOGEVENT
	// ------------------------
	// user_key and run_id are both python uuid.uuid4() completely anonymous
	// and random hexadecimal strings. We cannot use normal google ids
	// (zwieback, gaia, etc) because we are an open source command line tool.
	// uuid4 provides us a unique key for grouping, which is all we need.
	UserKey     *string   `protobuf:"bytes,1,opt,name=user_key,json=userKey" json:"user_key,omitempty"`
	RunId       *string   `protobuf:"bytes,2,opt,name=run_id,json=runId" json:"run_id,omitempty"`
	UserType    *UserType `protobuf:"varint,3,opt,name=user_type,json=userType,enum=metrics.UserType" json:"user_type,omitempty"`
	ToolName    *string   `protobuf:"bytes,10,opt,name=tool_name,json=toolName" json:"tool_name,omitempty"`
	SubToolName *string   `protobuf:"bytes,12,opt,name=sub_tool_name,json=subToolName" json:"sub_tool_name,omitempty"`
	// Types that are valid to be assigned to Event:
	//
	//	*AtestLogEventInternal_AtestStartEvent_
	//	*AtestLogEventInternal_AtestExitEvent_
	//	*AtestLogEventInternal_FindTestFinishEvent_
	//	*AtestLogEventInternal_BuildFinishEvent_
	//	*AtestLogEventInternal_RunnerFinishEvent_
	//	*AtestLogEventInternal_RunTestsFinishEvent_
	//	*AtestLogEventInternal_LocalDetectEvent_
	Event                isAtestLogEventInternal_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AtestLogEventInternal) Reset()         { *m = AtestLogEventInternal{} }
func (m *AtestLogEventInternal) String() string { return proto.CompactTextString(m) }
func (*AtestLogEventInternal) ProtoMessage()    {}
func (*AtestLogEventInternal) Descriptor() ([]byte, []int) {
	return fileDescriptor_36532fe67711fe38, []int{1}
}

func (m *AtestLogEventInternal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AtestLogEventInternal.Unmarshal(m, b)
}
func (m *AtestLogEventInternal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AtestLogEventInternal.Marshal(b, m, deterministic)
}
func (m *AtestLogEventInternal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AtestLogEventInternal.Merge(m, src)
}
func (m *AtestLogEventInternal) XXX_Size() int {
	return xxx_messageInfo_AtestLogEventInternal.Size(m)
}
func (m *AtestLogEventInternal) XXX_DiscardUnknown() {
	xxx_messageInfo_AtestLogEventInternal.DiscardUnknown(m)
}

var xxx_messageInfo_AtestLogEventInternal proto.InternalMessageInfo

func (m *AtestLogEventInternal) GetUserKey() string {
	if m != nil && m.UserKey != nil {
		return *m.UserKey
	}
	return ""
}

func (m *AtestLogEventInternal) GetRunId() string {
	if m != nil && m.RunId != nil {
		return *m.RunId
	}
	return ""
}

func (m *AtestLogEventInternal) GetUserType() UserType {
	if m != nil && m.UserType != nil {
		return *m.UserType
	}
	return UserType_GOOGLE
}

func (m *AtestLogEventInternal) GetToolName() string {
	if m != nil && m.ToolName != nil {
		return *m.ToolName
	}
	return ""
}

func (m *AtestLogEventInternal) GetSubToolName() string {
	if m != nil && m.SubToolName != nil {
		return *m.SubToolName
	}
	return ""
}

type isAtestLogEventInternal_Event interface {
	isAtestLogEventInternal_Event()
}

type AtestLogEventInternal_AtestStartEvent_ struct {
	AtestStartEvent *AtestLogEventInternal_AtestStartEvent `protobuf:"bytes,4,opt,name=atest_start_event,json=atestStartEvent,oneof"`
}

type AtestLogEventInternal_AtestExitEvent_ struct {
	AtestExitEvent *AtestLogEventInternal_AtestExitEvent `protobuf:"bytes,5,opt,name=atest_exit_event,json=atestExitEvent,oneof"`
}

type AtestLogEventInternal_FindTestFinishEvent_ struct {
	FindTestFinishEvent *AtestLogEventInternal_FindTestFinishEvent `protobuf:"bytes,6,opt,name=find_test_finish_event,json=findTestFinishEvent,oneof"`
}

type AtestLogEventInternal_BuildFinishEvent_ struct {
	BuildFinishEvent *AtestLogEventInternal_BuildFinishEvent `protobuf:"bytes,7,opt,name=build_finish_event,json=buildFinishEvent,oneof"`
}

type AtestLogEventInternal_RunnerFinishEvent_ struct {
	RunnerFinishEvent *AtestLogEventInternal_RunnerFinishEvent `protobuf:"bytes,8,opt,name=runner_finish_event,json=runnerFinishEvent,oneof"`
}

type AtestLogEventInternal_RunTestsFinishEvent_ struct {
	RunTestsFinishEvent *AtestLogEventInternal_RunTestsFinishEvent `protobuf:"bytes,9,opt,name=run_tests_finish_event,json=runTestsFinishEvent,oneof"`
}

type AtestLogEventInternal_LocalDetectEvent_ struct {
	LocalDetectEvent *AtestLogEventInternal_LocalDetectEvent `protobuf:"bytes,11,opt,name=local_detect_event,json=localDetectEvent,oneof"`
}

func (*AtestLogEventInternal_AtestStartEvent_) isAtestLogEventInternal_Event() {}

func (*AtestLogEventInternal_AtestExitEvent_) isAtestLogEventInternal_Event() {}

func (*AtestLogEventInternal_FindTestFinishEvent_) isAtestLogEventInternal_Event() {}

func (*AtestLogEventInternal_BuildFinishEvent_) isAtestLogEventInternal_Event() {}

func (*AtestLogEventInternal_RunnerFinishEvent_) isAtestLogEventInternal_Event() {}

func (*AtestLogEventInternal_RunTestsFinishEvent_) isAtestLogEventInternal_Event() {}

func (*AtestLogEventInternal_LocalDetectEvent_) isAtestLogEventInternal_Event() {}

func (m *AtestLogEventInternal) GetEvent() isAtestLogEventInternal_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *AtestLogEventInternal) GetAtestStartEvent() *AtestLogEventInternal_AtestStartEvent {
	if x, ok := m.GetEvent().(*AtestLogEventInternal_AtestStartEvent_); ok {
		return x.AtestStartEvent
	}
	return nil
}

func (m *AtestLogEventInternal) GetAtestExitEvent() *AtestLogEventInternal_AtestExitEvent {
	if x, ok := m.GetEvent().(*AtestLogEventInternal_AtestExitEvent_); ok {
		return x.AtestExitEvent
	}
	return nil
}

func (m *AtestLogEventInternal) GetFindTestFinishEvent() *AtestLogEventInternal_FindTestFinishEvent {
	if x, ok := m.GetEvent().(*AtestLogEventInternal_FindTestFinishEvent_); ok {
		return x.FindTestFinishEvent
	}
	return nil
}

func (m *AtestLogEventInternal) GetBuildFinishEvent() *AtestLogEventInternal_BuildFinishEvent {
	if x, ok := m.GetEvent().(*AtestLogEventInternal_BuildFinishEvent_); ok {
		return x.BuildFinishEvent
	}
	return nil
}

func (m *AtestLogEventInternal) GetRunnerFinishEvent() *AtestLogEventInternal_RunnerFinishEvent {
	if x, ok := m.GetEvent().(*AtestLogEventInternal_RunnerFinishEvent_); ok {
		return x.RunnerFinishEvent
	}
	return nil
}

func (m *AtestLogEventInternal) GetRunTestsFinishEvent() *AtestLogEventInternal_RunTestsFinishEvent {
	if x, ok := m.GetEvent().(*AtestLogEventInternal_RunTestsFinishEvent_); ok {
		return x.RunTestsFinishEvent
	}
	return nil
}

func (m *AtestLogEventInternal) GetLocalDetectEvent() *AtestLogEventInternal_LocalDetectEvent {
	if x, ok := m.GetEvent().(*AtestLogEventInternal_LocalDetectEvent_); ok {
		return x.LocalDetectEvent
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AtestLogEventInternal) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AtestLogEventInternal_AtestStartEvent_)(nil),
		(*AtestLogEventInternal_AtestExitEvent_)(nil),
		(*AtestLogEventInternal_FindTestFinishEvent_)(nil),
		(*AtestLogEventInternal_BuildFinishEvent_)(nil),
		(*AtestLogEventInternal_RunnerFinishEvent_)(nil),
		(*AtestLogEventInternal_RunTestsFinishEvent_)(nil),
		(*AtestLogEventInternal_LocalDetectEvent_)(nil),
	}
}

// ------------------------
// EVENT DEFINITIONS
// ------------------------
// Occurs immediately upon execution of atest
type AtestLogEventInternal_AtestStartEvent struct {
	CommandLine          *string  `protobuf:"bytes,1,opt,name=command_line,json=commandLine" json:"command_line,omitempty"`
	TestReferences       []string `protobuf:"bytes,2,rep,name=test_references,json=testReferences" json:"test_references,omitempty"`
	Cwd                  *string  `protobuf:"bytes,3,opt,name=cwd" json:"cwd,omitempty"`
	Os                   *string  `protobuf:"bytes,4,opt,name=os" json:"os,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AtestLogEventInternal_AtestStartEvent) Reset()         { *m = AtestLogEventInternal_AtestStartEvent{} }
func (m *AtestLogEventInternal_AtestStartEvent) String() string { return proto.CompactTextString(m) }
func (*AtestLogEventInternal_AtestStartEvent) ProtoMessage()    {}
func (*AtestLogEventInternal_AtestStartEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_36532fe67711fe38, []int{1, 0}
}

func (m *AtestLogEventInternal_AtestStartEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AtestLogEventInternal_AtestStartEvent.Unmarshal(m, b)
}
func (m *AtestLogEventInternal_AtestStartEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AtestLogEventInternal_AtestStartEvent.Marshal(b, m, deterministic)
}
func (m *AtestLogEventInternal_AtestStartEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AtestLogEventInternal_AtestStartEvent.Merge(m, src)
}
func (m *AtestLogEventInternal_AtestStartEvent) XXX_Size() int {
	return xxx_messageInfo_AtestLogEventInternal_AtestStartEvent.Size(m)
}
func (m *AtestLogEventInternal_AtestStartEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AtestLogEventInternal_AtestStartEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AtestLogEventInternal_AtestStartEvent proto.InternalMessageInfo

func (m *AtestLogEventInternal_AtestStartEvent) GetCommandLine() string {
	if m != nil && m.CommandLine != nil {
		return *m.CommandLine
	}
	return ""
}

func (m *AtestLogEventInternal_AtestStartEvent) GetTestReferences() []string {
	if m != nil {
		return m.TestReferences
	}
	return nil
}

func (m *AtestLogEventInternal_AtestStartEvent) GetCwd() string {
	if m != nil && m.Cwd != nil {
		return *m.Cwd
	}
	return ""
}

func (m *AtestLogEventInternal_AtestStartEvent) GetOs() string {
	if m != nil && m.Os != nil {
		return *m.Os
	}
	return ""
}

// Occurs when atest exits for any reason
type AtestLogEventInternal_AtestExitEvent struct {
	Duration             *Duration `protobuf:"bytes,1,opt,name=duration" json:"duration,omitempty"`
	ExitCode             *int32    `protobuf:"varint,2,opt,name=exit_code,json=exitCode" json:"exit_code,omitempty"`
	Stacktrace           *string   `protobuf:"bytes,3,opt,name=stacktrace" json:"stacktrace,omitempty"`
	Logs                 *string   `protobuf:"bytes,4,opt,name=logs" json:"logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AtestLogEventInternal_AtestExitEvent) Reset()         { *m = AtestLogEventInternal_AtestExitEvent{} }
func (m *AtestLogEventInternal_AtestExitEvent) String() string { return proto.CompactTextString(m) }
func (*AtestLogEventInternal_AtestExitEvent) ProtoMessage()    {}
func (*AtestLogEventInternal_AtestExitEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_36532fe67711fe38, []int{1, 1}
}

func (m *AtestLogEventInternal_AtestExitEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AtestLogEventInternal_AtestExitEvent.Unmarshal(m, b)
}
func (m *AtestLogEventInternal_AtestExitEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AtestLogEventInternal_AtestExitEvent.Marshal(b, m, deterministic)
}
func (m *AtestLogEventInternal_AtestExitEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AtestLogEventInternal_AtestExitEvent.Merge(m, src)
}
func (m *AtestLogEventInternal_AtestExitEvent) XXX_Size() int {
	return xxx_messageInfo_AtestLogEventInternal_AtestExitEvent.Size(m)
}
func (m *AtestLogEventInternal_AtestExitEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AtestLogEventInternal_AtestExitEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AtestLogEventInternal_AtestExitEvent proto.InternalMessageInfo

func (m *AtestLogEventInternal_AtestExitEvent) GetDuration() *Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *AtestLogEventInternal_AtestExitEvent) GetExitCode() int32 {
	if m != nil && m.ExitCode != nil {
		return *m.ExitCode
	}
	return 0
}

func (m *AtestLogEventInternal_AtestExitEvent) GetStacktrace() string {
	if m != nil && m.Stacktrace != nil {
		return *m.Stacktrace
	}
	return ""
}

func (m *AtestLogEventInternal_AtestExitEvent) GetLogs() string {
	if m != nil && m.Logs != nil {
		return *m.Logs
	}
	return ""
}

// Occurs after a SINGLE test reference has been resolved to a test or
// not found
type AtestLogEventInternal_FindTestFinishEvent struct {
	Duration             *Duration `protobuf:"bytes,1,opt,name=duration" json:"duration,omitempty"`
	Success              *bool     `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
	TestReference        *string   `protobuf:"bytes,3,opt,name=test_reference,json=testReference" json:"test_reference,omitempty"`
	TestFinders          []string  `protobuf:"bytes,4,rep,name=test_finders,json=testFinders" json:"test_finders,omitempty"`
	TestInfo             *string   `protobuf:"bytes,5,opt,name=test_info,json=testInfo" json:"test_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AtestLogEventInternal_FindTestFinishEvent) Reset() {
	*m = AtestLogEventInternal_FindTestFinishEvent{}
}
func (m *AtestLogEventInternal_FindTestFinishEvent) String() string {
	return proto.CompactTextString(m)
}
func (*AtestLogEventInternal_FindTestFinishEvent) ProtoMessage() {}
func (*AtestLogEventInternal_FindTestFinishEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_36532fe67711fe38, []int{1, 2}
}

func (m *AtestLogEventInternal_FindTestFinishEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AtestLogEventInternal_FindTestFinishEvent.Unmarshal(m, b)
}
func (m *AtestLogEventInternal_FindTestFinishEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AtestLogEventInternal_FindTestFinishEvent.Marshal(b, m, deterministic)
}
func (m *AtestLogEventInternal_FindTestFinishEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AtestLogEventInternal_FindTestFinishEvent.Merge(m, src)
}
func (m *AtestLogEventInternal_FindTestFinishEvent) XXX_Size() int {
	return xxx_messageInfo_AtestLogEventInternal_FindTestFinishEvent.Size(m)
}
func (m *AtestLogEventInternal_FindTestFinishEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AtestLogEventInternal_FindTestFinishEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AtestLogEventInternal_FindTestFinishEvent proto.InternalMessageInfo

func (m *AtestLogEventInternal_FindTestFinishEvent) GetDuration() *Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *AtestLogEventInternal_FindTestFinishEvent) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *AtestLogEventInternal_FindTestFinishEvent) GetTestReference() string {
	if m != nil && m.TestReference != nil {
		return *m.TestReference
	}
	return ""
}

func (m *AtestLogEventInternal_FindTestFinishEvent) GetTestFinders() []string {
	if m != nil {
		return m.TestFinders
	}
	return nil
}

func (m *AtestLogEventInternal_FindTestFinishEvent) GetTestInfo() string {
	if m != nil && m.TestInfo != nil {
		return *m.TestInfo
	}
	return ""
}

// Occurs after the build finishes, either successfully or not.
type AtestLogEventInternal_BuildFinishEvent struct {
	Duration             *Duration `protobuf:"bytes,1,opt,name=duration" json:"duration,omitempty"`
	Success              *bool     `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
	Targets              []string  `protobuf:"bytes,3,rep,name=targets" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AtestLogEventInternal_BuildFinishEvent) Reset() {
	*m = AtestLogEventInternal_BuildFinishEvent{}
}
func (m *AtestLogEventInternal_BuildFinishEvent) String() string { return proto.CompactTextString(m) }
func (*AtestLogEventInternal_BuildFinishEvent) ProtoMessage()    {}
func (*AtestLogEventInternal_BuildFinishEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_36532fe67711fe38, []int{1, 3}
}

func (m *AtestLogEventInternal_BuildFinishEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AtestLogEventInternal_BuildFinishEvent.Unmarshal(m, b)
}
func (m *AtestLogEventInternal_BuildFinishEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AtestLogEventInternal_BuildFinishEvent.Marshal(b, m, deterministic)
}
func (m *AtestLogEventInternal_BuildFinishEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AtestLogEventInternal_BuildFinishEvent.Merge(m, src)
}
func (m *AtestLogEventInternal_BuildFinishEvent) XXX_Size() int {
	return xxx_messageInfo_AtestLogEventInternal_BuildFinishEvent.Size(m)
}
func (m *AtestLogEventInternal_BuildFinishEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AtestLogEventInternal_BuildFinishEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AtestLogEventInternal_BuildFinishEvent proto.InternalMessageInfo

func (m *AtestLogEventInternal_BuildFinishEvent) GetDuration() *Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *AtestLogEventInternal_BuildFinishEvent) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *AtestLogEventInternal_BuildFinishEvent) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

// Occurs when a single test runner has completed
type AtestLogEventInternal_RunnerFinishEvent struct {
	Duration             *Duration                                       `protobuf:"bytes,1,opt,name=duration" json:"duration,omitempty"`
	Success              *bool                                           `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
	RunnerName           *string                                         `protobuf:"bytes,3,opt,name=runner_name,json=runnerName" json:"runner_name,omitempty"`
	Test                 []*AtestLogEventInternal_RunnerFinishEvent_Test `protobuf:"bytes,4,rep,name=test" json:"test,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *AtestLogEventInternal_RunnerFinishEvent) Reset() {
	*m = AtestLogEventInternal_RunnerFinishEvent{}
}
func (m *AtestLogEventInternal_RunnerFinishEvent) String() string { return proto.CompactTextString(m) }
func (*AtestLogEventInternal_RunnerFinishEvent) ProtoMessage()    {}
func (*AtestLogEventInternal_RunnerFinishEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_36532fe67711fe38, []int{1, 4}
}

func (m *AtestLogEventInternal_RunnerFinishEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AtestLogEventInternal_RunnerFinishEvent.Unmarshal(m, b)
}
func (m *AtestLogEventInternal_RunnerFinishEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AtestLogEventInternal_RunnerFinishEvent.Marshal(b, m, deterministic)
}
func (m *AtestLogEventInternal_RunnerFinishEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AtestLogEventInternal_RunnerFinishEvent.Merge(m, src)
}
func (m *AtestLogEventInternal_RunnerFinishEvent) XXX_Size() int {
	return xxx_messageInfo_AtestLogEventInternal_RunnerFinishEvent.Size(m)
}
func (m *AtestLogEventInternal_RunnerFinishEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AtestLogEventInternal_RunnerFinishEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AtestLogEventInternal_RunnerFinishEvent proto.InternalMessageInfo

func (m *AtestLogEventInternal_RunnerFinishEvent) GetDuration() *Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *AtestLogEventInternal_RunnerFinishEvent) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *AtestLogEventInternal_RunnerFinishEvent) GetRunnerName() string {
	if m != nil && m.RunnerName != nil {
		return *m.RunnerName
	}
	return ""
}

func (m *AtestLogEventInternal_RunnerFinishEvent) GetTest() []*AtestLogEventInternal_RunnerFinishEvent_Test {
	if m != nil {
		return m.Test
	}
	return nil
}

type AtestLogEventInternal_RunnerFinishEvent_Test struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Result               *int32   `protobuf:"varint,2,opt,name=result" json:"result,omitempty"`
	Stacktrace           *string  `protobuf:"bytes,3,opt,name=stacktrace" json:"stacktrace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AtestLogEventInternal_RunnerFinishEvent_Test) Reset() {
	*m = AtestLogEventInternal_RunnerFinishEvent_Test{}
}
func (m *AtestLogEventInternal_RunnerFinishEvent_Test) String() string {
	return proto.CompactTextString(m)
}
func (*AtestLogEventInternal_RunnerFinishEvent_Test) ProtoMessage() {}
func (*AtestLogEventInternal_RunnerFinishEvent_Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_36532fe67711fe38, []int{1, 4, 0}
}

func (m *AtestLogEventInternal_RunnerFinishEvent_Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AtestLogEventInternal_RunnerFinishEvent_Test.Unmarshal(m, b)
}
func (m *AtestLogEventInternal_RunnerFinishEvent_Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AtestLogEventInternal_RunnerFinishEvent_Test.Marshal(b, m, deterministic)
}
func (m *AtestLogEventInternal_RunnerFinishEvent_Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AtestLogEventInternal_RunnerFinishEvent_Test.Merge(m, src)
}
func (m *AtestLogEventInternal_RunnerFinishEvent_Test) XXX_Size() int {
	return xxx_messageInfo_AtestLogEventInternal_RunnerFinishEvent_Test.Size(m)
}
func (m *AtestLogEventInternal_RunnerFinishEvent_Test) XXX_DiscardUnknown() {
	xxx_messageInfo_AtestLogEventInternal_RunnerFinishEvent_Test.DiscardUnknown(m)
}

var xxx_messageInfo_AtestLogEventInternal_RunnerFinishEvent_Test proto.InternalMessageInfo

func (m *AtestLogEventInternal_RunnerFinishEvent_Test) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *AtestLogEventInternal_RunnerFinishEvent_Test) GetResult() int32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

func (m *AtestLogEventInternal_RunnerFinishEvent_Test) GetStacktrace() string {
	if m != nil && m.Stacktrace != nil {
		return *m.Stacktrace
	}
	return ""
}

// Occurs after all test runners and tests have finished
type AtestLogEventInternal_RunTestsFinishEvent struct {
	Duration             *Duration `protobuf:"bytes,1,opt,name=duration" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AtestLogEventInternal_RunTestsFinishEvent) Reset() {
	*m = AtestLogEventInternal_RunTestsFinishEvent{}
}
func (m *AtestLogEventInternal_RunTestsFinishEvent) String() string {
	return proto.CompactTextString(m)
}
func (*AtestLogEventInternal_RunTestsFinishEvent) ProtoMessage() {}
func (*AtestLogEventInternal_RunTestsFinishEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_36532fe67711fe38, []int{1, 5}
}

func (m *AtestLogEventInternal_RunTestsFinishEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AtestLogEventInternal_RunTestsFinishEvent.Unmarshal(m, b)
}
func (m *AtestLogEventInternal_RunTestsFinishEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AtestLogEventInternal_RunTestsFinishEvent.Marshal(b, m, deterministic)
}
func (m *AtestLogEventInternal_RunTestsFinishEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AtestLogEventInternal_RunTestsFinishEvent.Merge(m, src)
}
func (m *AtestLogEventInternal_RunTestsFinishEvent) XXX_Size() int {
	return xxx_messageInfo_AtestLogEventInternal_RunTestsFinishEvent.Size(m)
}
func (m *AtestLogEventInternal_RunTestsFinishEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AtestLogEventInternal_RunTestsFinishEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AtestLogEventInternal_RunTestsFinishEvent proto.InternalMessageInfo

func (m *AtestLogEventInternal_RunTestsFinishEvent) GetDuration() *Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

// Occurs after detection of catching bug by atest have finished
type AtestLogEventInternal_LocalDetectEvent struct {
	DetectType           *int32   `protobuf:"varint,1,opt,name=detect_type,json=detectType" json:"detect_type,omitempty"`
	Result               *int32   `protobuf:"varint,2,opt,name=result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AtestLogEventInternal_LocalDetectEvent) Reset() {
	*m = AtestLogEventInternal_LocalDetectEvent{}
}
func (m *AtestLogEventInternal_LocalDetectEvent) String() string { return proto.CompactTextString(m) }
func (*AtestLogEventInternal_LocalDetectEvent) ProtoMessage()    {}
func (*AtestLogEventInternal_LocalDetectEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_36532fe67711fe38, []int{1, 6}
}

func (m *AtestLogEventInternal_LocalDetectEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AtestLogEventInternal_LocalDetectEvent.Unmarshal(m, b)
}
func (m *AtestLogEventInternal_LocalDetectEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AtestLogEventInternal_LocalDetectEvent.Marshal(b, m, deterministic)
}
func (m *AtestLogEventInternal_LocalDetectEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AtestLogEventInternal_LocalDetectEvent.Merge(m, src)
}
func (m *AtestLogEventInternal_LocalDetectEvent) XXX_Size() int {
	return xxx_messageInfo_AtestLogEventInternal_LocalDetectEvent.Size(m)
}
func (m *AtestLogEventInternal_LocalDetectEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AtestLogEventInternal_LocalDetectEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AtestLogEventInternal_LocalDetectEvent proto.InternalMessageInfo

func (m *AtestLogEventInternal_LocalDetectEvent) GetDetectType() int32 {
	if m != nil && m.DetectType != nil {
		return *m.DetectType
	}
	return 0
}

func (m *AtestLogEventInternal_LocalDetectEvent) GetResult() int32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

func init() {
	proto.RegisterEnum("metrics.UserType", UserType_name, UserType_value)
	proto.RegisterType((*Duration)(nil), "metrics.Duration")
	proto.RegisterType((*AtestLogEventInternal)(nil), "metrics.AtestLogEventInternal")
	proto.RegisterType((*AtestLogEventInternal_AtestStartEvent)(nil), "metrics.AtestLogEventInternal.AtestStartEvent")
	proto.RegisterType((*AtestLogEventInternal_AtestExitEvent)(nil), "metrics.AtestLogEventInternal.AtestExitEvent")
	proto.RegisterType((*AtestLogEventInternal_FindTestFinishEvent)(nil), "metrics.AtestLogEventInternal.FindTestFinishEvent")
	proto.RegisterType((*AtestLogEventInternal_BuildFinishEvent)(nil), "metrics.AtestLogEventInternal.BuildFinishEvent")
	proto.RegisterType((*AtestLogEventInternal_RunnerFinishEvent)(nil), "metrics.AtestLogEventInternal.RunnerFinishEvent")
	proto.RegisterType((*AtestLogEventInternal_RunnerFinishEvent_Test)(nil), "metrics.AtestLogEventInternal.RunnerFinishEvent.Test")
	proto.RegisterType((*AtestLogEventInternal_RunTestsFinishEvent)(nil), "metrics.AtestLogEventInternal.RunTestsFinishEvent")
	proto.RegisterType((*AtestLogEventInternal_LocalDetectEvent)(nil), "metrics.AtestLogEventInternal.LocalDetectEvent")
}

func init() {
	proto.RegisterFile("internal_user_log.proto", fileDescriptor_36532fe67711fe38)
}

var fileDescriptor_36532fe67711fe38 = []byte{
	// 794 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x6d, 0x6f, 0xe3, 0x44,
	0x10, 0x3e, 0xe7, 0xd5, 0x19, 0xf7, 0xd2, 0x64, 0xc3, 0x1d, 0xc6, 0x48, 0x10, 0x2a, 0x10, 0x11,
	0xd2, 0x05, 0x14, 0x89, 0x2f, 0x7c, 0xbb, 0xa3, 0xb9, 0x23, 0xba, 0xe8, 0x2a, 0x2d, 0x41, 0x02,
	0x09, 0xc9, 0x72, 0xec, 0x49, 0xb0, 0xea, 0xec, 0x56, 0xbb, 0x6b, 0x68, 0x3e, 0xf1, 0x17, 0xf8,
	0x01, 0xfc, 0x1d, 0x7e, 0x16, 0x12, 0xda, 0x89, 0x1d, 0xea, 0xb4, 0xa7, 0xb4, 0x52, 0xbf, 0xed,
	0x3c, 0xd9, 0x99, 0x67, 0xf2, 0xec, 0x3c, 0x63, 0xf8, 0x30, 0x15, 0x06, 0x95, 0x88, 0xb2, 0x30,
	0xd7, 0xa8, 0xc2, 0x4c, 0xae, 0xc7, 0x57, 0x4a, 0x1a, 0xc9, 0xda, 0x1b, 0x34, 0x2a, 0x8d, 0xf5,
	0xd9, 0x77, 0xe0, 0x9e, 0xe7, 0x2a, 0x32, 0xa9, 0x14, 0xcc, 0x87, 0xb6, 0xc6, 0x58, 0x8a, 0x44,
	0xfb, 0xce, 0xb0, 0x36, 0xaa, 0xf3, 0x32, 0x64, 0x1f, 0x40, 0x53, 0x44, 0x42, 0x6a, 0xbf, 0x36,
	0xac, 0x8d, 0x9a, 0x7c, 0x17, 0x9c, 0xfd, 0xdb, 0x85, 0x67, 0x2f, 0x0d, 0x6a, 0x33, 0x97, 0xeb,
	0xe9, 0xef, 0x28, 0xcc, 0xac, 0x60, 0x63, 0x1f, 0x81, 0x4b, 0x84, 0x97, 0xb8, 0xf5, 0x9d, 0xa1,
	0x33, 0xea, 0xf0, 0xb6, 0x8d, 0xdf, 0xe2, 0x96, 0x3d, 0x83, 0x96, 0xca, 0x45, 0x98, 0x26, 0x7e,
	0x8d, 0x7e, 0x68, 0xaa, 0x5c, 0xcc, 0x12, 0x36, 0x86, 0x0e, 0x65, 0x98, 0xed, 0x15, 0xfa, 0xf5,
	0xa1, 0x33, 0xea, 0x4e, 0xfa, 0xe3, 0xa2, 0xc9, 0xf1, 0x4f, 0x1a, 0xd5, 0x62, 0x7b, 0x85, 0x9c,
	0xaa, 0xda, 0x13, 0xfb, 0x18, 0x3a, 0x46, 0xca, 0x2c, 0x14, 0xd1, 0x06, 0x7d, 0xa0, 0x4a, 0xae,
	0x05, 0xde, 0x45, 0x1b, 0x64, 0x67, 0xf0, 0x54, 0xe7, 0xcb, 0xf0, 0xff, 0x0b, 0x27, 0x74, 0xc1,
	0xd3, 0xf9, 0x72, 0x51, 0xde, 0xf9, 0x15, 0xfa, 0x91, 0xed, 0x3d, 0xd4, 0x26, 0x52, 0x26, 0x44,
	0xdb, 0xbf, 0xdf, 0x18, 0x3a, 0x23, 0x6f, 0x32, 0xde, 0x13, 0xdf, 0xf9, 0xef, 0x76, 0xe8, 0x8f,
	0x36, 0x8d, 0xf0, 0x1f, 0x9e, 0xf0, 0xd3, 0xa8, 0x0a, 0xb1, 0x5f, 0xa0, 0xb7, 0xab, 0x8e, 0xd7,
	0x69, 0x59, 0xbc, 0x49, 0xc5, 0x5f, 0xdc, 0xa7, 0xf8, 0xf4, 0x3a, 0xdd, 0xd7, 0xee, 0x46, 0x15,
	0x84, 0xa5, 0xf0, 0x7c, 0x95, 0x8a, 0x24, 0xa4, 0xf2, 0xab, 0x54, 0xa4, 0xfa, 0xb7, 0x82, 0xa0,
	0x45, 0x04, 0x93, 0x23, 0x04, 0xaf, 0x53, 0x91, 0x2c, 0x50, 0x9b, 0xd7, 0x94, 0x5a, 0xb2, 0x0c,
	0x56, 0xb7, 0x61, 0x16, 0x02, 0x5b, 0xe6, 0x69, 0x96, 0x54, 0x69, 0xda, 0x44, 0xf3, 0xf5, 0x11,
	0x9a, 0x57, 0x36, 0xb1, 0xca, 0xd1, 0x5b, 0x1e, 0x60, 0x6c, 0x09, 0x03, 0x95, 0x0b, 0x81, 0xaa,
	0xca, 0xe0, 0x12, 0xc3, 0x37, 0x47, 0x18, 0x38, 0x65, 0x56, 0x29, 0xfa, 0xea, 0x10, 0xb4, 0x7a,
	0xd9, 0x81, 0xb3, 0x25, 0x74, 0x95, 0xa6, 0x73, 0x2f, 0xbd, 0x78, 0x2e, 0xac, 0x2e, 0xfa, 0x40,
	0x2f, 0x75, 0x1b, 0xb6, 0x7a, 0x65, 0x32, 0x8e, 0xb2, 0x30, 0x41, 0x83, 0x71, 0xf9, 0xee, 0xde,
	0xbd, 0xf4, 0x9a, 0xdb, 0xc4, 0x73, 0xca, 0xdb, 0xeb, 0x95, 0x1d, 0x60, 0xc1, 0x9f, 0x70, 0x7a,
	0x30, 0x7c, 0xec, 0x33, 0x38, 0x89, 0xe5, 0x66, 0x13, 0x89, 0x24, 0xcc, 0x52, 0x81, 0x85, 0xdd,
	0xbc, 0x02, 0x9b, 0xa7, 0x02, 0xd9, 0x97, 0x70, 0x4a, 0xc3, 0xa2, 0x70, 0x85, 0x0a, 0x45, 0x8c,
	0xd6, 0xc7, 0xf5, 0x51, 0x87, 0x77, 0x2d, 0xcc, 0xf7, 0x28, 0xeb, 0x41, 0x3d, 0xfe, 0x23, 0x21,
	0xfb, 0x75, 0xb8, 0x3d, 0xb2, 0x2e, 0xd4, 0xa4, 0x26, 0x5b, 0x74, 0x78, 0x4d, 0xea, 0xe0, 0x2f,
	0x07, 0xba, 0xd5, 0x09, 0x65, 0x2f, 0xc0, 0x4d, 0x8a, 0x0d, 0x42, 0xe4, 0xde, 0x0d, 0xe3, 0x96,
	0xab, 0x85, 0xef, 0xaf, 0x58, 0xe3, 0x92, 0x27, 0x62, 0x99, 0x20, 0xad, 0x80, 0x26, 0x77, 0x2d,
	0xf0, 0xbd, 0x4c, 0x90, 0x7d, 0x02, 0xa0, 0x4d, 0x14, 0x5f, 0x1a, 0x15, 0xc5, 0x58, 0xf4, 0x71,
	0x03, 0x61, 0x0c, 0x1a, 0x99, 0x5c, 0x97, 0x0d, 0xd1, 0x39, 0xf8, 0xc7, 0x81, 0xc1, 0x1d, 0x33,
	0xfd, 0xd0, 0xbe, 0xec, 0xf2, 0xcb, 0xe3, 0x18, 0xb5, 0xa6, 0xae, 0x5c, 0x5e, 0x86, 0xec, 0x0b,
	0xe8, 0x56, 0xe5, 0x2b, 0x1a, 0x7b, 0x5a, 0x51, 0xcf, 0x3e, 0x44, 0x69, 0xc9, 0x04, 0x95, 0xed,
	0xd1, 0x4a, 0xec, 0x99, 0x5d, 0x5b, 0x16, 0xa2, 0xa5, 0x65, 0xaf, 0xa4, 0x62, 0x25, 0x69, 0x1d,
	0xd8, 0xa5, 0x85, 0xda, 0xcc, 0xc4, 0x4a, 0x06, 0x39, 0xf4, 0x0e, 0x3d, 0xf3, 0x78, 0xff, 0xc1,
	0x87, 0xb6, 0x89, 0xd4, 0x1a, 0x8d, 0xf6, 0xeb, 0xd4, 0x57, 0x19, 0x06, 0x7f, 0xd7, 0xa0, 0x7f,
	0xcb, 0x49, 0x8f, 0x47, 0xfc, 0x29, 0x78, 0x85, 0xc3, 0x69, 0x11, 0x17, 0x4f, 0xba, 0x83, 0x68,
	0x0f, 0xcf, 0xa0, 0x61, 0x25, 0x20, 0xb9, 0xbc, 0xc9, 0xb7, 0x0f, 0xf5, 0xfc, 0xd8, 0x3e, 0x3b,
	0xa7, 0x12, 0x01, 0x87, 0x86, 0x8d, 0xec, 0x94, 0x10, 0xd9, 0xce, 0x0a, 0x74, 0x66, 0xcf, 0xa1,
	0xa5, 0x50, 0xe7, 0x99, 0x29, 0x66, 0xae, 0x88, 0x8e, 0x4d, 0x5c, 0x70, 0x0e, 0x83, 0x3b, 0x16,
	0xc0, 0x03, 0xf5, 0x09, 0xde, 0x42, 0xef, 0xd0, 0xdf, 0x56, 0x99, 0x62, 0x4d, 0xd0, 0x37, 0xcf,
	0xa1, 0xb6, 0x60, 0x07, 0xd1, 0x27, 0xee, 0x3d, 0x2d, 0xbf, 0x6a, 0x43, 0x93, 0x16, 0xcb, 0x57,
	0x9f, 0x83, 0x5b, 0x7e, 0x19, 0x19, 0x40, 0xeb, 0xcd, 0xc5, 0xc5, 0x9b, 0xf9, 0xb4, 0xf7, 0x84,
	0x9d, 0x80, 0x3b, 0xfd, 0x79, 0x31, 0xe5, 0xef, 0x5e, 0xce, 0x7b, 0xce, 0x7f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x0b, 0xd9, 0xb8, 0x91, 0x04, 0x08, 0x00, 0x00,
}
