// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: grpc_user/user.proto

package grpc_user

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

// The user type
type UserType int32

const (
	UserType_USER_TYPE_UNSPECIFIED UserType = 0
	UserType_USER_TYPE_OA          UserType = 1
	UserType_USER_TYPE_NORMAL      UserType = 2
)

// Enum value maps for UserType.
var (
	UserType_name = map[int32]string{
		0: "USER_TYPE_UNSPECIFIED",
		1: "USER_TYPE_OA",
		2: "USER_TYPE_NORMAL",
	}
	UserType_value = map[string]int32{
		"USER_TYPE_UNSPECIFIED": 0,
		"USER_TYPE_OA":          1,
		"USER_TYPE_NORMAL":      2,
	}
)

func (x UserType) Enum() *UserType {
	p := new(UserType)
	*p = x
	return p
}

func (x UserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserType) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_user_user_proto_enumTypes[0].Descriptor()
}

func (UserType) Type() protoreflect.EnumType {
	return &file_grpc_user_user_proto_enumTypes[0]
}

func (x UserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserType.Descriptor instead.
func (UserType) EnumDescriptor() ([]byte, []int) {
	return file_grpc_user_user_proto_rawDescGZIP(), []int{0}
}

// The user status
type UserStatus int32

const (
	UserStatus_USER_STATUS_UNSPECIFIED UserStatus = 0
	UserStatus_USER_STATUS_ENABLE      UserStatus = 1
	UserStatus_USER_STATUS_DISABLE     UserStatus = 2
)

// Enum value maps for UserStatus.
var (
	UserStatus_name = map[int32]string{
		0: "USER_STATUS_UNSPECIFIED",
		1: "USER_STATUS_ENABLE",
		2: "USER_STATUS_DISABLE",
	}
	UserStatus_value = map[string]int32{
		"USER_STATUS_UNSPECIFIED": 0,
		"USER_STATUS_ENABLE":      1,
		"USER_STATUS_DISABLE":     2,
	}
)

func (x UserStatus) Enum() *UserStatus {
	p := new(UserStatus)
	*p = x
	return p
}

func (x UserStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_user_user_proto_enumTypes[1].Descriptor()
}

func (UserStatus) Type() protoreflect.EnumType {
	return &file_grpc_user_user_proto_enumTypes[1]
}

func (x UserStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserStatus.Descriptor instead.
func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return file_grpc_user_user_proto_rawDescGZIP(), []int{1}
}

// The request message containing the user's id.
type UserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`          // 用户ID
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"` // 用户名称
}

func (x *UserInfoRequest) Reset() {
	*x = UserInfoRequest{}
	mi := &file_grpc_user_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoRequest) ProtoMessage() {}

func (x *UserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_user_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoRequest.ProtoReflect.Descriptor instead.
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return file_grpc_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfoRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

// The response message containing the user info
type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                               // 用户ID
	Account   string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`                      // 用户账号
	Password  string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`                    // 用户密码
	Name      string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`                            // 用户名称
	Avatar    string `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`                        // 用户头像地址
	IsEnable  bool   `protobuf:"varint,6,opt,name=is_enable,json=isEnable,proto3" json:"is_enable,omitempty"`   // 可用状态(0:禁用,1:启用)
	CreatedAt string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // 添加时间
	UpdatedAt string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"` // 修改时间
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	mi := &file_grpc_user_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_user_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_grpc_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfo) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *UserInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserInfo) GetIsEnable() bool {
	if x != nil {
		return x.IsEnable
	}
	return false
}

func (x *UserInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UserInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// The request message containing the user's list
type UserListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64      `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int64      `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Name     string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type     UserType   `protobuf:"varint,4,opt,name=type,proto3,enum=grpc_user.UserType" json:"type,omitempty"`
	Status   UserStatus `protobuf:"varint,5,opt,name=status,proto3,enum=grpc_user.UserStatus" json:"status,omitempty"`
}

