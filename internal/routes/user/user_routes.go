package routes

import (
	"github.com/gofiber/fiber/v2"
	handler "github.com/sarthak7509/goftodo/internal/handler/User"
	"github.com/sarthak7509/goftodo/internal/middleware"
)

func UserRoutes(api fiber.Router){
	userRoutes := api.Group("/user")
	// sign up route
	userRoutes.Post("/signup",handler.CreateUser)
	//sign in route
	userRoutes.Post("/signin",handler.SingIn)
	//Get User basic detail
	userRoutes.Get("/",middleware.UserAuthMiddleWare(handler.ProtectedViewSet))
}