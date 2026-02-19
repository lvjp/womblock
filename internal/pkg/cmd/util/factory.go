package util

import (
	"github.com/spf13/cobra"
)

type Factory struct {
	configPath *string
	verbose    *bool
}

func NewFactory(configPath *string, verbose *bool) *Factory {
	return &Factory{
		configPath: configPath,
		verbose:    verbose,
	}
}

func (f *Factory) NewContext(cmd *cobra.Command) *Context {
	return NewContext(cmd, *f.configPath, *f.verbose)
}
