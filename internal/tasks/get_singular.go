package tasks

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/xrexy/togo/pkg/database"
)

type GetSingularTaskRequest struct {
	UUID string `json:"uuid"`
}

// GetTask
// @Summary Get a task
// @Description Returns a task by uuid
// @Tags tasks
// @Accept json
// @Produce json
// @Param uuid path string true "Task UUID"
// @Success 200 {object} database.Task
// @Failure 500 {object} database.MessageStruct "Internal server error while getting task"
// @Router /api/v1/task/{uuid} [get]
func (c *TaskController) GetTask(ctx *fiber.Ctx) error {
	var request GetSingularTaskRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(database.MessageStruct{
			ErrorCode:    fiber.StatusBadRequest,
			ErrorMessage: "Invalid body format",
			CreatedAt:    time.Now(),
		})
	}

	task := database.Task{}
	tx := database.PostgesClient.Where("uuid = ?", request.UUID).First(&task)
	if tx.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(database.MessageStruct{
			ErrorCode:    fiber.StatusInternalServerError,
			ErrorMessage: "Error while getting task",
			CreatedAt:    time.Now(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(task)
}
