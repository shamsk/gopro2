package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shamsk/gopro2/tree/master/database"
	
	
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("now it time ")
	})

	app.Listen(":3000")
}






