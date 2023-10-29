package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"os"
)

var (
	Version = "v0.0.1"
	rootCmd = &cobra.Command{
		Use:   "exchange-rate",
		Short: "Services for work with exchange rate currencies",
	}
)

func init() {
	cfgFile := os.Getenv("EXCHANGE_RATE")
	if cfgFile != "" {
		err := godotenv.Load(cfgFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	rootCmd.AddCommand(runConsoleCmd)
}
