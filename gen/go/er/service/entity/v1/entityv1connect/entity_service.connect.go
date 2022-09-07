// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: er/service/entity/v1/entity_service.proto

package entityv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/xmlking/entity-resolution/gen/go/er/schema/entity/v1"
	v11 "github.com/xmlking/entity-resolution/gen/go/er/service/entity/v1"
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
	// EntityServiceName is the fully-qualified name of the EntityService service.
	EntityServiceName = "er.service.entity.v1.EntityService"
)

// EntityServiceClient is a client for the er.service.entity.v1.EntityService service.
type EntityServiceClient interface {
	Ingest(context.Context, *connect_go.Request[v1.Member]) (*connect_go.Response[v11.Response], error)
	// rpc IngestStream(stream er.schema.entity.v1.Member) returns (stream Response) {}
	Inquiry(context.Context, *connect_go.Request[v1.Member]) (*connect_go.Response[v11.Response], error)
}

// NewEntityServiceClient constructs a client for the er.service.entity.v1.EntityService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewEntityServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) EntityServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &entityServiceClient{
		ingest: connect_go.NewClient[v1.Member, v11.Response](
			httpClient,
			baseURL+"/er.service.entity.v1.EntityService/Ingest",
			opts...,
		),
		inquiry: connect_go.NewClient[v1.Member, v11.Response](
			httpClient,
			baseURL+"/er.service.entity.v1.EntityService/Inquiry",
			opts...,
		),
	}
}

// entityServiceClient implements EntityServiceClient.
type entityServiceClient struct {
	ingest  *connect_go.Client[v1.Member, v11.Response]
	inquiry *connect_go.Client[v1.Member, v11.Response]
}

// Ingest calls er.service.entity.v1.EntityService.Ingest.
func (c *entityServiceClient) Ingest(ctx context.Context, req *connect_go.Request[v1.Member]) (*connect_go.Response[v11.Response], error) {
	return c.ingest.CallUnary(ctx, req)
}

// Inquiry calls er.service.entity.v1.EntityService.Inquiry.
func (c *entityServiceClient) Inquiry(ctx context.Context, req *connect_go.Request[v1.Member]) (*connect_go.Response[v11.Response], error) {
	return c.inquiry.CallUnary(ctx, req)
}

// EntityServiceHandler is an implementation of the er.service.entity.v1.EntityService service.
type EntityServiceHandler interface {
	Ingest(context.Context, *connect_go.Request[v1.Member]) (*connect_go.Response[v11.Response], error)
	// rpc IngestStream(stream er.schema.entity.v1.Member) returns (stream Response) {}
	Inquiry(context.Context, *connect_go.Request[v1.Member]) (*connect_go.Response[v11.Response], error)
}

// NewEntityServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewEntityServiceHandler(svc EntityServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/er.service.entity.v1.EntityService/Ingest", connect_go.NewUnaryHandler(
		"/er.service.entity.v1.EntityService/Ingest",
		svc.Ingest,
		opts...,
	))
	mux.Handle("/er.service.entity.v1.EntityService/Inquiry", connect_go.NewUnaryHandler(
		"/er.service.entity.v1.EntityService/Inquiry",
		svc.Inquiry,
		opts...,
	))
	return "/er.service.entity.v1.EntityService/", mux
}

// UnimplementedEntityServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedEntityServiceHandler struct{}

func (UnimplementedEntityServiceHandler) Ingest(context.Context, *connect_go.Request[v1.Member]) (*connect_go.Response[v11.Response], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("er.service.entity.v1.EntityService.Ingest is not implemented"))
}

func (UnimplementedEntityServiceHandler) Inquiry(context.Context, *connect_go.Request[v1.Member]) (*connect_go.Response[v11.Response], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("er.service.entity.v1.EntityService.Inquiry is not implemented"))
}
