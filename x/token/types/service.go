package types

import "context"

type MsgServer interface {
	MintTokens(context.Context, *MsgMintTokens) (*MsgMintTokensResponse, error)
	BurnTokens(context.Context, *MsgBurnTokens) (*MsgBurnTokensResponse, error)
}

type MsgMintTokensResponse struct {
	Success bool `json:"success"`
}

type MsgBurnTokensResponse struct {
	Success bool `json:"success"`
}

type QueryServer interface {
	GetTokenSupply(context.Context, *QueryGetTokenSupplyRequest) (*QueryGetTokenSupplyResponse, error)
	GetTokenConfig(context.Context, *QueryGetTokenConfigRequest) (*QueryGetTokenConfigResponse, error)
}

type QueryGetTokenSupplyRequest struct{}

type QueryGetTokenSupplyResponse struct {
	TotalSupply string `json:"total_supply"`
}

type QueryGetTokenConfigRequest struct{}

type QueryGetTokenConfigResponse struct {
	Config TokenConfig `json:"config"`
}

func RegisterMsgServer(server interface{}, impl MsgServer)           {}
func RegisterQueryServer(server interface{}, impl QueryServer)       {}
func RegisterQueryHandlerClient(ctx context.Context, mux, client interface{}) error { return nil }
func NewQueryClient(ctx interface{}) QueryClient                     { return nil }

type QueryClient interface{}
var _Msg_serviceDesc = struct{}{}

