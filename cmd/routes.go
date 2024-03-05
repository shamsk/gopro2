package main

import (
	//"github.com/divrhino/divrhino-trivia/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/shamsk/gopro2/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	//app.Get("/fact", handlers.NewFactView) // Add new route for new view

	//app.Post("/fact", handlers.CreateFact)
}
