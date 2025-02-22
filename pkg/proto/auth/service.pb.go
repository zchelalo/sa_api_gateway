// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: sa_proto/auth/service.proto

package authProto

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

type AuthUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Verified bool   `protobuf:"varint,4,opt,name=verified,proto3" json:"verified,omitempty"`
}

func (x *AuthUser) Reset() {
	*x = AuthUser{}
	mi := &file_sa_proto_auth_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUser) ProtoMessage() {}

func (x *AuthUser) ProtoReflect() protoreflect.Message {
	mi := &file_sa_proto_auth_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUser.ProtoReflect.Descriptor instead.
func (*AuthUser) Descriptor() ([]byte, []int) {
	return file_sa_proto_auth_service_proto_rawDescGZIP(), []int{0}
}

func (x *AuthUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuthUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AuthUser) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string    `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken string    `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	ExpiresAt    int64     `protobuf:"varint,3,opt,name=expiresAt,proto3" json:"expiresAt,omitempty"`
	User         *AuthUser `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	mi := &file_sa_proto_auth_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_sa_proto_auth_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_sa_proto_auth_service_proto_rawDescGZIP(), []int{1}
}

func (x *Auth) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Auth) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *Auth) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *Auth) GetUser() *AuthUser {
	if x != nil {
		return x.User
	}
	return nil
}

type Tokens struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	ExpiresAt    *int64 `protobuf:"varint,3,opt,name=expiresAt,proto3,oneof" json:"expiresAt,omitempty"`
}

func (x *Tokens) Reset() {
	*x = Tokens{}
	mi := &file_sa_proto_auth_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tokens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tokens) ProtoMessage() {}

func (x *Tokens) ProtoReflect() protoreflect.Message {
	mi := &file_sa_proto_auth_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tokens.ProtoReflect.Descriptor instead.
func (*Tokens) Descriptor() ([]byte, []int) {
	return file_sa_proto_auth_service_proto_rawDescGZIP(), []int{2}
}

func (x *Tokens) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Tokens) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *Tokens) GetExpiresAt() int64 {
	if x != nil && x.ExpiresAt != nil {
		return *x.ExpiresAt
	}
	return 0
}

type IsAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAuthorized bool    `protobuf:"varint,1,opt,name=isAuthorized,proto3" json:"isAuthorized,omitempty"`
	UserId       *string `protobuf:"bytes,2,opt,name=userId,proto3,oneof" json:"userId,omitempty"`
	Tokens       *Tokens `protobuf:"bytes,3,opt,name=tokens,proto3,oneof" json:"tokens,omitempty"`
}

func (x *IsAuth) Reset() {
	*x = IsAuth{}
	mi := &file_sa_proto_auth_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAuth) ProtoMessage() {}

func (x *IsAuth) ProtoReflect() protoreflect.Message {
	mi := &file_sa_proto_auth_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAuth.ProtoReflect.Descriptor instead.
func (*IsAuth) Descriptor() ([]byte, []int) {
	return file_sa_proto_auth_service_proto_rawDescGZIP(), []int{3}
}

func (x *IsAuth) GetIsAuthorized() bool {
	if x != nil {
		return x.IsAuthorized
	}
	return false
}

func (x *IsAuth) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *IsAuth) GetTokens() *Tokens {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type AuthError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      // Código de error, como los códigos de estado HTTP
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // Mensaje de error descriptivo
}

func (x *AuthError) Reset() {
	*x = AuthError{}
	mi := &file_sa_proto_auth_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthError) ProtoMessage() {}

func (x *AuthError) ProtoReflect() protoreflect.Message {
	mi := &file_sa_proto_auth_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthError.ProtoReflect.Descriptor instead.
func (*AuthError) Descriptor() ([]byte, []int) {
	return file_sa_proto_auth_service_proto_rawDescGZIP(), []int{4}
}

func (x *AuthError) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AuthError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignInRequest) Reset() {
	*x = SignInRequest{}
	mi := &file_sa_proto_auth_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRequest) ProtoMessage() {}

func (x *SignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sa_proto_auth_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRequest.ProtoReflect.Descriptor instead.
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return file_sa_proto_auth_service_proto_rawDescGZIP(), []int{5}
}

func (x *SignInRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//
	//	*SignInResponse_Auth
	//	*SignInResponse_Error
	Result isSignInResponse_Result `protobuf_oneof:"result"`
}

func (x *SignInResponse) Reset() {
	*x = SignInResponse{}
	mi := &file_sa_proto_auth_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResponse) ProtoMessage() {}

func (x *SignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sa_proto_auth_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResponse.ProtoReflect.Descriptor instead.
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return file_sa_proto_auth_service_proto_rawDescGZIP(), []int{6}
}

func (m *SignInResponse) GetResult() isSignInResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *SignInResponse) GetAuth() *Auth {
	if x, ok := x.GetResult().(*SignInResponse_Auth); ok {
		return x.Auth
	}
	return nil
}

func (x *SignInResponse) GetError() *AuthError {
	if x, ok := x.GetResult().(*SignInResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isSignInResponse_Result interface {
	isSignInResponse_Result()
}

type SignInResponse_Auth struct {
	Auth *Auth `protobuf:"bytes,1,opt,name=auth,proto3,oneof"`
}

type SignInResponse_Error struct {
	Error *AuthError `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*SignInResponse_Auth) isSignInResponse_Result() {}

