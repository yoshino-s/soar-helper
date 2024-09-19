// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: v1/icp.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_icp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_icp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_v1_icp_proto_rawDescGZIP(), []int{0}
}

func (x *QueryRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

type BatchQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hosts []string `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
}

func (x *BatchQueryRequest) Reset() {
	*x = BatchQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_icp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchQueryRequest) ProtoMessage() {}

func (x *BatchQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_icp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchQueryRequest.ProtoReflect.Descriptor instead.
func (*BatchQueryRequest) Descriptor() ([]byte, []int) {
	return file_v1_icp_proto_rawDescGZIP(), []int{1}
}

func (x *BatchQueryRequest) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

type QueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *IcpRecord `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_icp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_icp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResponse.ProtoReflect.Descriptor instead.
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return file_v1_icp_proto_rawDescGZIP(), []int{2}
}

func (x *QueryResponse) GetData() *IcpRecord {
	if x != nil {
		return x.Data
	}
	return nil
}

type BatchQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*IcpRecord `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *BatchQueryResponse) Reset() {
	*x = BatchQueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_icp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchQueryResponse) ProtoMessage() {}

func (x *BatchQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_icp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchQueryResponse.ProtoReflect.Descriptor instead.
func (*BatchQueryResponse) Descriptor() ([]byte, []int) {
	return file_v1_icp_proto_rawDescGZIP(), []int{3}
}

func (x *BatchQueryResponse) GetData() []*IcpRecord {
	if x != nil {
		return x.Data
	}
	return nil
}

type IcpRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Host      string                 `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	City      string                 `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Province  string                 `protobuf:"bytes,4,opt,name=province,proto3" json:"province,omitempty"`
	Company   string                 `protobuf:"bytes,5,opt,name=company,proto3" json:"company,omitempty"`
	Owner     string                 `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	Type      string                 `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	Homepage  string                 `protobuf:"bytes,8,opt,name=homepage,proto3" json:"homepage,omitempty"`
	Permit    string                 `protobuf:"bytes,9,opt,name=permit,proto3" json:"permit,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Cached    bool                   `protobuf:"varint,12,opt,name=cached,proto3" json:"cached,omitempty"`
	Input     string                 `protobuf:"bytes,13,opt,name=input,proto3" json:"input,omitempty"`
	WebName   string                 `protobuf:"bytes,14,opt,name=web_name,json=webName,proto3" json:"web_name,omitempty"`
}

func (x *IcpRecord) Reset() {
	*x = IcpRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_icp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IcpRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IcpRecord) ProtoMessage() {}

func (x *IcpRecord) ProtoReflect() protoreflect.Message {
	mi := &file_v1_icp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IcpRecord.ProtoReflect.Descriptor instead.
func (*IcpRecord) Descriptor() ([]byte, []int) {
	return file_v1_icp_proto_rawDescGZIP(), []int{4}
}

func (x *IcpRecord) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IcpRecord) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *IcpRecord) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *IcpRecord) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *IcpRecord) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *IcpRecord) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *IcpRecord) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *IcpRecord) GetHomepage() string {
	if x != nil {
		return x.Homepage
	}
	return ""
}

func (x *IcpRecord) GetPermit() string {
	if x != nil {
		return x.Permit
	}
	return ""
}

func (x *IcpRecord) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *IcpRecord) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *IcpRecord) GetCached() bool {
	if x != nil {
		return x.Cached
	}
	return false
}

func (x *IcpRecord) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *IcpRecord) GetWebName() string {
	if x != nil {
		return x.WebName
	}
	return ""
}

type StatisticResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *StatisticResponse) Reset() {
	*x = StatisticResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_icp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatisticResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatisticResponse) ProtoMessage() {}

