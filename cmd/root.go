package main

import (
	"encoding/json"
	"os"

	"github.com/cpliakas/cliutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// EnvPrefix is the environment variable prefix for configuration.
const EnvPrefix = "FINANCE"

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
	rootCfg = cliutil.InitConfig(EnvPrefix)
}

func printJSON(cmd *cobra.Command, v interface{}) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	err := enc.Encode(v)
	cliutil.HandleError(cmd, err)
}
