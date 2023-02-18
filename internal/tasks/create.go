package tasks

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/xrexy/togo/pkg/database"
	"github.com/xrexy/togo/pkg/validation"
)

type CreateTaskRequest struct {
	Title   string `json:"title" validate:"required,min=3,max=64"`
	Content string `json:"content" validate:"required,min=16,max=1024"`
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

	user := ctx.Locals("user").(*database.User)
	if user.Role != database.RoleAdmin && user.TaskCount >= database.PlanMaxTasks[user.Plan] {
		return ctx.Status(fiber.StatusForbidden).JSON(database.MessageStruct{
			ErrorMessage: fmt.Sprintf("You have reached the maximum number of tasks for your plan (%d)", database.PlanMaxTasks[user.Plan]),
			CreatedAt:    time.Now().Unix(),
		})
	}

	response, _ := validation.ValidateStruct(request)
	if len(response) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	task := database.Task{
		UUID:      uuid.NewV4().String(),
		Title:     request.Title,
		Content:   request.Content,
		UserID:    user.UUID,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	tx := database.PostgesClient.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(database.MessageStruct{
			ErrorMessage: "Something went wrong while creating transaction",
			CreatedAt:    time.Now().Unix(),
		})
	}

	if err := tx.Create(&task).Error; err != nil {
		tx.Rollback()
		return ctx.Status(fiber.StatusNotModified).JSON(database.MessageStruct{
			ErrorMessage: "[Transaction] Couldn't create task",
			CreatedAt:    time.Now().Unix(),
		})
	}

	if err := tx.Model(&user).Update("task_count", user.TaskCount+1).Error; err != nil {
		tx.Rollback()
		return ctx.Status(fiber.StatusNotModified).JSON(database.MessageStruct{
			ErrorMessage: "[Transaction] Couldn't update user task count",
			CreatedAt:    time.Now().Unix(),
		})
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return ctx.Status(fiber.StatusInternalServerError).JSON(database.MessageStruct{
			ErrorMessage: "Something went wrong while committing transaction",
			CreatedAt:    time.Now().Unix(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(task)
}
