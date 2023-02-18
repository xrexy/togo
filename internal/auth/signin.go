package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/xrexy/togo/pkg/authentication"
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
// @Param credentials body database.UserCredentials true "Credentials"
// @Success 200 {object} SignInOKResponse
// @Failure 400 {object} database.MessageStruct "Invalid credentials format"
// @Failure 401 {object} database.MessageStruct "Unauthorized"
// @Failure 500 {object} database.MessageStruct "Internal server error while signing token"
// @Router /api/v1/auth/signin [post]
func (ac *AuthController) Signin(ctx *fiber.Ctx) error {
	var creds database.UserCredentials
	if err := ctx.BodyParser(&creds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(database.MessageStruct{
			ErrorMessage: "Invalid credentials format",
			CreatedAt:    time.Now().Unix(),
		})
	}

	user := database.User{}
	database.PostgesClient.Find(&user, "email = ?", creds.Email)
	if user.Email == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(database.MessageStruct{
			ErrorMessage: "Unauthorized",
			CreatedAt:    time.Now().Unix(),
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(database.MessageStruct{
			ErrorMessage: "Unauthorized",
			CreatedAt:    time.Now().Unix(),
		})
	}

	token, exp, err := authentication.New().CreateJWT(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(database.MessageStruct{
			ErrorMessage: "Internal server error while signing token",
			CreatedAt:    time.Now().Unix(),
		})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    ac.config.JWT_COOKIE_KEY,
		Value:   token,
		Expires: exp,
	})

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
		"exp":   exp.Unix(),
	})
}
