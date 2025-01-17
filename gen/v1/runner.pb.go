// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: v1/runner.proto

package v1

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

type RunStreamData_Type int32

const (
	RunStreamData_Type_STDOUT    RunStreamData_Type = 0
	RunStreamData_Type_STDERR    RunStreamData_Type = 1
	RunStreamData_Type_EXIT_CODE RunStreamData_Type = 2
)

// Enum value maps for RunStreamData_Type.
var (
	RunStreamData_Type_name = map[int32]string{
		0: "STDOUT",
		1: "STDERR",
		2: "EXIT_CODE",
	}
	RunStreamData_Type_value = map[string]int32{
		"STDOUT":    0,
		"STDERR":    1,
		"EXIT_CODE": 2,
	}
)

func (x RunStreamData_Type) Enum() *RunStreamData_Type {
	p := new(RunStreamData_Type)
	*p = x
	return p
}

func (x RunStreamData_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RunStreamData_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_runner_proto_enumTypes[0].Descriptor()
}

func (RunStreamData_Type) Type() protoreflect.EnumType {
	return &file_v1_runner_proto_enumTypes[0]
}

func (x RunStreamData_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RunStreamData_Type.Descriptor instead.
func (RunStreamData_Type) EnumDescriptor() ([]byte, []int) {
	return file_v1_runner_proto_rawDescGZIP(), []int{0}
}

type ReadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ReadFileRequest) Reset() {
	*x = ReadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_runner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadFileRequest) ProtoMessage() {}

