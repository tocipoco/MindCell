package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrNoRewards           = sdkerrors.Register(ModuleName, 1, "no rewards available")
	ErrInvalidRewardAmount = sdkerrors.Register(ModuleName, 2, "invalid reward amount")
	ErrInsufficientPool    = sdkerrors.Register(ModuleName, 3, "insufficient reward pool")
)

