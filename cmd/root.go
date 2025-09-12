package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/samber/lo"
	"github.com/spf13/cobra"

	"github.com/struckchure/sentinel"
)

var ignoreForRootConfig = []string{"playground"}

var rootCmd = &cobra.Command{
	Use:   "sentinel",
	Short: "A simple and lightweight api gateway built with Golang.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if lo.Contains(ignoreForRootConfig, cmd.Name()) {
			return
		}

		loader := sentinel.NewConfigLoader()

		_configFile, _ := cmd.Flags().GetString("config")
		_configType := filepath.Ext(_configFile)[1:]
		_config, err := loader.Load(_configFile, sentinel.ConfigType(_configType))
		if err != nil {
			log.Panic(err)
		}

		config = _config
		gateway = sentinel.NewGateway(*config)
	},
}

var (
	configFile string
	config     *sentinel.Config
	gateway    sentinel.IGateway
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "sentinel.yaml", "sentinel config file")

	if version == "dev" {
		rootCmd.AddCommand(&cobra.Command{
			Use: "playground",
			Run: func(cmd *cobra.Command, args []string) { sentinel.Play() },
		})
	}
}
