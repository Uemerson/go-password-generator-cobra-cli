package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "password-generator",
	Short: "password-generator - a simple CLI to generates random passwords",
	Long: `password-generator is a simple CLI
   	
	Generates random passwords with customizable length and complexity options provided via command-line flags`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
