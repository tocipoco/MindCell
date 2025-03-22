package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidSlashType    = sdkerrors.Register(ModuleName, 1, "invalid slash type")
	ErrInsufficientStake   = sdkerrors.Register(ModuleName, 2, "insufficient stake")
	ErrSlashingDisabled    = sdkerrors.Register(ModuleName, 3, "slashing is disabled")
)

