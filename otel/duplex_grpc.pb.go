// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: duplex.proto

package otel

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Otel_Create_FullMethodName = "/Otel/Create"
)

// OtelClient is the client API for Otel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OtelClient interface {
	Create(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type otelClient struct {
	cc grpc.ClientConnInterface
}

func NewOtelClient(cc grpc.ClientConnInterface) OtelClient {
	return &otelClient{cc}
}

func (c *otelClient) Create(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Otel_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OtelServer is the server API for Otel service.
// All implementations must embed UnimplementedOtelServer
// for forward compatibility
type OtelServer interface {
	Create(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedOtelServer()
}

// UnimplementedOtelServer must be embedded to have forward compatible implementations.
type UnimplementedOtelServer struct {
}

func (UnimplementedOtelServer) Create(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOtelServer) mustEmbedUnimplementedOtelServer() {}

// UnsafeOtelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OtelServer will
// result in compilation errors.
type UnsafeOtelServer interface {
	mustEmbedUnimplementedOtelServer()
}

func RegisterOtelServer(s grpc.ServiceRegistrar, srv OtelServer) {
	s.RegisterService(&Otel_ServiceDesc, srv)
}

func _Otel_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OtelServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Otel_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OtelServer).Create(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Otel_ServiceDesc is the grpc.ServiceDesc for Otel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Otel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Otel",
	HandlerType: (*OtelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Otel_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "duplex.proto",
}
