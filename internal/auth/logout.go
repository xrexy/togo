package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func (ac *AuthController) Logout(ctx *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)
	ctx.Cookie(&fiber.Cookie{
		Name:    ac.config.JWT_COOKIE_KEY,
		Value:   "",
		Expires: expired,
	})
	return ctx.SendStatus(fiber.StatusOK)
}
