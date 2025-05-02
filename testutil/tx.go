package testutil

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)

// TxBuilder is a helper for building test transactions
type TxBuilder struct {
	txConfig client.TxConfig
	msgs     []sdk.Msg
	memo     string
	fees     sdk.Coins
	gasLimit uint64
}

// NewTxBuilder creates a new transaction builder
func NewTxBuilder(txConfig client.TxConfig) *TxBuilder {
	return &TxBuilder{
		txConfig: txConfig,
		msgs:     []sdk.Msg{},
		gasLimit: 200000,
	}
}

// WithMsgs adds messages to the transaction
func (b *TxBuilder) WithMsgs(msgs ...sdk.Msg) *TxBuilder {
	b.msgs = append(b.msgs, msgs...)
	return b
}

// WithMemo sets the transaction memo
func (b *TxBuilder) WithMemo(memo string) *TxBuilder {
	b.memo = memo
	return b
}

// WithFees sets the transaction fees
func (b *TxBuilder) WithFees(fees string) *TxBuilder {
	parsedFees, err := sdk.ParseCoinsNormalized(fees)
	if err != nil {
		panic(err)
	}
	b.fees = parsedFees
	return b
}

// WithGasLimit sets the gas limit
func (b *TxBuilder) WithGasLimit(gasLimit uint64) *TxBuilder {
	b.gasLimit = gasLimit
	return b
}

// Build builds and signs the transaction
func (b *TxBuilder) Build(privKey cryptotypes.PrivKey, accNum, accSeq uint64) (sdk.Tx, error) {
	txBuilder := b.txConfig.NewTxBuilder()
	
	if err := txBuilder.SetMsgs(b.msgs...); err != nil {
		return nil, err
	}
	
	txBuilder.SetMemo(b.memo)
	txBuilder.SetFeeAmount(b.fees)
	txBuilder.SetGasLimit(b.gasLimit)
	
	// Sign transaction
	sigV2 := signing.SignatureV2{
		PubKey: privKey.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  signing.SignMode_SIGN_MODE_DIRECT,
			Signature: nil,
		},
		Sequence: accSeq,
	}
	
	if err := txBuilder.SetSignatures(sigV2); err != nil {
		return nil, err
	}
	
	signerData := signing.SignerData{
		ChainID:       TestChainID,
		AccountNumber: accNum,
		Sequence:      accSeq,
	}
	
	sigV2, err := tx.SignWithPrivKey(
		signing.SignMode_SIGN_MODE_DIRECT,
		signerData,
		txBuilder,
		privKey,
		b.txConfig,
		accSeq,
	)
	if err != nil {
		return nil, err
	}
	
	if err := txBuilder.SetSignatures(sigV2); err != nil {
		return nil, err
	}
	
	return txBuilder.GetTx(), nil
}

// CreateTestTx creates a simple test transaction
func CreateTestTx(msg sdk.Msg, privKey cryptotypes.PrivKey) sdk.Tx {
	builder := NewTxBuilder(nil) // In real test, pass actual txConfig
	builder.WithMsgs(msg).WithFees("1000mcell")
	
	tx, err := builder.Build(privKey, 0, 0)
	if err != nil {
		panic(err)
	}
	
	return tx
}

