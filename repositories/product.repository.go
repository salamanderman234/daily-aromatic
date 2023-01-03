package repository

import (
	"context"

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

func (p *productRepository) CreateProducts(c context.Context, products []model.Product) error {
	result := p.conn.WithContext(c).Create(&products)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *productRepository) GetProductTotal(c context.Context, filter model.Product) int64 {
	result := int64(0)
	var product model.Product
	p.conn.WithContext(c).
		Model(&product).
		Where(&model.Product{Pabrikan: filter.Pabrikan}).
		Or(&model.Product{Aroma: filter.Aroma}).
		Or(&model.Product{Nama: filter.Nama}).
		Count(&result)

	return result
}

func (p *productRepository) GetProductByID(c context.Context, id uint) (model.Product, error) {
	var product model.Product
	result := p.conn.WithContext(c).Preload("Reviews").Where("id = ?", id).First(&product)
	if result.Error != nil {
		return model.Product{}, result.Error
	}
	return product, nil
}
func (p *productRepository) GetProducts(c context.Context, limit int, skip int, filter model.Product) ([]model.Product, error) {
	// querying product table using or statement
	var products []model.Product
	result := p.conn.WithContext(c).
		Where(&model.Product{Pabrikan: filter.Pabrikan}).
		Or(&model.Product{Aroma: filter.Aroma}).
		Or(&model.Product{Nama: filter.Nama}).
		Offset(skip).
		Limit(limit).
		Order("created_at desc").
		Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
func (p *productRepository) UpdateProduct(c context.Context, id uint, updatedField model.Product) error {
	product := model.Product{}
	result := p.conn.WithContext(c).Model(&product).Where("id = ?", id).Updates(updatedField)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (p *productRepository) DeleteProduct(c context.Context, id uint) error {
	product := model.Product{}
	result := p.conn.WithContext(c).Where("id = ?", id).Delete(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
