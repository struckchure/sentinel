package main

import (
	"github.com/struckchure/sentinel"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start gateway service",
	Run: func(cmd *cobra.Command, args []string) {
		if err := gateway.Run(); err != nil {
			new(sentinel.Logger).Error(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
