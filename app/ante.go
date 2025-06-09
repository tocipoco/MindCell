package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
)

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func NewAnteHandler(
	ak ante.AccountKeeper,
	bankKeeper ante.BankKeeper,
	signModeHandler signing.SignModeHandler,
	feegrantKeeper ante.FeegrantKeeper,
) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(ak),
		ante.NewConsumeGasForTxSizeDecorator(ak),
		ante.NewDeductFeeDecorator(ak, bankKeeper, feegrantKeeper, nil),
		ante.NewSetPubKeyDecorator(ak),
		ante.NewValidateSigCountDecorator(ak),
		ante.NewSigGasConsumeDecorator(ak, DefaultSigVerificationGasConsumer),
		ante.NewSigVerificationDecorator(ak, signModeHandler),
		ante.NewIncrementSequenceDecorator(ak),
	)
}

// DefaultSigVerificationGasConsumer is the default gas consumer for signature verification
func DefaultSigVerificationGasConsumer(
	meter sdk.GasMeter, sig signing.SignatureV2, params types.Params,
) error {
	switch pubkey := sig.PubKey.(type) {
	case *ed25519.PubKey:
		meter.ConsumeGas(params.SigVerifyCostED25519, "ante verify: ed25519")
		return nil
	case *secp256k1.PubKey:
		meter.ConsumeGas(params.SigVerifyCostSecp256k1, "ante verify: secp256k1")
		return nil
	default:
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "unrecognized public key type: %T", pubkey)
	}
}
