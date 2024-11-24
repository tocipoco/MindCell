package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tocipoco/MindCell/x/slashing/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) GetSlashingRecord(ctx context.Context, req *types.QueryGetSlashingRecordRequest) (*types.QueryGetSlashingRecordResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	record, found := k.GetSlashingRecord(sdkCtx, req.RecordID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "record not found")
	}

	return &types.QueryGetSlashingRecordResponse{Record: record}, nil
}

func (k Keeper) GetSlashingParams(ctx context.Context, req *types.QueryGetSlashingParamsRequest) (*types.QueryGetSlashingParamsResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	params := k.GetSlashingParams(sdkCtx)

	return &types.QueryGetSlashingParamsResponse{Params: params}, nil
}

func (k Keeper) ListSlashingRecords(ctx context.Context, req *types.QueryListSlashingRecordsRequest) (*types.QueryListSlashingRecordsResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	count := k.GetRecordCount(sdkCtx)

	var records []types.SlashingRecord
	for i := uint64(1); i <= count; i++ {
		if record, found := k.GetSlashingRecord(sdkCtx, i); found {
			if req.NodeAddress != "" && record.NodeAddress != req.NodeAddress {
				continue
			}
			records = append(records, record)
		}
	}

	return &types.QueryListSlashingRecordsResponse{Records: records}, nil
}

