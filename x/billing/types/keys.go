package types

const (
	ModuleName   = "billing"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
	MemStoreKey  = "mem_billing"
)

var (
	BillingRecordKey   = []byte{0x01}
	FeeConfigKey       = []byte{0x02}
	AccountBalanceKey  = []byte{0x03}
)

func GetBillingRecordKey(requestID uint64) []byte {
	bz := make([]byte, 8)
	for i := 0; i < 8; i++ {
		bz[i] = byte(requestID >> (8 * i))
	}
	return append(BillingRecordKey, bz...)
}

func GetAccountBalanceKey(address string) []byte {
	return append(AccountBalanceKey, []byte(address)...)
}

