package handler

import (
	"hxhs-playground/component"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/google/uuid"
)

var handlerHXGet = adaptor.HTTPHandler(templ.Handler(component.GenerateEmptyDivGET(uuid.New().String())))
var handlerHXPost = adaptor.HTTPHandler(templ.Handler(component.GenerateEmptyDivPOST(uuid.New().String())))
var handlerHXGetTarget = adaptor.HTTPHandler(templ.Handler(component.GenerateEmptyToReplaceDivWithComponenet(uuid.New().String())))
var handlerHXGetSwapAfternend = adaptor.HTTPHandler(templ.Handler(component.GenerateEmptyDivGETSwapAfterend(uuid.New().String())))

func InitHXHandler(api *fiber.App) error {

	hxAPI := api.Group("/api")
	hxAPI.Post("/postButtonComponentHX", handlerHXPost)
	hxAPI.Get("/getButtonComponentHX", handlerHXGet)
	hxAPI.Get("/getButtonComponentHXTarget", handlerHXGetTarget)
	hxAPI.Get("/getButtonComponentHXSwapAfterend", handlerHXGetSwapAfternend)
	hxAPI.Post("/postButtonComponentHXInclude", createTemplHandler(func(c *fiber.Ctx) templ.Component {
		body := string(c.Body())
		return component.GenerateEmptyDivGETWithParam(uuid.New().String(), body)
	}))
	return nil
}
func createTemplHandler(fn func(*fiber.Ctx) templ.Component) fiber.Handler {
	return func(c *fiber.Ctx) error {
		component := fn(c)
		return component.Render(c.Context(), c.Response().BodyWriter())
	}
}
