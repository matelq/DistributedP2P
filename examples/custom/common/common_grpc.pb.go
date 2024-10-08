// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--rc1
// source: common.proto

package common

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	P2PManager_SendMessage_FullMethodName = "/common.P2PManager/SendMessage"
)

// P2PManagerClient is the client API for P2PManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type P2PManagerClient interface {
	SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Echo, error)
}

type p2PManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewP2PManagerClient(cc grpc.ClientConnInterface) P2PManagerClient {
	return &p2PManagerClient{cc}
}

func (c *p2PManagerClient) SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Echo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Echo)
	err := c.cc.Invoke(ctx, P2PManager_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// P2PManagerServer is the server API for P2PManager service.
// All implementations must embed UnimplementedP2PManagerServer
// for forward compatibility.
type P2PManagerServer interface {
	SendMessage(context.Context, *Message) (*Echo, error)
	mustEmbedUnimplementedP2PManagerServer()
}

// UnimplementedP2PManagerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedP2PManagerServer struct{}

func (UnimplementedP2PManagerServer) SendMessage(context.Context, *Message) (*Echo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedP2PManagerServer) mustEmbedUnimplementedP2PManagerServer() {}
func (UnimplementedP2PManagerServer) testEmbeddedByValue()                    {}

// UnsafeP2PManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to P2PManagerServer will
// result in compilation errors.
type UnsafeP2PManagerServer interface {
	mustEmbedUnimplementedP2PManagerServer()
}

func RegisterP2PManagerServer(s grpc.ServiceRegistrar, srv P2PManagerServer) {
	// If the following call pancis, it indicates UnimplementedP2PManagerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&P2PManager_ServiceDesc, srv)
}

func _P2PManager_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PManagerServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: P2PManager_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PManagerServer).SendMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// P2PManager_ServiceDesc is the grpc.ServiceDesc for P2PManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var P2PManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "common.P2PManager",
	HandlerType: (*P2PManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _P2PManager_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common.proto",
}
