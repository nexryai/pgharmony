package server

import "github.com/gofiber/fiber/v3"

func createApiRouter(app *fiber.App) {
	app.Get("/api", func(ctx fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})
}
