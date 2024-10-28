package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tocipoco/MindCell/x/inferencegateway/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) GetInferenceRequest(ctx context.Context, req *types.QueryGetInferenceRequestRequest) (*types.QueryGetInferenceRequestResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	request, found := k.GetInferenceRequest(sdkCtx, req.RequestID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "request %d not found", req.RequestID)
	}

	return &types.QueryGetInferenceRequestResponse{Request: request}, nil
}

func (k Keeper) ListInferenceRequests(ctx context.Context, req *types.QueryListInferenceRequestsRequest) (*types.QueryListInferenceRequestsResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	count := k.GetRequestCount(sdkCtx)

	var requests []types.InferenceRequest
	for i := uint64(1); i <= count; i++ {
		if request, found := k.GetInferenceRequest(sdkCtx, i); found {
			// Apply filters
			if req.Requester != "" && request.Requester != req.Requester {
				continue
			}
			if req.Status != "" && request.Status != req.Status {
				continue
			}
			requests = append(requests, request)
		}
	}

	return &types.QueryListInferenceRequestsResponse{Requests: requests}, nil
}

