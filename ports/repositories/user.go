package repositories

import "github.com/nade-harlow/E-commerce/core/models"

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	SignUpUser(user *models.User) error
	SignInUser(user *models.User) error
}
