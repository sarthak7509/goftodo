package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sarthak7509/goftodo/database"
	"github.com/sarthak7509/goftodo/internal/auth"
	"github.com/sarthak7509/goftodo/internal/models"
)

type Handler func(c *fiber.Ctx,user models.User) error

func UserAuthMiddleWare(handler Handler) fiber.Handler{
	return func (c *fiber.Ctx) error{
		db := database.DB
		tokenString,err := auth.GetApiKey(c)

		if err != nil{
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err})
		}
		user := models.User{}
		err = db.Find(&user,"api_key=?",tokenString).Error
		if err!=nil{
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "User Not Found"})
		}
		
		return handler(c,user)

	}
}