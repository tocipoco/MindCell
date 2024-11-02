package types

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		BillingRecords: []BillingRecord{},
		FeeConfig:      DefaultFeeConfig(),
	}
}

func (gs GenesisState) Validate() error {
	return nil
}

type GenesisState struct {
	BillingRecords []BillingRecord `json:"billing_records"`
	FeeConfig      FeeConfig       `json:"fee_config"`
}

type BillingRecord struct {
	RequestID       uint64 `json:"request_id"`
	Requester       string `json:"requester"`
	TotalFee        string `json:"total_fee"`
	ModelOwnerShare string `json:"model_owner_share"`
	NodeShare       string `json:"node_share"`
	ProtocolShare   string `json:"protocol_share"`
	Timestamp       int64  `json:"timestamp"`
	Status          string `json:"status"` // pending, settled, refunded
}

type FeeConfig struct {
	BaseFee           string  `json:"base_fee"`
	ComputeUnitPrice  string  `json:"compute_unit_price"`
	ModelOwnerPercent float64 `json:"model_owner_percent"`
	NodePercent       float64 `json:"node_percent"`
	ProtocolPercent   float64 `json:"protocol_percent"`
}

func DefaultFeeConfig() FeeConfig {
	return FeeConfig{
		BaseFee:           "1000",
		ComputeUnitPrice:  "10",
		ModelOwnerPercent: 0.60,
		NodePercent:       0.30,
		ProtocolPercent:   0.10,
	}
}

