package testutil

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// TestKeeperContext creates a test context and keeper setup
type TestKeeperContext struct {
	Ctx      sdk.Context
	StoreKey storetypes.StoreKey
	Cdc      codec.BinaryCodec
}

// NewTestKeeperContext creates a new test keeper context
func NewTestKeeperContext(t *testing.T, moduleName string) TestKeeperContext {
	storeKey := sdk.NewKVStoreKey(moduleName)
	
	db := NewInMemoryDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	
	err := cms.LoadLatestVersion()
	require.NoError(t, err)
	
	ctx := sdk.NewContext(cms, sdk.NewBlockHeader(), false, nil)
	
	return TestKeeperContext{
		Ctx:      ctx,
		StoreKey: storeKey,
		Cdc:      codec.NewProtoCodec(nil),
	}
}

// SetBlockHeight sets the block height in the context
func (tkc *TestKeeperContext) SetBlockHeight(height int64) {
	header := tkc.Ctx.BlockHeader()
	header.Height = height
	tkc.Ctx = tkc.Ctx.WithBlockHeader(header)
}

// SetBlockTime sets the block time in the context
func (tkc *TestKeeperContext) SetBlockTime(timestamp int64) {
	header := tkc.Ctx.BlockHeader()
	header.Time = sdk.NewTime(timestamp)
	tkc.Ctx = tkc.Ctx.WithBlockHeader(header)
}

// CreateTestStore creates an isolated test store
func CreateTestStore(t *testing.T) sdk.KVStore {
	storeKey := sdk.NewKVStoreKey("test")
	db := NewInMemoryDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	
	err := cms.LoadLatestVersion()
	require.NoError(t, err)
	
	ctx := sdk.NewContext(cms, sdk.NewBlockHeader(), false, nil)
	return ctx.KVStore(storeKey)
}

