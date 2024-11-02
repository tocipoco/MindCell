package types

import "context"

type MsgServer interface {
	ProcessPayment(context.Context, *MsgProcessPayment) (*MsgProcessPaymentResponse, error)
	SettleBilling(context.Context, *MsgSettleBilling) (*MsgSettleBillingResponse, error)
	UpdateFeeConfig(context.Context, *MsgUpdateFeeConfig) (*MsgUpdateFeeConfigResponse, error)
}

type MsgProcessPaymentResponse struct {
	Success bool `json:"success"`
}

type MsgSettleBillingResponse struct {
	Success bool `json:"success"`
}

type MsgUpdateFeeConfigResponse struct {
	Success bool `json:"success"`
}

type QueryServer interface {
	GetBillingRecord(context.Context, *QueryGetBillingRecordRequest) (*QueryGetBillingRecordResponse, error)
	GetFeeConfig(context.Context, *QueryGetFeeConfigRequest) (*QueryGetFeeConfigResponse, error)
}

type QueryGetBillingRecordRequest struct {
	RequestID uint64 `json:"request_id"`
}

type QueryGetBillingRecordResponse struct {
	Record BillingRecord `json:"record"`
}

type QueryGetFeeConfigRequest struct{}

type QueryGetFeeConfigResponse struct {
	Config FeeConfig `json:"config"`
}

func RegisterMsgServer(server interface{}, impl MsgServer)           {}
func RegisterQueryServer(server interface{}, impl QueryServer)       {}
func RegisterQueryHandlerClient(ctx context.Context, mux, client interface{}) error { return nil }
func NewQueryClient(ctx interface{}) QueryClient                     { return nil }

type QueryClient interface{}
var _Msg_serviceDesc = struct{}{}

