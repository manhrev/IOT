// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/info.proto

package info

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

// InfoClient is the client API for Info service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InfoClient interface {
	// Group
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupReply, error)
	ListGroup(ctx context.Context, in *ListGroupRequest, opts ...grpc.CallOption) (*ListGroupReply, error)
	DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteGroupReply, error)
	// Feed
	AddFeedsToGroup(ctx context.Context, in *AddFeedsToGroupRequest, opts ...grpc.CallOption) (*AddFeedsToGroupReply, error)
	CreateFeed(ctx context.Context, in *CreateFeedRequest, opts ...grpc.CallOption) (*CreateFeedReply, error)
	GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*GetFeedReply, error)
	ListFeed(ctx context.Context, in *ListFeedRequest, opts ...grpc.CallOption) (*ListFeedReply, error)
	DeleteFeed(ctx context.Context, in *DeleteFeedRequest, opts ...grpc.CallOption) (*DeleteFeedReply, error)
	// Data
	CreateData(ctx context.Context, in *CreateDataRequest, opts ...grpc.CallOption) (*CreateDataReply, error)
	ListData(ctx context.Context, in *ListDataRequest, opts ...grpc.CallOption) (*ListDataReply, error)
}

type infoClient struct {
	cc grpc.ClientConnInterface
}

func NewInfoClient(cc grpc.ClientConnInterface) InfoClient {
	return &infoClient{cc}
}

