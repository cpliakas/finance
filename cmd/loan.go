package main

import (
	"errors"

	"github.com/cpliakas/cliutil"
	"github.com/cpliakas/finance/alphavantage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var loanCfg *viper.Viper

var loanCmd = &cobra.Command{
	Use:     "loan",
	Aliases: []string{"mortgage"},
	Short:   "calculate montly and total payments of loan",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(loanCmd)
	loanCfg = cliutil.InitConfig(alphavantage.EnvPrefix)

	flags := cliutil.NewFlagger(loanCmd, loanCfg)
	flags.PersistentFloat64("amount", "a", 0, "principal amount of the loan")
	flags.PersistentFloat64("rate", "r", 0, "annual intrest rate")
	flags.PersistentInt("years", "y", 0, "loan term in years")
}

func mortgageCmdValidate(cmd *cobra.Command, args []string) error {
	if loanCfg.GetFloat64("amount") == 0 {
		return errors.New("missing required option: amount")
	}
	if loanCfg.GetFloat64("rate") == 0 {
		return errors.New("missing required option: rate")
	}
	if loanCfg.GetInt("years") == 0 {
		return errors.New("missing required option: years")
	}
	return nil
}
