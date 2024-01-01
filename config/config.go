package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string{
	err := godotenv.Load(".env")
	if err!=nil{
		log.Fatal("Something happend cant load config")
		panic("Something happend cant load config")
	}
	return os.Getenv(key)
}