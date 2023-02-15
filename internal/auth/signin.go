package auth

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/xrexy/togo/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type SignInOKResponse struct {
	Token string `json:"token"`
}

// Signin
// @Summary Sign in
// @Description Authenticates a user and returns a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body Credentials true "Credentials"
// @Success 200 {object} SignInOKResponse
// @Failure 400 {object} AuthMessageStruct "Invalid credentials format"
// @Failure 401 {object} AuthMessageStruct "Unauthorized"
// @Failure 500 {object} AuthMessageStruct "Internal server error while signing token"
// @Router /signin [post]
func (ac *AuthController) Signin(ctx *fiber.Ctx) error {
	var creds Credentials
	if err := ctx.BodyParser(&creds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(AuthMessageStruct{
			ErrorCode:    fiber.StatusBadRequest,
			ErrorMessage: "Invalid credentials format",
			CreatedAt:    time.Now(),
		})
	}

	user := User{}
	database.PostgesClient.Find(&user, "email = ?", creds.Email)
	if user.Email == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(AuthMessageStruct{
			ErrorCode:    fiber.StatusUnauthorized,
			ErrorMessage: "Unauthorized",
			CreatedAt:    time.Now(),
		})
	}

	log.Default().Print(user.Password, creds.Password)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(AuthMessageStruct{
			ErrorCode:    fiber.StatusUnauthorized,
			ErrorMessage: "Unauthorized",
			CreatedAt:    time.Now(),
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
		return ctx.Status(fiber.StatusInternalServerError).JSON(AuthMessageStruct{
			ErrorCode:    fiber.StatusInternalServerError,
			ErrorMessage: "Internal server error while signing token",
			CreatedAt:    time.Now(),
		})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expTime,
	})

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": tokenString,
	})
}
