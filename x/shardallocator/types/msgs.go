package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgRegisterNode{}
	_ sdk.Msg = &MsgAssignShard{}
	_ sdk.Msg = &MsgReplaceShard{}
	_ sdk.Msg = &MsgUpdateNodeReputation{}
)

// MsgRegisterNode represents a message to register a new shard node
type MsgRegisterNode struct {
	Address     string `json:"address"`
	Endpoint    string `json:"endpoint"`
	StakeAmount string `json:"stake_amount"`
	MaxShards   uint32 `json:"max_shards"`
}

// MsgAssignShard represents a message to assign a shard to a node
type MsgAssignShard struct {
	Authority   string `json:"authority"`
	ModelID     uint64 `json:"model_id"`
	ShardID     uint32 `json:"shard_id"`
	NodeAddress string `json:"node_address"`
}

// MsgReplaceShard represents a message to replace a shard assignment
type MsgReplaceShard struct {
	Authority      string `json:"authority"`
	ModelID        uint64 `json:"model_id"`
	ShardID        uint32 `json:"shard_id"`
	OldNodeAddress string `json:"old_node_address"`
	NewNodeAddress string `json:"new_node_address"`
}

// MsgUpdateNodeReputation represents a message to update node reputation
type MsgUpdateNodeReputation struct {
	Authority   string `json:"authority"`
	NodeAddress string `json:"node_address"`
	Success     bool   `json:"success"`
}

// NewMsgRegisterNode creates a new MsgRegisterNode instance
func NewMsgRegisterNode(address, endpoint, stakeAmount string, maxShards uint32) *MsgRegisterNode {
	return &MsgRegisterNode{
		Address:     address,
		Endpoint:    endpoint,
		StakeAmount: stakeAmount,
		MaxShards:   maxShards,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgRegisterNode) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgRegisterNode) Type() string { return "register_node" }

// GetSigners implements the sdk.Msg interface.
func (msg MsgRegisterNode) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{address}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgRegisterNode) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgRegisterNode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}
	if msg.Endpoint == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "endpoint cannot be empty")
	}
	if msg.MaxShards == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "max shards must be greater than 0")
	}
	return nil
}

// NewMsgAssignShard creates a new MsgAssignShard instance
func NewMsgAssignShard(authority string, modelID uint64, shardID uint32, nodeAddress string) *MsgAssignShard {
	return &MsgAssignShard{
		Authority:   authority,
		ModelID:     modelID,
		ShardID:     shardID,
		NodeAddress: nodeAddress,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgAssignShard) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgAssignShard) Type() string { return "assign_shard" }

// GetSigners implements the sdk.Msg interface.
func (msg MsgAssignShard) GetSigners() []sdk.AccAddress {
	authority, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{authority}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgAssignShard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgAssignShard) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}
	if msg.ModelID == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "model ID must be greater than 0")
	}
	_, err = sdk.AccAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid node address (%s)", err)
	}
	return nil
}

var ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())

