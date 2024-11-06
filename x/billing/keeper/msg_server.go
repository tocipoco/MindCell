package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tocipoco/MindCell/x/billing/types"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) ProcessPayment(goCtx context.Context, msg *types.MsgProcessPayment) (*types.MsgProcessPaymentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Calculate total fee based on compute units
	totalFee, err := m.CalculateFee(ctx, msg.ComputeUnits)
	if err != nil {
		return nil, err
	}

	config := m.GetFeeConfig(ctx)

	// Calculate shares
	totalAmount := totalFee[0].Amount
	modelOwnerShare := totalAmount.ToDec().Mul(sdk.NewDec(int64(config.ModelOwnerPercent * 100)).Quo(sdk.NewDec(100))).TruncateInt()
	nodeShare := totalAmount.ToDec().Mul(sdk.NewDec(int64(config.NodePercent * 100)).Quo(sdk.NewDec(100))).TruncateInt()
	protocolShare := totalAmount.Sub(modelOwnerShare).Sub(nodeShare)

	// Create billing record
	record := types.BillingRecord{
		RequestID:       msg.RequestID,
		Requester:       msg.Payer,
		TotalFee:        totalAmount.String(),
		ModelOwnerShare: modelOwnerShare.String(),
		NodeShare:       nodeShare.String(),
		ProtocolShare:   protocolShare.String(),
		Timestamp:       ctx.BlockTime().Unix(),
		Status:          "pending",
	}

	m.SetBillingRecord(ctx, record)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"payment_processed",
			sdk.NewAttribute("request_id", sdk.Uint64ToBigEndian(msg.RequestID).String()),
			sdk.NewAttribute("total_fee", totalAmount.String()),
		),
	)

	return &types.MsgProcessPaymentResponse{Success: true}, nil
}

func (m msgServer) SettleBilling(goCtx context.Context, msg *types.MsgSettleBilling) (*types.MsgSettleBillingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	record, found := m.GetBillingRecord(ctx, msg.RequestID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "billing record %d not found", msg.RequestID)
	}

	if record.Status != "pending" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "billing already settled")
	}

	record.Status = "settled"
	m.SetBillingRecord(ctx, record)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"billing_settled",
			sdk.NewAttribute("request_id", sdk.Uint64ToBigEndian(msg.RequestID).String()),
		),
	)

	return &types.MsgSettleBillingResponse{Success: true}, nil
}

func (m msgServer) UpdateFeeConfig(goCtx context.Context, msg *types.MsgUpdateFeeConfig) (*types.MsgUpdateFeeConfigResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	m.SetFeeConfig(ctx, msg.FeeConfig)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"fee_config_updated",
			sdk.NewAttribute("base_fee", msg.FeeConfig.BaseFee),
		),
	)

	return &types.MsgUpdateFeeConfigResponse{Success: true}, nil
}

