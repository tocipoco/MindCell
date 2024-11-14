package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tocipoco/MindCell/x/reward/types"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) DistributeReward(goCtx context.Context, msg *types.MsgDistributeReward) (*types.MsgDistributeRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	amount, ok := sdk.NewIntFromString(msg.Amount)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount")
	}

	err := m.Keeper.DistributeReward(ctx, msg.NodeAddress, amount)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"reward_distributed",
			sdk.NewAttribute("node_address", msg.NodeAddress),
			sdk.NewAttribute("amount", msg.Amount),
			sdk.NewAttribute("type", msg.RewardType),
		),
	)

	return &types.MsgDistributeRewardResponse{Success: true}, nil
}

func (m msgServer) ClaimReward(goCtx context.Context, msg *types.MsgClaimReward) (*types.MsgClaimRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	reward, found := m.GetNodeReward(ctx, msg.Claimer)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "no rewards found")
	}

	pendingAmount, ok := sdk.NewIntFromString(reward.PendingReward)
	if !ok || pendingAmount.IsZero() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no pending rewards")
	}

	// Update reward record
	claimedAmount, _ := sdk.NewIntFromString(reward.ClaimedReward)
	reward.ClaimedReward = claimedAmount.Add(pendingAmount).String()
	reward.PendingReward = "0"
	reward.LastClaimTime = ctx.BlockTime().Unix()

	m.SetNodeReward(ctx, reward)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"reward_claimed",
			sdk.NewAttribute("claimer", msg.Claimer),
			sdk.NewAttribute("amount", pendingAmount.String()),
		),
	)

	return &types.MsgClaimRewardResponse{Amount: pendingAmount.String()}, nil
}

func (m msgServer) AddToPool(goCtx context.Context, msg *types.MsgAddToPool) (*types.MsgAddToPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	amount, ok := sdk.NewIntFromString(msg.Amount)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount")
	}

	pool := m.GetRewardPool(ctx)
	newPool := pool.Add(amount)
	m.SetRewardPool(ctx, newPool)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"pool_funded",
			sdk.NewAttribute("amount", msg.Amount),
			sdk.NewAttribute("new_total", newPool.String()),
		),
	)

	return &types.MsgAddToPoolResponse{Success: true}, nil
}

