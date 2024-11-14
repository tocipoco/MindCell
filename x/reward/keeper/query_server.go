package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tocipoco/MindCell/x/reward/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) GetNodeReward(ctx context.Context, req *types.QueryGetNodeRewardRequest) (*types.QueryGetNodeRewardResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	reward, found := k.GetNodeReward(sdkCtx, req.NodeAddress)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "reward not found")
	}

	return &types.QueryGetNodeRewardResponse{Reward: reward}, nil
}

func (k Keeper) GetRewardPool(ctx context.Context, req *types.QueryGetRewardPoolRequest) (*types.QueryGetRewardPoolResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	pool := k.GetRewardPool(sdkCtx)

	return &types.QueryGetRewardPoolResponse{TotalPool: pool.String()}, nil
}

