package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tocipoco/MindCell/x/modelregistry/types"
)

var _ types.QueryServer = Keeper{}

// GetModel returns a specific model by ID
func (k Keeper) GetModel(ctx context.Context, req *types.QueryGetModelRequest) (*types.QueryGetModelResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	model, found := k.GetModel(sdkCtx, req.ModelID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "model with ID %d not found", req.ModelID)
	}

	return &types.QueryGetModelResponse{Model: model}, nil
}

// ListModels returns all models or models filtered by owner
func (k Keeper) ListModels(ctx context.Context, req *types.QueryListModelsRequest) (*types.QueryListModelsResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	var models []types.Model

	if req.Owner != "" {
		// Filter by owner
		models = k.GetModelsByOwner(sdkCtx, req.Owner)
	} else {
		// Return all models
		count := k.GetModelsCount(sdkCtx)
		for i := uint64(1); i <= count; i++ {
			if model, found := k.GetModel(sdkCtx, i); found {
				models = append(models, model)
			}
		}
	}

	return &types.QueryListModelsResponse{Models: models}, nil
}

// ModelsCount returns the total number of registered models
func (k Keeper) ModelsCount(ctx context.Context, req *types.QueryModelsCountRequest) (*types.QueryModelsCountResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	count := k.GetModelsCount(sdkCtx)

	return &types.QueryModelsCountResponse{Count: count}, nil
}

