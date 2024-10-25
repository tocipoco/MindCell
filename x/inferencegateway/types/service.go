package types

import "context"

type MsgServer interface {
	SubmitInference(context.Context, *MsgSubmitInference) (*MsgSubmitInferenceResponse, error)
	VerifyProof(context.Context, *MsgVerifyProof) (*MsgVerifyProofResponse, error)
	CompleteInference(context.Context, *MsgCompleteInference) (*MsgCompleteInferenceResponse, error)
}

type MsgSubmitInferenceResponse struct {
	RequestID uint64 `json:"request_id"`
}

type MsgVerifyProofResponse struct {
	Valid bool `json:"valid"`
}

type MsgCompleteInferenceResponse struct {
	Success bool `json:"success"`
}

type QueryServer interface {
	GetInferenceRequest(context.Context, *QueryGetInferenceRequestRequest) (*QueryGetInferenceRequestResponse, error)
	ListInferenceRequests(context.Context, *QueryListInferenceRequestsRequest) (*QueryListInferenceRequestsResponse, error)
}

type QueryGetInferenceRequestRequest struct {
	RequestID uint64 `json:"request_id"`
}

type QueryGetInferenceRequestResponse struct {
	Request InferenceRequest `json:"request"`
}

type QueryListInferenceRequestsRequest struct {
	Requester string `json:"requester,omitempty"`
	Status    string `json:"status,omitempty"`
}

type QueryListInferenceRequestsResponse struct {
	Requests []InferenceRequest `json:"requests"`
}

func RegisterMsgServer(server interface{}, impl MsgServer)           {}
func RegisterQueryServer(server interface{}, impl QueryServer)       {}
func RegisterQueryHandlerClient(ctx context.Context, mux, client interface{}) error { return nil }
func NewQueryClient(ctx interface{}) QueryClient                     { return nil }

type QueryClient interface{}
var _Msg_serviceDesc = struct{}{}

