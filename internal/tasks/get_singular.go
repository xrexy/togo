package tasks

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/xrexy/togo/pkg/database"
	"github.com/xrexy/togo/pkg/validation"
)

type GetSingularTaskRequest struct {
	UUID string `json:"uuid" validate:"required,uuid4"`
}

// GetTask
// @Summary Get a task
// @Description Returns a task by uuid. UNFINISHED (No auth atm so can't verify if the user can access this task)
// @Tags tasks
// @Accept json
// @Produce json
// @Param uuid path string true "Task UUID"
// @Success 200 {object} database.Task\
// @Failure 400 {array} validation.ErrorResponse "Validation error"
// @Failure 500 {object} database.MessageStruct "Internal server error while getting task"
// @Router /api/v1/task/u/{uuid} [get]
func (c *TaskController) GetTask(ctx *fiber.Ctx) error {
	var request GetSingularTaskRequest
	request.UUID = ctx.Params("uuid")

	errors, _ := validation.ValidateStruct(request)
	if len(errors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	task := database.Task{}
	tx := database.PostgesClient.Where("uuid = ?", request.UUID).First(&task)
	if tx.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(database.MessageStruct{
			ErrorMessage: "Internal server error while getting task",
			CreatedAt:    time.Now().Unix(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(task)
}
