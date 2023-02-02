package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xrexy/togo/pkg/database"
	"github.com/xrexy/togo/pkg/database/models"
)

func (a *AuthController) Signup(c *fiber.Ctx) error {
	var creds models.Credentials
	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	client := database.PostgesClient
	result := client.Create(&models.User{
		Email:    creds.Email,
		Password: creds.Password,
	})

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create user",
		})
	}

	return c.SendString("User created")
}
