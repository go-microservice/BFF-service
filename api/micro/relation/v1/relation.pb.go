// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.3
// source: api/micro/relation/v1/relation.proto

package v1

import (
	v1 "github.com/go-microservice/ins-api/api/micro/user/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type FollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FollowedUid int64 `protobuf:"varint,2,opt,name=followed_uid,json=followedUid,proto3" json:"followed_uid,omitempty"`
}

func (x *FollowRequest) Reset() {
	*x = FollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_micro_relation_v1_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowRequest) ProtoMessage() {}

func (x *FollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_micro_relation_v1_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowRequest.ProtoReflect.Descriptor instead.
func (*FollowRequest) Descriptor() ([]byte, []int) {
	return file_api_micro_relation_v1_relation_proto_rawDescGZIP(), []int{0}
}

func (x *FollowRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FollowRequest) GetFollowedUid() int64 {
	if x != nil {
		return x.FollowedUid
	}
	return 0
}

type FollowReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FollowReply) Reset() {
	*x = FollowReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_micro_relation_v1_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowReply) ProtoMessage() {}

func (x *FollowReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_micro_relation_v1_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowReply.ProtoReflect.Descriptor instead.
func (*FollowReply) Descriptor() ([]byte, []int) {
	return file_api_micro_relation_v1_relation_proto_rawDescGZIP(), []int{1}
}

type UnfollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FollowedUid int64 `protobuf:"varint,2,opt,name=followed_uid,json=followedUid,proto3" json:"followed_uid,omitempty"`
}

func (x *UnfollowRequest) Reset() {
	*x = UnfollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_micro_relation_v1_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnfollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnfollowRequest) ProtoMessage() {}

func (x *UnfollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_micro_relation_v1_relation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnfollowRequest.ProtoReflect.Descriptor instead.
func (*UnfollowRequest) Descriptor() ([]byte, []int) {
	return file_api_micro_relation_v1_relation_proto_rawDescGZIP(), []int{2}
}

func (x *UnfollowRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UnfollowRequest) GetFollowedUid() int64 {
	if x != nil {
		return x.FollowedUid
	}
	return 0
}

type UnfollowReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnfollowReply) Reset() {
	*x = UnfollowReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_micro_relation_v1_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnfollowReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnfollowReply) ProtoMessage() {}

func (x *UnfollowReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_micro_relation_v1_relation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnfollowReply.ProtoReflect.Descriptor instead.
func (*UnfollowReply) Descriptor() ([]byte, []int) {
	return file_api_micro_relation_v1_relation_proto_rawDescGZIP(), []int{3}
}

type GetFollowingUserListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: form:"user_id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" form:"user_id"`
	// @gotags: form:"last_id"
	LastId string `protobuf:"bytes,2,opt,name=last_id,json=lastId,proto3" json:"last_id,omitempty" form:"last_id"`
	// @gotags: form:"limit"
	Limit string `protobuf:"bytes,3,opt,name=limit,proto3" json:"limit,omitempty" form:"limit"`
}

func (x *GetFollowingUserListRequest) Reset() {
	*x = GetFollowingUserListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_micro_relation_v1_relation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowingUserListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowingUserListRequest) ProtoMessage() {}

func (x *GetFollowingUserListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_micro_relation_v1_relation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowingUserListRequest.ProtoReflect.Descriptor instead.
func (*GetFollowingUserListRequest) Descriptor() ([]byte, []int) {
	return file_api_micro_relation_v1_relation_proto_rawDescGZIP(), []int{4}
}

func (x *GetFollowingUserListRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetFollowingUserListRequest) GetLastId() string {
	if x != nil {
		return x.LastId
	}
	return ""
}

func (x *GetFollowingUserListRequest) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

type GetFollowingUserListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasMore int32      `protobuf:"varint,1,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	LastId  string     `protobuf:"bytes,2,opt,name=last_id,json=lastId,proto3" json:"last_id,omitempty"`
	Items   []*v1.User `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetFollowingUserListReply) Reset() {
	*x = GetFollowingUserListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_micro_relation_v1_relation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowingUserListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowingUserListReply) ProtoMessage() {}

func (x *GetFollowingUserListReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_micro_relation_v1_relation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowingUserListReply.ProtoReflect.Descriptor instead.
func (*GetFollowingUserListReply) Descriptor() ([]byte, []int) {
	return file_api_micro_relation_v1_relation_proto_rawDescGZIP(), []int{5}
}

func (x *GetFollowingUserListReply) GetHasMore() int32 {
	if x != nil {
		return x.HasMore
	}
	return 0
}

func (x *GetFollowingUserListReply) GetLastId() string {
	if x != nil {
		return x.LastId
	}
	return ""
}

func (x *GetFollowingUserListReply) GetItems() []*v1.User {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetFollowerUserListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: form:"user_id"
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" form:"user_id"`
	// @gotags: form:"last_id"
	LastId string `protobuf:"bytes,2,opt,name=last_id,json=lastId,proto3" json:"last_id,omitempty" form:"last_id"`
	// @gotags: form:"limit"
	Limit string `protobuf:"bytes,3,opt,name=limit,proto3" json:"limit,omitempty" form:"limit"`
}

