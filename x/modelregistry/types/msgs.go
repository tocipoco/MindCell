package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgRegisterModel{}
	_ sdk.Msg = &MsgUpdateModel{}
	_ sdk.Msg = &MsgDeactivateModel{}
)

// Message types for the modelregistry module
type MsgRegisterModel struct {
	Owner       string `json:"owner"`
	MetadataCID string `json:"metadata_cid"`
	ShardCount  uint32 `json:"shard_count"`
}

type MsgUpdateModel struct {
	Owner       string `json:"owner"`
	ModelID     uint64 `json:"model_id"`
	MetadataCID string `json:"metadata_cid"`
}

type MsgDeactivateModel struct {
	Owner   string `json:"owner"`
	ModelID uint64 `json:"model_id"`
}

// NewMsgRegisterModel creates a new MsgRegisterModel instance
func NewMsgRegisterModel(owner, metadataCID string, shardCount uint32) *MsgRegisterModel {
	return &MsgRegisterModel{
		Owner:       owner,
		MetadataCID: metadataCID,
		ShardCount:  shardCount,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgRegisterModel) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgRegisterModel) Type() string { return "register_model" }

// GetSigners implements the sdk.Msg interface.
func (msg MsgRegisterModel) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgRegisterModel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgRegisterModel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	if msg.MetadataCID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "metadata CID cannot be empty")
	}
	if msg.ShardCount == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "shard count must be greater than 0")
	}
	return nil
}

// NewMsgUpdateModel creates a new MsgUpdateModel instance
func NewMsgUpdateModel(owner string, modelID uint64, metadataCID string) *MsgUpdateModel {
	return &MsgUpdateModel{
		Owner:       owner,
		ModelID:     modelID,
		MetadataCID: metadataCID,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgUpdateModel) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgUpdateModel) Type() string { return "update_model" }

// GetSigners implements the sdk.Msg interface.
func (msg MsgUpdateModel) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgUpdateModel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgUpdateModel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	if msg.ModelID == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "model ID must be greater than 0")
	}
	if msg.MetadataCID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "metadata CID cannot be empty")
	}
	return nil
}

// NewMsgDeactivateModel creates a new MsgDeactivateModel instance
func NewMsgDeactivateModel(owner string, modelID uint64) *MsgDeactivateModel {
	return &MsgDeactivateModel{
		Owner:   owner,
		ModelID: modelID,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgDeactivateModel) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgDeactivateModel) Type() string { return "deactivate_model" }

// GetSigners implements the sdk.Msg interface.
func (msg MsgDeactivateModel) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgDeactivateModel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgDeactivateModel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	if msg.ModelID == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "model ID must be greater than 0")
	}
	return nil
}

var ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())

