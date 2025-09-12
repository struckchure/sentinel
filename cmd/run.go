package main

import (
	"log"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start gateway service",
	Run: func(cmd *cobra.Command, args []string) {
		if err := gateway.Run(); err != nil {
			log.Panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
