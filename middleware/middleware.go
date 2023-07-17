package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Session *session.Store

func InitializeSession() {
	// Create a new session config
	sessionConfig := session.Config{
		Expiration:   1 * time.Minute,
		CookieSecure: true,
	}

	// Initialize the session middleware
	Session = session.New(sessionConfig)
}

func isLoggedIn(c *fiber.Ctx) bool {
	// Check if the user is logged in based on your application's logic
	// For example, you can check if a session or authentication token exists
	sess, err := Session.Get(c)
	if err != nil {
		log.Println("error in getting session.")
		return false
	}

	// Assuming you have a session value named "loggedInUser"
	loggedInUser := sess.Get("isLoggedIn")

	// Check if the user is logged in or not
	if loggedInUser == nil {
		return false
	}

	// Check if the user is logged in or not
	if !loggedInUser.(bool) {
		return false
	}

	return true
}

func isLoggedInWithRole(c *fiber.Ctx, role string) bool {
	// Check if the user is logged in based on your application's logic
	// For example, you can check if a session or authentication token exists
	sess, err := Session.Get(c)
	if err != nil {
		log.Println("error in getting session.")
		return false
	}

	// Assuming you have a session value named "loggedInUser"
	loggedInUser := sess.Get("isLoggedIn")
	loggedInRole := sess.Get("userType")

	// Check if the user is logged in or not
	if loggedInUser == nil || loggedInRole == nil {
		return false
	}

	// Check if the user is logged in and with proper role
	if !loggedInUser.(bool) || loggedInRole.(string) != role {
		return false
	}

	return true
}

func UserLoggedIn() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check if the user is logged in
		if !isLoggedIn(c) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
		}

		// Continue to the next middleware or route handler
		return c.Next()
	}
}

func UserLoggedInWithRole(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check if the user is logged in
		if !isLoggedInWithRole(c, role) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
		}

		// Continue to the next middleware or route handler
		return c.Next()
	}
}
