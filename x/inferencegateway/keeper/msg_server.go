package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tocipoco/MindCell/x/inferencegateway/types"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) SubmitInference(goCtx context.Context, msg *types.MsgSubmitInference) (*types.MsgSubmitInferenceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	count := m.GetRequestCount(ctx)
	newRequestID := count + 1

	request := types.InferenceRequest{
		RequestID:   newRequestID,
		ModelID:     msg.ModelID,
		Requester:   msg.Requester,
		InputData:   msg.InputData,
		Status:      "pending",
		SubmittedAt: ctx.BlockTime().Unix(),
		Fee:         msg.Fee,
	}

	m.SetInferenceRequest(ctx, request)
	m.SetRequestCount(ctx, newRequestID)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"inference_submitted",
			sdk.NewAttribute("request_id", sdk.Uint64ToBigEndian(newRequestID).String()),
			sdk.NewAttribute("model_id", sdk.Uint64ToBigEndian(msg.ModelID).String()),
			sdk.NewAttribute("requester", msg.Requester),
		),
	)

	return &types.MsgSubmitInferenceResponse{RequestID: newRequestID}, nil
}

func (m msgServer) VerifyProof(goCtx context.Context, msg *types.MsgVerifyProof) (*types.MsgVerifyProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	request, found := m.GetInferenceRequest(ctx, msg.RequestID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "inference request %d not found", msg.RequestID)
	}

	if request.Status != "processing" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "request is not in processing state")
	}

	// Simplified proof verification (in production, use actual zkML verification)
	valid := len(msg.ProofData) > 0

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"proof_verified",
			sdk.NewAttribute("request_id", sdk.Uint64ToBigEndian(msg.RequestID).String()),
			sdk.NewAttribute("valid", sdk.NewBool(valid).String()),
		),
	)

	return &types.MsgVerifyProofResponse{Valid: valid}, nil
}

func (m msgServer) CompleteInference(goCtx context.Context, msg *types.MsgCompleteInference) (*types.MsgCompleteInferenceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	request, found := m.GetInferenceRequest(ctx, msg.RequestID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "inference request %d not found", msg.RequestID)
	}

	request.Status = "completed"
	request.Result = msg.Result
	request.ProofHash = msg.ProofHash
	request.CompletedAt = ctx.BlockTime().Unix()

	m.SetInferenceRequest(ctx, request)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"inference_completed",
			sdk.NewAttribute("request_id", sdk.Uint64ToBigEndian(msg.RequestID).String()),
		),
	)

	return &types.MsgCompleteInferenceResponse{Success: true}, nil
}

