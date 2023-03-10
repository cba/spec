// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: file.proto

package file

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

// FileClient is the client API for File service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileClient interface {
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

type fileClient struct {
	cc grpc.ClientConnInterface
}

func NewFileClient(cc grpc.ClientConnInterface) FileClient {
	return &fileClient{cc}
}

func (c *fileClient) AppList(ctx context.Context, in *base.ListReq, opts ...grpc.CallOption) (*AppListResp, error) {
	out := new(AppListResp)
	err := c.cc.Invoke(ctx, "/file.file/AppList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) AppDetail(ctx context.Context, in *base.IdReq, opts ...grpc.CallOption) (*AppItem, error) {
	out := new(AppItem)
	err := c.cc.Invoke(ctx, "/file.file/AppDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) AppUpdateOrCreate(ctx context.Context, in *AppItem, opts ...grpc.CallOption) (*base.Empty, error) {
	out := new(base.Empty)
	err := c.cc.Invoke(ctx, "/file.file/AppUpdateOrCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) AppDelete(ctx context.Context, in *base.IdReq, opts ...grpc.CallOption) (*base.Empty, error) {
	out := new(base.Empty)
	err := c.cc.Invoke(ctx, "/file.file/AppDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) AppReset(ctx context.Context, in *base.IdReq, opts ...grpc.CallOption) (*AppResetResp, error) {
	out := new(AppResetResp)
	err := c.cc.Invoke(ctx, "/file.file/AppReset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServer is the server API for File service.
// All implementations must embed UnimplementedFileServer
// for forward compatibility
type FileServer interface {
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
	mustEmbedUnimplementedFileServer()
}

// UnimplementedFileServer must be embedded to have forward compatible implementations.
type UnimplementedFileServer struct {
}

func (UnimplementedFileServer) AppList(context.Context, *base.ListReq) (*AppListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppList not implemented")
}
func (UnimplementedFileServer) AppDetail(context.Context, *base.IdReq) (*AppItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppDetail not implemented")
}
func (UnimplementedFileServer) AppUpdateOrCreate(context.Context, *AppItem) (*base.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppUpdateOrCreate not implemented")
}
func (UnimplementedFileServer) AppDelete(context.Context, *base.IdReq) (*base.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppDelete not implemented")
}
func (UnimplementedFileServer) AppReset(context.Context, *base.IdReq) (*AppResetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppReset not implemented")
}
func (UnimplementedFileServer) mustEmbedUnimplementedFileServer() {}

// UnsafeFileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServer will
// result in compilation errors.
type UnsafeFileServer interface {
	mustEmbedUnimplementedFileServer()
}

func RegisterFileServer(s grpc.ServiceRegistrar, srv FileServer) {
	s.RegisterService(&File_ServiceDesc, srv)
}

func _File_AppList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(base.ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).AppList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.file/AppList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).AppList(ctx, req.(*base.ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_AppDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(base.IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).AppDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.file/AppDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).AppDetail(ctx, req.(*base.IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_AppUpdateOrCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).AppUpdateOrCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.file/AppUpdateOrCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).AppUpdateOrCreate(ctx, req.(*AppItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_AppDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(base.IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).AppDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.file/AppDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).AppDelete(ctx, req.(*base.IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_AppReset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(base.IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).AppReset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.file/AppReset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).AppReset(ctx, req.(*base.IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// File_ServiceDesc is the grpc.ServiceDesc for File service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var File_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "file.file",
	HandlerType: (*FileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppList",
			Handler:    _File_AppList_Handler,
		},
		{
			MethodName: "AppDetail",
			Handler:    _File_AppDetail_Handler,
		},
		{
			MethodName: "AppUpdateOrCreate",
			Handler:    _File_AppUpdateOrCreate_Handler,
		},
		{
			MethodName: "AppDelete",
			Handler:    _File_AppDelete_Handler,
		},
		{
			MethodName: "AppReset",
			Handler:    _File_AppReset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "file.proto",
}
