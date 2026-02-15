package cmd

import (
	"github.com/lvjp/womblock/internal/app/cmd/serve"
	"github.com/lvjp/womblock/internal/pkg/cmd/util"

	"github.com/spf13/cobra"
)

func NewServerCmd(factory *util.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Serve the Womblock API",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := factory.NewContext(cmd)
			return serve.Run(ctx)
		},
	}
}
