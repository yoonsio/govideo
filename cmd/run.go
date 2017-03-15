package cmd

import (
	"fmt"

	"github.com/sickyoon/govideo/govideo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run Server",
	Long:  `Runs web server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
		****************************************
		* RUNNING GOVIDEO                      *
		****************************************
		`)
		config := viper.GetString("config")
		app := govideo.NewApp(config)
		app.Run()
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
}
