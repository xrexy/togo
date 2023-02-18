package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/xrexy/togo/pkg/authentication"
	"github.com/xrexy/togo/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

type SignUpOKResponse struct {
	Token string `json:"token" example:"uuidv4"`
}

// Signup
// @Summary Sign up
// @Description Creates a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body database.UserCredentials true "Credentials"
// @Success 200 {object} SignUpOKResponse
// @Failure 400 {object} database.MessageStruct "Invalid credentials format"
// @Failure 400 {object} database.MessageStruct "User already exists"
// @Router /api/v1/auth/signup [post]
func (ac *AuthController) Signup(ctx *fiber.Ctx) error {
	var creds database.UserCredentials
	if err := ctx.BodyParser(&creds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(database.MessageStruct{
			ErrorMessage: "Invalid credentials format",
			CreatedAt:    time.Now().Unix(),
		})
	}

	hPass, err := hashPassword(creds.Password)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(database.MessageStruct{
			ErrorMessage: "We couldn't create your account. Please try again later.",
			CreatedAt:    time.Now().Unix(),
		})
	}

	user := database.User{
		UUID:      uuid.NewV4().String(),
		Email:     creds.Email,
		Password:  hPass,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		Role:      database.RoleUser,
		Plan:      database.PlanFree,
		TaskCount: 0,
		Tasks:     make([]database.Task, 0),
	}

	result := database.PostgesClient.Create(&user)
	if result.Error != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(database.MessageStruct{
			ErrorMessage: "User already exists",
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

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
