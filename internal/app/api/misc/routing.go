package misc

import (
	"github.com/gofiber/fiber/v3"
)

func Route(app fiber.Router, service Service) {
	app.Get("/version", VersionHandler(service))
	app.Get("/health", HealthHandler(service))
}
