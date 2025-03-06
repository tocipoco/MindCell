package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ExportAppStateAndValidators exports the state of the application for a genesis file.
func (app *MindCellApp) ExportAppStateAndValidators(forZeroHeight bool, jailAllowedAddrs []string) (servertypes.ExportedApp, error) {
	ctx := app.NewContext(true, sdk.NewBlockHeader())

	if forZeroHeight {
		app.prepareForZeroHeightExport(ctx)
	}

	genState := app.mm.ExportGenesis(ctx, app.appCodec)
	appState, err := app.appCodec.MarshalJSON(genState)
	if err != nil {
		return servertypes.ExportedApp{}, err
	}

	validators := []types.ValidatorUpdate{}
	return servertypes.ExportedApp{
		AppState:        appState,
		Validators:      validators,
		Height:          app.LastBlockHeight(),
		ConsensusParams: app.BaseApp.GetConsensusParams(ctx),
	}, nil
}

func (app *MindCellApp) prepareForZeroHeightExport(ctx sdk.Context) {
	// Prepare application state for export at zero height
}

