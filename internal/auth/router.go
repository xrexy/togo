package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xrexy/togo/config"
)

func CreateAuthGroup(app *fiber.App, authController *AuthController, config config.EnvVars) {
	auth := app.Group("/auth")
	auth.Post("/signin", authController.Signin)
	auth.Post("/signup", authController.Signup)
}
