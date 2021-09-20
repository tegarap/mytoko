package models

import "gorm.io/gorm"

// Category is data structure for category entity
type Category struct {
	gorm.Model
	Name       string `validate:"required" json:"name"`
}
