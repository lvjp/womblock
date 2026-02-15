package cmd

import (
	"fmt"
	"os"

	"github.com/lvjp/womblock/internal/app/config"
	"github.com/lvjp/womblock/internal/pkg/cmd/util"
	"github.com/lvjp/womblock/pkg/buildinfo"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "womblock",
	Short:         "Leveraging blockchain to help wombats.",
	SilenceErrors: true,
	SilenceUsage:  true,
	Version:       buildinfo.Get().String(),
}

func init() {
	flags := rootCmd.PersistentFlags()
	config := flags.String("config", config.DefaultConfigPath, "Path to the configuration file")
	verbose := flags.Bool("verbose", false, "Enable verbose logging (debug level)")

	factory := util.NewFactory(config, verbose)
	rootCmd.AddCommand(NewServerCmd(factory))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
