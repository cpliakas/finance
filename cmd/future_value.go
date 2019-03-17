package main

import (
	"fmt"

	"github.com/cpliakas/cliutil"
	"github.com/cpliakas/finance"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var futureValueCfg *viper.Viper

var futureValueCmd = &cobra.Command{
	Use:     "future-value",
	Aliases: []string{"fv"},
	Short:   "calculates the value of a cash flow at a later date than originally received",
	Run: func(cmd *cobra.Command, args []string) {
		amount := finance.FutureValue(
			futureValueCfg.GetFloat64("principal"),
			futureValueCfg.GetFloat64("interest-rate"),
			futureValueCfg.GetInt("times-per-year"),
			futureValueCfg.GetInt("years"),
		)

		fmt.Printf("$%v\n", amount)
	},
}

func init() {
	rootCmd.AddCommand(futureValueCmd)
	futureValueCfg = cliutil.InitConfig(EnvPrefix)

	flags := cliutil.NewFlagger(futureValueCmd, futureValueCfg)
	flags.Float64("principal", "p", 0.00, "amount of money that you have available to invest initially")
	flags.Float64("interest-rate", "r", 0.00, "estimated annual interest rate as a decimal")
	flags.Int("times-per-year", "t", finance.PeriodAnnual, "times per year that interest will be compounded")
	flags.Int("years", "y", 0, "length of time, in years, that you plan to save")
}
