package postgresql

import (
	"fmt"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type Config struct {
	User   string
	Pass   string
	Host   string
	Port   string
	DBName string
}

func SetupDatabase(db *gorm.DB, models ...interface{}) error {
	err := db.AutoMigrate(models...)
	return err
}

func New(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.Host, config.User, config.Pass, config.DBName, config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
