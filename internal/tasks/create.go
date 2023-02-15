package tasks

import (
	"time"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/xrexy/togo/pkg/database"
)

type CreateTaskRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Creator string `json:"creator"`
}

// CreateTask
// @Summary Create a task
// @Description Creates a task with title and content
// @Tags tasks
// @Accept json
// @Produce json
// @Param title body string true "Task title"
// @Param content body string true "Task content"
// @Success 200 {object} database.Task
// @Failure 500 {object} database.MessageStruct "Internal server error while creating task"
// @Router /api/v1/task [post]
func (c *TaskController) CreateTask(ctx *fiber.Ctx) error {
	var request CreateTaskRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(database.MessageStruct{
			ErrorCode:    fiber.StatusBadRequest,
			ErrorMessage: "Invalid body format",
			CreatedAt:    time.Now(),
		})
	}

	// TODO add authentication or auth layer/middleware

	task := database.Task{
		UUID:      uuid.NewV4().String(),
		Title:     request.Title,
		Content:   request.Content,
		UserUUID:  request.Creator,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tx := database.PostgesClient.Create(&task)
	if tx.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(database.MessageStruct{
			ErrorCode:    fiber.StatusInternalServerError,
			ErrorMessage: "Error while creating task",
			CreatedAt:    time.Now(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(task)
}
