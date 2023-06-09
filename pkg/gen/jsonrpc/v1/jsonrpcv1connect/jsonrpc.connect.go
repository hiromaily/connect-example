// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: jsonrpc/v1/jsonrpc.proto

package jsonrpcv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/hiromaily/connect-example/pkg/gen/jsonrpc/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// JSONRPCGreetServiceName is the fully-qualified name of the JSONRPCGreetService service.
	JSONRPCGreetServiceName = "jsonrpc.v1.JSONRPCGreetService"
)

// JSONRPCGreetServiceClient is a client for the jsonrpc.v1.JSONRPCGreetService service.
type JSONRPCGreetServiceClient interface {
	Greet(context.Context, *connect_go.Request[v1.JSONRPCRequest]) (*connect_go.Response[v1.JSONRPCResponse], error)
}

// NewJSONRPCGreetServiceClient constructs a client for the jsonrpc.v1.JSONRPCGreetService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewJSONRPCGreetServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) JSONRPCGreetServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &jSONRPCGreetServiceClient{
		greet: connect_go.NewClient[v1.JSONRPCRequest, v1.JSONRPCResponse](
			httpClient,
			baseURL+"/jsonrpc.v1.JSONRPCGreetService/Greet",
			opts...,
		),
	}
}

// jSONRPCGreetServiceClient implements JSONRPCGreetServiceClient.
type jSONRPCGreetServiceClient struct {
	greet *connect_go.Client[v1.JSONRPCRequest, v1.JSONRPCResponse]
}

// Greet calls jsonrpc.v1.JSONRPCGreetService.Greet.
func (c *jSONRPCGreetServiceClient) Greet(ctx context.Context, req *connect_go.Request[v1.JSONRPCRequest]) (*connect_go.Response[v1.JSONRPCResponse], error) {
	return c.greet.CallUnary(ctx, req)
}

// JSONRPCGreetServiceHandler is an implementation of the jsonrpc.v1.JSONRPCGreetService service.
type JSONRPCGreetServiceHandler interface {
	Greet(context.Context, *connect_go.Request[v1.JSONRPCRequest]) (*connect_go.Response[v1.JSONRPCResponse], error)
}

// NewJSONRPCGreetServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewJSONRPCGreetServiceHandler(svc JSONRPCGreetServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/jsonrpc.v1.JSONRPCGreetService/Greet", connect_go.NewUnaryHandler(
		"/jsonrpc.v1.JSONRPCGreetService/Greet",
		svc.Greet,
		opts...,
	))
	return "/jsonrpc.v1.JSONRPCGreetService/", mux
}

// UnimplementedJSONRPCGreetServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedJSONRPCGreetServiceHandler struct{}

func (UnimplementedJSONRPCGreetServiceHandler) Greet(context.Context, *connect_go.Request[v1.JSONRPCRequest]) (*connect_go.Response[v1.JSONRPCResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("jsonrpc.v1.JSONRPCGreetService.Greet is not implemented"))
}
