package models

import (
	"time"
)

// Admin is data structure for admin entity
type Admin struct {
	ID        uint       `gorm:"primarykey"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
	Username  string     `validate:"required" gorm:"type:varchar(50);unique_index" json:"username"`
	Password  string     `validate:"required" json:"password"`
}
