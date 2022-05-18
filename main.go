package main

import (
	"github.com/fitnessstack/web-service/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	// Allow cors because f*ck that shit
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	database.ConnectDatabase()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Fiber world!")
	})

	log.Fatal(app.Listen(":" + getenv("port", "3000")))
}
