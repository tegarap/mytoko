package database

import (
	"gorm.io/gorm"
	"mytoko/models"
)

type (
	GormProductModel struct {
		db *gorm.DB
	}
	ProductModel interface {
		Add(models.Product) error
		GetAll() ([]models.Product, error)
		GetByID(productID int) (models.Product, error)
		//AddMany([]models.Product) error
	}
)

// NewProductModel NewProductRepo method repo init
func NewProductModel(db *gorm.DB) ProductModel {
	return &GormProductModel{db: db}
}

// Add method to add new product
func (m *GormProductModel) Add(product models.Product) error {
	if err := m.db.Save(&product).Error; err != nil {
		return err
	}
	return nil
}

// GetAll method to get all data product
func (m *GormProductModel) GetAll() ([]models.Product, error) {
	var products []models.Product
	if err := m.db.Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

// GetByID method to retrieve a single data product by id
func (m *GormProductModel) GetByID(productID int) (models.Product, error) {
	product := models.Product{}
	if err := m.db.First(&product, productID).Error; err != nil {
		return product, err
	}
	return product, nil
}

// AddMany method to add many new products
//func (m *GormProductModel) AddMany(products []models.Product) error {
//	var sqlStr string
//	for _, product := range products {
//		sqlStr += fmt.Sprintf("INSERT INTO products (created_at, updated_at, sku, name) VALUES (datetime('now','localtime'), datetime('now','localtime'),'%s', '%s'); ", product.Sku, product.Name)
//	}
//	m.db.Exec(sqlStr)
//	return nil
//}
