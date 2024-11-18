package types

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		SlashingRecords: []SlashingRecord{},
		SlashingParams:  DefaultSlashingParams(),
	}
}

func (gs GenesisState) Validate() error {
	return nil
}

type GenesisState struct {
	SlashingRecords []SlashingRecord `json:"slashing_records"`
	SlashingParams  SlashingParams   `json:"slashing_params"`
}

type SlashingRecord struct {
	RecordID     uint64 `json:"record_id"`
	NodeAddress  string `json:"node_address"`
	SlashType    string `json:"slash_type"` // timeout, incorrect_proof, downtime
	Amount       string `json:"amount"`
	Timestamp    int64  `json:"timestamp"`
	Reason       string `json:"reason"`
	RequestID    uint64 `json:"request_id,omitempty"`
}

type SlashingParams struct {
	TimeoutSlashPercent        float64 `json:"timeout_slash_percent"`
	IncorrectProofSlashPercent float64 `json:"incorrect_proof_slash_percent"`
	DowntimeSlashPercent       float64 `json:"downtime_slash_percent"`
	MinSlashAmount             string  `json:"min_slash_amount"`
	MaxSlashAmount             string  `json:"max_slash_amount"`
}

func DefaultSlashingParams() SlashingParams {
	return SlashingParams{
		TimeoutSlashPercent:        0.05,  // 5%
		IncorrectProofSlashPercent: 0.20,  // 20%
		DowntimeSlashPercent:       0.10,  // 10%
		MinSlashAmount:             "100",
		MaxSlashAmount:             "10000",
	}
}

