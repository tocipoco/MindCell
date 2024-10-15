package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tocipoco/MindCell/x/modelregistry/types"
)

// Keeper of the modelregistry store
type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
}

// NewKeeper creates a new modelregistry Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

// GetModel returns a model by its ID
func (k Keeper) GetModel(ctx sdk.Context, modelID uint64) (types.Model, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetModelKey(modelID)

	bz := store.Get(key)
	if bz == nil {
		return types.Model{}, false
	}

	var model types.Model
	k.cdc.MustUnmarshal(bz, &model)
	return model, true
}

// SetModel stores a model
func (k Keeper) SetModel(ctx sdk.Context, model types.Model) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetModelKey(model.ID)
	bz := k.cdc.MustMarshal(&model)
	store.Set(key, bz)
}

// GetModelsCount returns the total number of models
func (k Keeper) GetModelsCount(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ModelsCountKey)
	if bz == nil {
		return 0
	}

	var count uint64
	k.cdc.MustUnmarshal(bz, &count)
	return count
}

// SetModelsCount sets the total number of models
func (k Keeper) SetModelsCount(ctx sdk.Context, count uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&count)
	store.Set(types.ModelsCountKey, bz)
}

// GetModelsByOwner returns all models owned by an address
func (k Keeper) GetModelsByOwner(ctx sdk.Context, owner string) []types.Model {
	store := ctx.KVStore(k.storeKey)
	prefix := types.GetModelsByOwnerKey(owner)

	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	var models []types.Model
	for ; iterator.Valid(); iterator.Next() {
		var modelID uint64
		k.cdc.MustUnmarshal(iterator.Value(), &modelID)

		if model, found := k.GetModel(ctx, modelID); found {
			models = append(models, model)
		}
	}

	return models
}

// SetModelOwnerIndex creates an index for models by owner
func (k Keeper) SetModelOwnerIndex(ctx sdk.Context, owner string, modelID uint64) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.GetModelsByOwnerKey(owner), sdk.Uint64ToBigEndian(modelID)...)
	bz := k.cdc.MustMarshal(&modelID)
	store.Set(key, bz)
}

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	k.SetModelsCount(ctx, genState.ModelsCount)
	for _, model := range genState.Models {
		k.SetModel(ctx, model)
		k.SetModelOwnerIndex(ctx, model.Owner, model.ID)
	}
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	count := k.GetModelsCount(ctx)
	var models []types.Model

	for i := uint64(1); i <= count; i++ {
		if model, found := k.GetModel(ctx, i); found {
			models = append(models, model)
		}
	}

	return &types.GenesisState{
		Models:      models,
		ModelsCount: count,
	}
}

