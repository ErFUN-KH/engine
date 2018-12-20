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
	return fileDescriptor_user_0fea83290166ff07, []int{0}
}

type User struct {
	Id                   int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"id" `
	Email                string           `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty" db:"email" `
	Password             string           `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty" db:"password" `
	Status               UserStatus       `protobuf:"varint,4,opt,name=status,proto3,enum=user.UserStatus" json:"status,omitempty" db:"status" `
	CreatedAt            *types.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" db:"created_at" `
	UpdatedAt            *types.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" db:"updated_at" `
	LastLogin            *types.Timestamp `protobuf:"bytes,7,opt,name=last_login,json=lastLogin,proto3" json:"last_login,omitempty" db:"last_login" `
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_0fea83290166ff07, []int{0}
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
	Id                   int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string     `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Token                string     `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Status               UserStatus `protobuf:"varint,4,opt,name=status,proto3,enum=user.UserStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_0fea83290166ff07, []int{1}
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

func (m *UserResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserResponse) GetToken() string {
	if m != nil {
		return m.Token
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
	return fileDescriptor_user_0fea83290166ff07, []int{2}
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
	return fileDescriptor_user_0fea83290166ff07, []int{3}
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
	return fileDescriptor_user_0fea83290166ff07, []int{4}
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
	return fileDescriptor_user_0fea83290166ff07, []int{5}
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

type PingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_0fea83290166ff07, []int{6}
}
func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (dst *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(dst, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*UserResponse)(nil), "user.UserResponse")
	proto.RegisterType((*LoginRequest)(nil), "user.LoginRequest")
	proto.RegisterType((*LogoutRequest)(nil), "user.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "user.LogoutResponse")
	proto.RegisterType((*RegisterRequest)(nil), "user.RegisterRequest")
	proto.RegisterType((*PingRequest)(nil), "user.PingRequest")
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
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*UserResponse, error)
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

func (c *userSystemClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserSystem/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSystemClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserSystem/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSystemServer is the server API for UserSystem service.
type UserSystemServer interface {
	Login(context.Context, *LoginRequest) (*UserResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	Register(context.Context, *RegisterRequest) (*UserResponse, error)
	Ping(context.Context, *PingRequest) (*UserResponse, error)
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

func _UserSystem_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSystemServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserSystem/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSystemServer).Ping(ctx, req.(*PingRequest))
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
		{
			MethodName: "Ping",
			Handler:    _UserSystem_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/user/proto/user.proto",
}

func init() { proto.RegisterFile("modules/user/proto/user.proto", fileDescriptor_user_0fea83290166ff07) }

var fileDescriptor_user_0fea83290166ff07 = []byte{
	// 752 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xcb, 0x6e, 0xdb, 0x38,
	0x14, 0x8d, 0xfc, 0x1a, 0x87, 0xf1, 0x43, 0x61, 0x3c, 0x19, 0x5b, 0x93, 0x19, 0x0b, 0x02, 0x66,
	0x60, 0x04, 0x33, 0x36, 0x92, 0x19, 0xcc, 0xc2, 0xc0, 0x2c, 0xec, 0xc6, 0x68, 0x5d, 0x04, 0xa9,
	0x21, 0x3b, 0x59, 0x74, 0x63, 0xc8, 0x11, 0xa3, 0x12, 0xb5, 0x45, 0x45, 0xa4, 0xd2, 0xc7, 0xaa,
	0x68, 0x77, 0xd9, 0xf6, 0x43, 0x0a, 0x74, 0xd7, 0x3f, 0x68, 0xd7, 0xdd, 0x75, 0xe1, 0x55, 0x3f,
	0x20, 0xf0, 0x17, 0x14, 0x24, 0xf5, 0x4a, 0x10, 0xa0, 0xd9, 0x75, 0x63, 0xf0, 0x5e, 0x9e, 0x73,
	0x78, 0x79, 0x0e, 0x65, 0xf0, 0xdb, 0x82, 0xd8, 0xc1, 0x1c, 0xd1, 0x4e, 0x40, 0x91, 0xdf, 0xf1,
	0x7c, 0xc2, 0x88, 0x58, 0xb6, 0xc5, 0x12, 0xe6, 0xf8, 0x5a, 0xdb, 0x73, 0x30, 0x7b, 0x12, 0xcc,
	0xda, 0xa7, 0x64, 0xd1, 0x39, 0x7b, 0x89, 0x7c, 0xe2, 0x07, 0x33, 0xec, 0xd8, 0x12, 0x3e, 0x0b,
	0xce, 0x3a, 0xe8, 0x39, 0xf3, 0x2d, 0xf9, 0x2b, 0x89, 0xda, 0xdf, 0x29, 0x8a, 0x43, 0x1c, 0x92,
	0x60, 0x79, 0x25, 0xcf, 0xe1, 0xab, 0x10, 0xbe, 0xe3, 0x10, 0xe2, 0xcc, 0x51, 0xc7, 0xf2, 0x70,
	0xc7, 0x72, 0x5d, 0xc2, 0x2c, 0x86, 0x89, 0x4b, 0xc3, 0xdd, 0x66, 0xb8, 0x1b, 0x6b, 0x30, 0xbc,
	0x40, 0x94, 0x59, 0x0b, 0x4f, 0x02, 0x8c, 0x0f, 0x59, 0x90, 0x3b, 0xa6, 0xc8, 0x87, 0x3b, 0x20,
	0x83, 0xed, 0xba, 0xa2, 0x2b, 0xad, 0x6c, 0xbf, 0xb4, 0x5a, 0x36, 0x8b, 0xf6, 0xac, 0x6b, 0x60,
	0xdb, 0xd0, 0xcd, 0x0c, 0xb6, 0xe1, 0x1f, 0x20, 0x8f, 0x16, 0x16, 0x9e, 0xd7, 0x33, 0xba, 0xd2,
	0x5a, 0xef, 0x57, 0x57, 0xcb, 0xe6, 0x06, 0x07, 0x88, 0xa6, 0xa1, 0x9b, 0x72, 0x17, 0xb6, 0x41,
	0xd1, 0xb3, 0x28, 0x7d, 0x46, 0x7c, 0xbb, 0x9e, 0x15, 0x48, 0xb8, 0x5a, 0x36, 0x2b, 0x1c, 0x19,
	0xf5, 0x0d, 0xdd, 0x8c, 0x31, 0xb0, 0x0b, 0x0a, 0x94, 0x59, 0x2c, 0xa0, 0xf5, 0x9c, 0xae, 0xb4,
	0x2a, 0xfb, 0x6a, 0x5b, 0x38, 0xc8, 0x07, 0x1a, 0x8b, 0x7e, 0x5f, 0x5d, 0x2d, 0x9b, 0x25, 0xce,
	0x97, 0x38, 0x43, 0x37, 0x43, 0x06, 0x7c, 0x04, 0xc0, 0xa9, 0x8f, 0x2c, 0x86, 0xec, 0xa9, 0xc5,
	0xea, 0x79, 0x5d, 0x69, 0x6d, 0xec, 0x6b, 0x6d, 0x79, 0xdf, 0x76, 0x74, 0xdf, 0xf6, 0x24, 0xba,
	0x6f, 0xbf, 0xb6, 0x5a, 0x36, 0x55, 0xae, 0x94, 0xb0, 0x0c, 0xdd, 0x5c, 0x0f, 0xab, 0x1e, 0xe3,
	0x82, 0x81, 0x67, 0x47, 0x82, 0x85, 0xbb, 0x0b, 0x26, 0x2c, 0x2e, 0x18, 0x56, 0x52, 0x70, 0x6e,
	0x51, 0x36, 0x9d, 0x13, 0x07, 0xbb, 0xf5, 0x9f, 0xee, 0x2e, 0x98, 0xb0, 0xb8, 0x20, 0xaf, 0x0e,
	0x79, 0xd1, 0x2d, 0x7f, 0xbc, 0x6a, 0x28, 0x5f, 0xae, 0x1a, 0x79, 0xee, 0x12, 0x35, 0x18, 0x28,
	0x71, 0xa7, 0x4c, 0x44, 0x3d, 0xe2, 0x52, 0x04, 0x2b, 0x49, 0x84, 0x22, 0xb4, 0xda, 0xb5, 0xd0,
	0xa2, 0x8c, 0x6a, 0x20, 0xcf, 0xc8, 0x53, 0xe4, 0xca, 0x80, 0x4c, 0x59, 0xc0, 0xd6, 0xf7, 0x92,
	0x88, 0x7c, 0x37, 0x5e, 0x29, 0xa0, 0x24, 0xc6, 0x31, 0xd1, 0x79, 0x80, 0x28, 0x83, 0xff, 0x46,
	0xc7, 0x28, 0x22, 0xf1, 0xdf, 0x57, 0xcb, 0xa6, 0x76, 0x61, 0xcd, 0x31, 0xb7, 0x21, 0x7c, 0x21,
	0x7f, 0xf9, 0xe8, 0x3c, 0xc0, 0x3e, 0xb2, 0x93, 0xa7, 0xd2, 0x4d, 0x3d, 0x95, 0xcc, 0x6d, 0x44,
	0x87, 0xa1, 0xff, 0xff, 0x4b, 0x13, 0x63, 0xbc, 0x51, 0x05, 0xe5, 0x43, 0xe2, 0x90, 0x80, 0x85,
	0x23, 0x18, 0x2a, 0xa8, 0x44, 0x0d, 0xe9, 0x85, 0xf1, 0x46, 0x01, 0x55, 0x13, 0x39, 0x98, 0x32,
	0x6e, 0xd0, 0x8f, 0x1a, 0xb4, 0x0c, 0x36, 0x46, 0xd8, 0x75, 0xc2, 0x01, 0x76, 0xcf, 0x01, 0x48,
	0x0c, 0x85, 0xbf, 0x80, 0xad, 0xe3, 0xf1, 0xc0, 0x9c, 0x8e, 0x27, 0xbd, 0xc9, 0xf1, 0x78, 0x3a,
	0x3c, 0x3a, 0xe9, 0x1d, 0x0e, 0x0f, 0xd4, 0x35, 0xa8, 0x81, 0xed, 0xf4, 0x86, 0x39, 0xb8, 0x3f,
	0x1c, 0x4f, 0x06, 0xe6, 0xe0, 0x40, 0x55, 0xe0, 0x36, 0x80, 0xe9, 0xbd, 0xde, 0xbd, 0xc9, 0xf0,
	0x64, 0xa0, 0x66, 0x6e, 0xf6, 0xfb, 0xbd, 0xa3, 0xa3, 0xc1, 0x81, 0x9a, 0xdd, 0xff, 0x94, 0x09,
	0xcf, 0x7c, 0x41, 0x19, 0x5a, 0xc0, 0x87, 0x20, 0x2f, 0xb2, 0x83, 0x50, 0xe6, 0x9b, 0x0e, 0x52,
	0x83, 0x49, 0xe6, 0xb1, 0x8f, 0x8d, 0xd7, 0x9f, 0xbf, 0xbe, 0xcd, 0x6c, 0x19, 0x95, 0xce, 0xc5,
	0x9e, 0xfc, 0xa7, 0x13, 0xef, 0xb2, 0xab, 0xec, 0xc2, 0x11, 0x28, 0x48, 0xd3, 0xe1, 0x56, 0x2c,
	0x96, 0x64, 0xa2, 0xd5, 0xae, 0x37, 0x43, 0xbd, 0x5f, 0x2f, 0xdf, 0x69, 0x6b, 0x42, 0x73, 0x13,
	0x56, 0xd3, 0x9a, 0x5c, 0x67, 0x0c, 0x8a, 0x51, 0x66, 0xf0, 0x67, 0x49, 0xbf, 0x91, 0xe1, 0xad,
	0x33, 0xee, 0x08, 0xbd, 0x6d, 0x63, 0x33, 0xd6, 0xf3, 0x43, 0x16, 0x1f, 0xf3, 0x01, 0xc8, 0xf1,
	0x0c, 0xe0, 0xa6, 0x64, 0xa6, 0xf2, 0xb8, 0xfd, 0xc2, 0xf1, 0x80, 0x55, 0x58, 0x8e, 0x05, 0x3d,
	0xec, 0x3a, 0xfd, 0x3f, 0x2f, 0xdf, 0x37, 0xb2, 0x96, 0x65, 0x81, 0xe2, 0x29, 0x59, 0x08, 0x6a,
	0x7f, 0x9d, 0x73, 0x47, 0xfc, 0x53, 0x1e, 0x29, 0x8f, 0x0b, 0xbc, 0xe5, 0xcd, 0x66, 0x05, 0xf1,
	0x6d, 0xff, 0xf3, 0x2d, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x72, 0x46, 0x70, 0x22, 0x06, 0x00, 0x00,
}
