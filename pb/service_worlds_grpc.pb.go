// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.0
// source: services/worlds/service_worlds.proto

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
	Worlds_CreateWorld_FullMethodName             = "/pb.Worlds/CreateWorld"
	Worlds_UpdateWorld_FullMethodName             = "/pb.Worlds/UpdateWorld"
	Worlds_UploadWorldImage_FullMethodName        = "/pb.Worlds/UploadWorldImage"
	Worlds_AddWorldTag_FullMethodName             = "/pb.Worlds/AddWorldTag"
	Worlds_RemoveWorldTag_FullMethodName          = "/pb.Worlds/RemoveWorldTag"
	Worlds_GetWorldAdmins_FullMethodName          = "/pb.Worlds/GetWorldAdmins"
	Worlds_CreateWorldAdmin_FullMethodName        = "/pb.Worlds/CreateWorldAdmin"
	Worlds_UpdateWorldAdmin_FullMethodName        = "/pb.Worlds/UpdateWorldAdmin"
	Worlds_DeleteWorldAdmin_FullMethodName        = "/pb.Worlds/DeleteWorldAdmin"
	Worlds_GetWorldDailyActivity_FullMethodName   = "/pb.Worlds/GetWorldDailyActivity"
	Worlds_GetWorldMonthlyActivity_FullMethodName = "/pb.Worlds/GetWorldMonthlyActivity"
	Worlds_GetWorlds_FullMethodName               = "/pb.Worlds/GetWorlds"
	Worlds_GetWorldById_FullMethodName            = "/pb.Worlds/GetWorldById"
	Worlds_UpdateWorldIntroduction_FullMethodName = "/pb.Worlds/UpdateWorldIntroduction"
)

