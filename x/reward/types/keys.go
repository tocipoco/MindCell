package types

const (
	ModuleName   = "reward"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	RewardPoolKey        = []byte{0x01}
	NodeRewardKey        = []byte{0x02}
	ModelOwnerRewardKey  = []byte{0x03}
	RewardDistributionKey = []byte{0x04}
)

func GetNodeRewardKey(nodeAddress string) []byte {
	return append(NodeRewardKey, []byte(nodeAddress)...)
}

func GetModelOwnerRewardKey(ownerAddress string) []byte {
	return append(ModelOwnerRewardKey, []byte(ownerAddress)...)
}

