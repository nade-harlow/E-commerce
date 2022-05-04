package services

import (
	"github.com/nade-harlow/E-commerce/adapter/repository/notification"
	"github.com/nade-harlow/E-commerce/core/models"
	"github.com/nade-harlow/E-commerce/core/requests"
	"github.com/nade-harlow/E-commerce/core/utils"
	repository2 "github.com/nade-harlow/E-commerce/ports/repositories"
)

type UserServices interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	SignUpUser(user *models.User) error
	SignInUser(user *requests.UserLoginRequest) (*models.User, error)
}

type UserService struct {
	repository repository2.UserRepository
}

func NewUserService(repository repository2.UserRepository) UserServices {
	return &UserService{
		repository: repository,
	}
}

func (user *UserService) GetUserByEmail(email string) (*models.User, error) {
	return user.repository.GetUserByEmail(email)
}

func (user *UserService) GetUserByID(id string) (*models.User, error) {
	return user.repository.GetUserByID(id)
}

func (user *UserService) GetUserByUsername(username string) (*models.User, error) {
	return user.repository.GetUserByUsername(username)
}

func (userr *UserService) SignUpUser(user *models.User) error {
	otp := utils.GenerateOTP()
	msg := "Please use the otp in verifying your account: " + otp

	err := notification.SendSms(user.Telephone, msg)
	if err != nil {
		return err
	}
	return userr.repository.SignUpUser(user)
}

func (userr UserService) SignInUser(user *requests.UserLoginRequest) (*models.User, error) {
	return userr.repository.SignInUser(user)
}
