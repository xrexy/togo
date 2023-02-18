package tasks

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/xrexy/togo/pkg/database"
	"github.com/xrexy/togo/pkg/validation"
	"gorm.io/gorm"
)

type UpdateTaskRequest struct {
	Creator string `json:"creator,omitempty" example:"uuidv4" validate:"omitempty,uuid4"`
	Title   string `json:"title,omitempty" example:"My updated title" validate:"omitempty,min=3,max=64"`
	Content string `json:"content,omitempty" example:"My updated content" validate:"omitempty,min=16,max=1024"`
}

func (c *TaskController) UpdateTask(ctx *fiber.Ctx) error {
	var request UpdateTaskRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(database.MessageStruct{
			ErrorMessage: "Invalid body format",
			CreatedAt:    time.Now().Unix(),
		})
	}

	if request.Title == "" && request.Content == "" && request.Creator == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(database.MessageStruct{
			ErrorMessage: "At least one of the fields must be set",
			CreatedAt:    time.Now().Unix(),
		})
	}

	if errors, _ := validation.ValidateStruct(request); errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	user := ctx.Locals("user").(*database.User)
	taskUUID := ctx.Params("uuid")

	task := database.Task{}
	tx := database.PostgesClient.Where("uuid = ?", taskUUID).First(&task)
	if tx.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(database.MessageStruct{
			ErrorMessage: "Internal server error while getting task",
			CreatedAt:    time.Now().Unix(),
		})
	}

	if task.UUID == "" {
		return ctx.Status(fiber.StatusNotFound).JSON(database.MessageStruct{
			ErrorMessage: "Task not found",
			CreatedAt:    time.Now().Unix(),
		})
	}

	if user.Role != database.RoleAdmin && task.UserID != user.UUID {
		return ctx.Status(fiber.StatusUnauthorized).JSON(database.MessageStruct{
			ErrorMessage: "You are not authorized to update this resource",
			CreatedAt:    time.Now().Unix(),
		})
	}

	if request.Title != "" {
		task.Title = request.Title
	}

	if request.Content != "" {
		task.Content = request.Content
	}

	if request.Creator != "" {
		requestedUser := database.User{}
		tx = database.PostgesClient.Where("uuid = ?", request.Creator).First(&requestedUser)
		if tx.Error != nil {
			if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
				return ctx.Status(fiber.StatusNotFound).JSON(database.MessageStruct{
					ErrorMessage: "User not found",
					CreatedAt:    time.Now().Unix(),
				})
			}

			return ctx.Status(fiber.StatusInternalServerError).JSON(database.MessageStruct{
				ErrorMessage: "Internal server error while getting user to assign task to",
				CreatedAt:    time.Now().Unix(),
			})
		}

		task.UserID = request.Creator
	}

	task.UpdatedAt = time.Now().Unix()

	tx = database.PostgesClient.Save(&task)
	if tx.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(database.MessageStruct{
			ErrorMessage: "Internal server error while updating task",
			CreatedAt:    time.Now().Unix(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(task)
}
