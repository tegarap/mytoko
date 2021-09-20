package database

import (
	"gorm.io/gorm"
	"mytoko/models"
)

type (
	GormAdminModel struct {
		db *gorm.DB
	}
	AdminModel interface {
		Add(models.Admin) error
		GetAll() ([]models.Admin, error)
		Delete(id int) (models.Admin, error)
	}
)

// NewAdminModel method repo init
func NewAdminModel(db *gorm.DB) AdminModel {
	return &GormAdminModel{db: db}
}

// Add method to add new admin
func (m *GormAdminModel) Add(admin models.Admin) error {
	if err := m.db.Save(&admin).Error; err != nil {
		return err
	}
	return nil
}

// GetAll method to get all data admin
func (m *GormAdminModel) GetAll() ([]models.Admin, error) {
	var admins []models.Admin
	if err := m.db.Select("id, username").Find(&admins).Error; err != nil {
		return admins, err
	}
	return admins, nil
}

// Delete method to delete data admin
func (m *GormAdminModel) Delete(id int) (models.Admin, error) {
	var admin models.Admin
	if err := m.db.Delete(&admin, id).Error; err != nil {
		return admin, err
	}
	return admin, nil
}
