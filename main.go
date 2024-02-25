package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nexryai/pgharmony/ui"
	"net/http"
)

func main() {
	app := fiber.New()

	embedFS := ui.StaticFiles
	/*staticFS, err := fs.Sub(embedFS, "dist")
	if err != nil {
		panic(err)
	}*/

	app.Get("/", func(ctx fiber.Ctx) error {
		rawIndexHtml, err := embedFS.ReadFile("dist/index.html")
		if err != nil {
			return ctx.Status(http.StatusNotFound).SendString("Not Found")
		}

		ctx.Set("Content-Type", "text/html")
		return ctx.Send(rawIndexHtml)
	})

	app.Get("/favicon.ico", func(ctx fiber.Ctx) error {
		rawFaviconIco, err := embedFS.ReadFile("dist/favicon.ico")
		if err != nil {
			return ctx.Status(http.StatusNotFound).SendString("Not Found")
		}

		ctx.Set("Content-Type", "image/x-icon")
		return ctx.Send(rawFaviconIco)
	})

	// /statuc/js/*.js
	app.Get("/static/js/*", func(ctx fiber.Ctx) error {
		rawStaticFile, err := embedFS.ReadFile("dist/static/js/" + ctx.Params("*"))
		if err != nil {
			return ctx.Status(http.StatusNotFound).SendString("Not Found")
		}

		ctx.Set("Content-Type", "application/javascript")
		return ctx.Send(rawStaticFile)
	})

	app.Get("/static/css/*", func(ctx fiber.Ctx) error {
		rawStaticFile, err := embedFS.ReadFile("dist/static/css/" + ctx.Params("*"))
		if err != nil {
			return ctx.Status(http.StatusNotFound).SendString("Not Found")
		}

		ctx.Set("Content-Type", "text/css")
		return ctx.Send(rawStaticFile)
	})

	app.Listen(":3000")
}
