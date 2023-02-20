package tasks

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xrexy/togo/pkg/database"
)

type InfoResponse struct {
	PlanMaxTasks map[database.Plan]int `json:"plan_tasks_limits"`
}

// Info
// @Summary Get info
// @Description Returns information about the tasks
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {object} InfoResponse
// @Router /api/v1/tasks/info [get]
func (a *TaskController) Info(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(InfoResponse{
		PlanMaxTasks: database.PlanMaxTasks,
	})
}
