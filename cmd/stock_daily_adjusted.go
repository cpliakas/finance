package main

import (
	"errors"

	"github.com/cpliakas/cliutil"
	"github.com/cpliakas/finance/alphavantage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var stockDailyAdjustedCfg *viper.Viper

var stockDailyAdjustedCmd = &cobra.Command{
	Use:   "daily-adjusted [SYMBOL]",
	Short: "returns daily adjusted time series of the global equity specified",
	Args:  stockDailyAdjustedCmdValidate,
	Run: func(cmd *cobra.Command, args []string) {
		client := newClient()

		input := &alphavantage.StockDailyAdjustedInput{Symbol: args[0]}
		if stockDailyAdjustedCfg.GetBool("full") {
			input.OutputSize = alphavantage.OutputSizeFull
		}

		output, err := client.StockDailyAdjusted(input)
		cliutil.HandleError(cmd, err)
		printJSON(cmd, output)
	},
}

func init() {
	stockCmd.AddCommand(stockDailyAdjustedCmd)
	stockDailyAdjustedCfg = cliutil.InitConfig(EnvPrefix)

	flags := cliutil.NewFlagger(stockDailyAdjustedCmd, stockDailyAdjustedCfg)
	flags.Bool("full", "f", false, "return the full-length time series of 20+ years of historical data")
}

func stockDailyAdjustedCmdValidate(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("missing required argument: [SYMBOL]")
	}
	return nil
}
