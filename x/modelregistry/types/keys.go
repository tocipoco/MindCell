package types

const (
	// ModuleName defines the module name
	ModuleName = "modelregistry"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for modelregistry
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_modelregistry"
)

var (
	// ModelsKey prefix for models by ID
	ModelsKey = []byte{0x01}

	// ModelsCountKey tracks the total number of models
	ModelsCountKey = []byte{0x02}

	// ModelsByOwnerKey prefix for indexing models by owner
	ModelsByOwnerKey = []byte{0x03}
)

// GetModelKey returns the store key to retrieve a Model from the index fields
func GetModelKey(modelID uint64) []byte {
	bz := make([]byte, 8)
	for i := 0; i < 8; i++ {
		bz[i] = byte(modelID >> (8 * i))
	}
	return append(ModelsKey, bz...)
}

// GetModelsByOwnerKey returns the store key for models by owner address
func GetModelsByOwnerKey(owner string) []byte {
	return append(ModelsByOwnerKey, []byte(owner)...)
}

