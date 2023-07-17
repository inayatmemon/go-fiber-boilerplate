package routes

import (
	"github.com/gofiber/fiber/v2"
)

var AdminGroup fiber.Router
var CustomerGroup fiber.Router

func RouteGroups(app *fiber.App) {

	AdminGroup = app.Group("/admin")
	CustomerGroup = app.Group("/customer")
	// AdminGroup = app.Group("/admin",middlerwareFunc)

}