func (x *GetFollowerUserListRequest) Reset() {
	*x = GetFollowerUserListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_micro_relation_v1_relation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowerUserListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowerUserListRequest) ProtoMessage() {}

func (x *GetFollowerUserListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_micro_relation_v1_relation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowerUserListRequest.ProtoReflect.Descriptor instead.
func (*GetFollowerUserListRequest) Descriptor() ([]byte, []int) {
	return file_api_micro_relation_v1_relation_proto_rawDescGZIP(), []int{6}
}

func (x *GetFollowerUserListRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetFollowerUserListRequest) GetLastId() string {
	if x != nil {
		return x.LastId
	}
	return ""
}

func (x *GetFollowerUserListRequest) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

type GetFollowerUserListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasMore int32      `protobuf:"varint,1,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	LastId  string     `protobuf:"bytes,2,opt,name=last_id,json=lastId,proto3" json:"last_id,omitempty"`
	Items   []*v1.User `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetFollowerUserListReply) Reset() {
	*x = GetFollowerUserListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_micro_relation_v1_relation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowerUserListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowerUserListReply) ProtoMessage() {}

func (x *GetFollowerUserListReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_micro_relation_v1_relation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowerUserListReply.ProtoReflect.Descriptor instead.
func (*GetFollowerUserListReply) Descriptor() ([]byte, []int) {
	return file_api_micro_relation_v1_relation_proto_rawDescGZIP(), []int{7}
}

func (x *GetFollowerUserListReply) GetHasMore() int32 {
	if x != nil {
		return x.HasMore
	}
	return 0
}

func (x *GetFollowerUserListReply) GetLastId() string {
	if x != nil {
		return x.LastId
	}
	return ""
}

func (x *GetFollowerUserListReply) GetItems() []*v1.User {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_api_micro_relation_v1_relation_proto protoreflect.FileDescriptor

var file_api_micro_relation_v1_relation_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x61, 0x70, 0x69,
	0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x0d, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f,
	0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x55, 0x69, 0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x4d, 0x0a, 0x0f, 0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x75, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x55, 0x69, 0x64, 0x22, 0x0f, 0x0a, 0x0d, 0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x5c, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x7e, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x64, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x7d, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xb8, 0x04, 0x0a, 0x0f, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x72, 0x0a, 0x06,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x3a, 0x01, 0x2a,
	0x12, 0x77, 0x0a, 0x08, 0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x26, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x17, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x75, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x9c, 0x01, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18,
	0x12, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x12, 0x98, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x42, 0x89, 0x02, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x69, 0x6e, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x92, 0x41, 0xc8, 0x01, 0x12, 0x9e, 0x01, 0x0a, 0x07, 0x49, 0x6e, 0x73, 0x20,
	0x41, 0x50, 0x49, 0x22, 0x3a, 0x0a, 0x0b, 0x49, 0x6e, 0x73, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x2b, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x73, 0x2d, 0x61, 0x70, 0x70, 0x2a,
	0x52, 0x0a, 0x0b, 0x4d, 0x49, 0x54, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x43,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x73, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x62, 0x6c, 0x6f, 0x62,
	0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x2e,
	0x74, 0x78, 0x74, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_micro_relation_v1_relation_proto_rawDescOnce sync.Once
	file_api_micro_relation_v1_relation_proto_rawDescData = file_api_micro_relation_v1_relation_proto_rawDesc
)

func file_api_micro_relation_v1_relation_proto_rawDescGZIP() []byte {
	file_api_micro_relation_v1_relation_proto_rawDescOnce.Do(func() {
		file_api_micro_relation_v1_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_micro_relation_v1_relation_proto_rawDescData)
	})
	return file_api_micro_relation_v1_relation_proto_rawDescData
}

