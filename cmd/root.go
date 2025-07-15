package cmd

import (
	"fmt"
	"os"

	"github.com/lucmahoux/eva/cmd/branch"
	initcmd "github.com/lucmahoux/eva/cmd/init"
	"github.com/lucmahoux/eva/cmd/open"
	"github.com/lucmahoux/eva/cmd/update"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "eva",
	Short: "🤖 EVA - CLI tool to automate dev workflows",
}

func Execute() {
	rootCmd.AddCommand(branch.NewCommand())
	rootCmd.AddCommand(open.NewCommand())
	rootCmd.AddCommand(update.NewCommand())
	rootCmd.AddCommand(initcmd.NewCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
