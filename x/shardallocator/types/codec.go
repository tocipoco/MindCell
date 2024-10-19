package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterNode{}, "shardallocator/RegisterNode", nil)
	cdc.RegisterConcrete(&MsgAssignShard{}, "shardallocator/AssignShard", nil)
	cdc.RegisterConcrete(&MsgReplaceShard{}, "shardallocator/ReplaceShard", nil)
	cdc.RegisterConcrete(&MsgUpdateNodeReputation{}, "shardallocator/UpdateNodeReputation", nil)
}

// RegisterInterfaces registers the interfaces types with the interface registry.
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterNode{},
		&MsgAssignShard{},
		&MsgReplaceShard{},
		&MsgUpdateNodeReputation{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

