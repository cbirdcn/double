// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: club.proto

package club_v1

import (
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

type CreateClubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateClubRequest) Reset() {
	*x = CreateClubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClubRequest) ProtoMessage() {}

func (x *CreateClubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_club_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClubRequest.ProtoReflect.Descriptor instead.
func (*CreateClubRequest) Descriptor() ([]byte, []int) {
	return file_club_proto_rawDescGZIP(), []int{0}
}

type CreateClubReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateClubReply) Reset() {
	*x = CreateClubReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClubReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClubReply) ProtoMessage() {}

func (x *CreateClubReply) ProtoReflect() protoreflect.Message {
	mi := &file_club_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClubReply.ProtoReflect.Descriptor instead.
func (*CreateClubReply) Descriptor() ([]byte, []int) {
	return file_club_proto_rawDescGZIP(), []int{1}
}

type UpdateClubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateClubRequest) Reset() {
	*x = UpdateClubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClubRequest) ProtoMessage() {}

func (x *UpdateClubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_club_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClubRequest.ProtoReflect.Descriptor instead.
func (*UpdateClubRequest) Descriptor() ([]byte, []int) {
	return file_club_proto_rawDescGZIP(), []int{2}
}

type UpdateClubReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateClubReply) Reset() {
	*x = UpdateClubReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClubReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClubReply) ProtoMessage() {}

func (x *UpdateClubReply) ProtoReflect() protoreflect.Message {
	mi := &file_club_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClubReply.ProtoReflect.Descriptor instead.
func (*UpdateClubReply) Descriptor() ([]byte, []int) {
	return file_club_proto_rawDescGZIP(), []int{3}
}

type DeleteClubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteClubRequest) Reset() {
	*x = DeleteClubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClubRequest) ProtoMessage() {}

func (x *DeleteClubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_club_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClubRequest.ProtoReflect.Descriptor instead.
func (*DeleteClubRequest) Descriptor() ([]byte, []int) {
	return file_club_proto_rawDescGZIP(), []int{4}
}

type DeleteClubReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteClubReply) Reset() {
	*x = DeleteClubReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClubReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClubReply) ProtoMessage() {}

func (x *DeleteClubReply) ProtoReflect() protoreflect.Message {
	mi := &file_club_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClubReply.ProtoReflect.Descriptor instead.
func (*DeleteClubReply) Descriptor() ([]byte, []int) {
	return file_club_proto_rawDescGZIP(), []int{5}
}

type GetClubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetClubRequest) Reset() {
	*x = GetClubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClubRequest) ProtoMessage() {}

