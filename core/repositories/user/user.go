package user

import (
	"errors"
	"github.com/nade-harlow/E-commerce/core/models"
	"github.com/nade-harlow/E-commerce/core/utils"
	"gorm.io/gorm"
	"strings"
)

type UserRepository struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) *UserRepository {
	return &UserRepository{DB}
}

func (repo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := repo.DB.Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetUserByID(id string) (*models.User, error) {
	user := &models.User{}
	err := repo.DB.Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := repo.DB.Where("username = ?", username).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) SignUpUser(user *models.User) error {
	user.Email = strings.ToLower(user.Email)
	user.FirstName = strings.ToLower(user.FirstName)
	user.LastName = strings.ToLower(user.LastName)
	user.Username = strings.ToLower(user.Username)

	userByEmail, err := repo.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if userByEmail != nil {
		return errors.New("user with this email already exists")
	}

	username, err := repo.GetUserByUsername(user.Username)
	if err != nil {
		return err
	}
	if username != nil {
		return errors.New("user with this username already exists")
	}
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		return errors.New("error hashing password")
	}
	err = repo.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) SignInUser(user *models.User) error {
	userByEmail, err := repo.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if userByEmail == nil {
		return errors.New("user with this email does not exist")
	}
	if ok := utils.CheckPasswordHash(user.Password, userByEmail.Password); !ok {
		return errors.New("incorrect password")
	}
	return nil
}
