package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tocipoco/MindCell/x/reward/types"
)

type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
}

func NewKeeper(cdc codec.BinaryCodec, storeKey storetypes.StoreKey) Keeper {
	return Keeper{cdc: cdc, storeKey: storeKey}
}

func (k Keeper) GetRewardPool(ctx sdk.Context) sdk.Int {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.RewardPoolKey)
	if bz == nil {
		return sdk.ZeroInt()
	}
	var amount sdk.Int
	k.cdc.MustUnmarshal(bz, &amount)
	return amount
}

func (k Keeper) SetRewardPool(ctx sdk.Context, amount sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&amount)
	store.Set(types.RewardPoolKey, bz)
}

func (k Keeper) GetNodeReward(ctx sdk.Context, nodeAddress string) (types.NodeReward, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetNodeRewardKey(nodeAddress)
	bz := store.Get(key)
	if bz == nil {
		return types.NodeReward{}, false
	}
	var reward types.NodeReward
	k.cdc.MustUnmarshal(bz, &reward)
	return reward, true
}

func (k Keeper) SetNodeReward(ctx sdk.Context, reward types.NodeReward) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetNodeRewardKey(reward.NodeAddress)
	bz := k.cdc.MustMarshal(&reward)
	store.Set(key, bz)
}

func (k Keeper) DistributeReward(ctx sdk.Context, nodeAddress string, amount sdk.Int) error {
	reward, found := k.GetNodeReward(ctx, nodeAddress)
	if !found {
		reward = types.NodeReward{
			NodeAddress:       nodeAddress,
			AccumulatedReward: "0",
			ClaimedReward:     "0",
			PendingReward:     "0",
		}
	}

	accumulated, _ := sdk.NewIntFromString(reward.AccumulatedReward)
	pending, _ := sdk.NewIntFromString(reward.PendingReward)

	reward.AccumulatedReward = accumulated.Add(amount).String()
	reward.PendingReward = pending.Add(amount).String()

	k.SetNodeReward(ctx, reward)
	return nil
}

func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	poolAmount, _ := sdk.NewIntFromString(genState.RewardPool)
	k.SetRewardPool(ctx, poolAmount)

	for _, reward := range genState.NodeRewards {
		k.SetNodeReward(ctx, reward)
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	pool := k.GetRewardPool(ctx)
	
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.NodeRewardKey)
	defer iterator.Close()

	var rewards []types.NodeReward
	for ; iterator.Valid(); iterator.Next() {
		var reward types.NodeReward
		k.cdc.MustUnmarshal(iterator.Value(), &reward)
		rewards = append(rewards, reward)
	}

	return &types.GenesisState{
		RewardPool:      pool.String(),
		NodeRewards:     rewards,
		DistributionLog: []types.RewardDistribution{},
	}
}

