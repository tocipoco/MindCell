package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgSubmitInference{}
	_ sdk.Msg = &MsgVerifyProof{}
	_ sdk.Msg = &MsgCompleteInference{}
)

// MsgSubmitInference submits a new inference request
type MsgSubmitInference struct {
	Requester string `json:"requester"`
	ModelID   uint64 `json:"model_id"`
	InputData string `json:"input_data"`
	Fee       string `json:"fee"`
}

// MsgVerifyProof verifies a zkML proof
type MsgVerifyProof struct {
	RequestID uint64 `json:"request_id"`
	ProofData string `json:"proof_data"`
	Verifier  string `json:"verifier"`
}

// MsgCompleteInference completes an inference request
type MsgCompleteInference struct {
	RequestID uint64 `json:"request_id"`
	Result    string `json:"result"`
	ProofHash string `json:"proof_hash"`
	Executor  string `json:"executor"`
}

func NewMsgSubmitInference(requester string, modelID uint64, inputData, fee string) *MsgSubmitInference {
	return &MsgSubmitInference{
		Requester: requester,
		ModelID:   modelID,
		InputData: inputData,
		Fee:       fee,
	}
}

func (msg MsgSubmitInference) Route() string { return RouterKey }
func (msg MsgSubmitInference) Type() string  { return "submit_inference" }

func (msg MsgSubmitInference) GetSigners() []sdk.AccAddress {
	requester, err := sdk.AccAddressFromBech32(msg.Requester)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{requester}
}

func (msg MsgSubmitInference) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgSubmitInference) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Requester)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid requester address (%s)", err)
	}
	if msg.ModelID == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "model ID must be greater than 0")
	}
	if msg.InputData == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "input data cannot be empty")
	}
	return nil
}

