package auth

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetApiKey(c *fiber.Ctx) (string,error){
	tokenString := ""
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == ""{
		return tokenString,errors.New("You are not Logged In")
	}

	return tokenString,nil
}