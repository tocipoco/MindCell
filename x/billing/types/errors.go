package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrBillingRecordNotFound = sdkerrors.Register(ModuleName, 1, "billing record not found")
	ErrInvalidAmount         = sdkerrors.Register(ModuleName, 2, "invalid amount")
	ErrAlreadySettled        = sdkerrors.Register(ModuleName, 3, "billing already settled")
	ErrInvalidFeeConfig      = sdkerrors.Register(ModuleName, 4, "invalid fee configuration")
)

