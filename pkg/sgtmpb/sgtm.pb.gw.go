// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: sgtm.proto

/*
Package sgtmpb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package sgtmpb

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage
var _ = metadata.Join

func request_WebAPI_UserList_0(ctx context.Context, marshaler runtime.Marshaler, client WebAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UserList_Request
	var metadata runtime.ServerMetadata

	msg, err := client.UserList(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_WebAPI_UserList_0(ctx context.Context, marshaler runtime.Marshaler, server WebAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UserList_Request
	var metadata runtime.ServerMetadata

	msg, err := server.UserList(ctx, &protoReq)
	return msg, metadata, err

}

func request_WebAPI_PostList_0(ctx context.Context, marshaler runtime.Marshaler, client WebAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PostList_Request
	var metadata runtime.ServerMetadata

	msg, err := client.PostList(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_WebAPI_PostList_0(ctx context.Context, marshaler runtime.Marshaler, server WebAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PostList_Request
	var metadata runtime.ServerMetadata

	msg, err := server.PostList(ctx, &protoReq)
	return msg, metadata, err

}

func request_WebAPI_Me_0(ctx context.Context, marshaler runtime.Marshaler, client WebAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Me_Request
	var metadata runtime.ServerMetadata

	msg, err := client.Me(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_WebAPI_Me_0(ctx context.Context, marshaler runtime.Marshaler, server WebAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Me_Request
	var metadata runtime.ServerMetadata

	msg, err := server.Me(ctx, &protoReq)
	return msg, metadata, err

}

func request_WebAPI_Ping_0(ctx context.Context, marshaler runtime.Marshaler, client WebAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Ping_Request
	var metadata runtime.ServerMetadata

	msg, err := client.Ping(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_WebAPI_Ping_0(ctx context.Context, marshaler runtime.Marshaler, server WebAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Ping_Request
	var metadata runtime.ServerMetadata

	msg, err := server.Ping(ctx, &protoReq)
	return msg, metadata, err

}

func request_WebAPI_Status_0