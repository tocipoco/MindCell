package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tocipoco/MindCell/x/token/types"
)

type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
}

func NewKeeper(cdc codec.BinaryCodec, storeKey storetypes.StoreKey) Keeper {
	return Keeper{cdc: cdc, storeKey: storeKey}
}

func (k Keeper) GetTokenSupply(ctx sdk.Context) sdk.Int {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.TokenSupplyKey)
	if bz == nil {
		return sdk.ZeroInt()
	}
	var supply sdk.Int
	k.cdc.MustUnmarshal(bz, &supply)
	return supply
}

func (k Keeper) SetTokenSupply(ctx sdk.Context, supply sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&supply)
	store.Set(types.TokenSupplyKey, bz)
}

func (k Keeper) GetTokenConfig(ctx sdk.Context) types.TokenConfig {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.TokenConfigKey)
	if bz == nil {
		return types.DefaultTokenConfig()
	}
	var config types.TokenConfig
	k.cdc.MustUnmarshal(bz, &config)
	return config
}

func (k Keeper) SetTokenConfig(ctx sdk.Context, config types.TokenConfig) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&config)
	store.Set(types.TokenConfigKey, bz)
}

func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	supply, _ := sdk.NewIntFromString(genState.TotalSupply)
	k.SetTokenSupply(ctx, supply)
	k.SetTokenConfig(ctx, genState.TokenConfig)
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	supply := k.GetTokenSupply(ctx)
	config := k.GetTokenConfig(ctx)

	return &types.GenesisState{
		TotalSupply: supply.String(),
		TokenConfig: config,
	}
}

