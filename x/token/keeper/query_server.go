package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tocipoco/MindCell/x/token/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) GetTokenSupply(ctx context.Context, req *types.QueryGetTokenSupplyRequest) (*types.QueryGetTokenSupplyResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	supply := k.GetTokenSupply(sdkCtx)

	return &types.QueryGetTokenSupplyResponse{TotalSupply: supply.String()}, nil
}

func (k Keeper) GetTokenConfig(ctx context.Context, req *types.QueryGetTokenConfigRequest) (*types.QueryGetTokenConfigResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	config := k.GetTokenConfig(sdkCtx)

	return &types.QueryGetTokenConfigResponse{Config: config}, nil
}

