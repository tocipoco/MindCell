package main

import (
	"errors"
	"io"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/spf13/cobra"

	"github.com/tocipoco/MindCell/app"
)

// NewRootCmd creates a new root command for mindcelld.
func NewRootCmd() *cobra.Command {
	initClientCtx := client.Context{}.
		WithCodec(app.MakeEncodingConfig().Codec).
		WithInterfaceRegistry(app.MakeEncodingConfig().InterfaceRegistry).
		WithTxConfig(app.MakeEncodingConfig().TxConfig).
		WithLegacyAmino(app.MakeEncodingConfig().Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(nil).
		WithHomeDir(app.DefaultNodeHome).
		WithViper("")

	rootCmd := &cobra.Command{
		Use:   "mindcelld",
		Short: "MindCell Daemon",
		Long:  "MindCell is a decentralized protocol for AI model sharding and inference.",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())

			initClientCtx, err := client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			initClientCtx, err = config.ReadFromClientConfig(initClientCtx)
			if err != nil {
				return err
			}

			if err := client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}

			customAppTemplate, customAppConfig := initAppConfig()
			return server.InterceptConfigsPreRunHandler(cmd, customAppTemplate, customAppConfig, nil)
		},
	}

	initRootCmd(rootCmd, initClientCtx)
	return rootCmd
}

func initRootCmd(rootCmd *cobra.Command, clientCtx client.Context) {
	rootCmd.AddCommand(
		InitCmd(app.ModuleBasics, app.DefaultNodeHome),
		server.NewStartCmd(createAppServer, nil),
		flags.NewCompletionCmd(rootCmd, true),
	)
}

func createAppServer(logger server.Logger, db server.Database, traceStore io.Writer, appOpts servertypes.AppOptions) servertypes.Application {
	return app.NewMindCellApp(
		logger,
		db,
		traceStore,
		true,
		appOpts,
	)
}

func initAppConfig() (string, interface{}) {
	type CustomAppConfig struct {
		serverconfig.Config
	}

	customAppConfig := CustomAppConfig{
		Config: *serverconfig.DefaultConfig(),
	}

	customAppTemplate := serverconfig.DefaultConfigTemplate

	return customAppTemplate, customAppConfig
}