func (x *ReadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_runner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadFileRequest.ProtoReflect.Descriptor instead.
func (*ReadFileRequest) Descriptor() ([]byte, []int) {
	return file_v1_runner_proto_rawDescGZIP(), []int{0}
}

func (x *ReadFileRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type ReadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ReadFileResponse) Reset() {
	*x = ReadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_runner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadFileResponse) ProtoMessage() {}

func (x *ReadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_runner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadFileResponse.ProtoReflect.Descriptor instead.
func (*ReadFileResponse) Descriptor() ([]byte, []int) {
	return file_v1_runner_proto_rawDescGZIP(), []int{1}
}

func (x *ReadFileResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type WriteFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path    string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *WriteFileRequest) Reset() {
	*x = WriteFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_runner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteFileRequest) ProtoMessage() {}

func (x *WriteFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_runner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteFileRequest.ProtoReflect.Descriptor instead.
func (*WriteFileRequest) Descriptor() ([]byte, []int) {
	return file_v1_runner_proto_rawDescGZIP(), []int{2}
}

func (x *WriteFileRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *WriteFileRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type WriteFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WriteFileResponse) Reset() {
	*x = WriteFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_runner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteFileResponse) ProtoMessage() {}

func (x *WriteFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_runner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteFileResponse.ProtoReflect.Descriptor instead.
func (*WriteFileResponse) Descriptor() ([]byte, []int) {
	return file_v1_runner_proto_rawDescGZIP(), []int{3}
}

type RunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commands []string `protobuf:"bytes,1,rep,name=commands,proto3" json:"commands,omitempty"`
	Stdin    string   `protobuf:"bytes,2,opt,name=stdin,proto3" json:"stdin,omitempty"`
}

func (x *RunRequest) Reset() {
	*x = RunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_runner_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunRequest) ProtoMessage() {}

func (x *RunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_runner_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunRequest.ProtoReflect.Descriptor instead.
func (*RunRequest) Descriptor() ([]byte, []int) {
	return file_v1_runner_proto_rawDescGZIP(), []int{4}
}

func (x *RunRequest) GetCommands() []string {
	if x != nil {
		return x.Commands
	}
	return nil
}

func (x *RunRequest) GetStdin() string {
	if x != nil {
		return x.Stdin
	}
	return ""
}

type RunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stdout   string `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr   string `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	ExitCode int32  `protobuf:"varint,3,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
}

func (x *RunResponse) Reset() {
	*x = RunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_runner_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunResponse) ProtoMessage() {}

func (x *RunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_runner_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunResponse.ProtoReflect.Descriptor instead.
func (*RunResponse) Descriptor() ([]byte, []int) {
	return file_v1_runner_proto_rawDescGZIP(), []int{5}
}

func (x *RunResponse) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *RunResponse) GetStderr() string {
	if x != nil {
		return x.Stderr
	}
	return ""
}

func (x *RunResponse) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

type RunStreamData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     RunStreamData_Type `protobuf:"varint,1,opt,name=type,proto3,enum=yoshino_s.soar_helper.v1.RunStreamData_Type" json:"type,omitempty"`
	Data     string             `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	ExitCode int32              `protobuf:"varint,3,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
}

func (x *RunStreamData) Reset() {
	*x = RunStreamData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_runner_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunStreamData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunStreamData) ProtoMessage() {}

func (x *RunStreamData) ProtoReflect() protoreflect.Message {
	mi := &file_v1_runner_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunStreamData.ProtoReflect.Descriptor instead.
func (*RunStreamData) Descriptor() ([]byte, []int) {
	return file_v1_runner_proto_rawDescGZIP(), []int{6}
}

func (x *RunStreamData) GetType() RunStreamData_Type {
	if x != nil {
		return x.Type
	}
	return RunStreamData_Type_STDOUT
}

func (x *RunStreamData) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *RunStreamData) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

var File_v1_runner_proto protoreflect.FileDescriptor

var file_v1_runner_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x18, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6e, 0x6f, 0x5f, 0x73, 0x2e, 0x73, 0x6f, 0x61,
	0x72, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x25, 0x0a, 0x0f, 0x52,
	0x65, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x22, 0x26, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x40, 0x0a, 0x10, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x13, 0x0a, 0x11,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x3e, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x64, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x64, 0x69,
	0x6e, 0x22, 0x5a, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x65,
	0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x82, 0x01,
	0x0a, 0x0d, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x40, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e,
	0x79, 0x6f, 0x73, 0x68, 0x69, 0x6e, 0x6f, 0x5f, 0x73, 0x2e, 0x73, 0x6f, 0x61, 0x72, 0x5f, 0x68,
	0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x5f, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x2a, 0x3b, 0x0a, 0x12, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44,
	0x61, 0x74, 0x61, 0x5f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x44, 0x4f,
	0x55, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x44, 0x45, 0x52, 0x52, 0x10, 0x01,
	0x12, 0x0d, 0x0a, 0x09, 0x45, 0x58, 0x49, 0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x02, 0x32,
	0x92, 0x03, 0x0a, 0x0d, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x54, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x24, 0x2e, 0x79, 0x6f, 0x73, 0x68, 0x69,
	0x6e, 0x6f, 0x5f, 0x73, 0x2e, 0x73, 0x6f, 0x61, 0x72, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6e, 0x6f, 0x5f, 0x73, 0x2e, 0x73, 0x6f, 0x61, 0x72, 0x5f,
	0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x09, 0x52, 0x75, 0x6e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x24, 0x2e, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6e, 0x6f, 0x5f, 0x73,
	0x2e, 0x73, 0x6f, 0x61, 0x72, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x79, 0x6f, 0x73,
	0x68, 0x69, 0x6e, 0x6f, 0x5f, 0x73, 0x2e, 0x73, 0x6f, 0x61, 0x72, 0x5f, 0x68, 0x65, 0x6c, 0x70,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x00, 0x30, 0x01, 0x12, 0x63, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x29, 0x2e, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6e, 0x6f, 0x5f, 0x73, 0x2e,
	0x73, 0x6f, 0x61, 0x72, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6e, 0x6f, 0x5f, 0x73, 0x2e, 0x73, 0x6f, 0x61, 0x72, 0x5f,
	0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x09,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2a, 0x2e, 0x79, 0x6f, 0x73, 0x68,
	0x69, 0x6e, 0x6f, 0x5f, 0x73, 0x2e, 0x73, 0x6f, 0x61, 0x72, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6e, 0x6f, 0x5f,
	0x73, 0x2e, 0x73, 0x6f, 0x61, 0x72, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6e, 0x6f, 0x2d, 0x73, 0x2f, 0x73, 0x6f, 0x61,
	0x72, 0x2d, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_runner_proto_rawDescOnce sync.Once
	file_v1_runner_proto_rawDescData = file_v1_runner_proto_rawDesc
)

func file_v1_runner_proto_rawDescGZIP() []byte {
	file_v1_runner_proto_rawDescOnce.Do(func() {
		file_v1_runner_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_runner_proto_rawDescData)
	})
	return file_v1_runner_proto_rawDescData
}

var file_v1_runner_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_runner_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v1_runner_proto_goTypes = []any{
	(RunStreamData_Type)(0),   // 0: yoshino_s.soar_helper.v1.RunStreamData_Type
	(*ReadFileRequest)(nil),   // 1: yoshino_s.soar_helper.v1.ReadFileRequest
	(*ReadFileResponse)(nil),  // 2: yoshino_s.soar_helper.v1.ReadFileResponse
	(*WriteFileRequest)(nil),  // 3: yoshino_s.soar_helper.v1.WriteFileRequest
	(*WriteFileResponse)(nil), // 4: yoshino_s.soar_helper.v1.WriteFileResponse
	(*RunRequest)(nil),        // 5: yoshino_s.soar_helper.v1.RunRequest
	(*RunResponse)(nil),       // 6: yoshino_s.soar_helper.v1.RunResponse
	(*RunStreamData)(nil),     // 7: yoshino_s.soar_helper.v1.RunStreamData
}
var file_v1_runner_proto_depIdxs = []int32{
	0, // 0: yoshino_s.soar_helper.v1.RunStreamData.type:type_name -> yoshino_s.soar_helper.v1.RunStreamData_Type
	5, // 1: yoshino_s.soar_helper.v1.RunnerService.Run:input_type -> yoshino_s.soar_helper.v1.RunRequest
	5, // 2: yoshino_s.soar_helper.v1.RunnerService.RunStream:input_type -> yoshino_s.soar_helper.v1.RunRequest
	1, // 3: yoshino_s.soar_helper.v1.RunnerService.ReadFile:input_type -> yoshino_s.soar_helper.v1.ReadFileRequest
	3, // 4: yoshino_s.soar_helper.v1.RunnerService.WriteFile:input_type -> yoshino_s.soar_helper.v1.WriteFileRequest
	6, // 5: yoshino_s.soar_helper.v1.RunnerService.Run:output_type -> yoshino_s.soar_helper.v1.RunResponse
	7, // 6: yoshino_s.soar_helper.v1.RunnerService.RunStream:output_type -> yoshino_s.soar_helper.v1.RunStreamData
	2, // 7: yoshino_s.soar_helper.v1.RunnerService.ReadFile:output_type -> yoshino_s.soar_helper.v1.ReadFileResponse
	4, // 8: yoshino_s.soar_helper.v1.RunnerService.WriteFile:output_type -> yoshino_s.soar_helper.v1.WriteFileResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_runner_proto_init() }
func file_v1_runner_proto_init() {
	if File_v1_runner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_runner_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ReadFileRequest); i {
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
		file_v1_runner_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ReadFileResponse); i {
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
		file_v1_runner_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*WriteFileRequest); i {
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
		file_v1_runner_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*WriteFileResponse); i {
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
		file_v1_runner_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*RunRequest); i {
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
		file_v1_runner_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*RunResponse); i {
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
		file_v1_runner_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*RunStreamData); i {
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
			RawDescriptor: file_v1_runner_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_runner_proto_goTypes,
		DependencyIndexes: file_v1_runner_proto_depIdxs,
		EnumInfos:         file_v1_runner_proto_enumTypes,
		MessageInfos:      file_v1_runner_proto_msgTypes,
	}.Build()
	File_v1_runner_proto = out.File
	file_v1_runner_proto_rawDesc = nil
	file_v1_runner_proto_goTypes = nil
	file_v1_runner_proto_depIdxs = nil
}
