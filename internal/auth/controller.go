package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xrexy/togo/config"
	"github.com/xrexy/togo/middleware"
)

type AuthController struct {
	config config.EnvVars
}

func NewAuthController(config config.EnvVars) *AuthController {
	return &AuthController{config: config}
}

func (c *AuthController) CreateGroup(base fiber.Router, config config.EnvVars) {
	auth := base.Group("/auth")
	auth.Post("/signin", c.Signin)
	auth.Post("/signup", c.Signup)

	auth.Get("/logout", c.Logout)
	auth.Get("/me", middleware.DeserializeUser, c.Me)
}
