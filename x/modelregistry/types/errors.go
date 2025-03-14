package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrModelNotFound      = sdkerrors.Register(ModuleName, 1, "model not found")
	ErrInvalidModelID     = sdkerrors.Register(ModuleName, 2, "invalid model ID")
	ErrUnauthorized       = sdkerrors.Register(ModuleName, 3, "unauthorized")
	ErrModelInactive      = sdkerrors.Register(ModuleName, 4, "model is inactive")
	ErrInvalidMetadata    = sdkerrors.Register(ModuleName, 5, "invalid metadata")
)

