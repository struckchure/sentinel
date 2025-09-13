package main

import (
	"github.com/struckchure/sentinel"

	"github.com/spf13/cobra"
)

var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Print config schema",
	Run: func(cmd *cobra.Command, args []string) {
		save, _ := cmd.Flags().GetBool("save")
		indentation, _ := cmd.Flags().GetInt("indentation")
		output, _ := cmd.Flags().GetString("output")

		config := sentinel.NewConfigLoader()
		if err := config.Schema(save, indentation, output); err != nil {
			new(sentinel.Logger).Error(err.Error())
		}
	},
}

func init() {
	ignoreForRootConfig = append(ignoreForRootConfig, schemaCmd.Name())

	schemaCmd.Flags().IntP("indentation", "i", 2, "indentation size")
	schemaCmd.Flags().BoolP("save", "s", false, "save scheme to file")
	schemaCmd.Flags().StringP("output", "o", "sentinel.schema.json", "schema outfile file")

	rootCmd.AddCommand(schemaCmd)
}
