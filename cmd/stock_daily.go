package main

import (
	"errors"

	"github.com/cpliakas/cliutil"
	"github.com/cpliakas/finance/alphavantage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var stockDailyCfg *viper.Viper

var stockDailyCmd = &cobra.Command{
	Use:   "daily [SYMBOL]",
	Short: "returns daily time series of the global equity specified",
	Args:  stockDailyCmdValidate,
	Run: func(cmd *cobra.Command, args []string) {
		client := newClient()

		input := &alphavantage.StockDailyInput{Symbol: args[0]}
		if stockDailyCfg.GetBool("full") {
			input.OutputSize = alphavantage.OutputSizeFull
		}

		output, err := client.StockDaily(input)
		cliutil.HandleError(cmd, err)
		printJSON(cmd, output)
	},
}

func init() {
	stockCmd.AddCommand(stockDailyCmd)
	stockDailyCfg = cliutil.InitConfig(EnvPrefix)

	flags := cliutil.NewFlagger(stockDailyCmd, stockDailyCfg)
	flags.Bool("full", "f", false, "return the full-length time series of 20+ years of historical data")
}

func stockDailyCmdValidate(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("missing required argument: [SYMBOL]")
	}
	return nil
}
