// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: sms.proto

package sms

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

// SmsClient is the client API for Sms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmsClient interface {
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

type smsClient struct {
	cc grpc.ClientConnInterface
}

func NewSmsClient(cc grpc.ClientConnInterface) SmsClient {
	return &smsClient{cc}
}

func (c *smsClient) AppList(ctx context.Context, in *base.ListReq, opts ...grpc.CallOption) (*AppListResp, error) {
	out := new(AppListResp)
	err := c.cc.Invoke(ctx, "/sms.sms/AppList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) AppDetail(ctx context.Context, in *base.IdReq, opts ...grpc.CallOption) (*AppItem, error) {
	out := new(AppItem)
	err := c.cc.Invoke(ctx, "/sms.sms/AppDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) AppUpdateOrCreate(ctx context.Context, in *AppItem, opts ...grpc.CallOption) (*base.Empty, error) {
	out := new(base.Empty)
	err := c.cc.Invoke(ctx, "/sms.sms/AppUpdateOrCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) AppDelete(ctx context.Context, in *base.IdReq, opts ...grpc.CallOption) (*base.Empty, error) {
	out := new(base.Empty)
	err := c.cc.Invoke(ctx, "/sms.sms/AppDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) AppReset(ctx context.Context, in *base.IdReq, opts ...grpc.CallOption) (*AppResetResp, error) {
	out := new(AppResetResp)
	err := c.cc.Invoke(ctx, "/sms.sms/AppReset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsServer is the server API for Sms service.
// All implementations must embed UnimplementedSmsServer
// for forward compatibility
type SmsServer interface {
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
	mustEmbedUnimplementedSmsServer()
}

// UnimplementedSmsServer must be embedded to have forward compatible implementations.
type UnimplementedSmsServer struct {
}

func (UnimplementedSmsServer) AppList(context.Context, *base.ListReq) (*AppListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppList not implemented")
}
func (UnimplementedSmsServer) AppDetail(context.Context, *base.IdReq) (*AppItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppDetail not implemented")
}
func (UnimplementedSmsServer) AppUpdateOrCreate(context.Context, *AppItem) (*base.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppUpdateOrCreate not implemented")
}
func (UnimplementedSmsServer) AppDelete(context.Context, *base.IdReq) (*base.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppDelete not implemented")
}
func (UnimplementedSmsServer) AppReset(context.Context, *base.IdReq) (*AppResetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppReset not implemented")
}
func (UnimplementedSmsServer) mustEmbedUnimplementedSmsServer() {}

// UnsafeSmsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmsServer will
// result in compilation errors.
type UnsafeSmsServer interface {
	mustEmbedUnimplementedSmsServer()
}

func RegisterSmsServer(s grpc.ServiceRegistrar, srv SmsServer) {
	s.RegisterService(&Sms_ServiceDesc, srv)
}

func _Sms_AppList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(base.ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).AppList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sms.sms/AppList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).AppList(ctx, req.(*base.ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_AppDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(base.IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).AppDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sms.sms/AppDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).AppDetail(ctx, req.(*base.IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_AppUpdateOrCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).AppUpdateOrCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sms.sms/AppUpdateOrCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).AppUpdateOrCreate(ctx, req.(*AppItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_AppDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(base.IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).AppDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sms.sms/AppDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).AppDelete(ctx, req.(*base.IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_AppReset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(base.IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).AppReset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sms.sms/AppReset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).AppReset(ctx, req.(*base.IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Sms_ServiceDesc is the grpc.ServiceDesc for Sms service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sms_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sms.sms",
	HandlerType: (*SmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppList",
			Handler:    _Sms_AppList_Handler,
		},
		{
			MethodName: "AppDetail",
			Handler:    _Sms_AppDetail_Handler,
		},
		{
			MethodName: "AppUpdateOrCreate",
			Handler:    _Sms_AppUpdateOrCreate_Handler,
		},
		{
			MethodName: "AppDelete",
			Handler:    _Sms_AppDelete_Handler,
		},
		{
			MethodName: "AppReset",
			Handler:    _Sms_AppReset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sms.proto",
}
