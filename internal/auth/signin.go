package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = map[string]string{
	"user1@gmail.com": "userpass",
	"user2@gmail.com": "userpass",
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (a *AuthController) Signin(c *fiber.Ctx) error {
	var creds Credentials
	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	expectedPassword, ok := users[creds.Email]
	if !ok || expectedPassword != creds.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	expTime := time.Now().Add(time.Hour * 24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Email: creds.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "localhost",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	})

	tokenString, err := token.SignedString(a.jwt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not sign the token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expTime,
	})

	return c.JSON(fiber.Map{
		"message": "success",
		"token":   tokenString,
	})
}
