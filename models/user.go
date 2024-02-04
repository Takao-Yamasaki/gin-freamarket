package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"not null;unique"`
	Password string `gormt:"not null"`
}
