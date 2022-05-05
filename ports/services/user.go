package services

import (
	"fmt"
	"github.com/nade-harlow/E-commerce/adapter/repository/database/redisql"
	"github.com/nade-harlow/E-commerce/adapter/repository/notification"
	"github.com/nade-harlow/E-commerce/core/models"
	"github.com/nade-harlow/E-commerce/core/requests"
	"github.com/nade-harlow/E-commerce/core/utils"
	repository2 "github.com/nade-harlow/E-commerce/ports/repositories"
	"log"
	"time"
)

type UserServices interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	SignUpUser(user *models.User) error
	SignInUser(user *requests.UserLoginRequest) (*models.User, error)
	VerifyUser(code string) error
	AddUserAddress(address *models.UserAddress) error
	UpdateUserAddress(user *models.UserAddress) error
}

type UserService struct {
	repository repository2.UserRepository
}

const SmsOtpMessage = "Please use the OTP Code: %s to complete your registration. Code expires in five minutes"

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
	msg := fmt.Sprintf(SmsOtpMessage, otp)
	err := userr.repository.SignUpUser(user)
	if err != nil {
		return err
	}
	log.Println("otp: ", otp)
	redisql.SetRedisKey(otp, user.ID, time.Minute*5)
	return notification.SendSms(user.Telephone, msg)
}

func (userr *UserService) SignInUser(user *requests.UserLoginRequest) (*models.User, error) {
	return userr.repository.SignInUser(user)
}

func (user *UserService) VerifyUser(code string) error {
	valid, value := redisql.ValidateRedisKey(code)
	if !valid {
		return fmt.Errorf("invalid OTP")
	}
	log.Println(value.(string))
	go redisql.RemoveRedisKey(code)
	err := user.repository.VerifyUser(value.(string))
	if err != nil {
		return err
	}
	return nil
}

func (user UserService) AddUserAddress(address *models.UserAddress) error {
	return user.repository.AddUserAddress(address)
}

func (user UserService) UpdateUserAddress(address *models.UserAddress) error {
	return user.repository.UpdateUserAddress(address)
}
