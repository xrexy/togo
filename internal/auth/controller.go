package auth

import (
	"time"

	"github.com/xrexy/togo/config"
)

type AuthMessageStruct struct {
	ErrorCode    int       `json:"error_code" example:"400"`
	ErrorMessage string    `json:"error_message" example:"User already exists"`
	CreatedAt    time.Time `json:"created_at" example:"1620000000"`
}

type Credentials struct {
	Email    string `json:"email" example:"example@togo.dev"`
	Password string `json:"password" example:"my_super_secret_password"`
}

type User struct {
	UUID     string `json:"uuid" example:"uuidv4" gorm:"primaryKey"`
	Email    string `json:"email" example:"example@togo.dev" gorm:"uniqueIndex"`
	Password string `json:"password" example:"my_super_secret_password" gorm:"not null"`
}

type AuthController struct {
	jwt []byte
}

func NewAuthController(config config.EnvVars) *AuthController {
	return &AuthController{
		jwt: []byte(config.JWT_KEY),
	}
}
