package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "govideo",
	Short: "Fast Video Streaming Server in Go",
	Long: `Fast Video Streaming Server in Go.
	
GoVideo is video streaming server that is designed with minimalistic approach.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize()
	RootCmd.PersistentFlags().String("config", "config.toml", "configuration file")
	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
}
