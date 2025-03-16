package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrNodeNotFound        = sdkerrors.Register(ModuleName, 1, "node not found")
	ErrShardNotFound       = sdkerrors.Register(ModuleName, 2, "shard not found")
	ErrInsufficientCapacity = sdkerrors.Register(ModuleName, 3, "insufficient node capacity")
	ErrNodeInactive        = sdkerrors.Register(ModuleName, 4, "node is inactive")
	ErrDuplicateNode       = sdkerrors.Register(ModuleName, 5, "node already registered")
)

