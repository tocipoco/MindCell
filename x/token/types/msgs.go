package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgMintTokens{}
	_ sdk.Msg = &MsgBurnTokens{}
)

type MsgMintTokens struct {
	Authority string `json:"authority"`
	Amount    string `json:"amount"`
	Recipient string `json:"recipient"`
}

type MsgBurnTokens struct {
	Burner string `json:"burner"`
	Amount string `json:"amount"`
}

func NewMsgMintTokens(authority, amount, recipient string) *MsgMintTokens {
	return &MsgMintTokens{
		Authority: authority,
		Amount:    amount,
		Recipient: recipient,
	}
}

func (msg MsgMintTokens) Route() string { return RouterKey }
func (msg MsgMintTokens) Type() string  { return "mint_tokens" }

func (msg MsgMintTokens) GetSigners() []sdk.AccAddress {
	authority, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{authority}
}

func (msg MsgMintTokens) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgMintTokens) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address (%s)", err)
	}
	return nil
}

