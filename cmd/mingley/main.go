package main

import (
	userRoutes "mingley/app/user/routes"
	db "mingley/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
)

func main() {
	err := db.InitDB()
	engine := jet.New("./views", ".jet")

	if err != nil {
		panic("Failed to connect to db")
	}

	var app *fiber.App = fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static(
		"/public",
		"static",
		fiber.Static{
			MaxAge: 0,
		},
	)
	userRoutes.SetupUserRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("not found")
	})

	app.Listen(":80")
}
