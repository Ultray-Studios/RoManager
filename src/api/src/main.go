package main

import (
	"log"

	"example.com/m/v2/src/config"
	"example.com/m/v2/src/db"
	"example.com/m/v2/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	app := fiber.New()

	// Load routes
	routes.SetupPlayerRoutes(app)

	// Start server
	port := config.GetEnv("PORT", "80")
	log.Fatal(app.Listen(":" + port))
}
