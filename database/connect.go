package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/sarthak7509/goftodo/config"
	"github.com/sarthak7509/goftodo/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var erro error

func Connect() {
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Printf("Not able to parse port: %v", err)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	DB, erro = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatalf("Something going wrong: %v\n", err)
		panic("Could not connect to database")
	}
	DB.AutoMigrate(&models.User{}, &models.Todo{})
	log.Printf("connected to database successfully. %v", DB.Name())

}
