package domain

import (
	entity "github.com/salamanderman234/daily-aromatic/entities"
	model "github.com/salamanderman234/daily-aromatic/models"
)

type ProductService interface {
	CreateProduct(product entity.Product) error
	CreateProductBatch(products []entity.Product) error
	GetProduct(id uint) (entity.Product, error)
	GetAllProducts(page int) ([]entity.Product, error)
	GetAllProductsWithReview(page int) ([]entity.Product, error)
	GetProductByFilter(page int, filter entity.Product) ([]entity.Product, error)
	UpdateProduct(id uint, updatedProduct entity.Product) error
	DeleteProduct(id uint) error
}

type ProductRepository interface {
	GetProductTotal() int64
	CreateProducts(products []model.Product) error
	GetProductByID(id uint) (model.Product, error)
	GetProducts(limit int, skip int, filter model.Product) ([]model.Product, error)
	UpdateProduct(id uint, updatedField model.Product) error
	DeleteProduct(id uint) error
}
