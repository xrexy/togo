package auth

import (
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
