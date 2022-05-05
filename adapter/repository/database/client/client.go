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
	err = postgresql.SetupDatabase(db, &models.User{}, &models.UserAddress{}, &models.UserPayment{}, &models.UserPayment{},
		&models.Product{}, &models.ProductCategory{}, &models.CartItem{}, &models.ProductImage{}, &models.OrderDetails{},
		&models.OrderItems{}, &models.PaymentDetails{}, &models.PasswordRecovery{})
	if err != nil {
		log.Fatal("Failed to setup database")
	}
	return db
}
