package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jun2900/refactoryTest/Q6/Testing"
)

type Log struct {
	Counter int `json:"counter"`
}

func main() {
	go Testing.Client()
	app := fiber.New()

	file, err := os.OpenFile("./server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status}: ${method} http://localhost:3000/ ${resBody} \n",
		TimeFormat: "2006-01-02T15:04:05Z",
		Output:     file,
	}))

	app.Post("/", func(c *fiber.Ctx) error {
		request := new(Log)

		if err := c.BodyParser(request); err != nil {
			return err
		}
		return c.Status(201).JSON(fiber.Map{"X-RANDOM": c.Get("X-RANDOM"), "counter": request.Counter})
	})

	app.Listen(":3000")
}
