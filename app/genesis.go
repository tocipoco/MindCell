package app

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
)

// GenesisState represents the genesis state of the blockchain
type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState generates the default starting genesis state for the application
func NewDefaultGenesisState(cdc codec.JSONCodec) GenesisState {
	return ModuleBasics.DefaultGenesis(cdc)
}

// ValidateGenesis validates the genesis state
func ValidateGenesis(genesisState GenesisState, cdc codec.JSONCodec) error {
	return ModuleBasics.ValidateGenesis(cdc, nil, genesisState)
}
