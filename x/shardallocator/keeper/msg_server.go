package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tocipoco/MindCell/x/shardallocator/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// RegisterNode handles node registration
func (m msgServer) RegisterNode(goCtx context.Context, msg *types.MsgRegisterNode) (*types.MsgRegisterNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if node already exists
	if _, found := m.GetNodeInfo(ctx, msg.Address); found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "node already registered")
	}

	// Create node info
	node := types.NodeInfo{
		Address:          msg.Address,
		Endpoint:         msg.Endpoint,
		StakeAmount:      msg.StakeAmount,
		MaxShards:        msg.MaxShards,
		CurrentShards:    0,
		Active:           true,
		RegistrationTime: ctx.BlockTime().Unix(),
	}

	// Store node info
	m.SetNodeInfo(ctx, node)

	// Initialize reputation
	reputation := types.NodeReputation{
		NodeAddress:      msg.Address,
		ReputationScore:  100.0,
		TotalInferences:  0,
		SuccessfulCount:  0,
		FailedCount:      0,
		LastActivityTime: ctx.BlockTime().Unix(),
	}
	m.SetNodeReputation(ctx, reputation)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"node_registered",
			sdk.NewAttribute("address", msg.Address),
			sdk.NewAttribute("max_shards", sdk.NewInt(int64(msg.MaxShards)).String()),
		),
	)

	return &types.MsgRegisterNodeResponse{Success: true}, nil
}

// AssignShard handles shard assignment to a node
func (m msgServer) AssignShard(goCtx context.Context, msg *types.MsgAssignShard) (*types.MsgAssignShardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if shard already assigned
	if _, found := m.GetShardAssignment(ctx, msg.ModelID, msg.ShardID); found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "shard already assigned")
	}

	// Get node info
	node, found := m.GetNodeInfo(ctx, msg.NodeAddress)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "node not found")
	}

	if !node.Active {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "node is not active")
	}

	if node.CurrentShards >= node.MaxShards {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "node has reached maximum shard capacity")
	}

	// Create shard assignment
	assignment := types.ShardAssignment{
		ModelID:     msg.ModelID,
		ShardID:     msg.ShardID,
		NodeAddress: msg.NodeAddress,
		AssignedAt:  ctx.BlockTime().Unix(),
		Status:      "active",
	}

	// Store assignment
	m.SetShardAssignment(ctx, assignment)

	// Update node's current shard count
	node.CurrentShards++
	m.SetNodeInfo(ctx, node)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"shard_assigned",
			sdk.NewAttribute("model_id", sdk.Uint64ToBigEndian(msg.ModelID).String()),
			sdk.NewAttribute("shard_id", sdk.NewInt(int64(msg.ShardID)).String()),
			sdk.NewAttribute("node_address", msg.NodeAddress),
		),
	)

	return &types.MsgAssignShardResponse{Success: true}, nil
}

// ReplaceShard handles replacing a shard assignment
func (m msgServer) ReplaceShard(goCtx context.Context, msg *types.MsgReplaceShard) (*types.MsgReplaceShardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get existing assignment
	assignment, found := m.GetShardAssignment(ctx, msg.ModelID, msg.ShardID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "shard assignment not found")
	}

	if assignment.NodeAddress != msg.OldNodeAddress {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "shard is not assigned to the specified old node")
	}

	// Get old node and decrease its shard count
	oldNode, found := m.GetNodeInfo(ctx, msg.OldNodeAddress)
	if found && oldNode.CurrentShards > 0 {
		oldNode.CurrentShards--
		m.SetNodeInfo(ctx, oldNode)
	}

	// Get new node and check capacity
	newNode, found := m.GetNodeInfo(ctx, msg.NewNodeAddress)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "new node not found")
	}

	if !newNode.Active {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "new node is not active")
	}

	if newNode.CurrentShards >= newNode.MaxShards {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "new node has reached maximum capacity")
	}

	// Update assignment
	assignment.NodeAddress = msg.NewNodeAddress
	assignment.AssignedAt = ctx.BlockTime().Unix()
	assignment.Status = "active"
	m.SetShardAssignment(ctx, assignment)

	// Increase new node's shard count
	newNode.CurrentShards++
	m.SetNodeInfo(ctx, newNode)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"shard_replaced",
			sdk.NewAttribute("model_id", sdk.Uint64ToBigEndian(msg.ModelID).String()),
			sdk.NewAttribute("shard_id", sdk.NewInt(int64(msg.ShardID)).String()),
			sdk.NewAttribute("old_node", msg.OldNodeAddress),
			sdk.NewAttribute("new_node", msg.NewNodeAddress),
		),
	)

	return &types.MsgReplaceShardResponse{Success: true}, nil
}

// UpdateNodeReputation handles updating node reputation
func (m msgServer) UpdateNodeReputation(goCtx context.Context, msg *types.MsgUpdateNodeReputation) (*types.MsgUpdateNodeReputationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Update reputation
	m.Keeper.UpdateNodeReputation(ctx, msg.NodeAddress, msg.Success)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"reputation_updated",
			sdk.NewAttribute("node_address", msg.NodeAddress),
			sdk.NewAttribute("success", sdk.NewBool(msg.Success).String()),
		),
	)

	return &types.MsgUpdateNodeReputationResponse{Success: true}, nil
}

