package cli

import (
	"fmt"

	"github.com/Uemerson/go-password-generator-cobra-cli/pkg/password"
	"github.com/spf13/cobra"
)

var length int
var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Generate a random password",
	Run: func(cmd *cobra.Command, args []string) {
		res := password.Generate(length)
		fmt.Println(res)
	},
}

func init() {
	generateCmd.Flags().IntVarP(&length, "length", "l", 10, "Password length")
	rootCmd.AddCommand(generateCmd)
}
