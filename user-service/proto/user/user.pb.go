// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package laracom_service_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	StripeId             string   `protobuf:"bytes,6,opt,name=stripe_id,json=stripeId,proto3" json:"stripe_id,omitempty"`
	CardBrand            string   `protobuf:"bytes,7,opt,name=card_brand,json=cardBrand,proto3" json:"card_brand,omitempty"`
	CardLastFour         string   `protobuf:"bytes,8,opt,name=card_last_four,json=cardLastFour,proto3" json:"card_last_four,omitempty"`
	TrialEndsAt          string   `protobuf:"bytes,9,opt,name=trial_ends_at,json=trialEndsAt,proto3" json:"trial_ends_at,omitempty"`
	DeletedAt            string   `protobuf:"bytes,10,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	RememberToken        string   `protobuf:"bytes,11,opt,name=remember_token,json=rememberToken,proto3" json:"remember_token,omitempty"`
	CreatedAt            string   `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
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

func (m *User) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *User) GetStripeId() string {
	if m != nil {
		return m.StripeId
	}
	return ""
}

func (m *User) GetCardBrand() string {
	if m != nil {
		return m.CardBrand
	}
	return ""
}

func (m *User) GetCardLastFour() string {
	if m != nil {
		return m.CardLastFour
	}
	return ""
}

func (m *User) GetTrialEndsAt() string {
	if m != nil {
		return m.TrialEndsAt
	}
	return ""
}

func (m *User) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

func (m *User) GetRememberToken() string {
	if m != nil {
		return m.RememberToken
	}
	return ""
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *User) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type PasswordReset struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	CreatedAt            string   `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PasswordReset) Reset()         { *m = PasswordReset{} }
func (m *PasswordReset) String() string { return proto.CompactTextString(m) }
func (*PasswordReset) ProtoMessage()    {}
func (*PasswordReset) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{5}
}

func (m *PasswordReset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasswordReset.Unmarshal(m, b)
}
func (m *PasswordReset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasswordReset.Marshal(b, m, deterministic)
}
func (m *PasswordReset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswordReset.Merge(m, src)
}
func (m *PasswordReset) XXX_Size() int {
	return xxx_messageInfo_PasswordReset.Size(m)
}
func (m *PasswordReset) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswordReset.DiscardUnknown(m)
}

var xxx_messageInfo_PasswordReset proto.InternalMessageInfo

func (m *PasswordReset) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *PasswordReset) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PasswordReset) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type PasswordResetResponse struct {
	PasswordReset        *PasswordReset `protobuf:"bytes,1,opt,name=passwordReset,proto3" json:"passwordReset,omitempty"`
	Errors               []*Error       `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PasswordResetResponse) Reset()         { *m = PasswordResetResponse{} }
func (m *PasswordResetResponse) String() string { return proto.CompactTextString(m) }
func (*PasswordResetResponse) ProtoMessage()    {}
func (*PasswordResetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{6}
}

