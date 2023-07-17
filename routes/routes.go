package routes

import (
	"go-fiber-boilerplate/middleware"
	"go-fiber-boilerplate/mvc/controllers/controller_home"
	"log"

	"github.com/gofiber/fiber/v2"
)

func InitializeAllRoutes(app *fiber.App) {
	DefaultRoutes(app)
	AdminRoutes()
	CustomerRoutes()
}

func DefaultRoutes(app *fiber.App) {
	//default or index route which doesn't need any groups
	app.Get("/", controller_home.IndexController)
}

func AdminRoutes() {
	AdminGroup.Get("/home", middleware.UserLoggedIn(), func(c *fiber.Ctx) error {
		return c.SendString("HelloWorld From Admin")
	})

	AdminGroup.Get("/login", func(c *fiber.Ctx) error {
		session, err := middleware.Session.Get(c)
		if err != nil {
			log.Println("error in saving session :", err)
			return err
		}
		log.Println("session: ", session)
		session.Set("isLoggedIn", true)
		session.Save()

		return c.SendString("Login Success...")
	})

	AdminGroup.Get("/logout", func(c *fiber.Ctx) error {
		session, err := middleware.Session.Get(c)
		if err != nil {
			log.Println("error in saving session :", err)
		}
		log.Println("session: ", session)
		session.Delete("isLoggedIn")

		// Destry session
		if err := session.Destroy(); err != nil {
			log.Println("error destroying session : ", err)
			return err
		}

		return c.SendString("Logout Success...")
	})
}

func CustomerRoutes() {
	CustomerGroup.Get("/home", middleware.UserLoggedInWithRole("customer"), func(c *fiber.Ctx) error {
		return c.SendString("HelloWorld From Customer")
	})

	CustomerGroup.Get("/login", func(c *fiber.Ctx) error {
		session, err := middleware.Session.Get(c)
		if err != nil {
			log.Println("error in saving session :", err)
			return err
		}
		log.Println("session: ", session)
		session.Set("isLoggedIn", true)
		session.Set("userType", "customer")
		session.Save()

		return c.SendString("Customer Login Success...")
	})

	CustomerGroup.Get("/logout", func(c *fiber.Ctx) error {
		session, err := middleware.Session.Get(c)
		if err != nil {
			log.Println("error in saving session :", err)
		}
		log.Println("session: ", session)
		session.Delete("isLoggedIn")
		session.Delete("userType")

		// Destry session
		if err := session.Destroy(); err != nil {
			log.Println("error destroying session : ", err)
			return err
		}

		return c.SendString("Customer Logout Success...")
	})
}
