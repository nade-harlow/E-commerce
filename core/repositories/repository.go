package repositories

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) *Repository {
	return &Repository{DB}
}
