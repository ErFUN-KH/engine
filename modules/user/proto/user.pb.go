// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/user/proto/user.proto

package userpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/fzerorubigd/protobuf/extra"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type UserStatus int32

const (
	UserStatus_USER_STATUS_INVALID    UserStatus = 0
	UserStatus_USER_STATUS_REGISTERED UserStatus = 1
	UserStatus_USER_STATUS_ACTIVE     UserStatus = 2
	UserStatus_USER_STATUS_BANNED     UserStatus = 3
)

var UserStatus_name = map[int32]string{
	0: "USER_STATUS_INVALID",
	1: "USER_STATUS_REGISTERED",
	2: "USER_STATUS_ACTIVE",
	3: "USER_STATUS_BANNED",
}
var UserStatus_value = map[string]int32{
	"USER_STATUS_INVALID":    0,
	"USER_STATUS_REGISTERED": 1,
	"USER_STATUS_ACTIVE":     2,
	"USER_STATUS_BANNED":     3,
}

func (x UserStatus) String() string {
	return proto.EnumName(UserStatus_name, int32(x))
}
func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_15606ede51b8395c, []int{0}
}

type User struct {
	Id                   int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"id" `
	Email                string           `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty" db:"email" `
	Password             string           `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty" db:"password" `
	Status               UserStatus       `protobuf:"varint,4,opt,name=status,proto3,enum=user.UserStatus" json:"status,omitempty" db:"status" `
	Token                string           `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty" db:"token" `
	CreatedAt            *types.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt" json:"created_at,omitempty" db:"created_at" `
	UpdatedAt            *types.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty" db:"updated_at" `
	LastLogin            *types.Timestamp `protobuf:"bytes,8,opt,name=last_login,json=lastLogin" json:"last_login,omitempty" db:"last_login" `
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_15606ede51b8395c, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetStatus() UserStatus {
	if m != nil {
		return m.Status
	}
	return UserStatus_USER_STATUS_INVALID
}

func (m *User) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *User) GetCreatedAt() *types.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *types.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *User) GetLastLogin() *types.Timestamp {
	if m != nil {
		return m.LastLogin
	}
	return nil
}

