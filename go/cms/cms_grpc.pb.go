// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: cms.proto

package cms

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

// CmsClient is the client API for Cms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CmsClient interface {
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

type cmsClient struct {
	cc grpc.ClientConnInterface
}

func NewCmsClient(cc grpc.ClientConnInterface) CmsClient {
	return &cmsClient{cc}
}

func (c *cmsClient) AppList(ctx context.Context, in *base.ListReq, opts ...grpc.CallOption) (*AppListResp, error) {
	out := new(AppListResp)
	err := c.cc.Invoke(ctx, "/cms.cms/AppList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsClient) AppDetail(ctx context.Context, in *base.IdReq, opts ...grpc.CallOption) (*AppItem, error) {
	out := new(AppItem)
	err := c.cc.Invoke(ctx, "/cms.cms/AppDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsClient) AppUpdateOrCreate(ctx context.Context, in *AppItem, opts ...grpc.CallOption) (*base.Empty, error) {
	out := new(base.Empty)
	err := c.cc.Invoke(ctx, "/cms.cms/AppUpdateOrCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsClient) AppDelete(ctx context.Context, in *base.IdReq, opts ...grpc.CallOption) (*base.Empty, error) {
	out := new(base.Empty)
	err := c.cc.Invoke(ctx, "/cms.cms/AppDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsClient) AppReset(ctx context.Context, in *base.IdReq, opts ...grpc.CallOption) (*AppResetResp, error) {
	out := new(AppResetResp)
	err := c.cc.Invoke(ctx, "/cms.cms/AppReset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CmsServer is the server API for Cms service.
// All implementations must embed UnimplementedCmsServer
// for forward compatibility
type CmsServer interface {
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
	mustEmbedUnimplementedCmsServer()
}

// UnimplementedCmsServer must be embedded to have forward compatible implementations.
type UnimplementedCmsServer struct {
}

func (UnimplementedCmsServer) AppList(context.Context, *base.ListReq) (*AppListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppList not implemented")
}
func (UnimplementedCmsServer) AppDetail(context.Context, *base.IdReq) (*AppItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppDetail not implemented")
}
func (UnimplementedCmsServer) AppUpdateOrCreate(context.Context, *AppItem) (*base.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppUpdateOrCreate not implemented")
}
func (UnimplementedCmsServer) AppDelete(context.Context, *base.IdReq) (*base.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppDelete not implemented")
}
func (UnimplementedCmsServer) AppReset(context.Context, *base.IdReq) (*AppResetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppReset not implemented")
}
func (UnimplementedCmsServer) mustEmbedUnimplementedCmsServer() {}

// UnsafeCmsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CmsServer will
// result in compilation errors.
type UnsafeCmsServer interface {
	mustEmbedUnimplementedCmsServer()
}

func RegisterCmsServer(s grpc.ServiceRegistrar, srv CmsServer) {
	s.RegisterService(&Cms_ServiceDesc, srv)
}

func _Cms_AppList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(base.ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServer).AppList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.cms/AppList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServer).AppList(ctx, req.(*base.ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cms_AppDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(base.IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServer).AppDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.cms/AppDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServer).AppDetail(ctx, req.(*base.IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cms_AppUpdateOrCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServer).AppUpdateOrCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.cms/AppUpdateOrCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServer).AppUpdateOrCreate(ctx, req.(*AppItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cms_AppDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(base.IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServer).AppDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.cms/AppDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServer).AppDelete(ctx, req.(*base.IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cms_AppReset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(base.IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServer).AppReset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.cms/AppReset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServer).AppReset(ctx, req.(*base.IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Cms_ServiceDesc is the grpc.ServiceDesc for Cms service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cms_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cms.cms",
	HandlerType: (*CmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppList",
			Handler:    _Cms_AppList_Handler,
		},
		{
			MethodName: "AppDetail",
			Handler:    _Cms_AppDetail_Handler,
		},
		{
			MethodName: "AppUpdateOrCreate",
			Handler:    _Cms_AppUpdateOrCreate_Handler,
		},
		{
			MethodName: "AppDelete",
			Handler:    _Cms_AppDelete_Handler,
		},
		{
			MethodName: "AppReset",
			Handler:    _Cms_AppReset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cms.proto",
}
