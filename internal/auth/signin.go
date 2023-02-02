package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/xrexy/togo/pkg/database/models"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// Signin godoc
// @Summary Sign in
// @Description Authenticates a user and returns a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.Credentials true "Credentials"
// @Success 200 {object} AuthOKResponse
// @Failure 401 {object} AuthUnauthorizedResponse
// @Failure 500 {object} AuthInternalServerErrorResponse
// @Router /signin [post]
func (ac *AuthController) Signin(ctx *fiber.Ctx) error {
	var creds models.Credentials
	if err := ctx.BodyParser(&creds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	expectedPassword, ok := ac.users[creds.Email]
	if !ok || expectedPassword != creds.Password {
		return ctx.Status(fiber.StatusUnauthorized).JSON(AuthUnauthorizedResponse{
			Message: "Invalid credentials",
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

	tokenString, err := token.SignedString(ac.jwt)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(AuthInternalServerErrorResponse{
			Message: "Error signing token",
		})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expTime,
	})

	return ctx.Status(fiber.StatusOK).JSON(AuthOKResponse{
		Message: "success",
		Token:   tokenString,
	})
}
