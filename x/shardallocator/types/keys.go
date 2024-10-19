package types

const (
	// ModuleName defines the module name
	ModuleName = "shardallocator"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for shardallocator
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_shardallocator"
)

var (
	// ShardAssignmentKey prefix for shard assignments
	ShardAssignmentKey = []byte{0x01}

	// NodeReputationKey prefix for node reputation scores
	NodeReputationKey = []byte{0x02}

	// ShardsByNodeKey prefix for indexing shards by node
	ShardsByNodeKey = []byte{0x03}

	// NodeRegistrationKey prefix for registered nodes
	NodeRegistrationKey = []byte{0x04}
)

// GetShardAssignmentKey returns the store key for a shard assignment
func GetShardAssignmentKey(modelID uint64, shardID uint32) []byte {
	key := make([]byte, 12)
	for i := 0; i < 8; i++ {
		key[i] = byte(modelID >> (8 * i))
	}
	for i := 0; i < 4; i++ {
		key[8+i] = byte(shardID >> (8 * i))
	}
	return append(ShardAssignmentKey, key...)
}

// GetNodeReputationKey returns the store key for node reputation
func GetNodeReputationKey(nodeAddress string) []byte {
	return append(NodeReputationKey, []byte(nodeAddress)...)
}

// GetShardsByNodeKey returns the store key for shards assigned to a node
func GetShardsByNodeKey(nodeAddress string) []byte {
	return append(ShardsByNodeKey, []byte(nodeAddress)...)
}

// GetNodeRegistrationKey returns the store key for node registration
func GetNodeRegistrationKey(nodeAddress string) []byte {
	return append(NodeRegistrationKey, []byte(nodeAddress)...)
}

