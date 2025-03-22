package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidTokenAmount  = sdkerrors.Register(ModuleName, 1, "invalid token amount")
	ErrInsufficientSupply  = sdkerrors.Register(ModuleName, 2, "insufficient token supply")
	ErrMintingDisabled     = sdkerrors.Register(ModuleName, 3, "minting is disabled")
)

