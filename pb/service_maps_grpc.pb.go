// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.0
// source: services/maps/service_maps.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Maps_UpdateMap_FullMethodName = "/pb.Maps/UpdateMap"
)

// MapsClient is the client API for Maps service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MapsClient interface {
	UpdateMap(ctx context.Context, in *UpdateMapRequest, opts ...grpc.CallOption) (*UpdateMapResponse, error)
}

type mapsClient struct {
	cc grpc.ClientConnInterface
}

func NewMapsClient(cc grpc.ClientConnInterface) MapsClient {
	return &mapsClient{cc}
}

func (c *mapsClient) UpdateMap(ctx context.Context, in *UpdateMapRequest, opts ...grpc.CallOption) (*UpdateMapResponse, error) {
	out := new(UpdateMapResponse)
	err := c.cc.Invoke(ctx, Maps_UpdateMap_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MapsServer is the server API for Maps service.
// All implementations must embed UnimplementedMapsServer
// for forward compatibility
type MapsServer interface {
	UpdateMap(context.Context, *UpdateMapRequest) (*UpdateMapResponse, error)
	mustEmbedUnimplementedMapsServer()
}

// UnimplementedMapsServer must be embedded to have forward compatible implementations.
type UnimplementedMapsServer struct {
}

func (UnimplementedMapsServer) UpdateMap(context.Context, *UpdateMapRequest) (*UpdateMapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMap not implemented")
}
func (UnimplementedMapsServer) mustEmbedUnimplementedMapsServer() {}

// UnsafeMapsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MapsServer will
// result in compilation errors.
type UnsafeMapsServer interface {
	mustEmbedUnimplementedMapsServer()
}

func RegisterMapsServer(s grpc.ServiceRegistrar, srv MapsServer) {
	s.RegisterService(&Maps_ServiceDesc, srv)
}

func _Maps_UpdateMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsServer).UpdateMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Maps_UpdateMap_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsServer).UpdateMap(ctx, req.(*UpdateMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Maps_ServiceDesc is the grpc.ServiceDesc for Maps service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Maps_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Maps",
	HandlerType: (*MapsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateMap",
			Handler:    _Maps_UpdateMap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/maps/service_maps.proto",
}