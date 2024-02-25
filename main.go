package main

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/nexryai/pgharmony/internal/core/config"
	"github.com/nexryai/pgharmony/internal/core/logger"
	"github.com/nexryai/pgharmony/internal/server"
)

func main() {
	app := fiber.New()
	log := logger.GetLogger("INIT")

	// ルーティング
	server.CreateRouter(app)

	err := app.Listen(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		log.FatalWithDetail("Failed to start server", err)
	}
}
