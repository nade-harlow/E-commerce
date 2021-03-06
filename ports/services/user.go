package services

import (
	"fmt"
	"github.com/nade-harlow/E-commerce/core/models"
	"github.com/nade-harlow/E-commerce/core/requests"
	"github.com/nade-harlow/E-commerce/core/utils"
	repository2 "github.com/nade-harlow/E-commerce/ports/repositories"
	"log"
	"os"
	"time"
)

type UserServices interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	SignUpUser(user requests.UserSignUpRequest) error
	SignInUser(user *requests.UserLoginRequest) (*models.User, error)
	VerifyUser(code string) error
	AddUserAddress(address requests.UserAddressRequest) error
	UpdateUserAddress(user requests.UserAddressRequest) error
	ForgotPassword(userID, email string) error
	ResetUserPassword(userID string, password string) error
	ClearUserCache()
}

type UserService struct {
	repository repository2.UserRepository
	mail       repository2.MailRepository
	sms        repository2.TwilloRepository
	redis      repository2.Redis
}

const SmsOtpMessage = "Please use the OTP Code: %s to complete your registration. Code expires in five minutes"

var BaseURL = os.Getenv("BASE_URL")

func NewUserService(repository repository2.UserRepository, mail repository2.MailRepository, sms repository2.TwilloRepository, redis repository2.Redis) UserServices {
	return &UserService{
		repository: repository,
		mail:       mail,
		sms:        sms,
		redis:      redis,
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

func (userr *UserService) SignUpUser(users requests.UserSignUpRequest) error {
	user := &models.User{
		FirstName: users.FirstName,
		LastName:  users.LastName,
		Role:      users.Role,
		Email:     users.Email,
		Password:  users.Password,
		Telephone: users.Telephone,
	}
	otp := utils.GenerateOTP()
	msg := fmt.Sprintf(SmsOtpMessage, otp)
	err := userr.repository.SignUpUser(user)
	if err != nil {
		return err
	}
	log.Println("otp: ", otp)
	userr.redis.SetRedisKey(otp, user.ID, time.Minute*5)
	return userr.sms.SendSms(user.Telephone, msg)
}

func (userr *UserService) SignInUser(user *requests.UserLoginRequest) (*models.User, error) {
	userData, err := userr.repository.SignInUser(user)
	if err != nil {
		return nil, err
	}
	userr.redis.SetRedisKey("userID", userData.ID, time.Hour*24)
	return userData, nil
}

func (user *UserService) VerifyUser(code string) error {
	valid, value := user.redis.ValidateRedisKey(code)
	if !valid {
		return fmt.Errorf("invalid OTP")
	}
	log.Println(value.(string))
	go user.redis.RemoveRedisKey(code)
	err := user.repository.VerifyUser(value.(string))
	if err != nil {
		return err
	}
	return nil
}

func (user *UserService) AddUserAddress(address requests.UserAddressRequest) error {
	userAddress := &models.UserAddress{
		UserID:       address.UserID,
		AddressLine1: address.AddressLine1,
		AddressLine2: address.AddressLine2,
		City:         address.City,
		PostalCode:   address.PostalCode,
		Country:      address.Country,
		Mobile:       address.Mobile,
	}
	return user.repository.AddUserAddress(userAddress)
}

func (user *UserService) UpdateUserAddress(address requests.UserAddressRequest) error {
	userAddress := &models.UserAddress{
		UserID:       address.UserID,
		AddressLine1: address.AddressLine1,
		AddressLine2: address.AddressLine2,
		City:         address.City,
		PostalCode:   address.PostalCode,
		Country:      address.Country,
		Mobile:       address.Mobile,
	}
	return user.repository.UpdateUserAddress(userAddress)
}

func (user *UserService) ResetUserPassword(userID string, password string) error {
	err := user.repository.ResetUserPassword(userID, password)
	if err != nil {
		return err
	}
	return user.repository.RemoveRecoveryPassword(userID)
}

func (user UserService) ForgotPassword(userID, email string) error {
	subject := "E-commerce: Reset Password"
	data := models.RecoverPassword{
		BaseUrl: os.Getenv("BASE_URL"),
		ID:      userID,
	}
	body := utils.ParseTemplate(data)
	err := user.mail.SendMail(email, subject, body)
	if err != nil {
		return err
	}
	return user.repository.AddRecoveryPassword(userID, email)
}

func (user UserService) ClearUserCache() {
	user.redis.RemoveRedisKey("userID")
}
