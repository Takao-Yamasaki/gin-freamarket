package repositories

import (
	"gin-fleamarket/models"

	"gorm.io/gorm"
)

type IAuthRepositoy interface {
	CreateUser(user models.User) error
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
