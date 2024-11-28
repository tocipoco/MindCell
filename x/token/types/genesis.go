package types

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TotalSupply:  "1000000000",
		TokenConfig:  DefaultTokenConfig(),
	}
}

func (gs GenesisState) Validate() error {
	return nil
}

type GenesisState struct {
	TotalSupply string      `json:"total_supply"`
	TokenConfig TokenConfig `json:"token_config"`
}

type TokenConfig struct {
	TokenDenom      string `json:"token_denom"`
	TokenDecimals   uint32 `json:"token_decimals"`
	MinStakeAmount  string `json:"min_stake_amount"`
	UnbondingPeriod int64  `json:"unbonding_period"` // in seconds
}

func DefaultTokenConfig() TokenConfig {
	return TokenConfig{
		TokenDenom:      "mcell",
		TokenDecimals:   18,
		MinStakeAmount:  "1000",
		UnbondingPeriod: 1814400, // 21 days
	}
}

