package user

import (
	userViews "mingley/views/pages/user"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func SetupUserRoutes(app *fiber.App) {
	userRoutes := app.Group("/users")

	userRoutes.Get("/sign-up", func(c *fiber.Ctx) error {
		return adaptor.HTTPHandler(templ.Handler(
			userViews.SignIn()))(c)
	})
	userRoutes.Post("/sign-up", func(c *fiber.Ctx) error {
		payload := struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		return c.JSON(payload)
	})
}
