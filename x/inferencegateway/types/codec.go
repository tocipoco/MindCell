package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSubmitInference{}, "inferencegateway/SubmitInference", nil)
	cdc.RegisterConcrete(&MsgVerifyProof{}, "inferencegateway/VerifyProof", nil)
	cdc.RegisterConcrete(&MsgCompleteInference{}, "inferencegateway/CompleteInference", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitInference{},
		&MsgVerifyProof{},
		&MsgCompleteInference{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())

