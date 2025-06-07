package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultParams returns default application parameters
func DefaultParams() Params {
	return Params{
		MaxBlockSize:     22020096, // 21MB
		EvidenceMaxBytes: 10000,
		MaxGas:           -1, // unlimited
	}
}

// Params defines application-level parameters
type Params struct {
	MaxBlockSize     int64
	EvidenceMaxBytes int64
	MaxGas           int64
}

// ConsensusParams returns consensus parameters for the application
func ConsensusParams() *tmproto.ConsensusParams {
	return &tmproto.ConsensusParams{
		Block: &tmproto.BlockParams{
			MaxBytes: 22020096,
			MaxGas:   -1,
		},
		Evidence: &tmproto.EvidenceParams{
			MaxAgeNumBlocks: 100000,
			MaxAgeDuration:  172800000000000, // 48 hours
			MaxBytes:        10000,
		},
		Validator: &tmproto.ValidatorParams{
			PubKeyTypes: []string{types.ABCIPubKeyTypeEd25519},
		},
	}
}
