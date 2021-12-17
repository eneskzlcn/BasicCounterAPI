package main

import (
	"github.com/gofiber/fiber/v2"
)

type Counter struct {
	Value int `json:"value""`
}

var counter = Counter{Value: 0}

func IncreaseHandler(c *fiber.Ctx) error {
	counter.Value++
	return c.Status(fiber.StatusOK).JSON(counter)
}
func DecreaseHandler(c *fiber.Ctx) error {
	counter.Value--
	return c.Status(fiber.StatusOK).JSON(counter)
}
func CounterHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(counter)
}
func ResetHandler(c *fiber.Ctx) error {
	counter.Value = 0
	return c.Status(fiber.StatusOK).JSON(counter)
}

// StartCounterApi starts the server
func StartCounterApi() {
	app := fiber.New()
	app.Get("/counter", CounterHandler)
	app.Get("/increase", IncreaseHandler)
	app.Get("/decrease", DecreaseHandler)
	app.Get("/reset", ResetHandler)
	app.Listen(":3003")
}
func main() {
	StartCounterApi()
}
