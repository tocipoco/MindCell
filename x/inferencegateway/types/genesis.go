package types

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		InferenceRequests: []InferenceRequest{},
		RequestCount:      0,
	}
}

// Validate performs basic genesis state validation
func (gs GenesisState) Validate() error {
	return nil
}

// GenesisState defines the inferencegateway module's genesis state
type GenesisState struct {
	InferenceRequests []InferenceRequest `json:"inference_requests"`
	RequestCount      uint64             `json:"request_count"`
}

// InferenceRequest represents an inference request
type InferenceRequest struct {
	RequestID   uint64 `json:"request_id"`
	ModelID     uint64 `json:"model_id"`
	Requester   string `json:"requester"`
	InputData   string `json:"input_data"`
	Status      string `json:"status"` // pending, processing, completed, failed
	Result      string `json:"result"`
	ProofHash   string `json:"proof_hash"`
	SubmittedAt int64  `json:"submitted_at"`
	CompletedAt int64  `json:"completed_at"`
	Fee         string `json:"fee"`
}

