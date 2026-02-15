package serve

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"

	"github.com/lvjp/womblock/internal/pkg/cmd/util"

	fiberzerolog "github.com/gofiber/contrib/v3/zerolog"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/rs/zerolog"
)

func Run(ctx *util.Context) error {
	var cancel context.CancelFunc
	ctx.Context, cancel = context.WithCancel(ctx.Context)
	defer cancel()

	server := newFiberApp(&ctx.Logger)
	server.All("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, Wombat!")
	})

	var serverErr error
	go func() {
		defer cancel()

		serverErr = server.Listen(*ctx.Config.Server.ListenAddress)
	}()

	<-ctx.Done()
	ctx.Logger.Info().Msg("Server shutdown sequence started")

	if serverErr != nil && !errors.Is(serverErr, http.ErrServerClosed) {
		return fmt.Errorf("ListenAndServe error: %v", serverErr)
	}

	if err := server.Shutdown(); err != nil {
		return fmt.Errorf("could not shutdown server: %v", err)
	}

	return nil
}

func newFiberApp(logger *zerolog.Logger) *fiber.App {
	app := fiber.New()

	app.Hooks().OnListen(func(listenData fiber.ListenData) error {
		u := url.URL{
			Scheme: "http",
			Host:   net.JoinHostPort(listenData.Host, listenData.Port),
		}

		if listenData.TLS {
			u.Scheme = "https"
		}

		logger.Info().
			Stringer("endpoint", &u).
			Msg("Listening")

		return nil
	})

	app.Hooks().OnPostShutdown(func(err error) error {
		logger.Info().Err(err).Msg("Fiber shutdown done")
		return nil
	})

	app.Use(requestid.New())

	app.Use(func(c fiber.Ctx) error {
		ctx := logger.With().
			Str("requestId", requestid.FromContext(c)).
			Logger().
			WithContext(c.Context())
		c.SetContext(ctx)

		return c.Next()
	})

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		GetLogger: func(c fiber.Ctx) zerolog.Logger {
			return *zerolog.Ctx(c.Context())
		},
	}))

	return app
}
