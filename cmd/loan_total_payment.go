package main

import (
	"github.com/cpliakas/cliutil"
	"github.com/cpliakas/finance"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var loanTotalPaymentCfg *viper.Viper

var loanTotalPaymentCmd = &cobra.Command{
	Use:   "total-payment",
	Short: "calculate the total payment for a loan over its entire term",
	Args:  mortgageCmdValidate,
	Run: func(cmd *cobra.Command, args []string) {
		a := loanCfg.GetFloat64("amount")
		r := loanCfg.GetFloat64("rate")
		y := loanCfg.GetInt("years")
		p := finance.TotalLoanPayment(a, r, y)

		finance.PrintlnDollars(p)
	},
}

func init() {
	loanCmd.AddCommand(loanTotalPaymentCmd)
	loanTotalPaymentCfg = cliutil.InitConfig(EnvPrefix)
}
