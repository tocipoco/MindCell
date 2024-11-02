package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgProcessPayment{}
	_ sdk.Msg = &MsgSettleBilling{}
	_ sdk.Msg = &MsgUpdateFeeConfig{}
)

type MsgProcessPayment struct {
	RequestID    uint64 `json:"request_id"`
	Payer        string `json:"payer"`
	Amount       string `json:"amount"`
	ComputeUnits uint64 `json:"compute_units"`
}

type MsgSettleBilling struct {
	RequestID uint64 `json:"request_id"`
	Authority string `json:"authority"`
}

type MsgUpdateFeeConfig struct {
	Authority string    `json:"authority"`
	FeeConfig FeeConfig `json:"fee_config"`
}

func NewMsgProcessPayment(requestID uint64, payer, amount string, computeUnits uint64) *MsgProcessPayment {
	return &MsgProcessPayment{
		RequestID:    requestID,
		Payer:        payer,
		Amount:       amount,
		ComputeUnits: computeUnits,
	}
}

func (msg MsgProcessPayment) Route() string { return RouterKey }
func (msg MsgProcessPayment) Type() string  { return "process_payment" }

func (msg MsgProcessPayment) GetSigners() []sdk.AccAddress {
	payer, err := sdk.AccAddressFromBech32(msg.Payer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{payer}
}

func (msg MsgProcessPayment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgProcessPayment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Payer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid payer address (%s)", err)
	}
	if msg.RequestID == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "request ID must be greater than 0")
	}
	return nil
}

