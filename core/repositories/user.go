package repositories

import (
	"errors"
	"github.com/nade-harlow/E-commerce/core/models"
	"github.com/nade-harlow/E-commerce/core/utils"
)

func (repo *Repository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := repo.DB.Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *Repository) GetUserByID(id string) (*models.User, error) {
	user := &models.User{}
	err := repo.DB.Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *Repository) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := repo.DB.Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *Repository) SignUpUser(user *models.User) error {
	userByEmail, err := repo.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if userByEmail != nil {
		return errors.New("user with this email already exists")
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

func (repo *Repository) SignInUser(user *models.User) error {
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
