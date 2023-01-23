package auth

import (
	"github.com/gofiber/fiber/v2"
)

func (a *AuthController) Signup(c *fiber.Ctx) error {
	return c.SendString("Signup")
}