func (x *StatisticResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_icp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatisticResponse.ProtoReflect.Descriptor instead.
func (*StatisticResponse) Descriptor() ([]byte, []int) {
	return file_v1_icp_proto_rawDescGZIP(), []int{5}
}

func (x *StatisticResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_v1_icp_proto protoreflect.FileDescriptor

var file_v1_icp_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x76, 0x31, 0x2f, 0x69, 0x63, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18,
	0x79, 0x6f, 0x73, 0x68, 0x69, 0x6e, 0x6f, 0x5f, 0x73, 0x2e, 0x73, 0x6f, 0x61, 0x72, 0x5f, 0x68,
	0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x22, 0x29, 0x0a, 0x11, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x68, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x48, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6e, 0x6f, 0x5f, 0x73,
	0x2e, 0x73, 0x6f, 0x61, 0x72, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x63, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x4d, 0x0a, 0x12, 0x42, 0x61, 0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6e, 0x6f, 0x5f, 0x73, 0x2e,
	0x73, 0x6f, 0x61, 0x72, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x63, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x96,
	0x03, 0x0a, 0x09, 0x49, 0x63, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x77, 0x65, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x77, 0x65, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x32, 0xac, 0x02, 0x0a, 0x0f, 0x49, 0x63, 0x70, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x26, 0x2e, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6e, 0x6f, 0x5f, 0x73, 0x2e, 0x73, 0x6f, 0x61, 0x72,
	0x5f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6e,
	0x6f, 0x5f, 0x73, 0x2e, 0x73, 0x6f, 0x61, 0x72, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x69, 0x0a, 0x0a, 0x42, 0x61, 0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x2b, 0x2e, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6e, 0x6f, 0x5f, 0x73, 0x2e, 0x73, 0x6f, 0x61,
	0x72, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x79, 0x6f, 0x73, 0x68, 0x69, 0x6e, 0x6f, 0x5f, 0x73, 0x2e, 0x73, 0x6f, 0x61, 0x72, 0x5f, 0x68,
	0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a,
	0x09, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x2b, 0x2e, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6e, 0x6f, 0x5f, 0x73, 0x2e, 0x73,
	0x6f, 0x61, 0x72, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x79, 0x6f, 0x73, 0x68,
	0x69, 0x6e, 0x6f, 0x2d, 0x73, 0x2e, 0x78, 0x79, 0x7a, 0x2f, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6e,
	0x6f, 0x2d, 0x73, 0x2f, 0x73, 0x6f, 0x61, 0x72, 0x2d, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_icp_proto_rawDescOnce sync.Once
	file_v1_icp_proto_rawDescData = file_v1_icp_proto_rawDesc
)

func file_v1_icp_proto_rawDescGZIP() []byte {
	file_v1_icp_proto_rawDescOnce.Do(func() {
		file_v1_icp_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_icp_proto_rawDescData)
	})
	return file_v1_icp_proto_rawDescData
}

var file_v1_icp_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_icp_proto_goTypes = []any{
	(*QueryRequest)(nil),          // 0: yoshino_s.soar_helper.v1.QueryRequest
	(*BatchQueryRequest)(nil),     // 1: yoshino_s.soar_helper.v1.BatchQueryRequest
	(*QueryResponse)(nil),         // 2: yoshino_s.soar_helper.v1.QueryResponse
	(*BatchQueryResponse)(nil),    // 3: yoshino_s.soar_helper.v1.BatchQueryResponse
	(*IcpRecord)(nil),             // 4: yoshino_s.soar_helper.v1.IcpRecord
	(*StatisticResponse)(nil),     // 5: yoshino_s.soar_helper.v1.StatisticResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_v1_icp_proto_depIdxs = []int32{
	4, // 0: yoshino_s.soar_helper.v1.QueryResponse.data:type_name -> yoshino_s.soar_helper.v1.IcpRecord
	4, // 1: yoshino_s.soar_helper.v1.BatchQueryResponse.data:type_name -> yoshino_s.soar_helper.v1.IcpRecord
	6, // 2: yoshino_s.soar_helper.v1.IcpRecord.created_at:type_name -> google.protobuf.Timestamp
	6, // 3: yoshino_s.soar_helper.v1.IcpRecord.updated_at:type_name -> google.protobuf.Timestamp
	0, // 4: yoshino_s.soar_helper.v1.IcpQueryService.Query:input_type -> yoshino_s.soar_helper.v1.QueryRequest
	1, // 5: yoshino_s.soar_helper.v1.IcpQueryService.BatchQuery:input_type -> yoshino_s.soar_helper.v1.BatchQueryRequest
	7, // 6: yoshino_s.soar_helper.v1.IcpQueryService.Statistic:input_type -> google.protobuf.Empty
	2, // 7: yoshino_s.soar_helper.v1.IcpQueryService.Query:output_type -> yoshino_s.soar_helper.v1.QueryResponse
	3, // 8: yoshino_s.soar_helper.v1.IcpQueryService.BatchQuery:output_type -> yoshino_s.soar_helper.v1.BatchQueryResponse
	5, // 9: yoshino_s.soar_helper.v1.IcpQueryService.Statistic:output_type -> yoshino_s.soar_helper.v1.StatisticResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_v1_icp_proto_init() }
func file_v1_icp_proto_init() {
	if File_v1_icp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_icp_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*QueryRequest); i {
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
		file_v1_icp_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BatchQueryRequest); i {
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
		file_v1_icp_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*QueryResponse); i {
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
		file_v1_icp_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*BatchQueryResponse); i {
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
		file_v1_icp_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*IcpRecord); i {
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
		file_v1_icp_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*StatisticResponse); i {
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
			RawDescriptor: file_v1_icp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_icp_proto_goTypes,
		DependencyIndexes: file_v1_icp_proto_depIdxs,
		MessageInfos:      file_v1_icp_proto_msgTypes,
	}.Build()
	File_v1_icp_proto = out.File
	file_v1_icp_proto_rawDesc = nil
	file_v1_icp_proto_goTypes = nil
	file_v1_icp_proto_depIdxs = nil
}
