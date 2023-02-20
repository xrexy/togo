package tasks

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/xrexy/togo/pkg/database"
	"gorm.io/gorm"
)

func (c *TaskController) DeleteTask(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*database.User)
	taskUUID := ctx.Params("uuid")

	tx := database.PostgesClient.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(database.MessageStruct{
			Message:   "Something went wrong while creating transaction",
			CreatedAt: time.Now().Unix(),
		})
	}

	task := database.Task{}
	if _tx := tx.Where("uuid = ?", taskUUID).First(&task); _tx.Error != nil {
		tx.Rollback()
		if errors.Is(_tx.Error, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(database.MessageStruct{
				Message:   "Task not found",
				CreatedAt: time.Now().Unix(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(database.MessageStruct{
			Message:   "Internal server error while getting task",
			CreatedAt: time.Now().Unix(),
		})
	}

	if user.Role != database.RoleAdmin && task.UserID != user.UUID {
		return ctx.Status(fiber.StatusUnauthorized).JSON(database.MessageStruct{
			Message:   "You are not authorized to delete this resource",
			CreatedAt: time.Now().Unix(),
		})
	}

	if err := tx.Delete(&task).Error; err != nil {
		tx.Rollback()
		return ctx.Status(fiber.StatusBadRequest).JSON(database.MessageStruct{
			Message:   "The requested task failed to delete",
			CreatedAt: time.Now().Unix(),
		})
	}

	var owner database.User
	if _tx := tx.Where("uuid = ?", task.UserID).First(&owner); _tx.Error != nil {
		tx.Rollback()
		if errors.Is(_tx.Error, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(database.MessageStruct{
				Message:   "User not found",
				CreatedAt: time.Now().Unix(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(database.MessageStruct{
			Message:   "Internal server error while getting user to assign task to",
			CreatedAt: time.Now().Unix(),
		})
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return ctx.Status(fiber.StatusInternalServerError).JSON(database.MessageStruct{
			Message:   "Internal server error while committing transaction",
			CreatedAt: time.Now().Unix(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(database.MessageStruct{
		Message:   "Task deleted successfully",
		CreatedAt: time.Now().Unix(),
	})
}
