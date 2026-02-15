package cmd

import (
	"github.com/lvjp/womblock/internal/app/cmd/healthcheck"
	"github.com/lvjp/womblock/internal/pkg/cmd/util"

	"github.com/spf13/cobra"
)

func NewHealthCheckCmd(factory *util.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "healthcheck",
		Short: "Check the health of the Womblock API server",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := factory.NewContext(cmd)
			return healthcheck.Run(ctx)
		},
	}
}
