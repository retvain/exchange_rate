package cmd

import "github.com/spf13/cobra"

var runConsoleCmd = &cobra.Command{
	Use:   "console",
	Short: "Run console",
	Long:  "Run application in console mode",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runConsole()
	},
}

func runConsole() error {
	return nil
}
