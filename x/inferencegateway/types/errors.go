package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrRequestNotFound     = sdkerrors.Register(ModuleName, 1, "inference request not found")
	ErrInvalidProof        = sdkerrors.Register(ModuleName, 2, "invalid proof")
	ErrInvalidInput        = sdkerrors.Register(ModuleName, 3, "invalid input data")
	ErrRequestProcessing   = sdkerrors.Register(ModuleName, 4, "request already processing")
	ErrInsufficientFee     = sdkerrors.Register(ModuleName, 5, "insufficient fee")
)

