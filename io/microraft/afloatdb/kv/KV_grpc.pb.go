// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: KV.proto

package kv

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

// KVRequestHandlerClient is the client API for KVRequestHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KVRequestHandlerClient interface {
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*KVResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*KVResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*KVResponse, error)
	Contains(ctx context.Context, in *ContainsRequest, opts ...grpc.CallOption) (*KVResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*KVResponse, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*KVResponse, error)
	Replace(ctx context.Context, in *ReplaceRequest, opts ...grpc.CallOption) (*KVResponse, error)
	Size(ctx context.Context, in *SizeRequest, opts ...grpc.CallOption) (*KVResponse, error)
	Clear(ctx context.Context, in *ClearRequest, opts ...grpc.CallOption) (*KVResponse, error)
}

type kVRequestHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewKVRequestHandlerClient(cc grpc.ClientConnInterface) KVRequestHandlerClient {
	return &kVRequestHandlerClient{cc}
}

func (c *kVRequestHandlerClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*KVResponse, error) {
	out := new(KVResponse)
	err := c.cc.Invoke(ctx, "/io.afloatdb.kv.proto.KVRequestHandler/put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVRequestHandlerClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*KVResponse, error) {
	out := new(KVResponse)
	err := c.cc.Invoke(ctx, "/io.afloatdb.kv.proto.KVRequestHandler/set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVRequestHandlerClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*KVResponse, error) {
	out := new(KVResponse)
	err := c.cc.Invoke(ctx, "/io.afloatdb.kv.proto.KVRequestHandler/get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVRequestHandlerClient) Contains(ctx context.Context, in *ContainsRequest, opts ...grpc.CallOption) (*KVResponse, error) {
	out := new(KVResponse)
	err := c.cc.Invoke(ctx, "/io.afloatdb.kv.proto.KVRequestHandler/contains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVRequestHandlerClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*KVResponse, error) {
	out := new(KVResponse)
	err := c.cc.Invoke(ctx, "/io.afloatdb.kv.proto.KVRequestHandler/delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVRequestHandlerClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*KVResponse, error) {
	out := new(KVResponse)
	err := c.cc.Invoke(ctx, "/io.afloatdb.kv.proto.KVRequestHandler/remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVRequestHandlerClient) Replace(ctx context.Context, in *ReplaceRequest, opts ...grpc.CallOption) (*KVResponse, error) {
	out := new(KVResponse)
	err := c.cc.Invoke(ctx, "/io.afloatdb.kv.proto.KVRequestHandler/replace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVRequestHandlerClient) Size(ctx context.Context, in *SizeRequest, opts ...grpc.CallOption) (*KVResponse, error) {
	out := new(KVResponse)
	err := c.cc.Invoke(ctx, "/io.afloatdb.kv.proto.KVRequestHandler/size", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVRequestHandlerClient) Clear(ctx context.Context, in *ClearRequest, opts ...grpc.CallOption) (*KVResponse, error) {
	out := new(KVResponse)
	err := c.cc.Invoke(ctx, "/io.afloatdb.kv.proto.KVRequestHandler/clear", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVRequestHandlerServer is the server API for KVRequestHandler service.
// All implementations must embed UnimplementedKVRequestHandlerServer
// for forward compatibility
type KVRequestHandlerServer interface {
	Put(context.Context, *PutRequest) (*KVResponse, error)
	Set(context.Context, *SetRequest) (*KVResponse, error)
	Get(context.Context, *GetRequest) (*KVResponse, error)
	Contains(context.Context, *ContainsRequest) (*KVResponse, error)
	Delete(context.Context, *DeleteRequest) (*KVResponse, error)
	Remove(context.Context, *RemoveRequest) (*KVResponse, error)
	Replace(context.Context, *ReplaceRequest) (*KVResponse, error)
	Size(context.Context, *SizeRequest) (*KVResponse, error)
	Clear(context.Context, *ClearRequest) (*KVResponse, error)
	mustEmbedUnimplementedKVRequestHandlerServer()
}

// UnimplementedKVRequestHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedKVRequestHandlerServer struct {
}

func (UnimplementedKVRequestHandlerServer) Put(context.Context, *PutRequest) (*KVResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedKVRequestHandlerServer) Set(context.Context, *SetRequest) (*KVResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedKVRequestHandlerServer) Get(context.Context, *GetRequest) (*KVResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedKVRequestHandlerServer) Contains(context.Context, *ContainsRequest) (*KVResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Contains not implemented")
}
func (UnimplementedKVRequestHandlerServer) Delete(context.Context, *DeleteRequest) (*KVResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedKVRequestHandlerServer) Remove(context.Context, *RemoveRequest) (*KVResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedKVRequestHandlerServer) Replace(context.Context, *ReplaceRequest) (*KVResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Replace not implemented")
}
func (UnimplementedKVRequestHandlerServer) Size(context.Context, *SizeRequest) (*KVResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Size not implemented")
}
func (UnimplementedKVRequestHandlerServer) Clear(context.Context, *ClearRequest) (*KVResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clear not implemented")
}
func (UnimplementedKVRequestHandlerServer) mustEmbedUnimplementedKVRequestHandlerServer() {}

// UnsafeKVRequestHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KVRequestHandlerServer will
// result in compilation errors.
type UnsafeKVRequestHandlerServer interface {
	mustEmbedUnimplementedKVRequestHandlerServer()
}

func RegisterKVRequestHandlerServer(s grpc.ServiceRegistrar, srv KVRequestHandlerServer) {
	s.RegisterService(&KVRequestHandler_ServiceDesc, srv)
}

func _KVRequestHandler_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVRequestHandlerServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.afloatdb.kv.proto.KVRequestHandler/put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVRequestHandlerServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVRequestHandler_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVRequestHandlerServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.afloatdb.kv.proto.KVRequestHandler/set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVRequestHandlerServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVRequestHandler_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVRequestHandlerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.afloatdb.kv.proto.KVRequestHandler/get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVRequestHandlerServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVRequestHandler_Contains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVRequestHandlerServer).Contains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.afloatdb.kv.proto.KVRequestHandler/contains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVRequestHandlerServer).Contains(ctx, req.(*ContainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVRequestHandler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVRequestHandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.afloatdb.kv.proto.KVRequestHandler/delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVRequestHandlerServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVRequestHandler_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVRequestHandlerServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.afloatdb.kv.proto.KVRequestHandler/remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVRequestHandlerServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVRequestHandler_Replace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVRequestHandlerServer).Replace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.afloatdb.kv.proto.KVRequestHandler/replace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVRequestHandlerServer).Replace(ctx, req.(*ReplaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVRequestHandler_Size_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVRequestHandlerServer).Size(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.afloatdb.kv.proto.KVRequestHandler/size",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVRequestHandlerServer).Size(ctx, req.(*SizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVRequestHandler_Clear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVRequestHandlerServer).Clear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.afloatdb.kv.proto.KVRequestHandler/clear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVRequestHandlerServer).Clear(ctx, req.(*ClearRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KVRequestHandler_ServiceDesc is the grpc.ServiceDesc for KVRequestHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KVRequestHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "io.afloatdb.kv.proto.KVRequestHandler",
	HandlerType: (*KVRequestHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "put",
			Handler:    _KVRequestHandler_Put_Handler,
		},
		{
			MethodName: "set",
			Handler:    _KVRequestHandler_Set_Handler,
		},
		{
			MethodName: "get",
			Handler:    _KVRequestHandler_Get_Handler,
		},
		{
			MethodName: "contains",
			Handler:    _KVRequestHandler_Contains_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _KVRequestHandler_Delete_Handler,
		},
		{
			MethodName: "remove",
			Handler:    _KVRequestHandler_Remove_Handler,
		},
		{
			MethodName: "replace",
			Handler:    _KVRequestHandler_Replace_Handler,
		},
		{
			MethodName: "size",
			Handler:    _KVRequestHandler_Size_Handler,
		},
		{
			MethodName: "clear",
			Handler:    _KVRequestHandler_Clear_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "KV.proto",
}
