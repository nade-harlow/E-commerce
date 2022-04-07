package client

import (
	"github.com/nade-harlow/E-commerce/adapter/repository/database/postgresql"
	"github.com/nade-harlow/E-commerce/core/repository"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DB   *gorm.DB
	REPO *repository.Repository
)

func InitializeConnection() {
	db, err := postgresql.New(&postgresql.Config{
		User:   os.Getenv("DB_USER"),
		Pass:   os.Getenv("DB_PASS"),
		Host:   os.Getenv("DB_HOST"),
		Port:   os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
	})
	if err != nil {
		log.Fatal("Failed to connect to mysql database")
	}
	err = postgresql.SetupDatabase(db)
	if err != nil {
		log.Fatal("Failed to setup database")
	}
	repo := repository.New(db)
	DB = db
	REPO = repo
}
