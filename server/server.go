package server

import (
	"fmt"
	"log"
	"strconv"

	"go-fiber-boilerplate/config"
	"go-fiber-boilerplate/middleware"
	"go-fiber-boilerplate/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func StartServer() {
	// Initialize standard Go html template engine
	engine := html.New("./mvc/views", ".html")
	// engine.Reload(true)
	// engine.Debug(true)
	// engine.Layout("embed")
	// engine.Delims("{{", "}}")
	app := fiber.New(
		fiber.Config{
			Views:       engine,
			ViewsLayout: "./mvc/views/*.html",
		},
	)

	app.Static("/static", "./static")

	// Logger middleware
	// Or extend your config for customization
	// Logging remote IP and Port
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	routes.RouteGroups(app)
	routes.InitializeAllRoutes(app)
	middleware.InitializeSession()
	// app.Use(&middleware.Session)

	log.Println("configuration loaded succesfully...")
	log.Println("server started successfully...")

	port := fmt.Sprintf(":%s", strconv.Itoa(config.AppConfig.Server.Port))
	app.Listen(port)
}
