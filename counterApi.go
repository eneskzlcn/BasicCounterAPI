package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
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
func StartCounterApi(port int) error {

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/counter", CounterHandler)
	app.Get("/increase", IncreaseHandler)
	app.Get("/decrease", DecreaseHandler)
	app.Get("/reset", ResetHandler)
	err := app.Listen(fmt.Sprintf(":%d", port))
	return err
}
func main() {
	err:= StartCounterApi(4000)
	if err != nil {
		log.Fatal(err)
	}
}
