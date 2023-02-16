package tasks

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/xrexy/togo/pkg/database"
	"github.com/xrexy/togo/pkg/validation"
)

type UpdateTaskRequest struct {
	UUID string `json:"uuid" validate:"required"`
}

func (c *TaskController) UpdateTask(ctx *fiber.Ctx) error {
	var request UpdateTaskRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(database.MessageStruct{
			ErrorCode:    fiber.StatusBadRequest,
			ErrorMessage: "Invalid body format",
			CreatedAt:    time.Now(),
		})
	}

	errors, _ := validation.ValidateStruct(request)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return ctx.Status(fiber.StatusOK).JSON("hey")
}
