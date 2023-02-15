package tasks

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xrexy/togo/config"
)

type TaskController struct {
}

func NewTaskController() *TaskController {
	return &TaskController{}
}

func (c *TaskController) CreateGroup(base fiber.Router, config config.EnvVars) {
	tasks := base.Group("/task")
	// -- GET
	tasks.Get("/", c.GetTasks)
	tasks.Get("/:uuid", c.GetTask)

	// -- POST
	tasks.Post("/", c.CreateTask)

	// -- PUT
	// tasks.Put("/:uuid", c.UpdateTask)

	// -- DELETE
	// tasks.Delete("/:uuid", c.DeleteTask)
}
