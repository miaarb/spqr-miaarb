// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: protos/coordinator.proto

package proto

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
	TopologyService_OpenRouter_FullMethodName      = "/spqr.TopologyService/OpenRouter"
	TopologyService_GetRouterStatus_FullMethodName = "/spqr.TopologyService/GetRouterStatus"
	TopologyService_CloseRouter_FullMethodName     = "/spqr.TopologyService/CloseRouter"
)

// TopologyServiceClient is the client API for TopologyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TopologyServiceClient interface {
	OpenRouter(ctx context.Context, in *OpenRouterRequest, opts ...grpc.CallOption) (*OpenRouterReply, error)
	GetRouterStatus(ctx context.Context, in *GetRouterStatusRequest, opts ...grpc.CallOption) (*GetRouterStatusReply, error)
	CloseRouter(ctx context.Context, in *CloseRouterRequest, opts ...grpc.CallOption) (*CloseRouterReply, error)
}

type topologyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTopologyServiceClient(cc grpc.ClientConnInterface) TopologyServiceClient {
	return &topologyServiceClient{cc}
}

func (c *topologyServiceClient) OpenRouter(ctx context.Context, in *OpenRouterRequest, opts ...grpc.CallOption) (*OpenRouterReply, error) {
	out := new(OpenRouterReply)
	err := c.cc.Invoke(ctx, TopologyService_OpenRouter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topologyServiceClient) GetRouterStatus(ctx context.Context, in *GetRouterStatusRequest, opts ...grpc.CallOption) (*GetRouterStatusReply, error) {
	out := new(GetRouterStatusReply)
	err := c.cc.Invoke(ctx, TopologyService_GetRouterStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topologyServiceClient) CloseRouter(ctx context.Context, in *CloseRouterRequest, opts ...grpc.CallOption) (*CloseRouterReply, error) {
	out := new(CloseRouterReply)
	err := c.cc.Invoke(ctx, TopologyService_CloseRouter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopologyServiceServer is the server API for TopologyService service.
// All implementations must embed UnimplementedTopologyServiceServer
// for forward compatibility
type TopologyServiceServer interface {
	OpenRouter(context.Context, *OpenRouterRequest) (*OpenRouterReply, error)
	GetRouterStatus(context.Context, *GetRouterStatusRequest) (*GetRouterStatusReply, error)
	CloseRouter(context.Context, *CloseRouterRequest) (*CloseRouterReply, error)
	mustEmbedUnimplementedTopologyServiceServer()
}

// UnimplementedTopologyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTopologyServiceServer struct {
}

func (UnimplementedTopologyServiceServer) OpenRouter(context.Context, *OpenRouterRequest) (*OpenRouterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenRouter not implemented")
}
func (UnimplementedTopologyServiceServer) GetRouterStatus(context.Context, *GetRouterStatusRequest) (*GetRouterStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRouterStatus not implemented")
}
func (UnimplementedTopologyServiceServer) CloseRouter(context.Context, *CloseRouterRequest) (*CloseRouterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseRouter not implemented")
}
func (UnimplementedTopologyServiceServer) mustEmbedUnimplementedTopologyServiceServer() {}

// UnsafeTopologyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TopologyServiceServer will
// result in compilation errors.
type UnsafeTopologyServiceServer interface {
	mustEmbedUnimplementedTopologyServiceServer()
}

func RegisterTopologyServiceServer(s grpc.ServiceRegistrar, srv TopologyServiceServer) {
	s.RegisterService(&TopologyService_ServiceDesc, srv)
}

func _TopologyService_OpenRouter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenRouterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopologyServiceServer).OpenRouter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopologyService_OpenRouter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopologyServiceServer).OpenRouter(ctx, req.(*OpenRouterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopologyService_GetRouterStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRouterStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopologyServiceServer).GetRouterStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopologyService_GetRouterStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopologyServiceServer).GetRouterStatus(ctx, req.(*GetRouterStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopologyService_CloseRouter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseRouterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopologyServiceServer).CloseRouter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopologyService_CloseRouter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopologyServiceServer).CloseRouter(ctx, req.(*CloseRouterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TopologyService_ServiceDesc is the grpc.ServiceDesc for TopologyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TopologyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spqr.TopologyService",
	HandlerType: (*TopologyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenRouter",
			Handler:    _TopologyService_OpenRouter_Handler,
		},
		{
			MethodName: "GetRouterStatus",
			Handler:    _TopologyService_GetRouterStatus_Handler,
		},
		{
			MethodName: "CloseRouter",
			Handler:    _TopologyService_CloseRouter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/coordinator.proto",
}
