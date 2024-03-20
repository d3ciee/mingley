package main

import (
	views "mingley/views/pages"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {

	var app *fiber.App = fiber.New()
	app.Static(
		"/static",
		"./public",
	)

	app.Get("/", func(c *fiber.Ctx) error {
		return adaptor.HTTPHandler(templ.Handler(views.Index()))(c)

	})
	app.Listen(":3000")
}
