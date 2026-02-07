package cmd

import (
	"fmt"
	"os"

	"github.com/lvjp/womblock/internal/pkg/cmd/util"
	"github.com/lvjp/womblock/pkg/buildinfo"

	"github.com/spf13/cobra"
)

var factory *util.Factory

var rootCmd = &cobra.Command{
	Use:           "womblock",
	Short:         "Leveraging blockchain to help wombats.",
	SilenceErrors: true,
	SilenceUsage:  true,
	Version:       buildinfo.Get().String(),
	Run: func(cmd *cobra.Command, _ []string) {
		ctx := factory.NewContext(cmd)
		ctx.Logger.Info().Msg("Hello, Womblock!")
	},
}

func init() {
	flags := rootCmd.PersistentFlags()
	verbose := flags.Bool("verbose", false, "Enable verbose logging (debug level)")

	factory = util.NewFactory(verbose)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
