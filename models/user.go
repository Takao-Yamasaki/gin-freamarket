package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	// ユーザーが削除されたら、商品も自動的に削除
	items []Item `gorm:"constraint:OnDelete:CASCADE"`
}
