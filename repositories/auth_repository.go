package repositories

import (
	"errors"
	"gin-fleamarket/models"

	"gorm.io/gorm"
)

type IAuthRepositoy interface {
	CreateUser(user models.User) error
	FindUser(email string) (*models.User, error)
}

type AuthReposiitory struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepositoy {
	return &AuthReposiitory{db: db}
}

func (r *AuthReposiitory) CreateUser(user models.User) error {
	result := r.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *AuthReposiitory) FindUser(email string) (*models.User, error) {
	var user models.User
	// "email = ?"は、SQLのWhere句に相当
	result := r.db.First(&user, "email = ?", email)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}
