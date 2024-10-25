package types

const (
	ModuleName   = "inferencegateway"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
	MemStoreKey  = "mem_inferencegateway"
)

var (
	InferenceRequestKey = []byte{0x01}
	ProofVerificationKey = []byte{0x02}
	InferenceHistoryKey = []byte{0x03}
)

// GetInferenceRequestKey returns the store key for an inference request
func GetInferenceRequestKey(requestID uint64) []byte {
	bz := make([]byte, 8)
	for i := 0; i < 8; i++ {
		bz[i] = byte(requestID >> (8 * i))
	}
	return append(InferenceRequestKey, bz...)
}

