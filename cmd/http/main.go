package main

import (
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

func buildServer(env config.EnvVars) *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "${time} - ${status} ${method} ${path}\n[${pid}] ${locals:requestid} ${latency}\n${ip}:${port} | ${ua}\n\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Europe/Sofia",
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/docs/*", swagger.HandlerDefault)

	// create auth domain
	authController := auth.NewAuthController(env)
	auth.CreateAuthGroup(app, authController, env)

	return app
}
