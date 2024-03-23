package user

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	userRoutes := app.Group("/users")

	userRoutes.Get("/sign-up", func(c *fiber.Ctx) error {
		return c.Render("/pages/user/auth", fiber.Map{
			"title": "Mingley | Sign up",
			"auth":  "sign-up",
		})
	})
	userRoutes.Post("/sign-up", func(c *fiber.Ctx) error {
		payload := struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}{}

		time.Sleep(2 * time.Second)

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		return c.JSON("An error occured whilst handling it")
	})
}
