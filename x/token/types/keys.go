package types

const (
	ModuleName   = "token"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	TokenSupplyKey = []byte{0x01}
	TokenConfigKey = []byte{0x02}
)

