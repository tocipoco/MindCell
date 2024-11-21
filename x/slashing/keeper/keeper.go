package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tocipoco/MindCell/x/slashing/types"
)

type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
}

func NewKeeper(cdc codec.BinaryCodec, storeKey storetypes.StoreKey) Keeper {
	return Keeper{cdc: cdc, storeKey: storeKey}
}

func (k Keeper) GetSlashingRecord(ctx sdk.Context, recordID uint64) (types.SlashingRecord, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetSlashingRecordKey(recordID)
	bz := store.Get(key)
	if bz == nil {
		return types.SlashingRecord{}, false
	}
	var record types.SlashingRecord
	k.cdc.MustUnmarshal(bz, &record)
	return record, true
}

func (k Keeper) SetSlashingRecord(ctx sdk.Context, record types.SlashingRecord) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetSlashingRecordKey(record.RecordID)
	bz := k.cdc.MustMarshal(&record)
	store.Set(key, bz)
}

func (k Keeper) GetSlashingParams(ctx sdk.Context) types.SlashingParams {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.SlashingParamsKey)
	if bz == nil {
		return types.DefaultSlashingParams()
	}
	var params types.SlashingParams
	k.cdc.MustUnmarshal(bz, &params)
	return params
}

func (k Keeper) SetSlashingParams(ctx sdk.Context, params types.SlashingParams) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&params)
	store.Set(types.SlashingParamsKey, bz)
}

func (k Keeper) GetRecordCount(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte("record_count"))
	if bz == nil {
		return 0
	}
	var count uint64
	k.cdc.MustUnmarshal(bz, &count)
	return count
}

func (k Keeper) SetRecordCount(ctx sdk.Context, count uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&count)
	store.Set([]byte("record_count"), bz)
}

func (k Keeper) SlashNode(ctx sdk.Context, nodeAddress, slashType, amount, reason string, requestID uint64) error {
	count := k.GetRecordCount(ctx)
	newRecordID := count + 1

	record := types.SlashingRecord{
		RecordID:    newRecordID,
		NodeAddress: nodeAddress,
		SlashType:   slashType,
		Amount:      amount,
		Timestamp:   ctx.BlockTime().Unix(),
		Reason:      reason,
		RequestID:   requestID,
	}

	k.SetSlashingRecord(ctx, record)
	k.SetRecordCount(ctx, newRecordID)

	return nil
}

func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	k.SetSlashingParams(ctx, genState.SlashingParams)
	for _, record := range genState.SlashingRecords {
		k.SetSlashingRecord(ctx, record)
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	params := k.GetSlashingParams(ctx)
	count := k.GetRecordCount(ctx)

	var records []types.SlashingRecord
	for i := uint64(1); i <= count; i++ {
		if record, found := k.GetSlashingRecord(ctx, i); found {
			records = append(records, record)
		}
	}

	return &types.GenesisState{
		SlashingRecords: records,
		SlashingParams:  params,
	}
}

