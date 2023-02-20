package internal

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/xrexy/togo/config"
	"github.com/xrexy/togo/internal/auth"
	"github.com/xrexy/togo/internal/tasks"
)

type MessageStruct struct {
	ErrorCode    int       `json:"error_code" example:"400"`
	ErrorMessage string    `json:"error_message" example:"User already exists"`
	CreatedAt    time.Time `json:"created_at" example:"1620000000"`
}

func RegisterRoutes(base fiber.Router, env config.EnvVars) {
	authController := auth.NewAuthController(env)
	authController.CreateGroup(base, env)

	tasksController := tasks.NewTaskController()
	tasksController.CreateGroup(base, env)
}
