package tasks

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/xrexy/togo/pkg/authentication"
	"github.com/xrexy/togo/pkg/database"
)

// GetTasks
// @Summary Get all user has access to
// @Description Returns all tasks the user has created. If the user is an admin, all tasks will be returned
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {array} database.Task
// @Failure 500 {object} database.MessageStruct "Internal server error while getting tasks"
// @Router /api/v1/task [get]
func (c *TaskController) GetTasks(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	auth := authentication.New()
	token := auth.GetTokenString(ctx)
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(database.MessageStruct{
			ErrorMessage: "Unauthorized",
			CreatedAt:    time.Now().Unix(),
		})
	}

	claims, err := auth.VerifyJWT(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(database.MessageStruct{
			ErrorMessage: "Unauthorized",
			CreatedAt:    time.Now().Unix(),
		})
	}

	fmt.Println(claims)

	uuid := fmt.Sprint(claims["sub"])

	var user database.User
	tx := database.PostgesClient.Where("uuid = ?", uuid).First(&user)
	if tx.Error != nil || user.UUID == "" {
		return ctx.Status(fiber.StatusInternalServerError).JSON(database.MessageStruct{
			ErrorMessage: "Error while getting user tasks",
			CreatedAt:    time.Now().Unix(),
		})
	}

	var tasks []database.Task

	// If user is admin, return all tasks
	// Else return only tasks that the user has created
	if user.Role == string(database.RoleAdmin) {
		tx = database.PostgesClient.Find(&tasks)
	} else {
		tx = database.PostgesClient.Where("user_id = ?", uuid).Find(&tasks)
	}

	if tx.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(database.MessageStruct{
			ErrorMessage: "Error while getting user tasks",
			CreatedAt:    time.Now().Unix(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(tasks)
}
