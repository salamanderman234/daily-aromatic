package repository

import (
	"github.com/salamanderman234/daily-aromatic/domain"
	model "github.com/salamanderman234/daily-aromatic/models"
	"gorm.io/gorm"
)

type productRepository struct {
	conn  *gorm.DB
	table string
}

func NewProductRepository(conn *gorm.DB) domain.ProductRepository {
	return &productRepository{
		conn:  conn,
		table: "products",
	}
}

func (p *productRepository) CreateProducts(products []model.Product) error {
	result := p.conn.Create(&products)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *productRepository) GetProductTotal() int64 {
	result := int64(0)
	var product model.Product
	p.conn.Model(&product).Count(&result)
	return result
}

func (p *productRepository) GetProductByID(id uint) (model.Product, error) {
	var product model.Product
	result := p.conn.Where("id = ?", id).First(&product)
	if result.Error != nil {
		return model.Product{}, result.Error
	}
	return product, nil
}
func (p *productRepository) GetProducts(limit int, skip int, filter model.Product) ([]model.Product, error) {
	var products []model.Product
	result := p.conn.Where(&filter).Offset(skip).Limit(limit).Order("age desc, name").Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
func (p *productRepository) UpdateProduct(id uint, updatedField model.Product) error {
	product := model.Product{}
	result := p.conn.Model(&product).Where("id = ?", id).Updates(updatedField)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (p *productRepository) DeleteProduct(id uint) error {
	product := model.Product{}
	result := p.conn.Where("id = ?", id).Delete(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
