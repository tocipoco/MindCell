package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tocipoco/MindCell/x/inferencegateway/types"
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

func (k Keeper) GetInferenceRequest(ctx sdk.Context, requestID uint64) (types.InferenceRequest, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetInferenceRequestKey(requestID)

	bz := store.Get(key)
	if bz == nil {
		return types.InferenceRequest{}, false
	}

	var request types.InferenceRequest
	k.cdc.MustUnmarshal(bz, &request)
	return request, true
}

func (k Keeper) SetInferenceRequest(ctx sdk.Context, request types.InferenceRequest) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetInferenceRequestKey(request.RequestID)
	bz := k.cdc.MustMarshal(&request)
	store.Set(key, bz)
}

func (k Keeper) GetRequestCount(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte("request_count"))
	if bz == nil {
		return 0
	}
	var count uint64
	k.cdc.MustUnmarshal(bz, &count)
	return count
}

func (k Keeper) SetRequestCount(ctx sdk.Context, count uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&count)
	store.Set([]byte("request_count"), bz)
}

func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	k.SetRequestCount(ctx, genState.RequestCount)
	for _, request := range genState.InferenceRequests {
		k.SetInferenceRequest(ctx, request)
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	count := k.GetRequestCount(ctx)
	var requests []types.InferenceRequest

	for i := uint64(1); i <= count; i++ {
		if request, found := k.GetInferenceRequest(ctx, i); found {
			requests = append(requests, request)
		}
	}

	return &types.GenesisState{
		InferenceRequests: requests,
		RequestCount:      count,
	}
}

