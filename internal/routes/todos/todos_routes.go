package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sarthak7509/goftodo/internal/handler/Todo"
	"github.com/sarthak7509/goftodo/internal/middleware"
)

func TodosRoute(api fiber.Router) {
	todoRoutes := api.Group("/todos")
	// create todo route
	todoRoutes.Post("/", middleware.UserAuthMiddleWare(Todo.CreateTodo))
	// Get todo List route
	todoRoutes.Get("/", middleware.UserAuthMiddleWare(Todo.GetTodoList))
	// Get todo route
	todoRoutes.Get("/:todoId", middleware.UserAuthMiddleWare(Todo.GetTodo))
	//Update route
	todoRoutes.Put("/:todoId", middleware.UserAuthMiddleWare(Todo.UpdateTodo))
	//Delete route
	todoRoutes.Delete("/:todoId", middleware.UserAuthMiddleWare(Todo.DeleteTodo))
}
