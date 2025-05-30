// Copyright © 2023 OpenIM open source community. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: chat/chat.proto

package chat

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Chat_UpdateUserInfo_FullMethodName          = "/openim.chat.chat/UpdateUserInfo"
	Chat_AddUserAccount_FullMethodName          = "/openim.chat.chat/AddUserAccount"
	Chat_SearchUserPublicInfo_FullMethodName    = "/openim.chat.chat/SearchUserPublicInfo"
	Chat_FindUserPublicInfo_FullMethodName      = "/openim.chat.chat/FindUserPublicInfo"
	Chat_SearchUserFullInfo_FullMethodName      = "/openim.chat.chat/SearchUserFullInfo"
	Chat_FindUserFullInfo_FullMethodName        = "/openim.chat.chat/FindUserFullInfo"
	Chat_SendVerifyCode_FullMethodName          = "/openim.chat.chat/SendVerifyCode"
	Chat_VerifyCode_FullMethodName              = "/openim.chat.chat/VerifyCode"
	Chat_RegisterUser_FullMethodName            = "/openim.chat.chat/RegisterUser"
	Chat_Login_FullMethodName                   = "/openim.chat.chat/Login"
	Chat_ResetPassword_FullMethodName           = "/openim.chat.chat/ResetPassword"
	Chat_ChangePassword_FullMethodName          = "/openim.chat.chat/ChangePassword"
	Chat_CheckUserExist_FullMethodName          = "/openim.chat.chat/CheckUserExist"
	Chat_DelUserAccount_FullMethodName          = "/openim.chat.chat/DelUserAccount"
	Chat_FindUserAccount_FullMethodName         = "/openim.chat.chat/FindUserAccount"
	Chat_FindAccountUser_FullMethodName         = "/openim.chat.chat/FindAccountUser"
	Chat_OpenIMCallback_FullMethodName          = "/openim.chat.chat/OpenIMCallback"
	Chat_UserLoginCount_FullMethodName          = "/openim.chat.chat/UserLoginCount"
	Chat_SearchUserInfo_FullMethodName          = "/openim.chat.chat/SearchUserInfo"
	Chat_GetTokenForVideoMeeting_FullMethodName = "/openim.chat.chat/GetTokenForVideoMeeting"
	Chat_SetAllowRegister_FullMethodName        = "/openim.chat.chat/SetAllowRegister"
	Chat_GetAllowRegister_FullMethodName        = "/openim.chat.chat/GetAllowRegister"
)

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	// Edit personal information - called by the user or an administrator
	UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*UpdateUserInfoResp, error)
	AddUserAccount(ctx context.Context, in *AddUserAccountReq, opts ...grpc.CallOption) (*AddUserAccountResp, error)
	// Get user's public information - called by strangers
	SearchUserPublicInfo(ctx context.Context, in *SearchUserPublicInfoReq, opts ...grpc.CallOption) (*SearchUserPublicInfoResp, error)
	FindUserPublicInfo(ctx context.Context, in *FindUserPublicInfoReq, opts ...grpc.CallOption) (*FindUserPublicInfoResp, error)
	// Search user information - called by administrators, other users get public fields
	SearchUserFullInfo(ctx context.Context, in *SearchUserFullInfoReq, opts ...grpc.CallOption) (*SearchUserFullInfoResp, error)
	FindUserFullInfo(ctx context.Context, in *FindUserFullInfoReq, opts ...grpc.CallOption) (*FindUserFullInfoResp, error)
	SendVerifyCode(ctx context.Context, in *SendVerifyCodeReq, opts ...grpc.CallOption) (*SendVerifyCodeResp, error)
	VerifyCode(ctx context.Context, in *VerifyCodeReq, opts ...grpc.CallOption) (*VerifyCodeResp, error)
	RegisterUser(ctx context.Context, in *RegisterUserReq, opts ...grpc.CallOption) (*RegisterUserResp, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	ResetPassword(ctx context.Context, in *ResetPasswordReq, opts ...grpc.CallOption) (*ResetPasswordResp, error)
	ChangePassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*ChangePasswordResp, error)
	CheckUserExist(ctx context.Context, in *CheckUserExistReq, opts ...grpc.CallOption) (*CheckUserExistResp, error)
	DelUserAccount(ctx context.Context, in *DelUserAccountReq, opts ...grpc.CallOption) (*DelUserAccountResp, error)
	FindUserAccount(ctx context.Context, in *FindUserAccountReq, opts ...grpc.CallOption) (*FindUserAccountResp, error)
	FindAccountUser(ctx context.Context, in *FindAccountUserReq, opts ...grpc.CallOption) (*FindAccountUserResp, error)
	OpenIMCallback(ctx context.Context, in *OpenIMCallbackReq, opts ...grpc.CallOption) (*OpenIMCallbackResp, error)
	// Statistics
	UserLoginCount(ctx context.Context, in *UserLoginCountReq, opts ...grpc.CallOption) (*UserLoginCountResp, error)
	SearchUserInfo(ctx context.Context, in *SearchUserInfoReq, opts ...grpc.CallOption) (*SearchUserInfoResp, error)
	// Audio/video call and video meeting
	GetTokenForVideoMeeting(ctx context.Context, in *GetTokenForVideoMeetingReq, opts ...grpc.CallOption) (*GetTokenForVideoMeetingResp, error)
	SetAllowRegister(ctx context.Context, in *SetAllowRegisterReq, opts ...grpc.CallOption) (*SetAllowRegisterResp, error)
	GetAllowRegister(ctx context.Context, in *GetAllowRegisterReq, opts ...grpc.CallOption) (*GetAllowRegisterResp, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*UpdateUserInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserInfoResp)
	err := c.cc.Invoke(ctx, Chat_UpdateUserInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) AddUserAccount(ctx context.Context, in *AddUserAccountReq, opts ...grpc.CallOption) (*AddUserAccountResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddUserAccountResp)
	err := c.cc.Invoke(ctx, Chat_AddUserAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) SearchUserPublicInfo(ctx context.Context, in *SearchUserPublicInfoReq, opts ...grpc.CallOption) (*SearchUserPublicInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchUserPublicInfoResp)
	err := c.cc.Invoke(ctx, Chat_SearchUserPublicInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) FindUserPublicInfo(ctx context.Context, in *FindUserPublicInfoReq, opts ...grpc.CallOption) (*FindUserPublicInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindUserPublicInfoResp)
	err := c.cc.Invoke(ctx, Chat_FindUserPublicInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) SearchUserFullInfo(ctx context.Context, in *SearchUserFullInfoReq, opts ...grpc.CallOption) (*SearchUserFullInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchUserFullInfoResp)
	err := c.cc.Invoke(ctx, Chat_SearchUserFullInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) FindUserFullInfo(ctx context.Context, in *FindUserFullInfoReq, opts ...grpc.CallOption) (*FindUserFullInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindUserFullInfoResp)
	err := c.cc.Invoke(ctx, Chat_FindUserFullInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) SendVerifyCode(ctx context.Context, in *SendVerifyCodeReq, opts ...grpc.CallOption) (*SendVerifyCodeResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendVerifyCodeResp)
	err := c.cc.Invoke(ctx, Chat_SendVerifyCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) VerifyCode(ctx context.Context, in *VerifyCodeReq, opts ...grpc.CallOption) (*VerifyCodeResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyCodeResp)
	err := c.cc.Invoke(ctx, Chat_VerifyCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) RegisterUser(ctx context.Context, in *RegisterUserReq, opts ...grpc.CallOption) (*RegisterUserResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterUserResp)
	err := c.cc.Invoke(ctx, Chat_RegisterUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, Chat_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) ResetPassword(ctx context.Context, in *ResetPasswordReq, opts ...grpc.CallOption) (*ResetPasswordResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResetPasswordResp)
	err := c.cc.Invoke(ctx, Chat_ResetPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) ChangePassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*ChangePasswordResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangePasswordResp)
	err := c.cc.Invoke(ctx, Chat_ChangePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) CheckUserExist(ctx context.Context, in *CheckUserExistReq, opts ...grpc.CallOption) (*CheckUserExistResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckUserExistResp)
	err := c.cc.Invoke(ctx, Chat_CheckUserExist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) DelUserAccount(ctx context.Context, in *DelUserAccountReq, opts ...grpc.CallOption) (*DelUserAccountResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DelUserAccountResp)
	err := c.cc.Invoke(ctx, Chat_DelUserAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) FindUserAccount(ctx context.Context, in *FindUserAccountReq, opts ...grpc.CallOption) (*FindUserAccountResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindUserAccountResp)
	err := c.cc.Invoke(ctx, Chat_FindUserAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) FindAccountUser(ctx context.Context, in *FindAccountUserReq, opts ...grpc.CallOption) (*FindAccountUserResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindAccountUserResp)
	err := c.cc.Invoke(ctx, Chat_FindAccountUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) OpenIMCallback(ctx context.Context, in *OpenIMCallbackReq, opts ...grpc.CallOption) (*OpenIMCallbackResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OpenIMCallbackResp)
	err := c.cc.Invoke(ctx, Chat_OpenIMCallback_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) UserLoginCount(ctx context.Context, in *UserLoginCountReq, opts ...grpc.CallOption) (*UserLoginCountResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserLoginCountResp)
	err := c.cc.Invoke(ctx, Chat_UserLoginCount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) SearchUserInfo(ctx context.Context, in *SearchUserInfoReq, opts ...grpc.CallOption) (*SearchUserInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchUserInfoResp)
	err := c.cc.Invoke(ctx, Chat_SearchUserInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetTokenForVideoMeeting(ctx context.Context, in *GetTokenForVideoMeetingReq, opts ...grpc.CallOption) (*GetTokenForVideoMeetingResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTokenForVideoMeetingResp)
	err := c.cc.Invoke(ctx, Chat_GetTokenForVideoMeeting_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) SetAllowRegister(ctx context.Context, in *SetAllowRegisterReq, opts ...grpc.CallOption) (*SetAllowRegisterResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetAllowRegisterResp)
	err := c.cc.Invoke(ctx, Chat_SetAllowRegister_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetAllowRegister(ctx context.Context, in *GetAllowRegisterReq, opts ...grpc.CallOption) (*GetAllowRegisterResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllowRegisterResp)
	err := c.cc.Invoke(ctx, Chat_GetAllowRegister_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility.
type ChatServer interface {
	// Edit personal information - called by the user or an administrator
	UpdateUserInfo(context.Context, *UpdateUserInfoReq) (*UpdateUserInfoResp, error)
	AddUserAccount(context.Context, *AddUserAccountReq) (*AddUserAccountResp, error)
	// Get user's public information - called by strangers
	SearchUserPublicInfo(context.Context, *SearchUserPublicInfoReq) (*SearchUserPublicInfoResp, error)
	FindUserPublicInfo(context.Context, *FindUserPublicInfoReq) (*FindUserPublicInfoResp, error)
	// Search user information - called by administrators, other users get public fields
	SearchUserFullInfo(context.Context, *SearchUserFullInfoReq) (*SearchUserFullInfoResp, error)
	FindUserFullInfo(context.Context, *FindUserFullInfoReq) (*FindUserFullInfoResp, error)
	SendVerifyCode(context.Context, *SendVerifyCodeReq) (*SendVerifyCodeResp, error)
	VerifyCode(context.Context, *VerifyCodeReq) (*VerifyCodeResp, error)
	RegisterUser(context.Context, *RegisterUserReq) (*RegisterUserResp, error)
	Login(context.Context, *LoginReq) (*LoginResp, error)
	ResetPassword(context.Context, *ResetPasswordReq) (*ResetPasswordResp, error)
	ChangePassword(context.Context, *ChangePasswordReq) (*ChangePasswordResp, error)
	CheckUserExist(context.Context, *CheckUserExistReq) (*CheckUserExistResp, error)
	DelUserAccount(context.Context, *DelUserAccountReq) (*DelUserAccountResp, error)
	FindUserAccount(context.Context, *FindUserAccountReq) (*FindUserAccountResp, error)
	FindAccountUser(context.Context, *FindAccountUserReq) (*FindAccountUserResp, error)
	OpenIMCallback(context.Context, *OpenIMCallbackReq) (*OpenIMCallbackResp, error)
	// Statistics
	UserLoginCount(context.Context, *UserLoginCountReq) (*UserLoginCountResp, error)
	SearchUserInfo(context.Context, *SearchUserInfoReq) (*SearchUserInfoResp, error)
	// Audio/video call and video meeting
	GetTokenForVideoMeeting(context.Context, *GetTokenForVideoMeetingReq) (*GetTokenForVideoMeetingResp, error)
	SetAllowRegister(context.Context, *SetAllowRegisterReq) (*SetAllowRegisterResp, error)
	GetAllowRegister(context.Context, *GetAllowRegisterReq) (*GetAllowRegisterResp, error)
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChatServer struct{}

func (UnimplementedChatServer) UpdateUserInfo(context.Context, *UpdateUserInfoReq) (*UpdateUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (UnimplementedChatServer) AddUserAccount(context.Context, *AddUserAccountReq) (*AddUserAccountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserAccount not implemented")
}
func (UnimplementedChatServer) SearchUserPublicInfo(context.Context, *SearchUserPublicInfoReq) (*SearchUserPublicInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUserPublicInfo not implemented")
}
func (UnimplementedChatServer) FindUserPublicInfo(context.Context, *FindUserPublicInfoReq) (*FindUserPublicInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserPublicInfo not implemented")
}
func (UnimplementedChatServer) SearchUserFullInfo(context.Context, *SearchUserFullInfoReq) (*SearchUserFullInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUserFullInfo not implemented")
}
func (UnimplementedChatServer) FindUserFullInfo(context.Context, *FindUserFullInfoReq) (*FindUserFullInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserFullInfo not implemented")
}
func (UnimplementedChatServer) SendVerifyCode(context.Context, *SendVerifyCodeReq) (*SendVerifyCodeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendVerifyCode not implemented")
}
func (UnimplementedChatServer) VerifyCode(context.Context, *VerifyCodeReq) (*VerifyCodeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyCode not implemented")
}
func (UnimplementedChatServer) RegisterUser(context.Context, *RegisterUserReq) (*RegisterUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedChatServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedChatServer) ResetPassword(context.Context, *ResetPasswordReq) (*ResetPasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedChatServer) ChangePassword(context.Context, *ChangePasswordReq) (*ChangePasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedChatServer) CheckUserExist(context.Context, *CheckUserExistReq) (*CheckUserExistResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserExist not implemented")
}
func (UnimplementedChatServer) DelUserAccount(context.Context, *DelUserAccountReq) (*DelUserAccountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelUserAccount not implemented")
}
func (UnimplementedChatServer) FindUserAccount(context.Context, *FindUserAccountReq) (*FindUserAccountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserAccount not implemented")
}
func (UnimplementedChatServer) FindAccountUser(context.Context, *FindAccountUserReq) (*FindAccountUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAccountUser not implemented")
}
func (UnimplementedChatServer) OpenIMCallback(context.Context, *OpenIMCallbackReq) (*OpenIMCallbackResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenIMCallback not implemented")
}
func (UnimplementedChatServer) UserLoginCount(context.Context, *UserLoginCountReq) (*UserLoginCountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLoginCount not implemented")
}
func (UnimplementedChatServer) SearchUserInfo(context.Context, *SearchUserInfoReq) (*SearchUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUserInfo not implemented")
}
func (UnimplementedChatServer) GetTokenForVideoMeeting(context.Context, *GetTokenForVideoMeetingReq) (*GetTokenForVideoMeetingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenForVideoMeeting not implemented")
}
func (UnimplementedChatServer) SetAllowRegister(context.Context, *SetAllowRegisterReq) (*SetAllowRegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAllowRegister not implemented")
}
func (UnimplementedChatServer) GetAllowRegister(context.Context, *GetAllowRegisterReq) (*GetAllowRegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllowRegister not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}
func (UnimplementedChatServer) testEmbeddedByValue()              {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	// If the following call pancis, it indicates UnimplementedChatServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_UpdateUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).UpdateUserInfo(ctx, req.(*UpdateUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_AddUserAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).AddUserAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_AddUserAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).AddUserAccount(ctx, req.(*AddUserAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_SearchUserPublicInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserPublicInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SearchUserPublicInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_SearchUserPublicInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SearchUserPublicInfo(ctx, req.(*SearchUserPublicInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_FindUserPublicInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserPublicInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).FindUserPublicInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_FindUserPublicInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).FindUserPublicInfo(ctx, req.(*FindUserPublicInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_SearchUserFullInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserFullInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SearchUserFullInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_SearchUserFullInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SearchUserFullInfo(ctx, req.(*SearchUserFullInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_FindUserFullInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserFullInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).FindUserFullInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_FindUserFullInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).FindUserFullInfo(ctx, req.(*FindUserFullInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_SendVerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendVerifyCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SendVerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_SendVerifyCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SendVerifyCode(ctx, req.(*SendVerifyCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_VerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).VerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_VerifyCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).VerifyCode(ctx, req.(*VerifyCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_RegisterUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).RegisterUser(ctx, req.(*RegisterUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_ResetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).ResetPassword(ctx, req.(*ResetPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).ChangePassword(ctx, req.(*ChangePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_CheckUserExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserExistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).CheckUserExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_CheckUserExist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).CheckUserExist(ctx, req.(*CheckUserExistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_DelUserAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelUserAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).DelUserAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_DelUserAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).DelUserAccount(ctx, req.(*DelUserAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_FindUserAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).FindUserAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_FindUserAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).FindUserAccount(ctx, req.(*FindUserAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_FindAccountUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAccountUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).FindAccountUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_FindAccountUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).FindAccountUser(ctx, req.(*FindAccountUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_OpenIMCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenIMCallbackReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).OpenIMCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_OpenIMCallback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).OpenIMCallback(ctx, req.(*OpenIMCallbackReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_UserLoginCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).UserLoginCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_UserLoginCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).UserLoginCount(ctx, req.(*UserLoginCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_SearchUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SearchUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_SearchUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SearchUserInfo(ctx, req.(*SearchUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetTokenForVideoMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenForVideoMeetingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetTokenForVideoMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_GetTokenForVideoMeeting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetTokenForVideoMeeting(ctx, req.(*GetTokenForVideoMeetingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_SetAllowRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAllowRegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SetAllowRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_SetAllowRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SetAllowRegister(ctx, req.(*SetAllowRegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetAllowRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllowRegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetAllowRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_GetAllowRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetAllowRegister(ctx, req.(*GetAllowRegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openim.chat.chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateUserInfo",
			Handler:    _Chat_UpdateUserInfo_Handler,
		},
		{
			MethodName: "AddUserAccount",
			Handler:    _Chat_AddUserAccount_Handler,
		},
		{
			MethodName: "SearchUserPublicInfo",
			Handler:    _Chat_SearchUserPublicInfo_Handler,
		},
		{
			MethodName: "FindUserPublicInfo",
			Handler:    _Chat_FindUserPublicInfo_Handler,
		},
		{
			MethodName: "SearchUserFullInfo",
			Handler:    _Chat_SearchUserFullInfo_Handler,
		},
		{
			MethodName: "FindUserFullInfo",
			Handler:    _Chat_FindUserFullInfo_Handler,
		},
		{
			MethodName: "SendVerifyCode",
			Handler:    _Chat_SendVerifyCode_Handler,
		},
		{
			MethodName: "VerifyCode",
			Handler:    _Chat_VerifyCode_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _Chat_RegisterUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Chat_Login_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _Chat_ResetPassword_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _Chat_ChangePassword_Handler,
		},
		{
			MethodName: "CheckUserExist",
			Handler:    _Chat_CheckUserExist_Handler,
		},
		{
			MethodName: "DelUserAccount",
			Handler:    _Chat_DelUserAccount_Handler,
		},
		{
			MethodName: "FindUserAccount",
			Handler:    _Chat_FindUserAccount_Handler,
		},
		{
			MethodName: "FindAccountUser",
			Handler:    _Chat_FindAccountUser_Handler,
		},
		{
			MethodName: "OpenIMCallback",
			Handler:    _Chat_OpenIMCallback_Handler,
		},
		{
			MethodName: "UserLoginCount",
			Handler:    _Chat_UserLoginCount_Handler,
		},
		{
			MethodName: "SearchUserInfo",
			Handler:    _Chat_SearchUserInfo_Handler,
		},
		{
			MethodName: "GetTokenForVideoMeeting",
			Handler:    _Chat_GetTokenForVideoMeeting_Handler,
		},
		{
			MethodName: "SetAllowRegister",
			Handler:    _Chat_SetAllowRegister_Handler,
		},
		{
			MethodName: "GetAllowRegister",
			Handler:    _Chat_GetAllowRegister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat/chat.proto",
}
