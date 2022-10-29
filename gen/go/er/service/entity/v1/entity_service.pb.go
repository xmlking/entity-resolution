// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: er/service/entity/v1/entity_service.proto

// provides api to create, read, update and delete policies.

package entityv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v1 "github.com/xmlking/entity-resolution/gen/go/er/schema/entity/v1"
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

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    *string   `protobuf:"bytes,1,opt,name=key,proto3,oneof" json:"key,omitempty"`
	Score  *uint32   `protobuf:"varint,2,opt,name=score,proto3,oneof" json:"score,omitempty"`
	Status v1.Status `protobuf:"varint,3,opt,name=status,proto3,enum=er.schema.entity.v1.Status" json:"status,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_er_service_entity_v1_entity_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_er_service_entity_v1_entity_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_er_service_entity_v1_entity_service_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *Response) GetScore() uint32 {
	if x != nil && x.Score != nil {
		return *x.Score
	}
	return 0
}

func (x *Response) GetStatus() v1.Status {
	if x != nil {
		return x.Status
	}
	return v1.Status(0)
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_er_service_entity_v1_entity_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_er_service_entity_v1_entity_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_er_service_entity_v1_entity_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search  *string `protobuf:"bytes,1,opt,name=search,proto3,oneof" json:"search,omitempty"`
	Filter  *string `protobuf:"bytes,2,opt,name=filter,proto3,oneof" json:"filter,omitempty"`
	Select  *string `protobuf:"bytes,3,opt,name=select,proto3,oneof" json:"select,omitempty"`
	OrderBy *string `protobuf:"bytes,4,opt,name=order_by,json=orderBy,proto3,oneof" json:"order_by,omitempty"`
	// specifically limit the page size in a request. Number of results to return per page.
	Top *uint32 `protobuf:"varint,5,opt,name=top,proto3,oneof" json:"top,omitempty"`
	// $skip is used to page through a result set.
	Skip *string `protobuf:"bytes,6,opt,name=skip,proto3,oneof" json:"skip,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_er_service_entity_v1_entity_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_er_service_entity_v1_entity_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_er_service_entity_v1_entity_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListRequest) GetSearch() string {
	if x != nil && x.Search != nil {
		return *x.Search
	}
	return ""
}

func (x *ListRequest) GetFilter() string {
	if x != nil && x.Filter != nil {
		return *x.Filter
	}
	return ""
}

func (x *ListRequest) GetSelect() string {
	if x != nil && x.Select != nil {
		return *x.Select
	}
	return ""
}

func (x *ListRequest) GetOrderBy() string {
	if x != nil && x.OrderBy != nil {
		return *x.OrderBy
	}
	return ""
}

func (x *ListRequest) GetTop() uint32 {
	if x != nil && x.Top != nil {
		return *x.Top
	}
	return 0
}

