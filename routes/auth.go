package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/punokawan/Simple-Web-Dompet-Kilat/controllers"
	"github.com/punokawan/Simple-Web-Dompet-Kilat/helpers"
)

func AuthRoute(route fiber.Router) {
	route.Get("/", func(c *fiber.Ctx) error {
		return helpers.ResponseMsg(c, 200, true, "Hello, this is routes with preffix auth 👋!", nil)
	})
	route.Post("/register", controllers.SignUp)
	route.Post("/login", controllers.SignIn)
}
