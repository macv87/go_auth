package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
	"github.com/macv/go_auth/database"
	routes "github.com/macv/go_auth/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")

}
