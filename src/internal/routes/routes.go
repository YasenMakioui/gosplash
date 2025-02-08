package routes

import (
	"fmt"
	"strings"

	"github.com/YasenMakioui/gosplash/src/internal/handlers"
	"github.com/YasenMakioui/gosplash/src/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		token := c.Get("Authorization", "")

		if token == "" {
			c.SendString("Missing auth")
		}

		fmt.Println(token)

		if !strings.HasPrefix(token, "Bearer ") {
			return c.SendString("Token is not Bearer compliant")
		}

		token = token[len("Bearer "):]

		if err := services.VerifyToken(token); err != nil {
			return c.SendString("Invalid token")
		}

		return c.Render("index", fiber.Map{
			"Title": "Rendered!",
		})
	})

	app.Get("/auth/login", handlers.Login)
}
