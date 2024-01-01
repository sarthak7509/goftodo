package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uuid.UUID `gorm:"Primarykey"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
	ApiKey   string
}

type UserSignUp struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserSignUpResponseStruct struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserSignInResponseStruct struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ApiKey string `json:"token"`
}

func UserSignUpResponse(user User) UserSignUpResponseStruct {
	return UserSignUpResponseStruct{
		Email: user.Email,
		Name:  user.Name,
	}
}

func UserSignInResponse(user User) UserSignInResponseStruct {
	return UserSignInResponseStruct{
		Email: user.Email,
		Name:  user.Name,
		ApiKey: user.ApiKey,
	}
}
