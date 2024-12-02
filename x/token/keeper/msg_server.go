package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tocipoco/MindCell/x/token/types"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) MintTokens(goCtx context.Context, msg *types.MsgMintTokens) (*types.MsgMintTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	amount, ok := sdk.NewIntFromString(msg.Amount)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount")
	}

	supply := m.GetTokenSupply(ctx)
	newSupply := supply.Add(amount)
	m.SetTokenSupply(ctx, newSupply)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"tokens_minted",
			sdk.NewAttribute("recipient", msg.Recipient),
			sdk.NewAttribute("amount", msg.Amount),
			sdk.NewAttribute("new_supply", newSupply.String()),
		),
	)

	return &types.MsgMintTokensResponse{Success: true}, nil
}

func (m msgServer) BurnTokens(goCtx context.Context, msg *types.MsgBurnTokens) (*types.MsgBurnTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	amount, ok := sdk.NewIntFromString(msg.Amount)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount")
	}

	supply := m.GetTokenSupply(ctx)
	if supply.LT(amount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "insufficient supply")
	}

	newSupply := supply.Sub(amount)
	m.SetTokenSupply(ctx, newSupply)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"tokens_burned",
			sdk.NewAttribute("burner", msg.Burner),
			sdk.NewAttribute("amount", msg.Amount),
			sdk.NewAttribute("new_supply", newSupply.String()),
		),
	)

	return &types.MsgBurnTokensResponse{Success: true}, nil
}

