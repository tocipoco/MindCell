package app

import (
	"io"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"github.com/cometbft/cometbft/abci/types"
)

const (
	AppName = "MindCell"
)

var (
	// DefaultNodeHome is the default home directory for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
	)
)

// MindCellApp extends an ABCI application
type MindCellApp struct {
	*baseapp.BaseApp

	cdc               *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	// keepers
	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.Keeper
	StakingKeeper *stakingkeeper.Keeper

	// module manager
	mm *module.Manager
}

// NewMindCellApp returns a reference to an initialized MindCellApp.
func NewMindCellApp(
	logger server.Logger,
	db server.Database,
	traceStore io.Writer,
	loadLatest bool,
	appOpts servertypes.AppOptions,
) *MindCellApp {
	encodingConfig := MakeEncodingConfig()

	bApp := baseapp.NewBaseApp(AppName, logger, db, encodingConfig.TxConfig.TxDecoder())
	bApp.SetCommitMultiStoreTracer(traceStore)

	app := &MindCellApp{
		BaseApp:           bApp,
		cdc:               encodingConfig.Amino,
		appCodec:          encodingConfig.Codec,
		interfaceRegistry: encodingConfig.InterfaceRegistry,
	}

	// Set module manager
	app.mm = module.NewManager()

	return app
}

// Name returns the name of the App
func (app *MindCellApp) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *MindCellApp) BeginBlocker(ctx sdk.Context, req types.RequestBeginBlock) types.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *MindCellApp) EndBlocker(ctx sdk.Context, req types.RequestEndBlock) types.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
func (app *MindCellApp) InitChainer(ctx sdk.Context, req types.RequestInitChain) types.ResponseInitChain {
	var genesisState map[string]interface{}
	if err := app.appCodec.UnmarshalJSON(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LegacyAmino returns SimApp's amino codec.
func (app *MindCellApp) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns MindCell's app codec.
func (app *MindCellApp) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns MindCell's InterfaceRegistry
func (app *MindCellApp) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