func (x *UserListRequest) Reset() {
	*x = UserListRequest{}
	mi := &file_grpc_user_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListRequest) ProtoMessage() {}

func (x *UserListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_user_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListRequest.ProtoReflect.Descriptor instead.
func (*UserListRequest) Descriptor() ([]byte, []int) {
	return file_grpc_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *UserListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *UserListRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserListRequest) GetType() UserType {
	if x != nil {
		return x.Type
	}
	return UserType_USER_TYPE_UNSPECIFIED
}

func (x *UserListRequest) GetStatus() UserStatus {
	if x != nil {
		return x.Status
	}
	return UserStatus_USER_STATUS_UNSPECIFIED
}

// The response message containing the user list
type UserListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*UserInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *UserListReply) Reset() {
	*x = UserListReply{}
	mi := &file_grpc_user_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListReply) ProtoMessage() {}

func (x *UserListReply) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_user_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListReply.ProtoReflect.Descriptor instead.
func (*UserListReply) Descriptor() ([]byte, []int) {
	return file_grpc_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserListReply) GetList() []*UserInfo {
	if x != nil {
		return x.List
	}
	return nil
}

// The request message containing the user's info
type UserSaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar   string     `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Type     UserType   `protobuf:"varint,4,opt,name=type,proto3,enum=grpc_user.UserType" json:"type,omitempty"`
	IsEnable bool       `protobuf:"varint,5,opt,name=is_enable,json=isEnable,proto3" json:"is_enable,omitempty"`
	Status   UserStatus `protobuf:"varint,6,opt,name=status,proto3,enum=grpc_user.UserStatus" json:"status,omitempty"`
}

func (x *UserSaveRequest) Reset() {
	*x = UserSaveRequest{}
	mi := &file_grpc_user_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserSaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSaveRequest) ProtoMessage() {}

func (x *UserSaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_user_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSaveRequest.ProtoReflect.Descriptor instead.
func (*UserSaveRequest) Descriptor() ([]byte, []int) {
	return file_grpc_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *UserSaveRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserSaveRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserSaveRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserSaveRequest) GetType() UserType {
	if x != nil {
		return x.Type
	}
	return UserType_USER_TYPE_UNSPECIFIED
}

func (x *UserSaveRequest) GetIsEnable() bool {
	if x != nil {
		return x.IsEnable
	}
	return false
}

func (x *UserSaveRequest) GetStatus() UserStatus {
	if x != nil {
		return x.Status
	}
	return UserStatus_USER_STATUS_UNSPECIFIED
}

// The response message containing the user save
type UserSaveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserSaveReply) Reset() {
	*x = UserSaveReply{}
	mi := &file_grpc_user_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserSaveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSaveReply) ProtoMessage() {}

func (x *UserSaveReply) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_user_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSaveReply.ProtoReflect.Descriptor instead.
func (*UserSaveReply) Descriptor() ([]byte, []int) {
	return file_grpc_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserSaveReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// The request message containing the user's id
type UserDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserDeleteRequest) Reset() {
	*x = UserDeleteRequest{}
	mi := &file_grpc_user_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDeleteRequest) ProtoMessage() {}

func (x *UserDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_user_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDeleteRequest.ProtoReflect.Descriptor instead.
func (*UserDeleteRequest) Descriptor() ([]byte, []int) {
	return file_grpc_user_user_proto_rawDescGZIP(), []int{6}
}

func (x *UserDeleteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// The response message containing the user delete
type UserDeleteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserDeleteReply) Reset() {
	*x = UserDeleteReply{}
	mi := &file_grpc_user_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserDeleteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDeleteReply) ProtoMessage() {}

func (x *UserDeleteReply) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_user_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDeleteReply.ProtoReflect.Descriptor instead.
func (*UserDeleteReply) Descriptor() ([]byte, []int) {
	return file_grpc_user_user_proto_rawDescGZIP(), []int{7}
}

func (x *UserDeleteReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_grpc_user_user_proto protoreflect.FileDescriptor

var file_grpc_user_user_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x22, 0x3b, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xd7,
	0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xae, 0x01, 0x0a, 0x0f, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x38, 0x0a, 0x0d, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x22, 0xc2, 0x01, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x53, 0x61, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1f, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x21,
	0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x2a, 0x4d, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x15, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x41, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x02,
	0x2a, 0x5a, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b,
	0x0a, 0x17, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x55,
	0x53, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x32, 0x8f, 0x02, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e,
	0x64, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x40, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x40, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x53, 0x61, 0x76, 0x65,
	0x12, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x61, 0x76,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x46, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x39,
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x50,
	0x01, 0x5a, 0x26, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x3b,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_grpc_user_user_proto_rawDescOnce sync.Once
	file_grpc_user_user_proto_rawDescData = file_grpc_user_user_proto_rawDesc
)

func file_grpc_user_user_proto_rawDescGZIP() []byte {
	file_grpc_user_user_proto_rawDescOnce.Do(func() {
		file_grpc_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_user_user_proto_rawDescData)
	})
	return file_grpc_user_user_proto_rawDescData
}

var file_grpc_user_user_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_grpc_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_grpc_user_user_proto_goTypes = []any{
	(UserType)(0),             // 0: grpc_user.UserType
	(UserStatus)(0),           // 1: grpc_user.UserStatus
	(*UserInfoRequest)(nil),   // 2: grpc_user.UserInfoRequest
	(*UserInfo)(nil),          // 3: grpc_user.UserInfo
	(*UserListRequest)(nil),   // 4: grpc_user.UserListRequest
	(*UserListReply)(nil),     // 5: grpc_user.UserListReply
	(*UserSaveRequest)(nil),   // 6: grpc_user.UserSaveRequest
	(*UserSaveReply)(nil),     // 7: grpc_user.UserSaveReply
	(*UserDeleteRequest)(nil), // 8: grpc_user.UserDeleteRequest
	(*UserDeleteReply)(nil),   // 9: grpc_user.UserDeleteReply
}
var file_grpc_user_user_proto_depIdxs = []int32{
	0, // 0: grpc_user.UserListRequest.type:type_name -> grpc_user.UserType
	1, // 1: grpc_user.UserListRequest.status:type_name -> grpc_user.UserStatus
	3, // 2: grpc_user.UserListReply.list:type_name -> grpc_user.UserInfo
	0, // 3: grpc_user.UserSaveRequest.type:type_name -> grpc_user.UserType
	1, // 4: grpc_user.UserSaveRequest.status:type_name -> grpc_user.UserStatus
	2, // 5: grpc_user.User.UserFind:input_type -> grpc_user.UserInfoRequest
	4, // 6: grpc_user.User.UserList:input_type -> grpc_user.UserListRequest
	6, // 7: grpc_user.User.UserSave:input_type -> grpc_user.UserSaveRequest
	8, // 8: grpc_user.User.UserDelete:input_type -> grpc_user.UserDeleteRequest
	3, // 9: grpc_user.User.UserFind:output_type -> grpc_user.UserInfo
	5, // 10: grpc_user.User.UserList:output_type -> grpc_user.UserListReply
	7, // 11: grpc_user.User.UserSave:output_type -> grpc_user.UserSaveReply
	9, // 12: grpc_user.User.UserDelete:output_type -> grpc_user.UserDeleteReply
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_grpc_user_user_proto_init() }
func file_grpc_user_user_proto_init() {
	if File_grpc_user_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_user_user_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_user_user_proto_goTypes,
		DependencyIndexes: file_grpc_user_user_proto_depIdxs,
		EnumInfos:         file_grpc_user_user_proto_enumTypes,
		MessageInfos:      file_grpc_user_user_proto_msgTypes,
	}.Build()
	File_grpc_user_user_proto = out.File
	file_grpc_user_user_proto_rawDesc = nil
	file_grpc_user_user_proto_goTypes = nil
	file_grpc_user_user_proto_depIdxs = nil
}
