package testutil

import (
	"github.com/cosmos/cosmos-sdk/db/memdb"
	"github.com/cosmos/cosmos-sdk/store/types"
)

// NewInMemoryDB creates a new in-memory database for testing
func NewInMemoryDB() types.CommitDB {
	return memdb.NewDB()
}

// DBStats returns statistics about a test database
type DBStats struct {
	KeyCount   int
	TotalSize  int64
	Iterations int
}

// GetDBStats returns statistics about the database contents
func GetDBStats(db types.KVStore) DBStats {
	stats := DBStats{}
	
	iterator := db.Iterator(nil, nil)
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		stats.KeyCount++
		stats.TotalSize += int64(len(iterator.Key()) + len(iterator.Value()))
		stats.Iterations++
	}
	
	return stats
}

// ClearDB removes all entries from a test database
func ClearDB(db types.KVStore) {
	iterator := db.Iterator(nil, nil)
	defer iterator.Close()
	
	keys := [][]byte{}
	for ; iterator.Valid(); iterator.Next() {
		keys = append(keys, iterator.Key())
	}
	
	for _, key := range keys {
		db.Delete(key)
	}
}

