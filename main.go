package main

import (
	"github.com/Hrit99/test_backend.git/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

	routes.setRoutes()
    app.Listen(":3000")
}