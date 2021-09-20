package models

import (
	"gorm.io/gorm"
	"time"
)

// Transaction is data structure for transaction entity
type Transaction struct {
	gorm.Model
	CashierID       uint      `validate:"required" json:"cashier_id"`
	CustomerName    string    `validate:"required" json:"name"`
	TotalPrice      int       `validate:"required" json:"price"`
	TransactionTime time.Time `validate:"required" json:"size"`
}