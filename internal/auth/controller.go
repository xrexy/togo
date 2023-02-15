package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xrexy/togo/config"
)

type AuthController struct {
	jwt []byte
}

func NewAuthController(config config.EnvVars) *AuthController {
	return &AuthController{
		jwt: []byte(config.JWT_KEY),
	}
}

func (c *AuthController) CreateGroup(base fiber.Router, config config.EnvVars) {
	auth := base.Group("/auth")
	auth.Post("/signin", c.Signin)
	auth.Post("/signup", c.Signup)
}
