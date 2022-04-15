package client

import (
	"github.com/nade-harlow/E-commerce/adapter/repository/database/postgresql"
	"github.com/nade-harlow/E-commerce/core/models"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitializeConnection() *gorm.DB {
	db, err := postgresql.New(&postgresql.Config{
		User:   os.Getenv("DB_USER"),
		Pass:   os.Getenv("DB_PASS"),
		Host:   os.Getenv("DB_HOST"),
		Port:   os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
	})
	if err != nil {
		log.Fatal("Failed to connect to postgresql database")
	}
	err = postgresql.SetupDatabase(db, &models.Product{}, &models.User{}, &models.Cart{})
	if err != nil {
		log.Fatal("Failed to setup database")
	}
	return db
}
