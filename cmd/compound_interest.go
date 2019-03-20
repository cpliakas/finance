package main

import (
	"github.com/cpliakas/cliutil"
	"github.com/cpliakas/finance"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var compoundInterestCfg *viper.Viper

var compoundInterestCmd = &cobra.Command{
	Use:     "compound-interest",
	Aliases: []string{"ci"},
	Short:   "interest calculated on the initial principal inclusive of the accumulated interest over a period of time",
	Run: func(cmd *cobra.Command, args []string) {
		amount := finance.CompoundInterest(
			compoundInterestCfg.GetFloat64("principal"),
			compoundInterestCfg.GetFloat64("interest-rate"),
			compoundInterestCfg.GetInt("times-per-year"),
			compoundInterestCfg.GetInt("years"),
		)

		finance.PrintlnDollars(amount)
	},
}

func init() {
	rootCmd.AddCommand(compoundInterestCmd)
	compoundInterestCfg = cliutil.InitConfig(EnvPrefix)

	flags := cliutil.NewFlagger(compoundInterestCmd, compoundInterestCfg)
	flags.Float64("principal", "p", 0.00, "amount of money that you have available to invest initially")
	flags.Float64("interest-rate", "r", 0.00, "estimated annual interest rate as a decimal")
	flags.Int("times-per-year", "t", finance.PeriodAnnual, "times per year that interest will be compounded")
	flags.Int("years", "y", 0, "length of time, in years, that you plan to save")
}
