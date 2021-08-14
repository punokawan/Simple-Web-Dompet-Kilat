package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/punokawan/Simple-Web-Dompet-Kilat/config"
	"github.com/punokawan/Simple-Web-Dompet-Kilat/routes"
)

func init() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":     true,
			"message":     "You are at the root endpoint 😉",
			"github_repo": "https://github.com/punokawan/Simple-Web-Dompet-Kilat",
		})
	})

	api := app.Group("/api")

	routes.AuthRoute(api.Group("/auth"))
}

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	config.ConnectDB()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
