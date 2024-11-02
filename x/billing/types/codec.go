package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgProcessPayment{}, "billing/ProcessPayment", nil)
	cdc.RegisterConcrete(&MsgSettleBilling{}, "billing/SettleBilling", nil)
	cdc.RegisterConcrete(&MsgUpdateFeeConfig{}, "billing/UpdateFeeConfig", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProcessPayment{},
		&MsgSettleBilling{},
		&MsgUpdateFeeConfig{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())

