// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: pb/svc/cache/cache.proto

package cache

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
	Cache_WriteMsg_FullMethodName = "/pb.svc.cache.Cache/WriteMsg"
	Cache_GetMsg_FullMethodName   = "/pb.svc.cache.Cache/GetMsg"
)

// CacheClient is the client API for Cache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CacheClient interface {
	WriteMsg(ctx context.Context, in *WriteMsgReq, opts ...grpc.CallOption) (*WriteMsgRes, error)
	GetMsg(ctx context.Context, in *GetMsgReq, opts ...grpc.CallOption) (*GetMsgRes, error)
}

type cacheClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheClient(cc grpc.ClientConnInterface) CacheClient {
	return &cacheClient{cc}
}

func (c *cacheClient) WriteMsg(ctx context.Context, in *WriteMsgReq, opts ...grpc.CallOption) (*WriteMsgRes, error) {
	out := new(WriteMsgRes)
	err := c.cc.Invoke(ctx, Cache_WriteMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) GetMsg(ctx context.Context, in *GetMsgReq, opts ...grpc.CallOption) (*GetMsgRes, error) {
	out := new(GetMsgRes)
	err := c.cc.Invoke(ctx, Cache_GetMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServer is the server API for Cache service.
// All implementations must embed UnimplementedCacheServer
// for forward compatibility
type CacheServer interface {
	WriteMsg(context.Context, *WriteMsgReq) (*WriteMsgRes, error)
	GetMsg(context.Context, *GetMsgReq) (*GetMsgRes, error)
	mustEmbedUnimplementedCacheServer()
}

// UnimplementedCacheServer must be embedded to have forward compatible implementations.
type UnimplementedCacheServer struct {
}

func (UnimplementedCacheServer) WriteMsg(context.Context, *WriteMsgReq) (*WriteMsgRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteMsg not implemented")
}
func (UnimplementedCacheServer) GetMsg(context.Context, *GetMsgReq) (*GetMsgRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMsg not implemented")
}
func (UnimplementedCacheServer) mustEmbedUnimplementedCacheServer() {}

// UnsafeCacheServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CacheServer will
// result in compilation errors.
type UnsafeCacheServer interface {
	mustEmbedUnimplementedCacheServer()
}

func RegisterCacheServer(s grpc.ServiceRegistrar, srv CacheServer) {
	s.RegisterService(&Cache_ServiceDesc, srv)
}

func _Cache_WriteMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).WriteMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cache_WriteMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).WriteMsg(ctx, req.(*WriteMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_GetMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).GetMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cache_GetMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).GetMsg(ctx, req.(*GetMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Cache_ServiceDesc is the grpc.ServiceDesc for Cache service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cache_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.svc.cache.Cache",
	HandlerType: (*CacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteMsg",
			Handler:    _Cache_WriteMsg_Handler,
		},
		{
			MethodName: "GetMsg",
			Handler:    _Cache_GetMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/svc/cache/cache.proto",
}