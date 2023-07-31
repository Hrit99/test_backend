package routes

import (
	"github.com/gofiber/fiber/v2"
)

func setRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/readers", func(c *fiber.Ctx) error {
		return c.SendString("on route readers")
	})
}
