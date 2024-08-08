package main

import (
	"hxhs-playground/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Static("/", "./assets")
	app.Use(logger.New())
	handler.InitApp(app)
	app.Listen(":3003")
}