func (m *PasswordResetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasswordResetResponse.Unmarshal(m, b)
}
func (m *PasswordResetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasswordResetResponse.Marshal(b, m, deterministic)
}
func (m *PasswordResetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswordResetResponse.Merge(m, src)
}
func (m *PasswordResetResponse) XXX_Size() int {
	return xxx_messageInfo_PasswordResetResponse.Size(m)
}
func (m *PasswordResetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswordResetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PasswordResetResponse proto.InternalMessageInfo

func (m *PasswordResetResponse) GetPasswordReset() *PasswordReset {
	if m != nil {
		return m.PasswordReset
	}
	return nil
}

func (m *PasswordResetResponse) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "laracom.service.user.User")
	proto.RegisterType((*Request)(nil), "laracom.service.user.Request")
	proto.RegisterType((*Response)(nil), "laracom.service.user.Response")
	proto.RegisterType((*Token)(nil), "laracom.service.user.Token")
	proto.RegisterType((*Error)(nil), "laracom.service.user.Error")
	proto.RegisterType((*PasswordReset)(nil), "laracom.service.user.PasswordReset")
	proto.RegisterType((*PasswordResetResponse)(nil), "laracom.service.user.PasswordResetResponse")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x6d, 0x9c, 0xd8, 0x4d, 0x26, 0x4d, 0x0f, 0xdb, 0x16, 0xad, 0x52, 0x15, 0x55, 0x06, 0x24,
	0x24, 0xa4, 0x80, 0xda, 0x33, 0x87, 0x50, 0xda, 0xaa, 0x82, 0x03, 0x32, 0x14, 0x21, 0x2e, 0xd6,
	0x36, 0x3b, 0xa8, 0x2b, 0x1c, 0xdb, 0xec, 0xae, 0xcb, 0xa7, 0x70, 0xe6, 0x03, 0xf9, 0x05, 0x84,
	0x76, 0xd6, 0x69, 0x63, 0x94, 0xb6, 0x11, 0x11, 0x97, 0xc4, 0xf3, 0xe6, 0xed, 0xcc, 0xce, 0xbc,
	0xe7, 0x04, 0x76, 0x4a, 0x5d, 0xd8, 0xe2, 0x79, 0x65, 0x50, 0xd3, 0xc7, 0x88, 0x62, 0xb6, 0x9d,
	0x09, 0x2d, 0x26, 0xc5, 0x74, 0x64, 0x50, 0x5f, 0xa9, 0x09, 0x8e, 0x5c, 0x2e, 0xfe, 0x1d, 0x40,
	0xe7, 0xdc, 0xa0, 0x66, 0x9b, 0x10, 0x28, 0xc9, 0x5b, 0xfb, 0xad, 0xa7, 0xbd, 0x24, 0x50, 0x92,
	0x31, 0xe8, 0xe4, 0x62, 0x8a, 0x3c, 0x20, 0x84, 0x9e, 0xd9, 0x36, 0x84, 0x38, 0x15, 0x2a, 0xe3,
	0x6d, 0x02, 0x7d, 0xc0, 0x86, 0xd0, 0x2d, 0x85, 0x31, 0xdf, 0x0b, 0x2d, 0x79, 0x87, 0x12, 0xd7,
	0x31, 0x7b, 0x00, 0x91, 0xb1, 0xc2, 0x56, 0x86, 0x87, 0x94, 0xa9, 0x23, 0xb6, 0x0b, 0x3d, 0x63,
	0xb5, 0x2a, 0x31, 0x55, 0x92, 0x47, 0xfe, 0x90, 0x07, 0xce, 0x24, 0xdb, 0x03, 0x98, 0x08, 0x2d,
	0xd3, 0x0b, 0x2d, 0x72, 0xc9, 0xd7, 0x29, 0xdb, 0x73, 0xc8, 0x2b, 0x07, 0xb0, 0xc7, 0xb0, 0x49,
	0xe9, 0x4c, 0x18, 0x9b, 0x7e, 0x29, 0x2a, 0xcd, 0xbb, 0x44, 0xd9, 0x70, 0xe8, 0x5b, 0x61, 0xec,
	0x49, 0x51, 0x69, 0x16, 0xc3, 0xc0, 0x6a, 0x25, 0xb2, 0x14, 0x73, 0x69, 0x52, 0x61, 0x79, 0x8f,
	0x48, 0x7d, 0x02, 0x8f, 0x73, 0x69, 0xc6, 0xd6, 0x35, 0x92, 0x98, 0xa1, 0x45, 0xe9, 0x08, 0xe0,
	0x1b, 0xd5, 0xc8, 0xd8, 0xb2, 0x27, 0xb0, 0xa9, 0x71, 0x8a, 0xd3, 0x0b, 0xd4, 0xa9, 0x2d, 0xbe,
	0x62, 0xce, 0xfb, 0x44, 0x19, 0xcc, 0xd0, 0x0f, 0x0e, 0xa4, 0xeb, 0x6a, 0x14, 0x75, 0x95, 0x8d,
	0xfa, 0xba, 0x1e, 0xf1, 0x4d, 0xaa, 0x52, 0xce, 0xd2, 0x03, 0x9f, 0xae, 0x91, 0xb1, 0x8d, 0x7b,
	0xb0, 0x9e, 0xe0, 0xb7, 0x0a, 0x8d, 0x8d, 0x7f, 0xb6, 0xa0, 0x9b, 0xa0, 0x29, 0x8b, 0xdc, 0x20,
	0x1b, 0x41, 0xc7, 0x09, 0x44, 0x8a, 0xf4, 0x0f, 0x86, 0xa3, 0x45, 0xea, 0x8d, 0x9c, 0x72, 0x09,
	0xf1, 0xd8, 0x0b, 0x08, 0xdd, 0xb7, 0xe1, 0xc1, 0x7e, 0xfb, 0x9e, 0x03, 0x9e, 0xc8, 0x0e, 0x21,
	0x42, 0xad, 0x0b, 0x6d, 0x78, 0x9b, 0x8e, 0xec, 0x2e, 0x3e, 0x72, 0xec, 0x38, 0x49, 0x4d, 0x8d,
	0x2f, 0x21, 0xf4, 0x53, 0x6f, 0x43, 0xe8, 0x77, 0xe2, 0x2d, 0xe3, 0x03, 0x87, 0x5e, 0x89, 0x4c,
	0x49, 0xb2, 0x4d, 0x37, 0xf1, 0xc1, 0xbf, 0x75, 0x7a, 0x09, 0x21, 0x01, 0xce, 0x89, 0x93, 0x42,
	0x22, 0x35, 0x0a, 0x13, 0x7a, 0x66, 0xfb, 0xd0, 0x97, 0x68, 0x26, 0x5a, 0x95, 0x56, 0x15, 0x79,
	0x6d, 0xd2, 0x79, 0x28, 0xfe, 0x0c, 0x83, 0x77, 0xb5, 0x0b, 0x13, 0x34, 0x68, 0x6f, 0xcc, 0xdb,
	0x9a, 0x37, 0xef, 0xf5, 0x18, 0xc1, 0xfc, 0x18, 0x4d, 0x49, 0xdb, 0x7f, 0x49, 0x1a, 0xff, 0x68,
	0xc1, 0x4e, 0xa3, 0xf8, 0xb5, 0x6a, 0x67, 0x30, 0x28, 0xe7, 0x13, 0xb5, 0x7c, 0x8f, 0x16, 0x0f,
	0xdc, 0xac, 0xd1, 0x3c, 0x39, 0xb7, 0xb4, 0x60, 0xe9, 0xa5, 0x1d, 0xfc, 0x0a, 0xa1, 0xef, 0x34,
	0x7e, 0xef, 0x29, 0xec, 0x04, 0xa2, 0x23, 0xba, 0x36, 0xbb, 0xc3, 0x10, 0xc3, 0x87, 0x8b, 0x73,
	0xb3, 0xa9, 0xe2, 0x35, 0x76, 0x04, 0xed, 0x53, 0xb4, 0x2b, 0x16, 0x39, 0x83, 0xe8, 0x14, 0xed,
	0x38, 0xcb, 0xd8, 0xde, 0x6d, 0x5c, 0x7a, 0x11, 0x96, 0x28, 0x35, 0x86, 0xce, 0xb8, 0xb2, 0x97,
	0x77, 0x5e, 0xe8, 0x96, 0x85, 0x91, 0x7d, 0xe3, 0x35, 0xf6, 0x06, 0x06, 0x1f, 0x9d, 0x3b, 0x85,
	0x45, 0xef, 0xe8, 0xbb, 0xf8, 0xf7, 0x15, 0x3b, 0x81, 0xe8, 0x9c, 0x5e, 0xe9, 0x15, 0x57, 0xa4,
	0x60, 0xcb, 0xeb, 0xd5, 0xf4, 0xee, 0x32, 0xfe, 0x19, 0x3e, 0x5b, 0xc6, 0x64, 0x37, 0xad, 0x3e,
	0xc1, 0x70, 0x36, 0x7f, 0x83, 0xb2, 0xfa, 0x32, 0x14, 0x6c, 0xbd, 0xa6, 0x1f, 0xd1, 0xff, 0x3e,
	0xc4, 0x45, 0x44, 0xff, 0x6d, 0x87, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x80, 0x76, 0xcc, 0x05,
	0xf4, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
	Update(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	CreatePasswordReset(ctx context.Context, in *PasswordReset, opts ...client.CallOption) (*PasswordResetResponse, error)
	ValidatePasswordResetToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
	DeletePasswordReset(ctx context.Context, in *PasswordReset, opts ...client.CallOption) (*PasswordResetResponse, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "laracom.service.user"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Auth", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreatePasswordReset(ctx context.Context, in *PasswordReset, opts ...client.CallOption) (*PasswordResetResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.CreatePasswordReset", in)
	out := new(PasswordResetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ValidatePasswordResetToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.ValidatePasswordResetToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeletePasswordReset(ctx context.Context, in *PasswordReset, opts ...client.CallOption) (*PasswordResetResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.DeletePasswordReset", in)
	out := new(PasswordResetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Create(context.Context, *User, *Response) error
	Get(context.Context, *User, *Response) error
	GetAll(context.Context, *Request, *Response) error
	Auth(context.Context, *User, *Token) error
	ValidateToken(context.Context, *Token, *Token) error
	Update(context.Context, *User, *Response) error
	CreatePasswordReset(context.Context, *PasswordReset, *PasswordResetResponse) error
	ValidatePasswordResetToken(context.Context, *Token, *Token) error
	DeletePasswordReset(context.Context, *PasswordReset, *PasswordResetResponse) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) Create(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *UserService) Get(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func (h *UserService) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.GetAll(ctx, in, out)
}

func (h *UserService) Auth(ctx context.Context, in *User, out *Token) error {
	return h.UserServiceHandler.Auth(ctx, in, out)
}

func (h *UserService) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.UserServiceHandler.ValidateToken(ctx, in, out)
}

func (h *UserService) Update(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Update(ctx, in, out)
}

func (h *UserService) CreatePasswordReset(ctx context.Context, in *PasswordReset, out *PasswordResetResponse) error {
	return h.UserServiceHandler.CreatePasswordReset(ctx, in, out)
}

func (h *UserService) ValidatePasswordResetToken(ctx context.Context, in *Token, out *Token) error {
	return h.UserServiceHandler.ValidatePasswordResetToken(ctx, in, out)
}

func (h *UserService) DeletePasswordReset(ctx context.Context, in *PasswordReset, out *PasswordResetResponse) error {
	return h.UserServiceHandler.DeletePasswordReset(ctx, in, out)
}
