// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: mcs.proto

package grpc

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

// AutherClient is the client API for Auther service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AutherClient interface {
	Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	Login(ctx context.Context, in *UserBasic, opts ...grpc.CallOption) (*User, error)
}

type autherClient struct {
	cc grpc.ClientConnInterface
}

func NewAutherClient(cc grpc.ClientConnInterface) AutherClient {
	return &autherClient{cc}
}

func (c *autherClient) Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/Auth.Auther/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autherClient) Login(ctx context.Context, in *UserBasic, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/Auth.Auther/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AutherServer is the server API for Auther service.
// All implementations must embed UnimplementedAutherServer
// for forward compatibility
type AutherServer interface {
	Register(context.Context, *User) (*User, error)
	Login(context.Context, *UserBasic) (*User, error)
	mustEmbedUnimplementedAutherServer()
}

// UnimplementedAutherServer must be embedded to have forward compatible implementations.
type UnimplementedAutherServer struct {
}

func (UnimplementedAutherServer) Register(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAutherServer) Login(context.Context, *UserBasic) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAutherServer) mustEmbedUnimplementedAutherServer() {}

// UnsafeAutherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AutherServer will
// result in compilation errors.
type UnsafeAutherServer interface {
	mustEmbedUnimplementedAutherServer()
}

func RegisterAutherServer(s grpc.ServiceRegistrar, srv AutherServer) {
	s.RegisterService(&Auther_ServiceDesc, srv)
}

func _Auther_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutherServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth.Auther/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutherServer).Register(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auther_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBasic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutherServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth.Auther/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutherServer).Login(ctx, req.(*UserBasic))
	}
	return interceptor(ctx, in, info, handler)
}

// Auther_ServiceDesc is the grpc.ServiceDesc for Auther service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auther_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Auth.Auther",
	HandlerType: (*AutherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Auther_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Auther_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mcs.proto",
}
