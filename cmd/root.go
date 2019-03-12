package main

import (
	"os"

	"github.com/cpliakas/cliutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCfg *viper.Viper

var rootCmd = &cobra.Command{
	Use:   "finance",
	Short: "A command line tool that executes finance calculations and operations.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCfg = initConfig()
}

func initConfig() *viper.Viper {
	return cliutil.InitConfig("FINANCE")
}
