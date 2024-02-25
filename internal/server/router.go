package server

import "github.com/gofiber/fiber/v3"

func CreateRouter(app *fiber.App) {
	createApiRouter(app)
	createEmbedFsRouter(app)
}
