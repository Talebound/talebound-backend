// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.0
// source: services/auth/service_auth.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Auth_VerifyEmail_FullMethodName                     = "/pb.Auth/VerifyEmail"
	Auth_LoginUser_FullMethodName                       = "/pb.Auth/LoginUser"
	Auth_LogoutUser_FullMethodName                      = "/pb.Auth/LogoutUser"
	Auth_ResetPasswordSendCode_FullMethodName           = "/pb.Auth/ResetPasswordSendCode"
	Auth_ResetPasswordVerifyCode_FullMethodName         = "/pb.Auth/ResetPasswordVerifyCode"
	Auth_ResetPasswordVerifyCodeValidity_FullMethodName = "/pb.Auth/ResetPasswordVerifyCodeValidity"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	// ============= VERIFY =================
	VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error)
	// ============= LOGIN & LOGOUT =================
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	LogoutUser(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ResetPasswordSendCode(ctx context.Context, in *ResetPasswordSendCodeRequest, opts ...grpc.CallOption) (*ResetPasswordSendCodeResponse, error)
	ResetPasswordVerifyCode(ctx context.Context, in *ResetPasswordVerifyCodeRequest, opts ...grpc.CallOption) (*ResetPasswordVerifyCodeResponse, error)
	ResetPasswordVerifyCodeValidity(ctx context.Context, in *ResetPasswordVerifyCodeValidityRequest, opts ...grpc.CallOption) (*ResetPasswordVerifyCodeValidityResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error) {
	out := new(VerifyEmailResponse)
	err := c.cc.Invoke(ctx, Auth_VerifyEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, Auth_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) LogoutUser(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Auth_LogoutUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ResetPasswordSendCode(ctx context.Context, in *ResetPasswordSendCodeRequest, opts ...grpc.CallOption) (*ResetPasswordSendCodeResponse, error) {
	out := new(ResetPasswordSendCodeResponse)
	err := c.cc.Invoke(ctx, Auth_ResetPasswordSendCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ResetPasswordVerifyCode(ctx context.Context, in *ResetPasswordVerifyCodeRequest, opts ...grpc.CallOption) (*ResetPasswordVerifyCodeResponse, error) {
	out := new(ResetPasswordVerifyCodeResponse)
	err := c.cc.Invoke(ctx, Auth_ResetPasswordVerifyCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ResetPasswordVerifyCodeValidity(ctx context.Context, in *ResetPasswordVerifyCodeValidityRequest, opts ...grpc.CallOption) (*ResetPasswordVerifyCodeValidityResponse, error) {
	out := new(ResetPasswordVerifyCodeValidityResponse)
	err := c.cc.Invoke(ctx, Auth_ResetPasswordVerifyCodeValidity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	// ============= VERIFY =================
	VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error)
	// ============= LOGIN & LOGOUT =================
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	LogoutUser(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	ResetPasswordSendCode(context.Context, *ResetPasswordSendCodeRequest) (*ResetPasswordSendCodeResponse, error)
	ResetPasswordVerifyCode(context.Context, *ResetPasswordVerifyCodeRequest) (*ResetPasswordVerifyCodeResponse, error)
	ResetPasswordVerifyCodeValidity(context.Context, *ResetPasswordVerifyCodeValidityRequest) (*ResetPasswordVerifyCodeValidityResponse, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyEmail not implemented")
}
func (UnimplementedAuthServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedAuthServer) LogoutUser(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogoutUser not implemented")
}
func (UnimplementedAuthServer) ResetPasswordSendCode(context.Context, *ResetPasswordSendCodeRequest) (*ResetPasswordSendCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPasswordSendCode not implemented")
}
func (UnimplementedAuthServer) ResetPasswordVerifyCode(context.Context, *ResetPasswordVerifyCodeRequest) (*ResetPasswordVerifyCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPasswordVerifyCode not implemented")
}
func (UnimplementedAuthServer) ResetPasswordVerifyCodeValidity(context.Context, *ResetPasswordVerifyCodeValidityRequest) (*ResetPasswordVerifyCodeValidityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPasswordVerifyCodeValidity not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_VerifyEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).VerifyEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_VerifyEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).VerifyEmail(ctx, req.(*VerifyEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_LogoutUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).LogoutUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_LogoutUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).LogoutUser(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ResetPasswordSendCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordSendCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ResetPasswordSendCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_ResetPasswordSendCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ResetPasswordSendCode(ctx, req.(*ResetPasswordSendCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ResetPasswordVerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordVerifyCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ResetPasswordVerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_ResetPasswordVerifyCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ResetPasswordVerifyCode(ctx, req.(*ResetPasswordVerifyCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ResetPasswordVerifyCodeValidity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordVerifyCodeValidityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ResetPasswordVerifyCodeValidity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_ResetPasswordVerifyCodeValidity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ResetPasswordVerifyCodeValidity(ctx, req.(*ResetPasswordVerifyCodeValidityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyEmail",
			Handler:    _Auth_VerifyEmail_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _Auth_LoginUser_Handler,
		},
		{
			MethodName: "LogoutUser",
			Handler:    _Auth_LogoutUser_Handler,
		},
		{
			MethodName: "ResetPasswordSendCode",
			Handler:    _Auth_ResetPasswordSendCode_Handler,
		},
		{
			MethodName: "ResetPasswordVerifyCode",
			Handler:    _Auth_ResetPasswordVerifyCode_Handler,
		},
		{
			MethodName: "ResetPasswordVerifyCodeValidity",
			Handler:    _Auth_ResetPasswordVerifyCodeValidity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/auth/service_auth.proto",
}