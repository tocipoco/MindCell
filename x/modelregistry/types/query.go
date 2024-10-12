package types

// QueryGetModelRequest is the request type for the Query/GetModel RPC method.
type QueryGetModelRequest struct {
	ModelID uint64 `json:"model_id"`
}

// QueryGetModelResponse is the response type for the Query/GetModel RPC method.
type QueryGetModelResponse struct {
	Model Model `json:"model"`
}

// QueryListModelsRequest is the request type for the Query/ListModels RPC method.
type QueryListModelsRequest struct {
	Owner string `json:"owner,omitempty"`
}

// QueryListModelsResponse is the response type for the Query/ListModels RPC method.
type QueryListModelsResponse struct {
	Models []Model `json:"models"`
}

// QueryModelsCountRequest is the request type for the Query/ModelsCount RPC method.
type QueryModelsCountRequest struct{}

// QueryModelsCountResponse is the response type for the Query/ModelsCount RPC method.
type QueryModelsCountResponse struct {
	Count uint64 `json:"count"`
}

