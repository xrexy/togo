package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xrexy/togo/config"
)

func CreateAuthGroup(router fiber.Router, authController *AuthController, config config.EnvVars) {
	auth := router.Group("/auth")
	auth.Post("/signin", authController.Signin)
	auth.Post("/signup", authController.Signup)
}
