package main

import (
	"github.com/cpliakas/cliutil"
	"github.com/cpliakas/finance"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var presentValueCfg *viper.Viper

var presentValueCmd = &cobra.Command{
	Use:     "present-value",
	Aliases: []string{"pv"},
	Short:   "calculates the present day value of an amount that is received at a future date",
	Run: func(cmd *cobra.Command, args []string) {
		principal := finance.PresentValue(
			presentValueCfg.GetFloat64("desired-amount"),
			presentValueCfg.GetFloat64("interest-rate"),
			presentValueCfg.GetInt("times-per-year"),
			presentValueCfg.GetInt("years"),
		)

		finance.PrintlnDollars(principal)
	},
}

func init() {
	rootCmd.AddCommand(presentValueCmd)
	presentValueCfg = cliutil.InitConfig(EnvPrefix)

	flags := cliutil.NewFlagger(presentValueCmd, presentValueCfg)
	flags.Float64("desired-amount", "a", 0.00, "the desired amount at the end of the time period")
	flags.Float64("interest-rate", "r", 0.00, "estimated annual interest rate as a decimal")
	flags.Int("times-per-year", "t", finance.PeriodAnnual, "times per year that interest will be compounded")
	flags.Int("years", "y", 0, "length of time, in years, that you plan to save")
}
