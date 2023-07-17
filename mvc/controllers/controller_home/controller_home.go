package controller_home

import (
	"go-fiber-boilerplate/mvc/models/models_home"

	"github.com/gofiber/fiber/v2"
)

func IndexController(c *fiber.Ctx) error {
	data := models_home.IndexTemplateModel{
		PageTitle: "This is page",
		BodyTitle: "BodyTitle",
	}
	return c.Render("index", data, "layouts/main")
	// return c.SendString("")
}
