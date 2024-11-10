package types

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		RewardPool:      "0",
		NodeRewards:     []NodeReward{},
		DistributionLog: []RewardDistribution{},
	}
}

func (gs GenesisState) Validate() error {
	return nil
}

type GenesisState struct {
	RewardPool      string               `json:"reward_pool"`
	NodeRewards     []NodeReward         `json:"node_rewards"`
	DistributionLog []RewardDistribution `json:"distribution_log"`
}

type NodeReward struct {
	NodeAddress       string `json:"node_address"`
	AccumulatedReward string `json:"accumulated_reward"`
	ClaimedReward     string `json:"claimed_reward"`
	PendingReward     string `json:"pending_reward"`
	LastClaimTime     int64  `json:"last_claim_time"`
}

type RewardDistribution struct {
	DistributionID uint64 `json:"distribution_id"`
	NodeAddress    string `json:"node_address"`
	Amount         string `json:"amount"`
	RequestID      uint64 `json:"request_id"`
	Timestamp      int64  `json:"timestamp"`
	Type           string `json:"type"` // inference, uptime, bonus
}

