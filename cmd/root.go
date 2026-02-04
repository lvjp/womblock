package cmd

import (
	"fmt"
	"os"

	"github.com/lvjp/womblock/pkg/buildinfo"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "womblock",
	Short:         "Leveraging blockchain to help wombats.",
	SilenceErrors: true,
	SilenceUsage:  true,
	Version:       buildinfo.Get().String(),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(cmd.OutOrStdout(), "Hello, Womblock!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