func (c *infoClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupReply, error) {
	out := new(CreateGroupReply)
	err := c.cc.Invoke(ctx, "/info.Info/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoClient) ListGroup(ctx context.Context, in *ListGroupRequest, opts ...grpc.CallOption) (*ListGroupReply, error) {
	out := new(ListGroupReply)
	err := c.cc.Invoke(ctx, "/info.Info/ListGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoClient) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteGroupReply, error) {
	out := new(DeleteGroupReply)
	err := c.cc.Invoke(ctx, "/info.Info/DeleteGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoClient) AddFeedsToGroup(ctx context.Context, in *AddFeedsToGroupRequest, opts ...grpc.CallOption) (*AddFeedsToGroupReply, error) {
	out := new(AddFeedsToGroupReply)
	err := c.cc.Invoke(ctx, "/info.Info/AddFeedsToGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoClient) CreateFeed(ctx context.Context, in *CreateFeedRequest, opts ...grpc.CallOption) (*CreateFeedReply, error) {
	out := new(CreateFeedReply)
	err := c.cc.Invoke(ctx, "/info.Info/CreateFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoClient) GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*GetFeedReply, error) {
	out := new(GetFeedReply)
	err := c.cc.Invoke(ctx, "/info.Info/GetFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoClient) ListFeed(ctx context.Context, in *ListFeedRequest, opts ...grpc.CallOption) (*ListFeedReply, error) {
	out := new(ListFeedReply)
	err := c.cc.Invoke(ctx, "/info.Info/ListFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoClient) DeleteFeed(ctx context.Context, in *DeleteFeedRequest, opts ...grpc.CallOption) (*DeleteFeedReply, error) {
	out := new(DeleteFeedReply)
	err := c.cc.Invoke(ctx, "/info.Info/DeleteFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoClient) CreateData(ctx context.Context, in *CreateDataRequest, opts ...grpc.CallOption) (*CreateDataReply, error) {
	out := new(CreateDataReply)
	err := c.cc.Invoke(ctx, "/info.Info/CreateData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoClient) ListData(ctx context.Context, in *ListDataRequest, opts ...grpc.CallOption) (*ListDataReply, error) {
	out := new(ListDataReply)
	err := c.cc.Invoke(ctx, "/info.Info/ListData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfoServer is the server API for Info service.
// All implementations must embed UnimplementedInfoServer
// for forward compatibility
type InfoServer interface {
	// Group
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupReply, error)
	ListGroup(context.Context, *ListGroupRequest) (*ListGroupReply, error)
	DeleteGroup(context.Context, *DeleteGroupRequest) (*DeleteGroupReply, error)
	// Feed
	AddFeedsToGroup(context.Context, *AddFeedsToGroupRequest) (*AddFeedsToGroupReply, error)
	CreateFeed(context.Context, *CreateFeedRequest) (*CreateFeedReply, error)
	GetFeed(context.Context, *GetFeedRequest) (*GetFeedReply, error)
	ListFeed(context.Context, *ListFeedRequest) (*ListFeedReply, error)
	DeleteFeed(context.Context, *DeleteFeedRequest) (*DeleteFeedReply, error)
	// Data
	CreateData(context.Context, *CreateDataRequest) (*CreateDataReply, error)
	ListData(context.Context, *ListDataRequest) (*ListDataReply, error)
	mustEmbedUnimplementedInfoServer()
}

// UnimplementedInfoServer must be embedded to have forward compatible implementations.
type UnimplementedInfoServer struct {
}

func (UnimplementedInfoServer) CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedInfoServer) ListGroup(context.Context, *ListGroupRequest) (*ListGroupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroup not implemented")
}
func (UnimplementedInfoServer) DeleteGroup(context.Context, *DeleteGroupRequest) (*DeleteGroupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedInfoServer) AddFeedsToGroup(context.Context, *AddFeedsToGroupRequest) (*AddFeedsToGroupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFeedsToGroup not implemented")
}
func (UnimplementedInfoServer) CreateFeed(context.Context, *CreateFeedRequest) (*CreateFeedReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFeed not implemented")
}
func (UnimplementedInfoServer) GetFeed(context.Context, *GetFeedRequest) (*GetFeedReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeed not implemented")
}
func (UnimplementedInfoServer) ListFeed(context.Context, *ListFeedRequest) (*ListFeedReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFeed not implemented")
}
func (UnimplementedInfoServer) DeleteFeed(context.Context, *DeleteFeedRequest) (*DeleteFeedReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFeed not implemented")
}
func (UnimplementedInfoServer) CreateData(context.Context, *CreateDataRequest) (*CreateDataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateData not implemented")
}
func (UnimplementedInfoServer) ListData(context.Context, *ListDataRequest) (*ListDataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListData not implemented")
}
func (UnimplementedInfoServer) mustEmbedUnimplementedInfoServer() {}

// UnsafeInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InfoServer will
// result in compilation errors.
type UnsafeInfoServer interface {
	mustEmbedUnimplementedInfoServer()
}

func RegisterInfoServer(s grpc.ServiceRegistrar, srv InfoServer) {
	s.RegisterService(&Info_ServiceDesc, srv)
}

func _Info_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.Info/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Info_ListGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).ListGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.Info/ListGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).ListGroup(ctx, req.(*ListGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Info_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.Info/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).DeleteGroup(ctx, req.(*DeleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Info_AddFeedsToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFeedsToGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).AddFeedsToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.Info/AddFeedsToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).AddFeedsToGroup(ctx, req.(*AddFeedsToGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Info_CreateFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).CreateFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.Info/CreateFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).CreateFeed(ctx, req.(*CreateFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Info_GetFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).GetFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.Info/GetFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).GetFeed(ctx, req.(*GetFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Info_ListFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).ListFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.Info/ListFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).ListFeed(ctx, req.(*ListFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Info_DeleteFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).DeleteFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.Info/DeleteFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).DeleteFeed(ctx, req.(*DeleteFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Info_CreateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).CreateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.Info/CreateData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).CreateData(ctx, req.(*CreateDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Info_ListData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).ListData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.Info/ListData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).ListData(ctx, req.(*ListDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Info_ServiceDesc is the grpc.ServiceDesc for Info service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Info_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "info.Info",
	HandlerType: (*InfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGroup",
			Handler:    _Info_CreateGroup_Handler,
		},
		{
			MethodName: "ListGroup",
			Handler:    _Info_ListGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _Info_DeleteGroup_Handler,
		},
		{
			MethodName: "AddFeedsToGroup",
			Handler:    _Info_AddFeedsToGroup_Handler,
		},
		{
			MethodName: "CreateFeed",
			Handler:    _Info_CreateFeed_Handler,
		},
		{
			MethodName: "GetFeed",
			Handler:    _Info_GetFeed_Handler,
		},
		{
			MethodName: "ListFeed",
			Handler:    _Info_ListFeed_Handler,
		},
		{
			MethodName: "DeleteFeed",
			Handler:    _Info_DeleteFeed_Handler,
		},
		{
			MethodName: "CreateData",
			Handler:    _Info_CreateData_Handler,
		},
		{
			MethodName: "ListData",
			Handler:    _Info_ListData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/info.proto",
}
