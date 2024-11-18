package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgSlashNode{}
	_ sdk.Msg = &MsgUpdateSlashingParams{}
)

type MsgSlashNode struct {
	Authority   string `json:"authority"`
	NodeAddress string `json:"node_address"`
	SlashType   string `json:"slash_type"`
	Amount      string `json:"amount"`
	Reason      string `json:"reason"`
	RequestID   uint64 `json:"request_id,omitempty"`
}

type MsgUpdateSlashingParams struct {
	Authority string         `json:"authority"`
	Params    SlashingParams `json:"params"`
}

func NewMsgSlashNode(authority, nodeAddress, slashType, amount, reason string, requestID uint64) *MsgSlashNode {
	return &MsgSlashNode{
		Authority:   authority,
		NodeAddress: nodeAddress,
		SlashType:   slashType,
		Amount:      amount,
		Reason:      reason,
		RequestID:   requestID,
	}
}

func (msg MsgSlashNode) Route() string { return RouterKey }
func (msg MsgSlashNode) Type() string  { return "slash_node" }

func (msg MsgSlashNode) GetSigners() []sdk.AccAddress {
	authority, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{authority}
}

func (msg MsgSlashNode) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgSlashNode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid node address (%s)", err)
	}
	if msg.SlashType == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "slash type cannot be empty")
	}
	return nil
}

