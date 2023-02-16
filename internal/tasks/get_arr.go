package tasks

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/xrexy/togo/pkg/database"
)

// GetTasks
// @Summary Get all tasks
// @Description Returns all tasks. UNFINISHED (No auth atm so can't get the uuid of user) - will return only tasks of user in the future
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {array} database.Task
// @Failure 500 {object} database.MessageStruct "Internal server error while getting tasks"
// @Router /api/v1/task [get]
func (c *TaskController) GetTasks(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	tasks := []database.Task{}
	tx := database.PostgesClient.Find(&tasks)
	if tx.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(database.MessageStruct{
			ErrorCode:    fiber.StatusInternalServerError,
			ErrorMessage: "Error while getting tasks",
			CreatedAt:    time.Now(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(tasks)
}
