// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sgtmpb

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// WebAPIClient is the client API for WebAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebAPIClient interface {
	//rpc Register(Register.Request) returns (Register.Response) { option (google.api.http) = {post: "/api/v1/Register", body: "*"}; }
	UserList(ctx context.Context, in *UserList_Request, opts ...grpc.CallOption) (*UserList_Response, error)
	PostList(ctx context.Context, in *PostList_Request, opts ...grpc.CallOption) (*PostList_Response, error)
	// rpc PostSync(PostSync.Request) returns (PostSync.Response) { option (google.api.http) = {get: "/api/v1/PostSync"}; }
	Me(ctx context.Context, in *Me_Request, opts ...grpc.CallOption) (*Me_Response, error)
	Ping(ctx context.Context, in *Ping_Request, opts ...grpc.CallOption) (*Ping_Response, error)
	Status(ctx context.Context, in *Status_Request, opts ...grpc.CallOption) (*Status_Response, error)
}

type webAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewWebAPIClient(cc grpc.ClientConnInterface) WebAPIClient {
	return &webAPIClient{cc}
}

func (c *webAPIClient) UserList(ctx context.Context, in *UserList_Request, opts ...grpc.CallOption) (*UserList_Response, error) {
	out := new(UserList_Response)
	err := c.cc.Invoke(ctx, "/sgtm.WebAPI/UserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webAPIClient) PostList(ctx context.Context, in *PostList_Request, opts ...grpc.CallOption) (*PostList_Response, error) {
	out := new(PostList_Response)
	err := c.cc.Invoke(ctx, "/sgtm.WebAPI/PostList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webAPIClient) Me(ctx context.Context, in *Me_Request, opts ...grpc.CallOption) (*Me_Response, error) {
	out := new(Me_Response)
	err := c.cc.Invoke(ctx, "/sgtm.WebAPI/Me", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webAPIClient) Ping(ctx context.Context, in *Ping_Request, opts ...grpc.CallOption) (*Ping_Response, error) {
	out := new(Ping_Response)
	err := c.cc.Invoke(ctx, "/sgtm.WebAPI/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webAPIClient) Status(ctx context.Context, in *Status_Request, opts ...grpc.CallOption) (*Status_Response, error) {
	out := new(Status_Response)
	err := c.cc.Invoke(ctx, "/sgtm.WebAPI/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebAPIServer is the server API for WebAPI service.
// All implementations must embed UnimplementedWebAPIServer
// for forward compatibility
type WebAPIServer interface {
	//rpc Register(Register.Request) returns (Register.Response) { option (google.api.http) = {post: "/api/v1/Register", body: "*"}; }
	UserList(context.Context, *UserList_Request) (*UserList_Response, error)
	PostList(context.Context, *PostList_Request) (*PostList_Response, error)
	// rpc PostSync(PostSync.Request) returns (PostSync.Response) { option (google.api.http) = {get: "/api/v1/PostSync"}; }
	Me(context.Context, *Me_Request) (*Me_Response, error)
	Ping(context.Context, *Ping_Request) (*Ping_Response, error)
	Status(context.Context, *Status_Request) (*Status_Response, error)
	mustEmbedUnimplementedWebAPIServer()
}

// UnimplementedWebAPIServer must be embedded to have forward compatible implementations.
type UnimplementedWebAPIServer struct {
}

func (UnimplementedWebAPIServer) UserList(context.Context, *UserList_Request) (*UserList_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (UnimplementedWebAPIServer) PostList(context.Context, *PostList_Request) (*PostList_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostList not implemented")
}
func (UnimplementedWebAPIServer) Me(context.Context, *Me_Request) (*Me_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Me not implemented")
}
func (UnimplementedWebAPIServer) Ping(context.Context, *Ping_Request) (*Ping_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedWebAPIServer) Status(context.Context, *Status_Request) (*Status_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedWebAPIServer) mustEmbedUnimplementedWebAPIServer() {}

// UnsafeWebAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebAPIServer will
// result in compilation errors.
type UnsafeWebAPIServer interface {
	mustEmbedUnimplementedWebAPIServer()
}

func RegisterWebAPIServer(s grpc.ServiceRegistrar, srv WebAPIServer) {
	s.RegisterService(&_WebAPI_serviceDesc, srv)
}

func _WebAPI_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserList_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebAPIServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sgtm.WebAPI/UserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebAPIServer).UserList(ctx, req.(*UserList_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebAPI_PostList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostList_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		re