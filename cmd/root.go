package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kraken",
	Short: "Kraken is a microservice generator",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
