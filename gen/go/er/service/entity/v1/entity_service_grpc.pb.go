// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: er/service/entity/v1/entity_service.proto

package entityv1

import (
	context "context"
	v1 "github.com/xmlking/entity-resolution/gen/go/er/schema/entity/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EntityServiceClient is the client API for EntityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntityServiceClient interface {
	Ingest(ctx context.Context, in *v1.Member, opts ...grpc.CallOption) (*Response, error)
	// rpc IngestStream(stream er.schema.entity.v1.Member) returns (stream Response) {}
	Inquiry(ctx context.Context, in *v1.Member, opts ...grpc.CallOption) (*Response, error)
}

type entityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEntityServiceClient(cc grpc.ClientConnInterface) EntityServiceClient {
	return &entityServiceClient{cc}
}

func (c *entityServiceClient) Ingest(ctx context.Context, in *v1.Member, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/er.service.entity.v1.EntityService/Ingest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityServiceClient) Inquiry(ctx context.Context, in *v1.Member, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/er.service.entity.v1.EntityService/Inquiry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntityServiceServer is the server API for EntityService service.
// All implementations should embed UnimplementedEntityServiceServer
// for forward compatibility
type EntityServiceServer interface {
	Ingest(context.Context, *v1.Member) (*Response, error)
	// rpc IngestStream(stream er.schema.entity.v1.Member) returns (stream Response) {}
	Inquiry(context.Context, *v1.Member) (*Response, error)
}

// UnimplementedEntityServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEntityServiceServer struct {
}

func (UnimplementedEntityServiceServer) Ingest(context.Context, *v1.Member) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ingest not implemented")
}
func (UnimplementedEntityServiceServer) Inquiry(context.Context, *v1.Member) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Inquiry not implemented")
}

// UnsafeEntityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntityServiceServer will
// result in compilation errors.
type UnsafeEntityServiceServer interface {
	mustEmbedUnimplementedEntityServiceServer()
}

func RegisterEntityServiceServer(s grpc.ServiceRegistrar, srv EntityServiceServer) {
	s.RegisterService(&EntityService_ServiceDesc, srv)
}

func _EntityService_Ingest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Member)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServiceServer).Ingest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/er.service.entity.v1.EntityService/Ingest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServiceServer).Ingest(ctx, req.(*v1.Member))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityService_Inquiry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Member)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServiceServer).Inquiry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/er.service.entity.v1.EntityService/Inquiry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServiceServer).Inquiry(ctx, req.(*v1.Member))
	}
	return interceptor(ctx, in, info, handler)
}

// EntityService_ServiceDesc is the grpc.ServiceDesc for EntityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EntityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "er.service.entity.v1.EntityService",
	HandlerType: (*EntityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ingest",
			Handler:    _EntityService_Ingest_Handler,
		},
		{
			MethodName: "Inquiry",
			Handler:    _EntityService_Inquiry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "er/service/entity/v1/entity_service.proto",
}
