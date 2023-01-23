package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

	"github.com/xrexy/togo/config"
	_ "github.com/xrexy/togo/docs"
	"github.com/xrexy/togo/internal/auth"
	"github.com/xrexy/togo/pkg/shutdown"
)

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
	app := buildServer(env)

	// start the server
	go func() {
		app.Listen(":" + env.PORT)
	}()

	// return a function to stop the server and db
	return func() {
		app.Shutdown()
	}, nil
}

func buildServer(env config.EnvVars) *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/docs/*", swagger.HandlerDefault) // default

	// create auth domain
	authController := auth.NewAuthController(env)
	auth.CreateAuthGroup(app, authController, env)

	return app
}
