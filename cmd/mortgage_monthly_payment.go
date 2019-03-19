package main

import (
	"fmt"

	"github.com/cpliakas/cliutil"
	"github.com/cpliakas/finance"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var mortgageMonthlyPaymentCfg *viper.Viper

var mortgageMonthlyPaymentCmd = &cobra.Command{
	Use:   "monthly-payment",
	Short: "calculate the monthly payment for a mortgage",
	Args:  mortgageCmdValidate,
	Run: func(cmd *cobra.Command, args []string) {
		a := mortgageCfg.GetFloat64("amount")
		r := mortgageCfg.GetFloat64("rate")
		y := mortgageCfg.GetInt("years")
		p := finance.MonthlyMortgagePayment(a, r, y)
		fmt.Printf("$%.2f\n", p)
	},
}

func init() {
	mortgageCmd.AddCommand(mortgageMonthlyPaymentCmd)
	mortgageMonthlyPaymentCfg = cliutil.InitConfig(EnvPrefix)
}