func (*SignInResponse_Error) isSignInResponse_Result() {}

type SignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignUpRequest) Reset() {
	*x = SignUpRequest{}
	mi := &file_sa_proto_auth_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRequest) ProtoMessage() {}

func (x *SignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sa_proto_auth_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRequest.ProtoReflect.Descriptor instead.
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return file_sa_proto_auth_service_proto_rawDescGZIP(), []int{7}
}

func (x *SignUpRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SignUpRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//
	//	*SignUpResponse_Auth
	//	*SignUpResponse_Error
	Result isSignUpResponse_Result `protobuf_oneof:"result"`
}

func (x *SignUpResponse) Reset() {
	*x = SignUpResponse{}
	mi := &file_sa_proto_auth_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResponse) ProtoMessage() {}

func (x *SignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sa_proto_auth_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResponse.ProtoReflect.Descriptor instead.
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return file_sa_proto_auth_service_proto_rawDescGZIP(), []int{8}
}

func (m *SignUpResponse) GetResult() isSignUpResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *SignUpResponse) GetAuth() *Auth {
	if x, ok := x.GetResult().(*SignUpResponse_Auth); ok {
		return x.Auth
	}
	return nil
}

func (x *SignUpResponse) GetError() *AuthError {
	if x, ok := x.GetResult().(*SignUpResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isSignUpResponse_Result interface {
	isSignUpResponse_Result()
}

type SignUpResponse_Auth struct {
	Auth *Auth `protobuf:"bytes,1,opt,name=auth,proto3,oneof"`
}

type SignUpResponse_Error struct {
	Error *AuthError `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*SignUpResponse_Auth) isSignUpResponse_Result() {}

func (*SignUpResponse_Error) isSignUpResponse_Result() {}

type SignOutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *SignOutRequest) Reset() {
	*x = SignOutRequest{}
	mi := &file_sa_proto_auth_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignOutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignOutRequest) ProtoMessage() {}

func (x *SignOutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sa_proto_auth_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignOutRequest.ProtoReflect.Descriptor instead.
func (*SignOutRequest) Descriptor() ([]byte, []int) {
	return file_sa_proto_auth_service_proto_rawDescGZIP(), []int{9}
}

func (x *SignOutRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type SignOutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//
	//	*SignOutResponse_Success
	//	*SignOutResponse_Error
	Result isSignOutResponse_Result `protobuf_oneof:"result"`
}

func (x *SignOutResponse) Reset() {
	*x = SignOutResponse{}
	mi := &file_sa_proto_auth_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignOutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignOutResponse) ProtoMessage() {}

func (x *SignOutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sa_proto_auth_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignOutResponse.ProtoReflect.Descriptor instead.
func (*SignOutResponse) Descriptor() ([]byte, []int) {
	return file_sa_proto_auth_service_proto_rawDescGZIP(), []int{10}
}

func (m *SignOutResponse) GetResult() isSignOutResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *SignOutResponse) GetSuccess() bool {
	if x, ok := x.GetResult().(*SignOutResponse_Success); ok {
		return x.Success
	}
	return false
}

func (x *SignOutResponse) GetError() *AuthError {
	if x, ok := x.GetResult().(*SignOutResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isSignOutResponse_Result interface {
	isSignOutResponse_Result()
}

type SignOutResponse_Success struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3,oneof"`
}

type SignOutResponse_Error struct {
	Error *AuthError `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*SignOutResponse_Success) isSignOutResponse_Result() {}

func (*SignOutResponse_Error) isSignOutResponse_Result() {}

type IsAuthorizedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *IsAuthorizedRequest) Reset() {
	*x = IsAuthorizedRequest{}
	mi := &file_sa_proto_auth_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsAuthorizedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAuthorizedRequest) ProtoMessage() {}

func (x *IsAuthorizedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sa_proto_auth_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAuthorizedRequest.ProtoReflect.Descriptor instead.
func (*IsAuthorizedRequest) Descriptor() ([]byte, []int) {
	return file_sa_proto_auth_service_proto_rawDescGZIP(), []int{11}
}

func (x *IsAuthorizedRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *IsAuthorizedRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type IsAuthorizedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//
	//	*IsAuthorizedResponse_Data
	//	*IsAuthorizedResponse_Error
	Result isIsAuthorizedResponse_Result `protobuf_oneof:"result"`
}

func (x *IsAuthorizedResponse) Reset() {
	*x = IsAuthorizedResponse{}
	mi := &file_sa_proto_auth_service_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsAuthorizedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAuthorizedResponse) ProtoMessage() {}

func (x *IsAuthorizedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sa_proto_auth_service_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAuthorizedResponse.ProtoReflect.Descriptor instead.
func (*IsAuthorizedResponse) Descriptor() ([]byte, []int) {
	return file_sa_proto_auth_service_proto_rawDescGZIP(), []int{12}
}

func (m *IsAuthorizedResponse) GetResult() isIsAuthorizedResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *IsAuthorizedResponse) GetData() *IsAuth {
	if x, ok := x.GetResult().(*IsAuthorizedResponse_Data); ok {
		return x.Data
	}
	return nil
}

func (x *IsAuthorizedResponse) GetError() *AuthError {
	if x, ok := x.GetResult().(*IsAuthorizedResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isIsAuthorizedResponse_Result interface {
	isIsAuthorizedResponse_Result()
}

type IsAuthorizedResponse_Data struct {
	Data *IsAuth `protobuf:"bytes,1,opt,name=data,proto3,oneof"`
}

type IsAuthorizedResponse_Error struct {
	Error *AuthError `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*IsAuthorizedResponse_Data) isIsAuthorizedResponse_Result() {}

func (*IsAuthorizedResponse_Error) isIsAuthorizedResponse_Result() {}

var File_sa_proto_auth_service_proto protoreflect.FileDescriptor

var file_sa_proto_auth_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a,
	0x08, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22,
	0x89, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x7f, 0x0a, 0x06, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x09, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00,
	0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x22, 0x85, 0x01, 0x0a,
	0x06, 0x49, 0x73, 0x41, 0x75, 0x74, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69,
	0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x48, 0x01, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x88, 0x01, 0x01, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x22, 0x39, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x41, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x5b, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x48, 0x00, 0x52, 0x04, 0x61, 0x75, 0x74,
	0x68, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x55, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x5b, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x48, 0x00, 0x52,
	0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x34, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5b, 0x0a, 0x0f, 0x53, 0x69, 0x67,
	0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x5b, 0x0a, 0x13, 0x49, 0x73, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x63, 0x0a, 0x14, 0x49, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x49, 0x73, 0x41, 0x75,
	0x74, 0x68, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x08,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xce, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e,
	0x49, 0x6e, 0x12, 0x0e, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x0e, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x12, 0x0f, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0c,
	0x69, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x14, 0x2e, 0x49,
	0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x49, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sa_proto_auth_service_proto_rawDescOnce sync.Once
	file_sa_proto_auth_service_proto_rawDescData = file_sa_proto_auth_service_proto_rawDesc
)

func file_sa_proto_auth_service_proto_rawDescGZIP() []byte {
	file_sa_proto_auth_service_proto_rawDescOnce.Do(func() {
		file_sa_proto_auth_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_sa_proto_auth_service_proto_rawDescData)
	})
	return file_sa_proto_auth_service_proto_rawDescData
}

