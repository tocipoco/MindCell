package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tocipoco/MindCell/x/shardallocator/types"
)

var _ types.QueryServer = Keeper{}

// GetShardAssignment returns shard assignment information
func (k Keeper) GetShardAssignment(ctx context.Context, req *types.QueryGetShardAssignmentRequest) (*types.QueryGetShardAssignmentResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	assignment, found := k.GetShardAssignment(sdkCtx, req.ModelID, req.ShardID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "shard assignment not found")
	}

	return &types.QueryGetShardAssignmentResponse{Assignment: assignment}, nil
}

// GetNodeInfo returns node information and reputation
func (k Keeper) GetNodeInfo(ctx context.Context, req *types.QueryGetNodeInfoRequest) (*types.QueryGetNodeInfoResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	node, found := k.GetNodeInfo(sdkCtx, req.NodeAddress)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "node not found")
	}

	reputation, _ := k.GetNodeReputation(sdkCtx, req.NodeAddress)

	return &types.QueryGetNodeInfoResponse{
		Node:       node,
		Reputation: reputation,
	}, nil
}

// ListNodes returns all registered nodes
func (k Keeper) ListNodes(ctx context.Context, req *types.QueryListNodesRequest) (*types.QueryListNodesResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.NodeRegistrationKey)
	defer iterator.Close()

	var nodes []types.NodeInfo
	for ; iterator.Valid(); iterator.Next() {
		var node types.NodeInfo
		k.cdc.MustUnmarshal(iterator.Value(), &node)

		if req.ActiveOnly && !node.Active {
			continue
		}

		nodes = append(nodes, node)
	}

	return &types.QueryListNodesResponse{Nodes: nodes}, nil
}

