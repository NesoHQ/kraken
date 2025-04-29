package cmd

import (
	"github.com/spf13/cobra"

	"github.com/NesoHQ/kraken/internal/generator"
)

var initCmd = &cobra.Command{
	Use:   "init [service-name]",
	Short: "Initialize a new microservice",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := args[0]
		generator.GenerateService(serviceName)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
