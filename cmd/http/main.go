package main

import (
	"html/template"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"

	"github.com/xrexy/togo/config"
	_ "github.com/xrexy/togo/docs"
	"github.com/xrexy/togo/internal/auth"
	"github.com/xrexy/togo/pkg/database"
	"github.com/xrexy/togo/pkg/shutdown"
)

// @title To-Go API
// @version 0.1-alpha
// @description To-Go is a simple API for managing tasks. It is built with Fiber and GORM.
// @host localhost:8080
// @BasePath /
func main() {
	env, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	cleanup, err := run(env)

	defer cleanup()

	if err != nil {
		panic(err)
	}

	shutdown.Gracefully()
}

func run(env config.EnvVars) (func(), error) {
	err := database.StartPostgresDB(&env)
	if err != nil {
		panic(err)
	}

	app := buildServer(env)

	// start the server
	go func() {
		app.Listen("0.0.0.0:" + env.PORT)
	}()

	// return a cleanup function to be called on shutdown
	return func() {
		app.Shutdown()
	}, nil
}

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

func buildServer(env config.EnvVars) *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "${time} - ${status} ${method} ${path}\n[${pid}] ${locals:requestid} ${latency}\n${ip}:${port} | ${ua}\n\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Europe/Sofia",
	}))

	app.Get("/docs/*", swagger.New(swagger.Config{
		Title:  "To-Go API Docs",
		Layout: "StandaloneLayout",
		Plugins: []template.JS{
			template.JS("SwaggerUIBundle.plugins.DownloadUrl"),
		},
		Presets: []template.JS{
			template.JS("SwaggerUIBundle.presets.apis"),
			template.JS("SwaggerUIStandalonePreset"),
		},
		DeepLinking:              true,
		DefaultModelsExpandDepth: 1,
		DefaultModelExpandDepth:  1,
		DefaultModelRendering:    "example",
		DocExpansion:             "list",
		SyntaxHighlight: &swagger.SyntaxHighlightConfig{
			Activate: true,
			Theme:    "nord",
		},
		ShowMutatedRequest: true,
	}))

	v1 := app.Group("/api/v1")

	// health check
	// @Summary Health check
	// @Description Checks if the server is up and running
	// @Tags health
	// @Accept json
	// @Produce json
	// @Success 200 {object} HealthResponse
	// @Router /health [get]
	v1.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(HealthResponse{
			Status:    "OK",
			Timestamp: time.Now().Unix(),
		})
	})

	// create auth domain
	authController := auth.NewAuthController(env)
	auth.CreateAuthGroup(v1, authController, env)

	return app
}
