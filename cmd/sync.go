package cmd

import (
	"log"

	"github.com/sickyoon/govideo/govideo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync Database with Media Files",
	Long:  `Sync Database with Media Files`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Sync started...")
		config := viper.GetString("config")
		app := govideo.NewApp(config)
		err := app.Sync()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(syncCmd)
}
