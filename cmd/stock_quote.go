package main

import (
	"errors"

	"github.com/cpliakas/cliutil"
	"github.com/cpliakas/finance/alphavantage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var stockQuoteCfg *viper.Viper

var stockQuoteCmd = &cobra.Command{
	Use:   "quote [SYMBOL]",
	Short: "returns the latest price and volume information for the global equity specified",
	Args:  stockQuoteCmdValidate,
	Run: func(cmd *cobra.Command, args []string) {
		client := newClient()
		input := &alphavantage.StockQuoteInput{Symbol: args[0]}
		output, err := client.StockQuote(input)
		cliutil.HandleError(cmd, err)
		printJSON(cmd, output)
	},
}

func init() {
	stockCmd.AddCommand(stockQuoteCmd)
	stockQuoteCfg = cliutil.InitConfig(EnvPrefix)
}

func stockQuoteCmdValidate(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("missing required argument: [SYMBOL]")
	}
	return nil
}
