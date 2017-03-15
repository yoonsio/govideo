package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed database with test data",
	Long: `Seed database with test data

* create user
`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: drop database
		// TODO: add database seed
		fmt.Println("Seeding database...")
	},
}

func init() {
	RootCmd.AddCommand(seedCmd)
}
