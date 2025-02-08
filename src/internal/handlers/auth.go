package handlers

import (
	"fmt"

	"github.com/YasenMakioui/gosplash/src/internal/services"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	u := new(services.User)

	if err := c.BodyParser(u); err != nil {
		return err
	}

	if u.Username != "test" || u.Password != "123456" {
		return fmt.Errorf("invalid credentials")
	}

	tokenString, err := services.NewToken(u.Username)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not authenticate user"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"JWT": tokenString})
}
