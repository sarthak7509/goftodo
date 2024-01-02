package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sarthak7509/goftodo/internal/handler/Todo"
	"github.com/sarthak7509/goftodo/internal/middleware"
)

func TodosRoute(api fiber.Router){
	todoRoutes := api.Group("/todos")
	// create todo route
	todoRoutes.Post("/",middleware.UserAuthMiddleWare(Todo.CreateTodo))
	// Get todo route
	todoRoutes.Get("/",middleware.UserAuthMiddleWare(Todo.GetTodoList))
	
}