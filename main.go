package main

import (
	"log"
	"os"
	"synapsis-test/app/service"

	"github.com/gofiber/fiber/v2"
)

// @title Synapsis Test
// @version 1.0
// @description Documentation for synapsys test
// @host localhost:8000
// @BasePath /api
func main() {
	service.InitDB()

	app := fiber.New()

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