var file_sa_proto_auth_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_sa_proto_auth_service_proto_goTypes = []any{
	(*AuthUser)(nil),             // 0: AuthUser
	(*Auth)(nil),                 // 1: Auth
	(*Tokens)(nil),               // 2: Tokens
	(*IsAuth)(nil),               // 3: IsAuth
	(*AuthError)(nil),            // 4: AuthError
	(*SignInRequest)(nil),        // 5: SignInRequest
	(*SignInResponse)(nil),       // 6: SignInResponse
	(*SignUpRequest)(nil),        // 7: SignUpRequest
	(*SignUpResponse)(nil),       // 8: SignUpResponse
	(*SignOutRequest)(nil),       // 9: SignOutRequest
	(*SignOutResponse)(nil),      // 10: SignOutResponse
	(*IsAuthorizedRequest)(nil),  // 11: IsAuthorizedRequest
	(*IsAuthorizedResponse)(nil), // 12: IsAuthorizedResponse
}
var file_sa_proto_auth_service_proto_depIdxs = []int32{
	0,  // 0: Auth.user:type_name -> AuthUser
	2,  // 1: IsAuth.tokens:type_name -> Tokens
	1,  // 2: SignInResponse.auth:type_name -> Auth
	4,  // 3: SignInResponse.error:type_name -> AuthError
	1,  // 4: SignUpResponse.auth:type_name -> Auth
	4,  // 5: SignUpResponse.error:type_name -> AuthError
	4,  // 6: SignOutResponse.error:type_name -> AuthError
	3,  // 7: IsAuthorizedResponse.data:type_name -> IsAuth
	4,  // 8: IsAuthorizedResponse.error:type_name -> AuthError
	5,  // 9: AuthService.signIn:input_type -> SignInRequest
	7,  // 10: AuthService.signUp:input_type -> SignUpRequest
	9,  // 11: AuthService.signOut:input_type -> SignOutRequest
	11, // 12: AuthService.isAuthorized:input_type -> IsAuthorizedRequest
	6,  // 13: AuthService.signIn:output_type -> SignInResponse
	8,  // 14: AuthService.signUp:output_type -> SignUpResponse
	10, // 15: AuthService.signOut:output_type -> SignOutResponse
	12, // 16: AuthService.isAuthorized:output_type -> IsAuthorizedResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_sa_proto_auth_service_proto_init() }
func file_sa_proto_auth_service_proto_init() {
	if File_sa_proto_auth_service_proto != nil {
		return
	}
	file_sa_proto_auth_service_proto_msgTypes[2].OneofWrappers = []any{}
	file_sa_proto_auth_service_proto_msgTypes[3].OneofWrappers = []any{}
	file_sa_proto_auth_service_proto_msgTypes[6].OneofWrappers = []any{
		(*SignInResponse_Auth)(nil),
		(*SignInResponse_Error)(nil),
	}
	file_sa_proto_auth_service_proto_msgTypes[8].OneofWrappers = []any{
		(*SignUpResponse_Auth)(nil),
		(*SignUpResponse_Error)(nil),
	}
	file_sa_proto_auth_service_proto_msgTypes[10].OneofWrappers = []any{
		(*SignOutResponse_Success)(nil),
		(*SignOutResponse_Error)(nil),
	}
	file_sa_proto_auth_service_proto_msgTypes[12].OneofWrappers = []any{
		(*IsAuthorizedResponse_Data)(nil),
		(*IsAuthorizedResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sa_proto_auth_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sa_proto_auth_service_proto_goTypes,
		DependencyIndexes: file_sa_proto_auth_service_proto_depIdxs,
		MessageInfos:      file_sa_proto_auth_service_proto_msgTypes,
	}.Build()
	File_sa_proto_auth_service_proto = out.File
	file_sa_proto_auth_service_proto_rawDesc = nil
	file_sa_proto_auth_service_proto_goTypes = nil
	file_sa_proto_auth_service_proto_depIdxs = nil
}
