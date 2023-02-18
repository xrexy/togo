package tasks

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xrexy/togo/config"
	"github.com/xrexy/togo/middleware"
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
	tasks.Get("/u/:uuid", c.GetTask)
	tasks.Get("/info", c.Info)

	// -- POST
	tasks.Post("/", middleware.DeserializeUser, c.CreateTask)

	// -- PUT
	tasks.Put("/u/:uuid", middleware.DeserializeUser, c.UpdateTask)

	// -- DELETE
	// tasks.Delete("/:uuid", c.DeleteTask)
}