var file_api_micro_relation_v1_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_micro_relation_v1_relation_proto_goTypes = []interface{}{
	(*FollowRequest)(nil),               // 0: api.micro.relation.v1.FollowRequest
	(*FollowReply)(nil),                 // 1: api.micro.relation.v1.FollowReply
	(*UnfollowRequest)(nil),             // 2: api.micro.relation.v1.UnfollowRequest
	(*UnfollowReply)(nil),               // 3: api.micro.relation.v1.UnfollowReply
	(*GetFollowingUserListRequest)(nil), // 4: api.micro.relation.v1.GetFollowingUserListRequest
	(*GetFollowingUserListReply)(nil),   // 5: api.micro.relation.v1.GetFollowingUserListReply
	(*GetFollowerUserListRequest)(nil),  // 6: api.micro.relation.v1.GetFollowerUserListRequest
	(*GetFollowerUserListReply)(nil),    // 7: api.micro.relation.v1.GetFollowerUserListReply
	(*v1.User)(nil),                     // 8: api.micro.user.v1.User
}
var file_api_micro_relation_v1_relation_proto_depIdxs = []int32{
	8, // 0: api.micro.relation.v1.GetFollowingUserListReply.items:type_name -> api.micro.user.v1.User
	8, // 1: api.micro.relation.v1.GetFollowerUserListReply.items:type_name -> api.micro.user.v1.User
	0, // 2: api.micro.relation.v1.RelationService.Follow:input_type -> api.micro.relation.v1.FollowRequest
	2, // 3: api.micro.relation.v1.RelationService.Unfollow:input_type -> api.micro.relation.v1.UnfollowRequest
	4, // 4: api.micro.relation.v1.RelationService.GetFollowingUserList:input_type -> api.micro.relation.v1.GetFollowingUserListRequest
	6, // 5: api.micro.relation.v1.RelationService.GetFollowerUserList:input_type -> api.micro.relation.v1.GetFollowerUserListRequest
	1, // 6: api.micro.relation.v1.RelationService.Follow:output_type -> api.micro.relation.v1.FollowReply
	3, // 7: api.micro.relation.v1.RelationService.Unfollow:output_type -> api.micro.relation.v1.UnfollowReply
	5, // 8: api.micro.relation.v1.RelationService.GetFollowingUserList:output_type -> api.micro.relation.v1.GetFollowingUserListReply
	7, // 9: api.micro.relation.v1.RelationService.GetFollowerUserList:output_type -> api.micro.relation.v1.GetFollowerUserListReply
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_micro_relation_v1_relation_proto_init() }
func file_api_micro_relation_v1_relation_proto_init() {
	if File_api_micro_relation_v1_relation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_micro_relation_v1_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowRequest); i {
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
		file_api_micro_relation_v1_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowReply); i {
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
		file_api_micro_relation_v1_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnfollowRequest); i {
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
		file_api_micro_relation_v1_relation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnfollowReply); i {
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
		file_api_micro_relation_v1_relation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowingUserListRequest); i {
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
		file_api_micro_relation_v1_relation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowingUserListReply); i {
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
		file_api_micro_relation_v1_relation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowerUserListRequest); i {
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
		file_api_micro_relation_v1_relation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowerUserListReply); i {
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
			RawDescriptor: file_api_micro_relation_v1_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_micro_relation_v1_relation_proto_goTypes,
		DependencyIndexes: file_api_micro_relation_v1_relation_proto_depIdxs,
		MessageInfos:      file_api_micro_relation_v1_relation_proto_msgTypes,
	}.Build()
	File_api_micro_relation_v1_relation_proto = out.File
	file_api_micro_relation_v1_relation_proto_rawDesc = nil
	file_api_micro_relation_v1_relation_proto_goTypes = nil
	file_api_micro_relation_v1_relation_proto_depIdxs = nil
}
