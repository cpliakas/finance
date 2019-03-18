package main

import (
	"github.com/cpliakas/cliutil"
	"github.com/cpliakas/finance/alphavantage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var stockCfg *viper.Viper

var stockCmd = &cobra.Command{
	Use:   "stock",
	Short: "realtime and historical global equity data",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(stockCmd)
	stockCfg = cliutil.InitConfig(alphavantage.EnvPrefix)

	flags := cliutil.NewFlagger(stockCmd, stockCfg)
	flags.PersistentString("api-key", "a", "", "api key used to authenticate requests to alpha vantage")
}

func newClient() alphavantage.Client {
	cfg := alphavantage.NewDefaultConfig(stockCfg.GetString("api-key"))
	return alphavantage.NewClient(cfg)
}