func (x *GetClubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_club_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClubRequest.ProtoReflect.Descriptor instead.
func (*GetClubRequest) Descriptor() ([]byte, []int) {
	return file_club_proto_rawDescGZIP(), []int{6}
}

func (x *GetClubRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetClubReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetClubReply) Reset() {
	*x = GetClubReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClubReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClubReply) ProtoMessage() {}

func (x *GetClubReply) ProtoReflect() protoreflect.Message {
	mi := &file_club_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClubReply.ProtoReflect.Descriptor instead.
func (*GetClubReply) Descriptor() ([]byte, []int) {
	return file_club_proto_rawDescGZIP(), []int{7}
}

type ListClubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListClubRequest) Reset() {
	*x = ListClubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClubRequest) ProtoMessage() {}

func (x *ListClubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_club_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClubRequest.ProtoReflect.Descriptor instead.
func (*ListClubRequest) Descriptor() ([]byte, []int) {
	return file_club_proto_rawDescGZIP(), []int{8}
}

type ListClubReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListClubReply) Reset() {
	*x = ListClubReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClubReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClubReply) ProtoMessage() {}

func (x *ListClubReply) ProtoReflect() protoreflect.Message {
	mi := &file_club_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClubReply.ProtoReflect.Descriptor instead.
func (*ListClubReply) Descriptor() ([]byte, []int) {
	return file_club_proto_rawDescGZIP(), []int{9}
}

var File_club_proto protoreflect.FileDescriptor

var file_club_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x13, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c,
	0x75, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x62,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x89, 0x03, 0x0a, 0x04, 0x43, 0x6c, 0x75, 0x62, 0x12, 0x4a,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x62, 0x12, 0x1e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4a, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x62, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6c, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6c, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75,
	0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4a, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6c, 0x75, 0x62, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x57, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x62, 0x12, 0x1b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x62,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x76,
	0x31, 0x2f, 0x63, 0x6c, 0x75, 0x62, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x44, 0x0a, 0x08, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x62, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c,
	0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x75, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x42, 0x1c, 0x5a, 0x1a, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6c, 0x75, 0x62, 0x5f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_club_proto_rawDescOnce sync.Once
	file_club_proto_rawDescData = file_club_proto_rawDesc
)

func file_club_proto_rawDescGZIP() []byte {
	file_club_proto_rawDescOnce.Do(func() {
		file_club_proto_rawDescData = protoimpl.X.CompressGZIP(file_club_proto_rawDescData)
	})
	return file_club_proto_rawDescData
}

var file_club_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_club_proto_goTypes = []interface{}{
	(*CreateClubRequest)(nil), // 0: api.club.v1.CreateClubRequest
	(*CreateClubReply)(nil),   // 1: api.club.v1.CreateClubReply
	(*UpdateClubRequest)(nil), // 2: api.club.v1.UpdateClubRequest
	(*UpdateClubReply)(nil),   // 3: api.club.v1.UpdateClubReply
	(*DeleteClubRequest)(nil), // 4: api.club.v1.DeleteClubRequest
	(*DeleteClubReply)(nil),   // 5: api.club.v1.DeleteClubReply
	(*GetClubRequest)(nil),    // 6: api.club.v1.GetClubRequest
	(*GetClubReply)(nil),      // 7: api.club.v1.GetClubReply
	(*ListClubRequest)(nil),   // 8: api.club.v1.ListClubRequest
	(*ListClubReply)(nil),     // 9: api.club.v1.ListClubReply
}
var file_club_proto_depIdxs = []int32{
	0, // 0: api.club.v1.Club.CreateClub:input_type -> api.club.v1.CreateClubRequest
	2, // 1: api.club.v1.Club.UpdateClub:input_type -> api.club.v1.UpdateClubRequest
	4, // 2: api.club.v1.Club.DeleteClub:input_type -> api.club.v1.DeleteClubRequest
	6, // 3: api.club.v1.Club.GetClub:input_type -> api.club.v1.GetClubRequest
	8, // 4: api.club.v1.Club.ListClub:input_type -> api.club.v1.ListClubRequest
	1, // 5: api.club.v1.Club.CreateClub:output_type -> api.club.v1.CreateClubReply
	3, // 6: api.club.v1.Club.UpdateClub:output_type -> api.club.v1.UpdateClubReply
	5, // 7: api.club.v1.Club.DeleteClub:output_type -> api.club.v1.DeleteClubReply
	7, // 8: api.club.v1.Club.GetClub:output_type -> api.club.v1.GetClubReply
	9, // 9: api.club.v1.Club.ListClub:output_type -> api.club.v1.ListClubReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_club_proto_init() }
func file_club_proto_init() {
	if File_club_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_club_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClubRequest); i {
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
		file_club_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClubReply); i {
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
		file_club_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClubRequest); i {
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
		file_club_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClubReply); i {
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
		file_club_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClubRequest); i {
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
		file_club_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClubReply); i {
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
		file_club_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClubRequest); i {
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
		file_club_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClubReply); i {
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
		file_club_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClubRequest); i {
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
		file_club_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClubReply); i {
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
			RawDescriptor: file_club_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_club_proto_goTypes,
		DependencyIndexes: file_club_proto_depIdxs,
		MessageInfos:      file_club_proto_msgTypes,
	}.Build()
	File_club_proto = out.File
	file_club_proto_rawDesc = nil
	file_club_proto_goTypes = nil
	file_club_proto_depIdxs = nil
}