func (x *ListRequest) GetSkip() string {
	if x != nil && x.Skip != nil {
		return *x.Skip
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*v1.Member `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   uint32       `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	// $skip is used to page through a result set.
	Skip *string `protobuf:"bytes,3,opt,name=skip,proto3,oneof" json:"skip,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_er_service_entity_v1_entity_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_er_service_entity_v1_entity_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_er_service_entity_v1_entity_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListResponse) GetResults() []*v1.Member {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ListResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListResponse) GetSkip() string {
	if x != nil && x.Skip != nil {
		return *x.Skip
	}
	return ""
}

var File_er_service_entity_v1_entity_service_proto protoreflect.FileDescriptor

var file_er_service_entity_v1_entity_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x31, 0x1a, 0x20, 0x65, 0x72, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x88, 0x01, 0x01,
	0x12, 0x19, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48,
	0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x65, 0x72,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6b, 0x65, 0x79, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xfe, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x06, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x79, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x18, 0x64, 0x28, 0x01, 0x48,
	0x04, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x73, 0x6b, 0x69,
	0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x88,
	0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79,
	0x42, 0x06, 0x0a, 0x04, 0x5f, 0x74, 0x6f, 0x70, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x6b, 0x69,
	0x70, 0x22, 0x82, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x72, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x17, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x04, 0x73, 0x6b, 0x69, 0x70, 0x88, 0x01, 0x01, 0x3a, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x32, 0xb3, 0x02, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x49, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x2e, 0x65, 0x72, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x1a,
	0x1e, 0x2e, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x46, 0x0a, 0x07, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x12, 0x1b, 0x2e, 0x65, 0x72, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x1e, 0x2e, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x20,
	0x2e, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x65, 0x72, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x4d, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xec, 0x01, 0x0a,
	0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x6d, 0x6c, 0x6b,
	0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2d, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x72, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76,
	0x31, 0x3b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x53, 0x45,
	0xaa, 0x02, 0x14, 0x45, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x14, 0x45, 0x72, 0x5c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x20, 0x45, 0x72, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x17, 0x45, 0x72, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a,
	0x3a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_er_service_entity_v1_entity_service_proto_rawDescOnce sync.Once
	file_er_service_entity_v1_entity_service_proto_rawDescData = file_er_service_entity_v1_entity_service_proto_rawDesc
)

func file_er_service_entity_v1_entity_service_proto_rawDescGZIP() []byte {
	file_er_service_entity_v1_entity_service_proto_rawDescOnce.Do(func() {
		file_er_service_entity_v1_entity_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_er_service_entity_v1_entity_service_proto_rawDescData)
	})
	return file_er_service_entity_v1_entity_service_proto_rawDescData
}

var file_er_service_entity_v1_entity_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_er_service_entity_v1_entity_service_proto_goTypes = []interface{}{
	(*Response)(nil),     // 0: er.service.entity.v1.Response
	(*GetRequest)(nil),   // 1: er.service.entity.v1.GetRequest
	(*ListRequest)(nil),  // 2: er.service.entity.v1.ListRequest
	(*ListResponse)(nil), // 3: er.service.entity.v1.ListResponse
	(v1.Status)(0),       // 4: er.schema.entity.v1.Status
	(*v1.Member)(nil),    // 5: er.schema.entity.v1.Member
}
var file_er_service_entity_v1_entity_service_proto_depIdxs = []int32{
	4, // 0: er.service.entity.v1.Response.status:type_name -> er.schema.entity.v1.Status
	5, // 1: er.service.entity.v1.ListResponse.results:type_name -> er.schema.entity.v1.Member
	5, // 2: er.service.entity.v1.EntityService.Ingest:input_type -> er.schema.entity.v1.Member
	5, // 3: er.service.entity.v1.EntityService.Inquiry:input_type -> er.schema.entity.v1.Member
	1, // 4: er.service.entity.v1.EntityService.Get:input_type -> er.service.entity.v1.GetRequest
	2, // 5: er.service.entity.v1.EntityService.List:input_type -> er.service.entity.v1.ListRequest
	0, // 6: er.service.entity.v1.EntityService.Ingest:output_type -> er.service.entity.v1.Response
	0, // 7: er.service.entity.v1.EntityService.Inquiry:output_type -> er.service.entity.v1.Response
	5, // 8: er.service.entity.v1.EntityService.Get:output_type -> er.schema.entity.v1.Member
	3, // 9: er.service.entity.v1.EntityService.List:output_type -> er.service.entity.v1.ListResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_er_service_entity_v1_entity_service_proto_init() }
func file_er_service_entity_v1_entity_service_proto_init() {
	if File_er_service_entity_v1_entity_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_er_service_entity_v1_entity_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_er_service_entity_v1_entity_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_er_service_entity_v1_entity_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_er_service_entity_v1_entity_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
	file_er_service_entity_v1_entity_service_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_er_service_entity_v1_entity_service_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_er_service_entity_v1_entity_service_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_er_service_entity_v1_entity_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_er_service_entity_v1_entity_service_proto_goTypes,
		DependencyIndexes: file_er_service_entity_v1_entity_service_proto_depIdxs,
		MessageInfos:      file_er_service_entity_v1_entity_service_proto_msgTypes,
	}.Build()
	File_er_service_entity_v1_entity_service_proto = out.File
	file_er_service_entity_v1_entity_service_proto_rawDesc = nil
	file_er_service_entity_v1_entity_service_proto_goTypes = nil
	file_er_service_entity_v1_entity_service_proto_depIdxs = nil
}
