package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	fmt.Println("Simple crud application using fiber")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, World!"})
		// return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}