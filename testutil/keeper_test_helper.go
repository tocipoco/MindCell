package testutil

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// CreateTestContext creates a test context for keeper tests
func CreateTestContext(t *testing.T) sdk.Context {
	ctx := sdk.NewContext(nil, sdk.NewBlockHeader(), false, nil)
	return ctx
}

// NewTestCodec creates a test codec
func NewTestCodec() codec.Codec {
	return codec.NewProtoCodec(nil)
}

// RequireNoError is a helper to check for no errors
func RequireNoError(t *testing.T, err error) {
	require.NoError(t, err)
}

