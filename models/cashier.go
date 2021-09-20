package models

import "gorm.io/gorm"

// Cashier is data structure for admin entity
type Cashier struct {
	gorm.Model
	Name     string `validate:"required" json:"name"`
	Username string `validate:"required" gorm:"type:varchar(50);unique_index" json:"username"`
	Password string `validate:"required" json:"password"`
}
