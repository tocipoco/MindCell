package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tocipoco/MindCell/x/slashing/types"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) SlashNode(goCtx context.Context, msg *types.MsgSlashNode) (*types.MsgSlashNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := m.Keeper.SlashNode(ctx, msg.NodeAddress, msg.SlashType, msg.Amount, msg.Reason, msg.RequestID)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"node_slashed",
			sdk.NewAttribute("node_address", msg.NodeAddress),
			sdk.NewAttribute("slash_type", msg.SlashType),
			sdk.NewAttribute("amount", msg.Amount),
		),
	)

	return &types.MsgSlashNodeResponse{Success: true}, nil
}

func (m msgServer) UpdateSlashingParams(goCtx context.Context, msg *types.MsgUpdateSlashingParams) (*types.MsgUpdateSlashingParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	m.SetSlashingParams(ctx, msg.Params)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"slashing_params_updated",
			sdk.NewAttribute("timeout_slash_percent", sdk.NewDec(int64(msg.Params.TimeoutSlashPercent*100)).String()),
		),
	)

	return &types.MsgUpdateSlashingParamsResponse{Success: true}, nil
}

