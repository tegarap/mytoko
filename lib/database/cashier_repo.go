package database

import (
	"gorm.io/gorm"
	"mytoko/models"
)

type (
	GormCashierModel struct {
		db *gorm.DB
	}
	CashierModel interface {
		Add(models.Cashier) error
		GetAll() ([]models.Cashier, error)
		GetByID(cashierID int) (models.Cashier, error)
	}
)

// NewCashierModel method repo init
func NewCashierModel(db *gorm.DB) CashierModel {
	return &GormCashierModel{db: db}
}

// Add method to add new cashier
func (m *GormCashierModel) Add(cashier models.Cashier) error {
	if err := m.db.Save(&cashier).Error; err != nil {
		return err
	}
	return nil
}

// GetAll method to get all data cashier
func (m *GormCashierModel) GetAll() ([]models.Cashier, error) {
	var cashiers []models.Cashier
	if err := m.db.Find(&cashiers).Error; err != nil {
		return cashiers, err
	}
	return cashiers, nil
}

// GetByID method to retrieve a single data cashier by id
func (m *GormCashierModel) GetByID(cashierID int) (models.Cashier, error) {
	cashier := models.Cashier{}
	if err := m.db.First(&cashier, cashierID).Error; err != nil {
		return cashier, err
	}
	return cashier, nil
}