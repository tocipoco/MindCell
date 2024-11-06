package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tocipoco/MindCell/x/billing/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) GetBillingRecord(ctx context.Context, req *types.QueryGetBillingRecordRequest) (*types.QueryGetBillingRecordResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	record, found := k.GetBillingRecord(sdkCtx, req.RequestID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "billing record %d not found", req.RequestID)
	}

	return &types.QueryGetBillingRecordResponse{Record: record}, nil
}

func (k Keeper) GetFeeConfig(ctx context.Context, req *types.QueryGetFeeConfigRequest) (*types.QueryGetFeeConfigResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	config := k.GetFeeConfig(sdkCtx)

	return &types.QueryGetFeeConfigResponse{Config: config}, nil
}

