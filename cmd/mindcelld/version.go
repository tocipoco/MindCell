package main

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

const AppVersion = "v1.0.0"

// versionCmd represents the version command
func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the application version",
		Long:  "Print the version information for mindcelld binary",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("MindCell Version: %s\n", AppVersion)
			fmt.Printf("Git Commit: %s\n", getGitCommit())
			fmt.Printf("Go Version: %s\n", runtime.Version())
			fmt.Printf("OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
		},
	}
}

func getGitCommit() string {
	// In production, this would be set via ldflags during build
	return "dev"
}
