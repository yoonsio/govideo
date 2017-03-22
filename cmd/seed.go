package cmd

import (
	"log"

	"github.com/sickyoon/govideo/govideo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed database with test data",
	Long: `Seed database with test data

* create user
`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Seeding started...")
		config := viper.GetString("config")
		app := govideo.NewApp(config)
		err := app.Seed()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(seedCmd)
}
