package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/xrexy/togo/pkg/authentication"
	"github.com/xrexy/togo/pkg/database"
)

// DeserializeUser is a middleware to deserialize user from token
// and set it to the context
//
// Before using it make sure the handler is using the middleware
// To get the user from the context, use:
// user := c.Locals("user").(*database.User)
func DeserializeUser(c *fiber.Ctx) error {
	createdAt := time.Now().Unix()

	authentication := authentication.New()
	tokenString := authentication.GetTokenString(c)

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(database.MessageStruct{
			Message:   "Not logged in",
			CreatedAt: createdAt,
		})
	}

	claims, err := authentication.VerifyJWT(tokenString)
	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return c.Status(fiber.StatusUnauthorized).JSON(database.MessageStruct{
				Message:   "Token is expired",
				CreatedAt: createdAt,
			})
		}

		return c.Status(fiber.StatusUnauthorized).JSON(database.MessageStruct{
			Message:   "Invalid Token",
			CreatedAt: createdAt,
		})
	}

	var user database.User
	database.PostgesClient.Model(&database.User{}).Preload(
		"Tasks",
	).Where("uuid = ?", fmt.Sprint(claims["sub"])).First(&user)

	if user.UUID != claims["sub"] {
		return c.Status(fiber.StatusForbidden).JSON(database.MessageStruct{
			Message:   "The user belonging to this token no logger exists",
			CreatedAt: createdAt,
		})
	}

	c.Locals("user", &user)

	return c.Next()
}
