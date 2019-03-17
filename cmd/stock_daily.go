package main

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/cpliakas/cliutil"
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
		output, err := client.StockTimeSeriesDaily(args[0])
		cliutil.HandleError(cmd, err)

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "    ")
		err = enc.Encode(output)
		cliutil.HandleError(cmd, err)
	},
}

func init() {
	stockCmd.AddCommand(stockDailyCmd)
	stockDailyCfg = cliutil.InitConfig(EnvPrefix)

	// flags := cliutil.NewFlagger(stockDailyCmd, stockDailyCfg)
}

func stockDailyCmdValidate(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("missing required argument: [SYMBOL]")
	}

	return nil
}
