package database

import (
	"gorm.io/gorm"
	"mytoko/models"
)

type (
	GormCategoryModel struct {
		db *gorm.DB
	}
	CategoryModel interface {
		Add(category models.Category) error
		GetAll() ([]models.Category, error)
		Update(category models.Category, id int) (models.Category, error)
		Delete(id int) (models.Category, error)
	}
)

// NewCategoryModel method repo init
func NewCategoryModel(db *gorm.DB) CategoryModel {
	return &GormCategoryModel{db: db}
}

func (m GormCategoryModel) Add(category models.Category) error {
	if err := m.db.Save(&category).Error; err != nil {
		return err
	}
	return nil
}

func (m GormCategoryModel) GetAll() ([]models.Category, error) {
	var categories []models.Category
	if err := m.db.Select("id, username").Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}

func (m GormCategoryModel) Update(categoryTemp models.Category, id int) (models.Category, error) {
	var category models.Category
	if err := m.db.First(&category, id).Error; err != nil {
		return category, err
	}
	if categoryTemp.Name != "" {
		category.Name = categoryTemp.Name
	}
	if err := m.db.Save(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (m GormCategoryModel) Delete(id int) (models.Category, error) {
	var category models.Category
	if err := m.db.Delete(&category, id).Error; err != nil {
		return category, err
	}
	return category, nil
}