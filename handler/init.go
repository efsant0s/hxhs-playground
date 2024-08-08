package handler

import "github.com/gofiber/fiber/v2"

func InitApp(app *fiber.App) {
	app.Get("/test", HandleTestPage)
	app.Get("/", DemoPageView)
	InitHXHandler(app)

}
