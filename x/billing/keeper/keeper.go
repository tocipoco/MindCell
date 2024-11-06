package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tocipoco/MindCell/x/billing/types"
)

type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
}

func NewKeeper(cdc codec.BinaryCodec, storeKey storetypes.StoreKey) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

func (k Keeper) GetBillingRecord(ctx sdk.Context, requestID uint64) (types.BillingRecord, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetBillingRecordKey(requestID)

	bz := store.Get(key)
	if bz == nil {
		return types.BillingRecord{}, false
	}

	var record types.BillingRecord
	k.cdc.MustUnmarshal(bz, &record)
	return record, true
}

func (k Keeper) SetBillingRecord(ctx sdk.Context, record types.BillingRecord) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetBillingRecordKey(record.RequestID)
	bz := k.cdc.MustMarshal(&record)
	store.Set(key, bz)
}

func (k Keeper) GetFeeConfig(ctx sdk.Context) types.FeeConfig {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.FeeConfigKey)
	if bz == nil {
		return types.DefaultFeeConfig()
	}

	var config types.FeeConfig
	k.cdc.MustUnmarshal(bz, &config)
	return config
}

func (k Keeper) SetFeeConfig(ctx sdk.Context, config types.FeeConfig) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&config)
	store.Set(types.FeeConfigKey, bz)
}

func (k Keeper) CalculateFee(ctx sdk.Context, computeUnits uint64) (sdk.Coins, error) {
	config := k.GetFeeConfig(ctx)
	
	baseFee, ok := sdk.NewIntFromString(config.BaseFee)
	if !ok {
		return nil, sdk.ErrInvalidCoins
	}

	computePrice, ok := sdk.NewIntFromString(config.ComputeUnitPrice)
	if !ok {
		return nil, sdk.ErrInvalidCoins
	}

	computeCost := computePrice.Mul(sdk.NewIntFromUint64(computeUnits))
	totalFee := baseFee.Add(computeCost)

	return sdk.NewCoins(sdk.NewCoin("mcell", totalFee)), nil
}

func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	k.SetFeeConfig(ctx, genState.FeeConfig)
	for _, record := range genState.BillingRecords {
		k.SetBillingRecord(ctx, record)
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	config := k.GetFeeConfig(ctx)
	
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.BillingRecordKey)
	defer iterator.Close()

	var records []types.BillingRecord
	for ; iterator.Valid(); iterator.Next() {
		var record types.BillingRecord
		k.cdc.MustUnmarshal(iterator.Value(), &record)
		records = append(records, record)
	}

	return &types.GenesisState{
		BillingRecords: records,
		FeeConfig:      config,
	}
}

