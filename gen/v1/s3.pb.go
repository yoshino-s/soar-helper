// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: v1/s3.proto

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

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Key  string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_s3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_s3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_v1_s3_proto_rawDescGZIP(), []int{0}
}

func (x *UploadRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *UploadRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_s3_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_s3_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_v1_s3_proto_rawDescGZIP(), []int{1}
}

func (x *UploadResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_v1_s3_proto protoreflect.FileDescriptor

var file_v1_s3_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x31, 0x2f, 0x73, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x79,
	0x6f, 0x73, 0x68, 0x69, 0x6e, 0x6f, 0x5f, 0x73, 0x2e, 0x73, 0x6f, 0x61, 0x72, 0x5f, 0x68, 0x65,
	0x6c, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x35, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x22,
	0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x32, 0x6a, 0x0a, 0x09, 0x53, 0x33, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5d, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x27, 0x2e, 0x79, 0x6f, 0x73, 0x68,
	0x69, 0x6e, 0x6f, 0x5f, 0x73, 0x2e, 0x73, 0x6f, 0x61, 0x72, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6e, 0x6f, 0x5f, 0x73, 0x2e, 0x73,
	0x6f, 0x61, 0x72, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x33,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6e, 0x6f,
	0x2d, 0x73, 0x2e, 0x78, 0x79, 0x7a, 0x2f, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6e, 0x6f, 0x2d, 0x73,
	0x2f, 0x73, 0x6f, 0x61, 0x72, 0x2d, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_s3_proto_rawDescOnce sync.Once
	file_v1_s3_proto_rawDescData = file_v1_s3_proto_rawDesc
)

func file_v1_s3_proto_rawDescGZIP() []byte {
	file_v1_s3_proto_rawDescOnce.Do(func() {
		file_v1_s3_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_s3_proto_rawDescData)
	})
	return file_v1_s3_proto_rawDescData
}

var file_v1_s3_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_s3_proto_goTypes = []any{
	(*UploadRequest)(nil),  // 0: yoshino_s.soar_helper.v1.UploadRequest
	(*UploadResponse)(nil), // 1: yoshino_s.soar_helper.v1.UploadResponse
}
var file_v1_s3_proto_depIdxs = []int32{
	0, // 0: yoshino_s.soar_helper.v1.S3Service.Upload:input_type -> yoshino_s.soar_helper.v1.UploadRequest
	1, // 1: yoshino_s.soar_helper.v1.S3Service.Upload:output_type -> yoshino_s.soar_helper.v1.UploadResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_s3_proto_init() }
func file_v1_s3_proto_init() {
	if File_v1_s3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_s3_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UploadRequest); i {
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
		file_v1_s3_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UploadResponse); i {
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
			RawDescriptor: file_v1_s3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_s3_proto_goTypes,
		DependencyIndexes: file_v1_s3_proto_depIdxs,
		MessageInfos:      file_v1_s3_proto_msgTypes,
	}.Build()
	File_v1_s3_proto = out.File
	file_v1_s3_proto_rawDesc = nil
	file_v1_s3_proto_goTypes = nil
	file_v1_s3_proto_depIdxs = nil
}
