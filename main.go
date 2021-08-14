package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/punokawan/Simple-Web-Dompet-Kilat/database"
	"github.com/punokawan/Simple-Web-Dompet-Kilat/helpers"
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
		return helpers.ResponseMsg(
			c,
			200,
			true,
			"You are at the root endpoint 😉, github_repo: https://github.com/punokawan/Simple-Web-Dompet-Kilat",
			nil)
	})

	api := app.Group("/api")

	routes.AuthRoute(api.Group("/auth"))
}

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	database.ConnectDB()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
