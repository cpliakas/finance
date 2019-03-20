package main

import (
	"errors"

	"github.com/cpliakas/cliutil"
	"github.com/cpliakas/finance/alphavantage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var stockIntradayCfg *viper.Viper

var stockIntradayCmd = &cobra.Command{
	Use:   "intraday [SYMBOL]",
	Short: "returns intraday time series of the global equity specified",
	Args:  stockIntradayCmdValidate,
	Run: func(cmd *cobra.Command, args []string) {
		client := newClient()

		input := &alphavantage.StockIntradayInput{
			Symbol:   args[0],
			Interval: stockIntradayCfg.GetString("interval"),
		}
		if stockIntradayCfg.GetBool("full") {
			input.OutputSize = alphavantage.OutputSizeFull
		}

		output, err := client.StockIntraday(input)
		cliutil.HandleError(cmd, err)
		printJSON(cmd, output)
	},
}

func init() {
	stockCmd.AddCommand(stockIntradayCmd)
	stockIntradayCfg = cliutil.InitConfig(EnvPrefix)

	flags := cliutil.NewFlagger(stockIntradayCmd, stockIntradayCfg)
	flags.Bool("full", "f", false, "return the full-length time series of 20+ years of historical data")
	flags.String("interval", "i", alphavantage.Interval5Min, "time interval between two consecutive data points in the time series")
}

func stockIntradayCmdValidate(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("missing required argument: [SYMBOL]")
	}

	// TODO validate interval

	return nil
}
