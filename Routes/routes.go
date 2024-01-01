package routes

import (
	"github.com/gofiber/fiber/v2"
	routes "github.com/sarthak7509/goftodo/internal/routes/user"
)

func Routes(app *fiber.App) {
	api := app.Group("/api")
	routes.UserRoutes(api)
}
