package models

import "gorm.io/gorm"

// Product is data structure for product entity
type Product struct {
	gorm.Model
	Sku        string `validate:"required" gorm:"type:varchar(100);unique_index" json:"sku"`
	Name       string `validate:"required" json:"name"`
	CategoryID uint   `validate:"required" json:"category_id"`
	Size       string `validate:"required" json:"size"`
	Price      int    `validate:"required" json:"price"`
	Stock      int    `json:"stock"`
}
