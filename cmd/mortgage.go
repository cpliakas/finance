package main

import (
	"errors"

	"github.com/cpliakas/cliutil"
	"github.com/cpliakas/finance/alphavantage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var mortgageCfg *viper.Viper

var mortgageCmd = &cobra.Command{
	Use:   "mortgage",
	Short: "realtime and historical global equity data",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(mortgageCmd)
	mortgageCfg = cliutil.InitConfig(alphavantage.EnvPrefix)

	flags := cliutil.NewFlagger(mortgageCmd, mortgageCfg)
	flags.PersistentFloat64("amount", "a", 0, "principal amount of the mortgage")
	flags.PersistentFloat64("rate", "r", 0, "annual intrest rate")
	flags.PersistentInt("years", "y", 0, "loan term in years")
}

func mortgageCmdValidate(cmd *cobra.Command, args []string) error {
	if mortgageCfg.GetFloat64("amount") == 0 {
		return errors.New("missing required option: amount")
	}
	if mortgageCfg.GetFloat64("rate") == 0 {
		return errors.New("missing required option: rate")
	}
	if mortgageCfg.GetInt("years") == 0 {
		return errors.New("missing required option: years")
	}
	return nil
}
