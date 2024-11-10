package types

import "context"

type MsgServer interface {
	DistributeReward(context.Context, *MsgDistributeReward) (*MsgDistributeRewardResponse, error)
	ClaimReward(context.Context, *MsgClaimReward) (*MsgClaimRewardResponse, error)
	AddToPool(context.Context, *MsgAddToPool) (*MsgAddToPoolResponse, error)
}

type MsgDistributeRewardResponse struct {
	Success bool `json:"success"`
}

type MsgClaimRewardResponse struct {
	Amount string `json:"amount"`
}

type MsgAddToPoolResponse struct {
	Success bool `json:"success"`
}

type QueryServer interface {
	GetNodeReward(context.Context, *QueryGetNodeRewardRequest) (*QueryGetNodeRewardResponse, error)
	GetRewardPool(context.Context, *QueryGetRewardPoolRequest) (*QueryGetRewardPoolResponse, error)
}

type QueryGetNodeRewardRequest struct {
	NodeAddress string `json:"node_address"`
}

type QueryGetNodeRewardResponse struct {
	Reward NodeReward `json:"reward"`
}

type QueryGetRewardPoolRequest struct{}

type QueryGetRewardPoolResponse struct {
	TotalPool string `json:"total_pool"`
}

func RegisterMsgServer(server interface{}, impl MsgServer)           {}
func RegisterQueryServer(server interface{}, impl QueryServer)       {}
func RegisterQueryHandlerClient(ctx context.Context, mux, client interface{}) error { return nil }
func NewQueryClient(ctx interface{}) QueryClient                     { return nil }

type QueryClient interface{}
var _Msg_serviceDesc = struct{}{}

