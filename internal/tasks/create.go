package tasks

import (
	"time"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/xrexy/togo/pkg/database"
	"github.com/xrexy/togo/pkg/validation"
)

type CreateTaskRequest struct {
	Title   string `json:"title" validate:"required,min=3,max=64"`
	Content string `json:"content" validate:"required,min=3,max=1024"`
	Creator string `json:"creator" validate:"required,uuid4"`
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
// @Failure 400 {object} database.MessageStruct "Invalid body format"
// @Failure 400 {object} validation.ErrorResponse "Validation failed"
// @Failure 500 {object} database.MessageStruct "Internal server error while creating task"
// @Router /api/v1/task [post]
func (c *TaskController) CreateTask(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	var request CreateTaskRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(database.MessageStruct{
			ErrorMessage: "Invalid body format",
			CreatedAt:    time.Now().Unix(),
		})
	}

	response, _ := validation.ValidateStruct(request)
	if len(response) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	// TODO add authentication or auth layer/middleware

	task := database.Task{
		UUID:      uuid.NewV4().String(),
		Title:     request.Title,
		Content:   request.Content,
		UserID:    request.Creator,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	tx := database.PostgesClient.Create(&task)
	if tx.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(database.MessageStruct{
			ErrorMessage: "Error while creating task",
			CreatedAt:    time.Now().Unix(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(task)
}
