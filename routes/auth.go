package routes

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/mikefmeyer/catchphrase-go-mongodb-rest-api/controllers" // replace
)

func AuthRoute(route fiber.Router) {
	route.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, this is routes with preffix auth 👋!")
	})
}
