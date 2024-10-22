package keeper

import (
	"math"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tocipoco/MindCell/x/shardallocator/types"
)

// Keeper of the shardallocator store
type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
}

// NewKeeper creates a new shardallocator Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

// GetShardAssignment returns a shard assignment
func (k Keeper) GetShardAssignment(ctx sdk.Context, modelID uint64, shardID uint32) (types.ShardAssignment, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetShardAssignmentKey(modelID, shardID)

	bz := store.Get(key)
	if bz == nil {
		return types.ShardAssignment{}, false
	}

	var assignment types.ShardAssignment
	k.cdc.MustUnmarshal(bz, &assignment)
	return assignment, true
}

// SetShardAssignment stores a shard assignment
func (k Keeper) SetShardAssignment(ctx sdk.Context, assignment types.ShardAssignment) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetShardAssignmentKey(assignment.ModelID, assignment.ShardID)
	bz := k.cdc.MustMarshal(&assignment)
	store.Set(key, bz)
}

// GetNodeReputation returns a node's reputation
func (k Keeper) GetNodeReputation(ctx sdk.Context, nodeAddress string) (types.NodeReputation, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetNodeReputationKey(nodeAddress)

	bz := store.Get(key)
	if bz == nil {
		return types.NodeReputation{}, false
	}

	var reputation types.NodeReputation
	k.cdc.MustUnmarshal(bz, &reputation)
	return reputation, true
}

// SetNodeReputation stores a node's reputation
func (k Keeper) SetNodeReputation(ctx sdk.Context, reputation types.NodeReputation) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetNodeReputationKey(reputation.NodeAddress)
	bz := k.cdc.MustMarshal(&reputation)
	store.Set(key, bz)
}

// GetNodeInfo returns node information
func (k Keeper) GetNodeInfo(ctx sdk.Context, nodeAddress string) (types.NodeInfo, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetNodeRegistrationKey(nodeAddress)

	bz := store.Get(key)
	if bz == nil {
		return types.NodeInfo{}, false
	}

	var node types.NodeInfo
	k.cdc.MustUnmarshal(bz, &node)
	return node, true
}

// SetNodeInfo stores node information
func (k Keeper) SetNodeInfo(ctx sdk.Context, node types.NodeInfo) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetNodeRegistrationKey(node.Address)
	bz := k.cdc.MustMarshal(&node)
	store.Set(key, bz)
}

// SelectBestNode selects the best node for shard assignment based on reputation and load
func (k Keeper) SelectBestNode(ctx sdk.Context) (string, error) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.NodeRegistrationKey)
	defer iterator.Close()

	var bestNode string
	bestScore := float64(-1)

	for ; iterator.Valid(); iterator.Next() {
		var node types.NodeInfo
		k.cdc.MustUnmarshal(iterator.Value(), &node)

		if !node.Active || node.CurrentShards >= node.MaxShards {
			continue
		}

		reputation, found := k.GetNodeReputation(ctx, node.Address)
		if !found {
			reputation = types.NodeReputation{
				NodeAddress:     node.Address,
				ReputationScore: 100.0,
			}
		}

		// Calculate selection score based on reputation and available capacity
		loadFactor := 1.0 - float64(node.CurrentShards)/float64(node.MaxShards)
		score := reputation.ReputationScore * loadFactor

		if score > bestScore {
			bestScore = score
			bestNode = node.Address
		}
	}

	if bestNode == "" {
		return "", sdk.ErrInsufficientFunds
	}

	return bestNode, nil
}

// UpdateNodeReputation updates a node's reputation based on performance
func (k Keeper) UpdateNodeReputation(ctx sdk.Context, nodeAddress string, success bool) {
	reputation, found := k.GetNodeReputation(ctx, nodeAddress)
	if !found {
		reputation = types.NodeReputation{
			NodeAddress:     nodeAddress,
			ReputationScore: 100.0,
		}
	}

	reputation.TotalInferences++
	if success {
		reputation.SuccessfulCount++
		// Increase reputation slightly on success (cap at 100)
		reputation.ReputationScore = math.Min(100.0, reputation.ReputationScore+0.1)
	} else {
		reputation.FailedCount++
		// Decrease reputation on failure
		reputation.ReputationScore = math.Max(0.0, reputation.ReputationScore-5.0)
	}

	reputation.LastActivityTime = ctx.BlockTime().Unix()
	k.SetNodeReputation(ctx, reputation)
}

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	for _, assignment := range genState.ShardAssignments {
		k.SetShardAssignment(ctx, assignment)
	}
	for _, reputation := range genState.NodeReputations {
		k.SetNodeReputation(ctx, reputation)
	}
	for _, node := range genState.RegisteredNodes {
		k.SetNodeInfo(ctx, node)
	}
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	var assignments []types.ShardAssignment
	var reputations []types.NodeReputation
	var nodes []types.NodeInfo

	store := ctx.KVStore(k.storeKey)

	// Export shard assignments
	iterator := sdk.KVStorePrefixIterator(store, types.ShardAssignmentKey)
	for ; iterator.Valid(); iterator.Next() {
		var assignment types.ShardAssignment
		k.cdc.MustUnmarshal(iterator.Value(), &assignment)
		assignments = append(assignments, assignment)
	}
	iterator.Close()

	// Export node reputations
	iterator = sdk.KVStorePrefixIterator(store, types.NodeReputationKey)
	for ; iterator.Valid(); iterator.Next() {
		var reputation types.NodeReputation
		k.cdc.MustUnmarshal(iterator.Value(), &reputation)
		reputations = append(reputations, reputation)
	}
	iterator.Close()

	// Export registered nodes
	iterator = sdk.KVStorePrefixIterator(store, types.NodeRegistrationKey)
	for ; iterator.Valid(); iterator.Next() {
		var node types.NodeInfo
		k.cdc.MustUnmarshal(iterator.Value(), &node)
		nodes = append(nodes, node)
	}
	iterator.Close()

	return &types.GenesisState{
		ShardAssignments: assignments,
		NodeReputations:  reputations,
		RegisteredNodes:  nodes,
	}
}