// WorldsClient is the client API for Worlds service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorldsClient interface {
	CreateWorld(ctx context.Context, in *CreateWorldRequest, opts ...grpc.CallOption) (*World, error)
	UpdateWorld(ctx context.Context, in *UpdateWorldRequest, opts ...grpc.CallOption) (*World, error)
	UploadWorldImage(ctx context.Context, in *UploadWorldImageRequest, opts ...grpc.CallOption) (*Image, error)
	AddWorldTag(ctx context.Context, in *AddWorldTagRequest, opts ...grpc.CallOption) (*Tag, error)
	RemoveWorldTag(ctx context.Context, in *RemoveWorldTagRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetWorldAdmins(ctx context.Context, in *GetWorldAdminsRequest, opts ...grpc.CallOption) (*GetWorldAdminsResponse, error)
	CreateWorldAdmin(ctx context.Context, in *CreateWorldAdminRequest, opts ...grpc.CallOption) (*WorldAdmin, error)
	UpdateWorldAdmin(ctx context.Context, in *UpdateWorldAdminRequest, opts ...grpc.CallOption) (*WorldAdmin, error)
	DeleteWorldAdmin(ctx context.Context, in *DeleteWorldAdminRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetWorldDailyActivity(ctx context.Context, in *GetWorldDailyActivityRequest, opts ...grpc.CallOption) (*GetWorldDailyActivityResponse, error)
	GetWorldMonthlyActivity(ctx context.Context, in *GetWorldMonthlyActivityRequest, opts ...grpc.CallOption) (*GetWorldMonthlyActivityResponse, error)
	GetWorlds(ctx context.Context, in *GetWorldsRequest, opts ...grpc.CallOption) (*GetWorldsResponse, error)
	GetWorldById(ctx context.Context, in *GetWorldByIdRequest, opts ...grpc.CallOption) (*World, error)
	UpdateWorldIntroduction(ctx context.Context, in *UpdateWorldIntroductionRequest, opts ...grpc.CallOption) (*Post, error)
}

type worldsClient struct {
	cc grpc.ClientConnInterface
}

func NewWorldsClient(cc grpc.ClientConnInterface) WorldsClient {
	return &worldsClient{cc}
}

func (c *worldsClient) CreateWorld(ctx context.Context, in *CreateWorldRequest, opts ...grpc.CallOption) (*World, error) {
	out := new(World)
	err := c.cc.Invoke(ctx, Worlds_CreateWorld_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *worldsClient) UpdateWorld(ctx context.Context, in *UpdateWorldRequest, opts ...grpc.CallOption) (*World, error) {
	out := new(World)
	err := c.cc.Invoke(ctx, Worlds_UpdateWorld_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *worldsClient) UploadWorldImage(ctx context.Context, in *UploadWorldImageRequest, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := c.cc.Invoke(ctx, Worlds_UploadWorldImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *worldsClient) AddWorldTag(ctx context.Context, in *AddWorldTagRequest, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, Worlds_AddWorldTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *worldsClient) RemoveWorldTag(ctx context.Context, in *RemoveWorldTagRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Worlds_RemoveWorldTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *worldsClient) GetWorldAdmins(ctx context.Context, in *GetWorldAdminsRequest, opts ...grpc.CallOption) (*GetWorldAdminsResponse, error) {
	out := new(GetWorldAdminsResponse)
	err := c.cc.Invoke(ctx, Worlds_GetWorldAdmins_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *worldsClient) CreateWorldAdmin(ctx context.Context, in *CreateWorldAdminRequest, opts ...grpc.CallOption) (*WorldAdmin, error) {
	out := new(WorldAdmin)
	err := c.cc.Invoke(ctx, Worlds_CreateWorldAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *worldsClient) UpdateWorldAdmin(ctx context.Context, in *UpdateWorldAdminRequest, opts ...grpc.CallOption) (*WorldAdmin, error) {
	out := new(WorldAdmin)
	err := c.cc.Invoke(ctx, Worlds_UpdateWorldAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *worldsClient) DeleteWorldAdmin(ctx context.Context, in *DeleteWorldAdminRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Worlds_DeleteWorldAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *worldsClient) GetWorldDailyActivity(ctx context.Context, in *GetWorldDailyActivityRequest, opts ...grpc.CallOption) (*GetWorldDailyActivityResponse, error) {
	out := new(GetWorldDailyActivityResponse)
	err := c.cc.Invoke(ctx, Worlds_GetWorldDailyActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *worldsClient) GetWorldMonthlyActivity(ctx context.Context, in *GetWorldMonthlyActivityRequest, opts ...grpc.CallOption) (*GetWorldMonthlyActivityResponse, error) {
	out := new(GetWorldMonthlyActivityResponse)
	err := c.cc.Invoke(ctx, Worlds_GetWorldMonthlyActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *worldsClient) GetWorlds(ctx context.Context, in *GetWorldsRequest, opts ...grpc.CallOption) (*GetWorldsResponse, error) {
	out := new(GetWorldsResponse)
	err := c.cc.Invoke(ctx, Worlds_GetWorlds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *worldsClient) GetWorldById(ctx context.Context, in *GetWorldByIdRequest, opts ...grpc.CallOption) (*World, error) {
	out := new(World)
	err := c.cc.Invoke(ctx, Worlds_GetWorldById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *worldsClient) UpdateWorldIntroduction(ctx context.Context, in *UpdateWorldIntroductionRequest, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, Worlds_UpdateWorldIntroduction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorldsServer is the server API for Worlds service.
// All implementations must embed UnimplementedWorldsServer
// for forward compatibility
type WorldsServer interface {
	CreateWorld(context.Context, *CreateWorldRequest) (*World, error)
	UpdateWorld(context.Context, *UpdateWorldRequest) (*World, error)
	UploadWorldImage(context.Context, *UploadWorldImageRequest) (*Image, error)
	AddWorldTag(context.Context, *AddWorldTagRequest) (*Tag, error)
	RemoveWorldTag(context.Context, *RemoveWorldTagRequest) (*emptypb.Empty, error)
	GetWorldAdmins(context.Context, *GetWorldAdminsRequest) (*GetWorldAdminsResponse, error)
	CreateWorldAdmin(context.Context, *CreateWorldAdminRequest) (*WorldAdmin, error)
	UpdateWorldAdmin(context.Context, *UpdateWorldAdminRequest) (*WorldAdmin, error)
	DeleteWorldAdmin(context.Context, *DeleteWorldAdminRequest) (*emptypb.Empty, error)
	GetWorldDailyActivity(context.Context, *GetWorldDailyActivityRequest) (*GetWorldDailyActivityResponse, error)
	GetWorldMonthlyActivity(context.Context, *GetWorldMonthlyActivityRequest) (*GetWorldMonthlyActivityResponse, error)
	GetWorlds(context.Context, *GetWorldsRequest) (*GetWorldsResponse, error)
	GetWorldById(context.Context, *GetWorldByIdRequest) (*World, error)
	UpdateWorldIntroduction(context.Context, *UpdateWorldIntroductionRequest) (*Post, error)
	mustEmbedUnimplementedWorldsServer()
}

// UnimplementedWorldsServer must be embedded to have forward compatible implementations.
type UnimplementedWorldsServer struct {
}

func (UnimplementedWorldsServer) CreateWorld(context.Context, *CreateWorldRequest) (*World, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorld not implemented")
}
func (UnimplementedWorldsServer) UpdateWorld(context.Context, *UpdateWorldRequest) (*World, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorld not implemented")
}
func (UnimplementedWorldsServer) UploadWorldImage(context.Context, *UploadWorldImageRequest) (*Image, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadWorldImage not implemented")
}
func (UnimplementedWorldsServer) AddWorldTag(context.Context, *AddWorldTagRequest) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWorldTag not implemented")
}
func (UnimplementedWorldsServer) RemoveWorldTag(context.Context, *RemoveWorldTagRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveWorldTag not implemented")
}
func (UnimplementedWorldsServer) GetWorldAdmins(context.Context, *GetWorldAdminsRequest) (*GetWorldAdminsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorldAdmins not implemented")
}
func (UnimplementedWorldsServer) CreateWorldAdmin(context.Context, *CreateWorldAdminRequest) (*WorldAdmin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorldAdmin not implemented")
}
func (UnimplementedWorldsServer) UpdateWorldAdmin(context.Context, *UpdateWorldAdminRequest) (*WorldAdmin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorldAdmin not implemented")
}
func (UnimplementedWorldsServer) DeleteWorldAdmin(context.Context, *DeleteWorldAdminRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorldAdmin not implemented")
}
func (UnimplementedWorldsServer) GetWorldDailyActivity(context.Context, *GetWorldDailyActivityRequest) (*GetWorldDailyActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorldDailyActivity not implemented")
}
func (UnimplementedWorldsServer) GetWorldMonthlyActivity(context.Context, *GetWorldMonthlyActivityRequest) (*GetWorldMonthlyActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorldMonthlyActivity not implemented")
}
func (UnimplementedWorldsServer) GetWorlds(context.Context, *GetWorldsRequest) (*GetWorldsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorlds not implemented")
}
func (UnimplementedWorldsServer) GetWorldById(context.Context, *GetWorldByIdRequest) (*World, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorldById not implemented")
}
func (UnimplementedWorldsServer) UpdateWorldIntroduction(context.Context, *UpdateWorldIntroductionRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorldIntroduction not implemented")
}
func (UnimplementedWorldsServer) mustEmbedUnimplementedWorldsServer() {}

// UnsafeWorldsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorldsServer will
// result in compilation errors.
type UnsafeWorldsServer interface {
	mustEmbedUnimplementedWorldsServer()
}

func RegisterWorldsServer(s grpc.ServiceRegistrar, srv WorldsServer) {
	s.RegisterService(&Worlds_ServiceDesc, srv)
}

func _Worlds_CreateWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorldsServer).CreateWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worlds_CreateWorld_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorldsServer).CreateWorld(ctx, req.(*CreateWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worlds_UpdateWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorldsServer).UpdateWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worlds_UpdateWorld_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorldsServer).UpdateWorld(ctx, req.(*UpdateWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worlds_UploadWorldImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadWorldImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorldsServer).UploadWorldImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worlds_UploadWorldImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorldsServer).UploadWorldImage(ctx, req.(*UploadWorldImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worlds_AddWorldTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddWorldTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorldsServer).AddWorldTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worlds_AddWorldTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorldsServer).AddWorldTag(ctx, req.(*AddWorldTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worlds_RemoveWorldTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveWorldTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorldsServer).RemoveWorldTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worlds_RemoveWorldTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorldsServer).RemoveWorldTag(ctx, req.(*RemoveWorldTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worlds_GetWorldAdmins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorldAdminsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorldsServer).GetWorldAdmins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worlds_GetWorldAdmins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorldsServer).GetWorldAdmins(ctx, req.(*GetWorldAdminsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worlds_CreateWorldAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorldAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorldsServer).CreateWorldAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worlds_CreateWorldAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorldsServer).CreateWorldAdmin(ctx, req.(*CreateWorldAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worlds_UpdateWorldAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorldAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorldsServer).UpdateWorldAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worlds_UpdateWorldAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorldsServer).UpdateWorldAdmin(ctx, req.(*UpdateWorldAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worlds_DeleteWorldAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorldAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorldsServer).DeleteWorldAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worlds_DeleteWorldAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorldsServer).DeleteWorldAdmin(ctx, req.(*DeleteWorldAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worlds_GetWorldDailyActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorldDailyActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorldsServer).GetWorldDailyActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worlds_GetWorldDailyActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorldsServer).GetWorldDailyActivity(ctx, req.(*GetWorldDailyActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worlds_GetWorldMonthlyActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorldMonthlyActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorldsServer).GetWorldMonthlyActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worlds_GetWorldMonthlyActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorldsServer).GetWorldMonthlyActivity(ctx, req.(*GetWorldMonthlyActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worlds_GetWorlds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorldsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorldsServer).GetWorlds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worlds_GetWorlds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorldsServer).GetWorlds(ctx, req.(*GetWorldsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worlds_GetWorldById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorldByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorldsServer).GetWorldById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worlds_GetWorldById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorldsServer).GetWorldById(ctx, req.(*GetWorldByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worlds_UpdateWorldIntroduction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorldIntroductionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorldsServer).UpdateWorldIntroduction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worlds_UpdateWorldIntroduction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorldsServer).UpdateWorldIntroduction(ctx, req.(*UpdateWorldIntroductionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Worlds_ServiceDesc is the grpc.ServiceDesc for Worlds service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Worlds_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Worlds",
	HandlerType: (*WorldsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWorld",
			Handler:    _Worlds_CreateWorld_Handler,
		},
		{
			MethodName: "UpdateWorld",
			Handler:    _Worlds_UpdateWorld_Handler,
		},
		{
			MethodName: "UploadWorldImage",
			Handler:    _Worlds_UploadWorldImage_Handler,
		},
		{
			MethodName: "AddWorldTag",
			Handler:    _Worlds_AddWorldTag_Handler,
		},
		{
			MethodName: "RemoveWorldTag",
			Handler:    _Worlds_RemoveWorldTag_Handler,
		},
		{
			MethodName: "GetWorldAdmins",
			Handler:    _Worlds_GetWorldAdmins_Handler,
		},
		{
			MethodName: "CreateWorldAdmin",
			Handler:    _Worlds_CreateWorldAdmin_Handler,
		},
		{
			MethodName: "UpdateWorldAdmin",
			Handler:    _Worlds_UpdateWorldAdmin_Handler,
		},
		{
			MethodName: "DeleteWorldAdmin",
			Handler:    _Worlds_DeleteWorldAdmin_Handler,
		},
		{
			MethodName: "GetWorldDailyActivity",
			Handler:    _Worlds_GetWorldDailyActivity_Handler,
		},
		{
			MethodName: "GetWorldMonthlyActivity",
			Handler:    _Worlds_GetWorldMonthlyActivity_Handler,
		},
		{
			MethodName: "GetWorlds",
			Handler:    _Worlds_GetWorlds_Handler,
		},
		{
			MethodName: "GetWorldById",
			Handler:    _Worlds_GetWorldById_Handler,
		},
		{
			MethodName: "UpdateWorldIntroduction",
			Handler:    _Worlds_UpdateWorldIntroduction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/worlds/service_worlds.proto",
}