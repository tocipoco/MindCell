package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgDistributeReward{}
	_ sdk.Msg = &MsgClaimReward{}
	_ sdk.Msg = &MsgAddToPool{}
)

type MsgDistributeReward struct {
	Authority   string `json:"authority"`
	NodeAddress string `json:"node_address"`
	Amount      string `json:"amount"`
	RequestID   uint64 `json:"request_id"`
	RewardType  string `json:"reward_type"`
}

type MsgClaimReward struct {
	Claimer string `json:"claimer"`
}

type MsgAddToPool struct {
	Authority string `json:"authority"`
	Amount    string `json:"amount"`
}

func NewMsgDistributeReward(authority, nodeAddress, amount string, requestID uint64, rewardType string) *MsgDistributeReward {
	return &MsgDistributeReward{
		Authority:   authority,
		NodeAddress: nodeAddress,
		Amount:      amount,
		RequestID:   requestID,
		RewardType:  rewardType,
	}
}

func (msg MsgDistributeReward) Route() string { return RouterKey }
func (msg MsgDistributeReward) Type() string  { return "distribute_reward" }

func (msg MsgDistributeReward) GetSigners() []sdk.AccAddress {
	authority, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{authority}
}

func (msg MsgDistributeReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgDistributeReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid node address (%s)", err)
	}
	return nil
}

func NewMsgClaimReward(claimer string) *MsgClaimReward {
	return &MsgClaimReward{Claimer: claimer}
}

func (msg MsgClaimReward) Route() string { return RouterKey }
func (msg MsgClaimReward) Type() string  { return "claim_reward" }

func (msg MsgClaimReward) GetSigners() []sdk.AccAddress {
	claimer, err := sdk.AccAddressFromBech32(msg.Claimer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{claimer}
}

func (msg MsgClaimReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgClaimReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Claimer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid claimer address (%s)", err)
	}
	return nil
}

