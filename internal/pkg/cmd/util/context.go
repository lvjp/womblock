package util

import (
	"context"
	"fmt"
	"io"
	stdlog "log"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type Context struct {
	context.Context

	Input  io.Reader
	Output io.Writer
	Error  io.Writer
	Logger zerolog.Logger
}

// NewContext will print error on os.Stderr and exit with code 1 if any error occurs during initialization.
func NewContext(cmd *cobra.Command, verbose bool) *Context {
	ret := &Context{
		Context: cmd.Context(),

		Input:  cmd.InOrStdin(),
		Output: cmd.OutOrStdout(),
		Error:  cmd.ErrOrStderr(),
	}

	ret.CheckErr(ret.initLogger(verbose), 1)

	return ret
}

func (ctx *Context) CheckErr(err error, code int) {
	if err == nil {
		return
	}

	fmt.Fprintln(ctx.Error, "Error:", err)
	os.Exit(code)
}

func (ctx *Context) initLogger(verbose bool) error {
	writer := zerolog.ConsoleWriter{
		Out:        ctx.Output,
		TimeFormat: time.RFC3339,
	}

	level := zerolog.InfoLevel
	if verbose {
		level = zerolog.DebugLevel
	}

	ctx.Logger = zerolog.New(writer).With().Timestamp().Logger()
	ctx.Context = ctx.Logger.WithContext(ctx.Context)

	log.Logger = ctx.Logger.With().Str("component", "default logger").Logger()
	zerolog.DefaultContextLogger = &log.Logger
	zerolog.SetGlobalLevel(level)

	// Remove date/time flags which are already present in zerolog output
	stdlog.SetFlags(stdlog.Flags() & ^(stdlog.Ldate | stdlog.Ltime | stdlog.Lmicroseconds))
	stdlog.SetOutput(ctx.Logger.With().Str("level", "stdlog").Logger())

	return nil
}
