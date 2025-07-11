package nyx

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nyx",
	Short: "Nyx is a Nyx is a GoLang demo app to use for validating K8s controllers",
	Long:  `Nyx is a Nyx is a GoLang demo app to use for validating K8s controllers`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
