package main

import (
    "github.com/gofiber/fiber/v2"
	"github.com/mrajnansky/go-test/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Facts API!")
	})
    app.Get("/fact", handlers.ListFacts)
	app.Post("/fact", handlers.CreateFact)
	app.Get("/fact/random", handlers.GetRandomFact)
}