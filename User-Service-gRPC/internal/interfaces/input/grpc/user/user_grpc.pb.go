// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.14.0
// source: internal/interfaces/input/grpc/user/user.proto

package user

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
	UserService_FindUserById_FullMethodName        = "/user.UserService/FindUserById"
	UserService_FindUsersListFromID_FullMethodName = "/user.UserService/FindUsersListFromID"
	UserService_FindUserByFilter_FullMethodName    = "/user.UserService/FindUserByFilter"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	FindUserById(ctx context.Context, in *FindUserByIdInput, opts ...grpc.CallOption) (*FindUserByIdResponse, error)
	FindUsersListFromID(ctx context.Context, in *FindUsersListFromIDReq, opts ...grpc.CallOption) (*FindUsersListFromIDResponse, error)
	FindUserByFilter(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*FindUserByFilterResp, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) FindUserById(ctx context.Context, in *FindUserByIdInput, opts ...grpc.CallOption) (*FindUserByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindUserByIdResponse)
	err := c.cc.Invoke(ctx, UserService_FindUserById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FindUsersListFromID(ctx context.Context, in *FindUsersListFromIDReq, opts ...grpc.CallOption) (*FindUsersListFromIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindUsersListFromIDResponse)
	err := c.cc.Invoke(ctx, UserService_FindUsersListFromID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FindUserByFilter(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*FindUserByFilterResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindUserByFilterResp)
	err := c.cc.Invoke(ctx, UserService_FindUserByFilter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	FindUserById(context.Context, *FindUserByIdInput) (*FindUserByIdResponse, error)
	FindUsersListFromID(context.Context, *FindUsersListFromIDReq) (*FindUsersListFromIDResponse, error)
	FindUserByFilter(context.Context, *Filter) (*FindUserByFilterResp, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) FindUserById(context.Context, *FindUserByIdInput) (*FindUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserById not implemented")
}
func (UnimplementedUserServiceServer) FindUsersListFromID(context.Context, *FindUsersListFromIDReq) (*FindUsersListFromIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUsersListFromID not implemented")
}
func (UnimplementedUserServiceServer) FindUserByFilter(context.Context, *Filter) (*FindUserByFilterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserByFilter not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_FindUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserByIdInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_FindUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindUserById(ctx, req.(*FindUserByIdInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FindUsersListFromID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUsersListFromIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindUsersListFromID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_FindUsersListFromID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindUsersListFromID(ctx, req.(*FindUsersListFromIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FindUserByFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindUserByFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_FindUserByFilter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindUserByFilter(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindUserById",
			Handler:    _UserService_FindUserById_Handler,
		},
		{
			MethodName: "FindUsersListFromID",
			Handler:    _UserService_FindUsersListFromID_Handler,
		},
		{
			MethodName: "FindUserByFilter",
			Handler:    _UserService_FindUserByFilter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/interfaces/input/grpc/user/user.proto",
}
