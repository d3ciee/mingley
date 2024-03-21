package main

import (
	userRoutes "mingley/app/user/routes"
	db "mingley/config"
	views "mingley/views/pages"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {
	err := db.InitDB()

	if err != nil {
		panic("Failed to connect to db")
	}

	var app *fiber.App = fiber.New()

	app.Static(
		"/public",
		"static",
	)
	userRoutes.SetupUserRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return adaptor.HTTPHandler(templ.Handler(views.Index()))(c)

	})

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("not found")
	})

	app.Listen(":80")
}