type UserResponse struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string     `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Status               UserStatus `protobuf:"varint,3,opt,name=status,proto3,enum=user.UserStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_15606ede51b8395c, []int{1}
}
func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (dst *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(dst, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserResponse) GetStatus() UserStatus {
	if m != nil {
		return m.Status
	}
	return UserStatus_USER_STATUS_INVALID
}

type LoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty" validate:"email,required" `
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" validate:"gte=6,required" `
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_15606ede51b8395c, []int{2}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (dst *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(dst, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LogoutRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_15606ede51b8395c, []int{3}
}
func (m *LogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRequest.Unmarshal(m, b)
}
func (m *LogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRequest.Marshal(b, m, deterministic)
}
func (dst *LogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRequest.Merge(dst, src)
}
func (m *LogoutRequest) XXX_Size() int {
	return xxx_messageInfo_LogoutRequest.Size(m)
}
func (m *LogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRequest proto.InternalMessageInfo

type LogoutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutResponse) Reset()         { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()    {}
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_15606ede51b8395c, []int{4}
}
func (m *LogoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutResponse.Unmarshal(m, b)
}
func (m *LogoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutResponse.Marshal(b, m, deterministic)
}
func (dst *LogoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutResponse.Merge(dst, src)
}
func (m *LogoutResponse) XXX_Size() int {
	return xxx_messageInfo_LogoutResponse.Size(m)
}
func (m *LogoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutResponse proto.InternalMessageInfo

type RegisterRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty" validate:"email,required" `
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" validate:"gte=6,required" `
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_15606ede51b8395c, []int{5}
}
func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(dst, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_15606ede51b8395c, []int{6}
}
func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (dst *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(dst, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*UserResponse)(nil), "user.UserResponse")
	proto.RegisterType((*LoginRequest)(nil), "user.LoginRequest")
	proto.RegisterType((*LogoutRequest)(nil), "user.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "user.LogoutResponse")
	proto.RegisterType((*RegisterRequest)(nil), "user.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "user.RegisterResponse")
	proto.RegisterEnum("user.UserStatus", UserStatus_name, UserStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserSystemClient is the client API for UserSystem service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserSystemClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type userSystemClient struct {
	cc *grpc.ClientConn
}

func NewUserSystemClient(cc *grpc.ClientConn) UserSystemClient {
	return &userSystemClient{cc}
}

func (c *userSystemClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserSystem/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSystemClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/user.UserSystem/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSystemClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/user.UserSystem/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSystemServer is the server API for UserSystem service.
type UserSystemServer interface {
	Login(context.Context, *LoginRequest) (*UserResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
}

func RegisterUserSystemServer(s *grpc.Server, srv UserSystemServer) {
	s.RegisterService(&_UserSystem_serviceDesc, srv)
}

func _UserSystem_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSystemServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserSystem/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSystemServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSystem_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSystemServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserSystem/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSystemServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSystem_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSystemServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserSystem/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSystemServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserSystem_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserSystem",
	HandlerType: (*UserSystemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserSystem_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _UserSystem_Logout_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _UserSystem_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/user/proto/user.proto",
}

func init() { proto.RegisterFile("modules/user/proto/user.proto", fileDescriptor_user_15606ede51b8395c) }

var fileDescriptor_user_15606ede51b8395c = []byte{
	// 741 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0x4f, 0x6f, 0x12, 0x4d,
	0x1c, 0x7e, 0x17, 0x0a, 0x2f, 0x4c, 0x29, 0x6c, 0xa7, 0x7d, 0x79, 0x61, 0x53, 0xcb, 0x66, 0x12,
	0x0d, 0x69, 0x14, 0xd2, 0x6a, 0x3c, 0x90, 0x78, 0x60, 0x2d, 0x31, 0x98, 0xa6, 0x36, 0x0b, 0xad,
	0x89, 0x07, 0xc9, 0xd0, 0x9d, 0xae, 0x1b, 0x81, 0xa1, 0x3b, 0xb3, 0xf5, 0xcf, 0xc9, 0x68, 0xe2,
	0xa1, 0x57, 0x3f, 0x88, 0x89, 0x9f, 0xc2, 0xbb, 0x37, 0x0f, 0xc4, 0x83, 0x1f, 0xa0, 0xe1, 0x13,
	0x98, 0x99, 0xd9, 0x3f, 0xb4, 0xe9, 0xa1, 0x37, 0x2f, 0xe4, 0xf7, 0xe7, 0xf9, 0x3d, 0xf3, 0x9b,
	0xe7, 0x99, 0x05, 0xdc, 0x1a, 0x53, 0x27, 0x18, 0x11, 0xd6, 0x0c, 0x18, 0xf1, 0x9b, 0x53, 0x9f,
	0x72, 0x2a, 0xc3, 0x86, 0x0c, 0xe1, 0x92, 0x88, 0x8d, 0x6d, 0xd7, 0xe3, 0xaf, 0x82, 0x61, 0xe3,
	0x98, 0x8e, 0x9b, 0x27, 0xef, 0x89, 0x4f, 0xfd, 0x60, 0xe8, 0xb9, 0x8e, 0x82, 0x0f, 0x83, 0x93,
	0x26, 0x79, 0xcb, 0x7d, 0xac, 0x7e, 0xd5, 0xa0, 0x71, 0x6f, 0x61, 0xc4, 0xa5, 0x2e, 0x4d, 0xb0,
	0x22, 0x53, 0xe7, 0x88, 0x28, 0x84, 0x6f, 0xb8, 0x94, 0xba, 0x23, 0xd2, 0xc4, 0x53, 0xaf, 0x89,
	0x27, 0x13, 0xca, 0x31, 0xf7, 0xe8, 0x84, 0x85, 0xdd, 0x5a, 0xd8, 0x8d, 0x39, 0xb8, 0x37, 0x26,
	0x8c, 0xe3, 0xf1, 0x54, 0x01, 0xd0, 0xaf, 0x34, 0x58, 0x3a, 0x64, 0xc4, 0x87, 0x1b, 0x20, 0xe5,
	0x39, 0x15, 0xcd, 0xd4, 0xea, 0x69, 0xab, 0x30, 0x9f, 0xd5, 0x72, 0xce, 0xb0, 0x85, 0x3c, 0x07,
	0x99, 0x76, 0xca, 0x73, 0xe0, 0x6d, 0x90, 0x21, 0x63, 0xec, 0x8d, 0x2a, 0x29, 0x53, 0xab, 0xe7,
	0xad, 0xd2, 0x7c, 0x56, 0x5b, 0x16, 0x00, 0x59, 0x44, 0xa6, 0xad, 0xba, 0xb0, 0x01, 0x72, 0x53,
	0xcc, 0xd8, 0x1b, 0xea, 0x3b, 0x95, 0xb4, 0x44, 0xc2, 0xf9, 0xac, 0x56, 0x14, 0xc8, 0xa8, 0x8e,
	0x4c, 0x3b, 0xc6, 0xc0, 0x16, 0xc8, 0x32, 0x8e, 0x79, 0xc0, 0x2a, 0x4b, 0xa6, 0x56, 0x2f, 0xee,
	0xe8, 0x0d, 0xa9, 0xa0, 0x58, 0xa8, 0x27, 0xeb, 0x96, 0x3e, 0x9f, 0xd5, 0x0a, 0x62, 0x5e, 0xe1,
	0x90, 0x69, 0x87, 0x13, 0x62, 0x25, 0x4e, 0x5f, 0x93, 0x49, 0x25, 0x73, 0x79, 0x25, 0x59, 0x14,
	0x2b, 0xc9, 0x00, 0x3e, 0x03, 0xe0, 0xd8, 0x27, 0x98, 0x13, 0x67, 0x80, 0x79, 0x25, 0x6b, 0x6a,
	0xf5, 0xe5, 0x1d, 0xa3, 0xa1, 0x64, 0x69, 0x44, 0xb2, 0x34, 0xfa, 0x91, 0x2c, 0xd6, 0xfa, 0x7c,
	0x56, 0xd3, 0x05, 0x4f, 0x32, 0x85, 0x4c, 0x3b, 0x1f, 0x66, 0x6d, 0x2e, 0x08, 0x83, 0xa9, 0x13,
	0x11, 0xfe, 0x7b, 0x73, 0xc2, 0x64, 0x4a, 0x10, 0x86, 0x99, 0x22, 0x1c, 0x61, 0xc6, 0x07, 0x23,
	0xea, 0x7a, 0x93, 0x4a, 0xee, 0xe6, 0x84, 0xc9, 0x94, 0x20, 0x14, 0xd9, 0x9e, 0x48, 0x5a, 0x2b,
	0xdf, 0x2f, 0xaa, 0xda, 0xcf, 0x8b, 0x6a, 0x46, 0x88, 0xc9, 0xd0, 0x4b, 0x50, 0x10, 0x82, 0xda,
	0x84, 0x4d, 0xe9, 0x84, 0x11, 0x58, 0x8c, 0x9d, 0xce, 0x4b, 0x6f, 0xd7, 0x2f, 0x79, 0x1b, 0x59,
	0x59, 0x8f, 0xad, 0x49, 0x5f, 0x6f, 0x4d, 0x64, 0x04, 0xfa, 0xa0, 0x81, 0x82, 0x3c, 0xd8, 0x26,
	0xa7, 0x01, 0x61, 0x1c, 0x3e, 0x88, 0x08, 0xe5, 0x19, 0xd6, 0xe6, 0x7c, 0x56, 0x33, 0xce, 0xf0,
	0xc8, 0x13, 0x17, 0x0e, 0x9f, 0xcc, 0x5d, 0x9f, 0x9c, 0x06, 0x9e, 0x4f, 0x9c, 0xe4, 0xed, 0xb4,
	0x16, 0xde, 0x4e, 0xea, 0xba, 0x41, 0x97, 0x93, 0x47, 0x0f, 0x17, 0x07, 0x63, 0x3c, 0x2a, 0x81,
	0x95, 0x3d, 0xea, 0xd2, 0x80, 0x87, 0x2b, 0x20, 0x1d, 0x14, 0xa3, 0x82, 0xba, 0x35, 0xfa, 0xa4,
	0x81, 0x92, 0x4d, 0x5c, 0x8f, 0x71, 0x21, 0xc5, 0xdf, 0x5a, 0x14, 0x02, 0x3d, 0x59, 0x42, 0x6d,
	0xb6, 0x75, 0x0a, 0x40, 0xa2, 0x2a, 0xfc, 0x1f, 0xac, 0x1d, 0xf6, 0x3a, 0xf6, 0xa0, 0xd7, 0x6f,
	0xf7, 0x0f, 0x7b, 0x83, 0xee, 0xfe, 0x51, 0x7b, 0xaf, 0xbb, 0xab, 0xff, 0x03, 0x0d, 0x50, 0x5e,
	0x6c, 0xd8, 0x9d, 0x27, 0xdd, 0x5e, 0xbf, 0x63, 0x77, 0x76, 0x75, 0x0d, 0x96, 0x01, 0x5c, 0xec,
	0xb5, 0x1f, 0xf7, 0xbb, 0x47, 0x1d, 0x3d, 0x75, 0xb5, 0x6e, 0xb5, 0xf7, 0xf7, 0x3b, 0xbb, 0x7a,
	0x7a, 0xe7, 0x73, 0x2a, 0x3c, 0xf3, 0x1d, 0xe3, 0x64, 0x0c, 0x9f, 0x82, 0x8c, 0x34, 0x10, 0x42,
	0x65, 0xf2, 0xa2, 0x9b, 0x06, 0x4c, 0x8c, 0x8f, 0xc5, 0xac, 0x7e, 0xfc, 0xf1, 0xfb, 0x4b, 0x6a,
	0x0d, 0x15, 0x9b, 0x67, 0xdb, 0xea, 0xff, 0x4f, 0x3e, 0xc3, 0x96, 0xb6, 0x05, 0x7b, 0x20, 0xab,
	0x94, 0x87, 0x6b, 0x31, 0x59, 0x62, 0x8c, 0xb1, 0x7e, 0xb9, 0x18, 0xf2, 0x6d, 0x9e, 0x7f, 0x35,
	0xd2, 0x8c, 0x31, 0x49, 0xbb, 0x0a, 0x4b, 0x8b, 0xb4, 0x82, 0xea, 0x39, 0xc8, 0x45, 0xb2, 0xc1,
	0xff, 0x14, 0xc3, 0x15, 0x2f, 0x8d, 0xf2, 0xd5, 0x72, 0x48, 0xbd, 0x21, 0x39, 0xcb, 0x68, 0x35,
	0xe6, 0xf4, 0x43, 0x48, 0x4b, 0xdb, 0xb2, 0xee, 0x9c, 0x7f, 0xab, 0xa6, 0x31, 0xc6, 0x20, 0x77,
	0x4c, 0xc7, 0x92, 0xc1, 0xca, 0x8b, 0x9b, 0x1e, 0x88, 0xcf, 0xee, 0x40, 0x7b, 0x91, 0x15, 0xa5,
	0xe9, 0x70, 0x98, 0x95, 0xdf, 0xe1, 0xfd, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xf3, 0x06,
	0x6a, 0xf5, 0x05, 0x00, 0x00,
}
