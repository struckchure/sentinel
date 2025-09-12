package main

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	version   = "dev" // default value
	commit    = "none"
	buildDate = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\nGo Version: %s\nCommit: %s\nBuild Date: %s\n", version, runtime.Version(), commit, buildDate)
	},
}

func init() {
	ignoreForRootConfig = append(ignoreForRootConfig, versionCmd.Name())

	rootCmd.AddCommand(versionCmd)
}
