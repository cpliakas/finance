package main

import (
	"github.com/cpliakas/cliutil"
	"github.com/cpliakas/finance"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var mortgageTotalPaymentCfg *viper.Viper

var mortgageTotalPaymentCmd = &cobra.Command{
	Use:   "total-payment",
	Short: "calculate the total payment for a mortgage over its entire term",
	Args:  mortgageCmdValidate,
	Run: func(cmd *cobra.Command, args []string) {
		a := mortgageCfg.GetFloat64("amount")
		r := mortgageCfg.GetFloat64("rate")
		y := mortgageCfg.GetInt("years")
		p := finance.TotalMortgagePayment(a, r, y)

		finance.PrintlnDollars(p)
	},
}

func init() {
	mortgageCmd.AddCommand(mortgageTotalPaymentCmd)
	mortgageTotalPaymentCfg = cliutil.InitConfig(EnvPrefix)
}
