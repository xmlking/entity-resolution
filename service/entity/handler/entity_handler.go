package handler

import (
	"context"
	"github.com/bufbuild/connect-go"
	schema_entity_v1 "github.com/xmlking/entity-resolution/gen/go/er/schema/entity/v1"
	service_entity_v1 "github.com/xmlking/entity-resolution/gen/go/er/service/entity/v1"
	"github.com/xmlking/entity-resolution/gen/go/er/service/entity/v1/entityv1connect"
)

// to force implement interface
var _ entityv1connect.EntityServiceHandler = (*entityServiceHandler)(nil)

// entityServiceHandler struct
type entityServiceHandler struct {
}

// NewEntityServiceHandler returns an instance of `EntityServiceHandler`.
func NewEntityServiceHandler() (entityv1connect.EntityServiceHandler, error) {
	// func NewEntityServiceHandler(i *do.Injector) (entityv1connect.EntityServiceHandler, error) {
	//	client, err := do.Invoke[directory.Client](i)
	//	if err != nil {
	//		return nil, err
	//	}
	return &entityServiceHandler{}, nil
}

func (e entityServiceHandler) Ingest(ctx context.Context, req *connect.Request[schema_entity_v1.Member]) (resp *connect.Response[service_entity_v1.Response], err error) {
	if err := req.Msg.ValidateAll(); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	//TODO implement me
	key := "12345"
	return connect.NewResponse(&service_entity_v1.Response{
		Key:    &key,
		Status: schema_entity_v1.Status_STATUS_NEW,
	}), nil
}

func (e entityServiceHandler) Inquiry(ctx context.Context, req *connect.Request[schema_entity_v1.Member]) (resp *connect.Response[service_entity_v1.Response], err error) {
	if err := req.Msg.ValidateAll(); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	//TODO implement me
	key := "12345"
	return connect.NewResponse(&service_entity_v1.Response{
		Key:    &key,
		Status: schema_entity_v1.Status_STATUS_NEW,
	}), nil
}
