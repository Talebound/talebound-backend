// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.0
// source: services/posts/service_posts.proto

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
	Posts_GetPostById_FullMethodName        = "/pb.Posts/GetPostById"
	Posts_GetPosts_FullMethodName           = "/pb.Posts/GetPosts"
	Posts_GetPostHistory_FullMethodName     = "/pb.Posts/GetPostHistory"
	Posts_GetPostHistoryById_FullMethodName = "/pb.Posts/GetPostHistoryById"
	Posts_CreatePost_FullMethodName         = "/pb.Posts/CreatePost"
	Posts_UpdatePost_FullMethodName         = "/pb.Posts/UpdatePost"
	Posts_DeletePost_FullMethodName         = "/pb.Posts/DeletePost"
)

// PostsClient is the client API for Posts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostsClient interface {
	GetPostById(ctx context.Context, in *GetPostByIdRequest, opts ...grpc.CallOption) (*Post, error)
	GetPosts(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (*GetPostsResponse, error)
	GetPostHistory(ctx context.Context, in *GetPostHistoryRequest, opts ...grpc.CallOption) (*GetPostHistoryResponse, error)
	GetPostHistoryById(ctx context.Context, in *GetPostHistoryByIdRequest, opts ...grpc.CallOption) (*PostHistory, error)
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*Post, error)
	UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*Post, error)
	DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error)
}

type postsClient struct {
	cc grpc.ClientConnInterface
}

func NewPostsClient(cc grpc.ClientConnInterface) PostsClient {
	return &postsClient{cc}
}

func (c *postsClient) GetPostById(ctx context.Context, in *GetPostByIdRequest, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, Posts_GetPostById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsClient) GetPosts(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (*GetPostsResponse, error) {
	out := new(GetPostsResponse)
	err := c.cc.Invoke(ctx, Posts_GetPosts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsClient) GetPostHistory(ctx context.Context, in *GetPostHistoryRequest, opts ...grpc.CallOption) (*GetPostHistoryResponse, error) {
	out := new(GetPostHistoryResponse)
	err := c.cc.Invoke(ctx, Posts_GetPostHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsClient) GetPostHistoryById(ctx context.Context, in *GetPostHistoryByIdRequest, opts ...grpc.CallOption) (*PostHistory, error) {
	out := new(PostHistory)
	err := c.cc.Invoke(ctx, Posts_GetPostHistoryById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, Posts_CreatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsClient) UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, Posts_UpdatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsClient) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error) {
	out := new(DeletePostResponse)
	err := c.cc.Invoke(ctx, Posts_DeletePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostsServer is the server API for Posts service.
// All implementations must embed UnimplementedPostsServer
// for forward compatibility
type PostsServer interface {
	GetPostById(context.Context, *GetPostByIdRequest) (*Post, error)
	GetPosts(context.Context, *GetPostsRequest) (*GetPostsResponse, error)
	GetPostHistory(context.Context, *GetPostHistoryRequest) (*GetPostHistoryResponse, error)
	GetPostHistoryById(context.Context, *GetPostHistoryByIdRequest) (*PostHistory, error)
	CreatePost(context.Context, *CreatePostRequest) (*Post, error)
	UpdatePost(context.Context, *UpdatePostRequest) (*Post, error)
	DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error)
	mustEmbedUnimplementedPostsServer()
}

// UnimplementedPostsServer must be embedded to have forward compatible implementations.
type UnimplementedPostsServer struct {
}

func (UnimplementedPostsServer) GetPostById(context.Context, *GetPostByIdRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostById not implemented")
}
func (UnimplementedPostsServer) GetPosts(context.Context, *GetPostsRequest) (*GetPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPosts not implemented")
}
func (UnimplementedPostsServer) GetPostHistory(context.Context, *GetPostHistoryRequest) (*GetPostHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostHistory not implemented")
}
func (UnimplementedPostsServer) GetPostHistoryById(context.Context, *GetPostHistoryByIdRequest) (*PostHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostHistoryById not implemented")
}
func (UnimplementedPostsServer) CreatePost(context.Context, *CreatePostRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedPostsServer) UpdatePost(context.Context, *UpdatePostRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedPostsServer) DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedPostsServer) mustEmbedUnimplementedPostsServer() {}

// UnsafePostsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostsServer will
// result in compilation errors.
type UnsafePostsServer interface {
	mustEmbedUnimplementedPostsServer()
}

func RegisterPostsServer(s grpc.ServiceRegistrar, srv PostsServer) {
	s.RegisterService(&Posts_ServiceDesc, srv)
}

func _Posts_GetPostById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).GetPostById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Posts_GetPostById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).GetPostById(ctx, req.(*GetPostByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posts_GetPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).GetPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Posts_GetPosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).GetPosts(ctx, req.(*GetPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posts_GetPostHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).GetPostHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Posts_GetPostHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).GetPostHistory(ctx, req.(*GetPostHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posts_GetPostHistoryById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostHistoryByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).GetPostHistoryById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Posts_GetPostHistoryById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).GetPostHistoryById(ctx, req.(*GetPostHistoryByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posts_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Posts_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posts_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Posts_UpdatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).UpdatePost(ctx, req.(*UpdatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posts_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Posts_DeletePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).DeletePost(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Posts_ServiceDesc is the grpc.ServiceDesc for Posts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Posts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Posts",
	HandlerType: (*PostsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPostById",
			Handler:    _Posts_GetPostById_Handler,
		},
		{
			MethodName: "GetPosts",
			Handler:    _Posts_GetPosts_Handler,
		},
		{
			MethodName: "GetPostHistory",
			Handler:    _Posts_GetPostHistory_Handler,
		},
		{
			MethodName: "GetPostHistoryById",
			Handler:    _Posts_GetPostHistoryById_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _Posts_CreatePost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _Posts_UpdatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _Posts_DeletePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/posts/service_posts.proto",
}
