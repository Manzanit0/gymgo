package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	setupRouter(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}

func setupRouter(app *fiber.App) {
	app.Put("/classes", createClassHandler)
}

func createClassHandler(c *fiber.Ctx) error {
	b := &class{}
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := createClass(b); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(b)
}
