package types

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Models:      []Model{},
		ModelsCount: 0,
	}
}

// Validate performs basic genesis state validation returning an error upon any failure.
func (gs GenesisState) Validate() error {
	return nil
}

// GenesisState defines the modelregistry module's genesis state.
type GenesisState struct {
	Models      []Model `json:"models"`
	ModelsCount uint64  `json:"models_count"`
}

// Model represents a registered AI model
type Model struct {
	ID          uint64 `json:"id"`
	Owner       string `json:"owner"`
	MetadataCID string `json:"metadata_cid"`
	ShardCount  uint32 `json:"shard_count"`
	Version     uint32 `json:"version"`
	Active      bool   `json:"active"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

