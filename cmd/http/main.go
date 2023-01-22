package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "github.com/xrexy/togo/docs"
)

// @title ToGO API
// @version 1.0
// @description This is a sample server ToGO server.
// @BasePath /
func main() {
	app := fiber.New()

	app.Get("/docs/*", swagger.HandlerDefault) // default

	app.Listen(":8080")
}
