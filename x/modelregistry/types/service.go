package types

import (
	context "context"
)

// MsgServer is the server API for Msg service.
type MsgServer interface {
	RegisterModel(context.Context, *MsgRegisterModel) (*MsgRegisterModelResponse, error)
	UpdateModel(context.Context, *MsgUpdateModel) (*MsgUpdateModelResponse, error)
	DeactivateModel(context.Context, *MsgDeactivateModel) (*MsgDeactivateModelResponse, error)
}

// MsgRegisterModelResponse defines the Msg/RegisterModel response type.
type MsgRegisterModelResponse struct {
	ModelID uint64 `json:"model_id"`
}

// MsgUpdateModelResponse defines the Msg/UpdateModel response type.
type MsgUpdateModelResponse struct {
	Success bool `json:"success"`
}

// MsgDeactivateModelResponse defines the Msg/DeactivateModel response type.
type MsgDeactivateModelResponse struct {
	Success bool `json:"success"`
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	GetModel(context.Context, *QueryGetModelRequest) (*QueryGetModelResponse, error)
	ListModels(context.Context, *QueryListModelsRequest) (*QueryListModelsResponse, error)
	ModelsCount(context.Context, *QueryModelsCountRequest) (*QueryModelsCountResponse, error)
}

// RegisterMsgServer registers the msg server
func RegisterMsgServer(server interface{}, impl MsgServer) {}

// RegisterQueryServer registers the query server
func RegisterQueryServer(server interface{}, impl QueryServer) {}

// RegisterQueryHandlerClient registers the query handler client
func RegisterQueryHandlerClient(ctx context.Context, mux interface{}, client interface{}) error {
	return nil
}

// NewQueryClient creates a new query client
func NewQueryClient(ctx interface{}) QueryClient {
	return nil
}

// QueryClient is the client API for Query service.
type QueryClient interface {
	GetModel(ctx context.Context, in *QueryGetModelRequest) (*QueryGetModelResponse, error)
	ListModels(ctx context.Context, in *QueryListModelsRequest) (*QueryListModelsResponse, error)
	ModelsCount(ctx context.Context, in *QueryModelsCountRequest) (*QueryModelsCountResponse, error)
}

var _Msg_serviceDesc = struct{}{}

