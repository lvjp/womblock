package util

import (
	"github.com/spf13/cobra"
)

type Factory struct {
	verbose *bool
}

func NewFactory(verbose *bool) *Factory {
	return &Factory{
		verbose: verbose,
	}
}

func (f *Factory) NewContext(cmd *cobra.Command) *Context {
	return NewContext(cmd, *f.verbose)
}
