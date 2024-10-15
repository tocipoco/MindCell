package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tocipoco/MindCell/x/modelregistry/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// RegisterModel handles the registration of a new AI model
func (m msgServer) RegisterModel(goCtx context.Context, msg *types.MsgRegisterModel) (*types.MsgRegisterModelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the next model ID
	count := m.GetModelsCount(ctx)
	newModelID := count + 1

	// Create the new model
	model := types.Model{
		ID:          newModelID,
		Owner:       msg.Owner,
		MetadataCID: msg.MetadataCID,
		ShardCount:  msg.ShardCount,
		Version:     1,
		Active:      true,
		CreatedAt:   ctx.BlockTime().Unix(),
		UpdatedAt:   ctx.BlockTime().Unix(),
	}

	// Store the model
	m.SetModel(ctx, model)
	m.SetModelsCount(ctx, newModelID)
	m.SetModelOwnerIndex(ctx, msg.Owner, newModelID)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"model_registered",
			sdk.NewAttribute("model_id", sdk.Uint64ToBigEndian(newModelID).String()),
			sdk.NewAttribute("owner", msg.Owner),
			sdk.NewAttribute("shard_count", sdk.NewInt(int64(msg.ShardCount)).String()),
		),
	)

	return &types.MsgRegisterModelResponse{
		ModelID: newModelID,
	}, nil
}

// UpdateModel handles updating an existing model's metadata
func (m msgServer) UpdateModel(goCtx context.Context, msg *types.MsgUpdateModel) (*types.MsgUpdateModelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the existing model
	model, found := m.GetModel(ctx, msg.ModelID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "model with ID %d not found", msg.ModelID)
	}

	// Check ownership
	if model.Owner != msg.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "only the model owner can update the model")
	}

	// Check if model is active
	if !model.Active {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "cannot update inactive model")
	}

	// Update the model
	model.MetadataCID = msg.MetadataCID
	model.Version++
	model.UpdatedAt = ctx.BlockTime().Unix()

	// Store the updated model
	m.SetModel(ctx, model)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"model_updated",
			sdk.NewAttribute("model_id", sdk.Uint64ToBigEndian(msg.ModelID).String()),
			sdk.NewAttribute("version", sdk.NewInt(int64(model.Version)).String()),
		),
	)

	return &types.MsgUpdateModelResponse{
		Success: true,
	}, nil
}

// DeactivateModel handles deactivating an existing model
func (m msgServer) DeactivateModel(goCtx context.Context, msg *types.MsgDeactivateModel) (*types.MsgDeactivateModelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the existing model
	model, found := m.GetModel(ctx, msg.ModelID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "model with ID %d not found", msg.ModelID)
	}

	// Check ownership
	if model.Owner != msg.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "only the model owner can deactivate the model")
	}

	// Deactivate the model
	model.Active = false
	model.UpdatedAt = ctx.BlockTime().Unix()

	// Store the updated model
	m.SetModel(ctx, model)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"model_deactivated",
			sdk.NewAttribute("model_id", sdk.Uint64ToBigEndian(msg.ModelID).String()),
		),
	)

	return &types.MsgDeactivateModelResponse{
		Success: true,
	}, nil
}

