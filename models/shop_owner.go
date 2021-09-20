package models

import "gorm.io/gorm"

// ShopOwner is data structure for shop owner entity
type ShopOwner struct {
	gorm.Model
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
