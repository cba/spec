// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: user.proto

package user

import (
	context "context"
	base "github.com/cba/spec/go/base"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	// APP列表
	AppList(ctx context.Context, in *base.ListReq, opts ...grpc.CallOption) (*AppListResp, error)
	// APP详情
	AppDetail(ctx context.Context, in *base.IdReq, opts ...grpc.CallOption) (*AppItem, error)
	// APP更新或创建
	AppUpdateOrCreate(ctx context.Context, in *AppItem, opts ...grpc.CallOption) (*base.Empty, error)
	// APP删除
	AppDelete(ctx context.Context, in *base.IdReq, opts ...grpc.CallOption) (*base.Empty, error)
	// APP重置Secret
	AppReset(ctx context.Context, in *base.IdReq, opts ...grpc.CallOption) (*AppResetResp, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) AppList(ctx context.Context, in *base.ListReq, opts ...grpc.CallOption) (*AppListResp, error) {
	out := new(AppListResp)
	err := c.cc.Invoke(ctx, "/user.user/AppList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AppDetail(ctx context.Context, in *base.IdReq, opts ...grpc.CallOption) (*AppItem, error) {
	out := new(AppItem)
	err := c.cc.Invoke(ctx, "/user.user/AppDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AppUpdateOrCreate(ctx context.Context, in *AppItem, opts ...grpc.CallOption) (*base.Empty, error) {
	out := new(base.Empty)
	err := c.cc.Invoke(ctx, "/user.user/AppUpdateOrCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AppDelete(ctx context.Context, in *base.IdReq, opts ...grpc.CallOption) (*base.Empty, error) {
	out := new(base.Empty)
	err := c.cc.Invoke(ctx, "/user.user/AppDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AppReset(ctx context.Context, in *base.IdReq, opts ...grpc.CallOption) (*AppResetResp, error) {
	out := new(AppResetResp)
	err := c.cc.Invoke(ctx, "/user.user/AppReset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	// APP列表
	AppList(context.Context, *base.ListReq) (*AppListResp, error)
	// APP详情
	AppDetail(context.Context, *base.IdReq) (*AppItem, error)
	// APP更新或创建
	AppUpdateOrCreate(context.Context, *AppItem) (*base.Empty, error)
	// APP删除
	AppDelete(context.Context, *base.IdReq) (*base.Empty, error)
	// APP重置Secret
	AppReset(context.Context, *base.IdReq) (*AppResetResp, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) AppList(context.Context, *base.ListReq) (*AppListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppList not implemented")
}
func (UnimplementedUserServer) AppDetail(context.Context, *base.IdReq) (*AppItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppDetail not implemented")
}
func (UnimplementedUserServer) AppUpdateOrCreate(context.Context, *AppItem) (*base.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppUpdateOrCreate not implemented")
}
func (UnimplementedUserServer) AppDelete(context.Context, *base.IdReq) (*base.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppDelete not implemented")
}
func (UnimplementedUserServer) AppReset(context.Context, *base.IdReq) (*AppResetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppReset not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_AppList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(base.ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AppList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.user/AppList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AppList(ctx, req.(*base.ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AppDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(base.IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AppDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.user/AppDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AppDetail(ctx, req.(*base.IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AppUpdateOrCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AppUpdateOrCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.user/AppUpdateOrCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AppUpdateOrCreate(ctx, req.(*AppItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AppDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(base.IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AppDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.user/AppDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AppDelete(ctx, req.(*base.IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AppReset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(base.IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AppReset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.user/AppReset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AppReset(ctx, req.(*base.IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.user",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppList",
			Handler:    _User_AppList_Handler,
		},
		{
			MethodName: "AppDetail",
			Handler:    _User_AppDetail_Handler,
		},
		{
			MethodName: "AppUpdateOrCreate",
			Handler:    _User_AppUpdateOrCreate_Handler,
		},
		{
			MethodName: "AppDelete",
			Handler:    _User_AppDelete_Handler,
		},
		{
			MethodName: "AppReset",
			Handler:    _User_AppReset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
