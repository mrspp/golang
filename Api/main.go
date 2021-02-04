package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger" // new
	"github.com/mrspp/golang-api/Algorithms/fibonacci"
	"github.com/mrspp/golang-api/Algorithms/isprime"
)

func main() {
	app := fiber.New()
	app.Use(logger.New()) // new

	// give response when at /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})
	})

	app.Get("/fibonacci", func(c *fiber.Ctx) error {
		c.Query("length")
		val := c.Query("length")
		s, err := strconv.Atoi(val)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Ok",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":   true,
			"message":   "Ok",
			"fibonacci": fibonacci.ResultAsString(s),
		})
	})

	app.Get("/check-prime", func(c *fiber.Ctx) error {
		c.Query("number")

		val := c.Query("number")
		s, err := strconv.Atoi(val)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "StatusBadRequest",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":  true,
			"isPrime?": isprime.IsPrime(s),
		})
	})

	// Listen on server 8000 and catch error if any
	err := app.Listen(":8000")

	// handle error
	if err != nil {
		panic(err)
	}
}
