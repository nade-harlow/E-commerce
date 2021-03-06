package repositories

import (
	"github.com/nade-harlow/E-commerce/core/models"
	"github.com/nade-harlow/E-commerce/core/requests"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	SignUpUser(user *models.User) error
	SignInUser(user *requests.UserLoginRequest) (*models.User, error)
	VerifyUser(userID string) error
	AddRecoveryPassword(userID, email string) error
	RemoveRecoveryPassword(userID string) error
	AddUserAddress(address *models.UserAddress) error
	UpdateUserAddress(user *models.UserAddress) error
	ResetUserPassword(userID string, password string) error
}
