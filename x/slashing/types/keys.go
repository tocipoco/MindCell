package types

const (
	ModuleName   = "slashing"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	SlashingRecordKey = []byte{0x01}
	NodePenaltyKey    = []byte{0x02}
	SlashingParamsKey = []byte{0x03}
)

func GetSlashingRecordKey(recordID uint64) []byte {
	bz := make([]byte, 8)
	for i := 0; i < 8; i++ {
		bz[i] = byte(recordID >> (8 * i))
	}
	return append(SlashingRecordKey, bz...)
}

func GetNodePenaltyKey(nodeAddress string) []byte {
	return append(NodePenaltyKey, []byte(nodeAddress)...)
}

