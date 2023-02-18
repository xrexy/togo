package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xrexy/togo/pkg/database"
)

func (ac *AuthController) Me(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*database.User)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": user,
	})
}
