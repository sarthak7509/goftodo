package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	routes "github.com/sarthak7509/goftodo/Routes"
	"github.com/sarthak7509/goftodo/database"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Use(logger.New())
	routes.Routes(app)
	app.Listen(":8000")
}
