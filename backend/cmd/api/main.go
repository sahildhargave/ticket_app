package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sahildhargave/ticket-project-v1/handlers"
	"github.com/sahildhargave/ticket-project-v1/repositories"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "TicketBooking",
		ServerHeader: "Fiber",
	})

	eventRepository := repositories.NewEventRepository(nil)

	server := app.Group("/api")

	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(":3000")
}
