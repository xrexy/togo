package auth

import (
	"github.com/xrexy/togo/config"
)

type AuthController struct {
	jwt   []byte
	users map[string]string
}

type AuthOKResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type AuthInternalServerErrorResponse struct {
	Message string `json:"message"`
}

type AuthUnauthorizedResponse struct {
	Message string `json:"message"`
}

var users = map[string]string{
	"user1@gmail.com": "userpass",
	"user2@gmail.com": "userpass",
}

func NewAuthController(config config.EnvVars) *AuthController {
	return &AuthController{
		jwt:   []byte(config.JWT_KEY),
		users: users,
	}
}
