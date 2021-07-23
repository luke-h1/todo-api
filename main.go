package main

import (
	"github.com/gofiber/fiber"
)

func testHandler(c *fiber.Ctx) {
	c.Send("Hello world 👋")
}

func main() {
	app := fiber.New()

	app.Get("/", testHandler)
	app.Listen(4000)
}
