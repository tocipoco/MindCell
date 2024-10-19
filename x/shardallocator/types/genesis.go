package types

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ShardAssignments: []ShardAssignment{},
		NodeReputations:  []NodeReputation{},
		RegisteredNodes:  []NodeInfo{},
	}
}

// Validate performs basic genesis state validation returning an error upon any failure.
func (gs GenesisState) Validate() error {
	return nil
}

// GenesisState defines the shardallocator module's genesis state.
type GenesisState struct {
	ShardAssignments []ShardAssignment `json:"shard_assignments"`
	NodeReputations  []NodeReputation  `json:"node_reputations"`
	RegisteredNodes  []NodeInfo        `json:"registered_nodes"`
}

// ShardAssignment represents the assignment of a shard to a node
type ShardAssignment struct {
	ModelID     uint64 `json:"model_id"`
	ShardID     uint32 `json:"shard_id"`
	NodeAddress string `json:"node_address"`
	AssignedAt  int64  `json:"assigned_at"`
	Status      string `json:"status"` // active, inactive, slashed
}

// NodeReputation tracks a node's reputation score
type NodeReputation struct {
	NodeAddress      string  `json:"node_address"`
	ReputationScore  float64 `json:"reputation_score"`
	TotalInferences  uint64  `json:"total_inferences"`
	SuccessfulCount  uint64  `json:"successful_count"`
	FailedCount      uint64  `json:"failed_count"`
	LastActivityTime int64   `json:"last_activity_time"`
}

// NodeInfo contains information about a registered node
type NodeInfo struct {
	Address         string `json:"address"`
	Endpoint        string `json:"endpoint"`
	StakeAmount     string `json:"stake_amount"`
	MaxShards       uint32 `json:"max_shards"`
	CurrentShards   uint32 `json:"current_shards"`
	Active          bool   `json:"active"`
	RegistrationTime int64  `json:"registration_time"`
}

