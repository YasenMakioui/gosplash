package main

import (
	"github.com/YasenMakioui/gosplash/src/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("static", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.SetupRoutes(app)

	app.Listen(":8080")
}
