package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
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
// @Param credentials body Credentials true "Credentials"
// @Success 200 {object} SignUpOKResponse
// @Failure 400 {object} AuthMessageStruct "Invalid credentials format"
// @Failure 400 {object} AuthMessageStruct "User already exists"
// @Router /signup [post]
func (a *AuthController) Signup(c *fiber.Ctx) error {
	var creds Credentials
	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(AuthMessageStruct{
			ErrorCode:    fiber.StatusBadRequest,
			ErrorMessage: "Invalid credentials format",
			CreatedAt:    time.Now(),
		})
	}

	hPass, err := hashPassword(creds.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(AuthMessageStruct{
			ErrorCode:    fiber.StatusInternalServerError,
			ErrorMessage: "We couldn't create your account. Please try again later.",
			CreatedAt:    time.Now(),
		})
	}

	user := User{
		UUID:     uuid.NewV4().String(),
		Email:    creds.Email,
		Password: hPass,
	}

	result := database.PostgesClient.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusConflict).JSON(AuthMessageStruct{
			ErrorCode:    fiber.StatusConflict,
			ErrorMessage: "User already exists",
			CreatedAt:    time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(SignUpOKResponse{
		Token: user.UUID,
	})
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
