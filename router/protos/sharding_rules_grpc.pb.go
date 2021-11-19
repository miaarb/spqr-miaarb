// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// ShardingRulesServiceClient is the client API for ShardingRulesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShardingRulesServiceClient interface {
	AddShardingRules(ctx context.Context, in *AddShardingRuleRequest, opts ...grpc.CallOption) (*AddShardingRuleReply, error)
	ListShardingRules(ctx context.Context, in *AddShardingRuleRequest, opts ...grpc.CallOption) (*ListShardingRuleReply, error)
}

type shardingRulesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShardingRulesServiceClient(cc grpc.ClientConnInterface) ShardingRulesServiceClient {
	return &shardingRulesServiceClient{cc}
}

func (c *shardingRulesServiceClient) AddShardingRules(ctx context.Context, in *AddShardingRuleRequest, opts ...grpc.CallOption) (*AddShardingRuleReply, error) {
	out := new(AddShardingRuleReply)
	err := c.cc.Invoke(ctx, "/yandex.spqr.ShardingRulesService/AddShardingRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shardingRulesServiceClient) ListShardingRules(ctx context.Context, in *AddShardingRuleRequest, opts ...grpc.CallOption) (*ListShardingRuleReply, error) {
	out := new(ListShardingRuleReply)
	err := c.cc.Invoke(ctx, "/yandex.spqr.ShardingRulesService/ListShardingRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShardingRulesServiceServer is the server API for ShardingRulesService service.
// All implementations must embed UnimplementedShardingRulesServiceServer
// for forward compatibility
type ShardingRulesServiceServer interface {
	AddShardingRules(context.Context, *AddShardingRuleRequest) (*AddShardingRuleReply, error)
	ListShardingRules(context.Context, *AddShardingRuleRequest) (*ListShardingRuleReply, error)
	mustEmbedUnimplementedShardingRulesServiceServer()
}

// UnimplementedShardingRulesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShardingRulesServiceServer struct {
}

func (UnimplementedShardingRulesServiceServer) AddShardingRules(context.Context, *AddShardingRuleRequest) (*AddShardingRuleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddShardingRules not implemented")
}
func (UnimplementedShardingRulesServiceServer) ListShardingRules(context.Context, *AddShardingRuleRequest) (*ListShardingRuleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShardingRules not implemented")
}
func (UnimplementedShardingRulesServiceServer) mustEmbedUnimplementedShardingRulesServiceServer() {}

// UnsafeShardingRulesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShardingRulesServiceServer will
// result in compilation errors.
type UnsafeShardingRulesServiceServer interface {
	mustEmbedUnimplementedShardingRulesServiceServer()
}

func RegisterShardingRulesServiceServer(s grpc.ServiceRegistrar, srv ShardingRulesServiceServer) {
	s.RegisterService(&ShardingRulesService_ServiceDesc, srv)
}

func _ShardingRulesService_AddShardingRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddShardingRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardingRulesServiceServer).AddShardingRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.spqr.ShardingRulesService/AddShardingRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardingRulesServiceServer).AddShardingRules(ctx, req.(*AddShardingRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShardingRulesService_ListShardingRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddShardingRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardingRulesServiceServer).ListShardingRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.spqr.ShardingRulesService/ListShardingRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardingRulesServiceServer).ListShardingRules(ctx, req.(*AddShardingRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShardingRulesService_ServiceDesc is the grpc.ServiceDesc for ShardingRulesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShardingRulesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.spqr.ShardingRulesService",
	HandlerType: (*ShardingRulesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddShardingRules",
			Handler:    _ShardingRulesService_AddShardingRules_Handler,
		},
		{
			MethodName: "ListShardingRules",
			Handler:    _ShardingRulesService_ListShardingRules_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/sharding_rules.proto",
}
