package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterModel{}, "modelregistry/RegisterModel", nil)
	cdc.RegisterConcrete(&MsgUpdateModel{}, "modelregistry/UpdateModel", nil)
	cdc.RegisterConcrete(&MsgDeactivateModel{}, "modelregistry/DeactivateModel", nil)
}

// RegisterInterfaces registers the interfaces types with the interface registry.
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterModel{},
		&MsgUpdateModel{},
		&MsgDeactivateModel{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

