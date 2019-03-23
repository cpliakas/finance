package main

import (
	"github.com/cpliakas/cliutil"
	"github.com/cpliakas/finance"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var loanMonthlyPaymentCfg *viper.Viper

var laonMonthlyPaymentCmd = &cobra.Command{
	Use:   "monthly-payment",
	Short: "calculate the monthly payment for a mortgage",
	Args:  mortgageCmdValidate,
	Run: func(cmd *cobra.Command, args []string) {
		a := loanCfg.GetFloat64("amount")
		r := loanCfg.GetFloat64("rate")
		y := loanCfg.GetInt("years")
		p := finance.MonthlyLoanPayment(a, r, y)

		finance.PrintlnDollars(p)
	},
}

func init() {
	loanCmd.AddCommand(laonMonthlyPaymentCmd)
	loanMonthlyPaymentCfg = cliutil.InitConfig(EnvPrefix)
}
