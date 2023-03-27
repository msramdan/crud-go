package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		// return json go fiber
		return c.JSON(fiber.Map{
			"hello": "Ramdan",
		})
	})

	app.Listen(":3000")
}
