package testutil

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TestAddress generates a test address
func TestAddress(seed int) string {
	return sdk.AccAddress([]byte{byte(seed)}).String()
}

// TestModelID returns a test model ID
func TestModelID() uint64 {
	return 1
}

// TestShardID returns a test shard ID
func TestShardID() uint32 {
	return 1
}

// TestAmount returns a test amount string
func TestAmount() string {
	return "1000"
}

