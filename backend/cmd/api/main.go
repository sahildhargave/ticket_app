package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sahildhargave/ticket-project-v1/config"
	"github.com/sahildhargave/ticket-project-v1/db"
	"github.com/sahildhargave/ticket-project-v1/handlers"
	"github.com/sahildhargave/ticket-project-v1/repositories"
)

func main() {
	// Load environment configuration
	envConfig := config.NewEnvConfig()

	// Initialize the database
	database, err := db.Init(envConfig, db.DBMigrator) // Capture both return values
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Create a new Fiber application
	app := fiber.New(fiber.Config{
		AppName:      "TicketBooking",
		ServerHeader: "Fiber",
	})

	// Initialize the event repository
	eventRepository := repositories.NewEventRepository(database)

	// Define API routes
	server := app.Group("/api")
	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	// Start the server
	if err := app.Listen(fmt.Sprintf(":%s", envConfig.ServerPort)); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
