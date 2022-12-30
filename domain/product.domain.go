package domain

import (
	entity "github.com/salamanderman234/daily-aromatic/entities"
	model "github.com/salamanderman234/daily-aromatic/models"
)

type ProductService interface {
	CreateProduct(product entity.Product) error
	CreateProductBatch(products []entity.Product) error
	GetProduct(id uint) (entity.Product, error)
	GetAllProducts(limit int) ([]entity.Product, error)
	GetProductByFilter(limit int, filter entity.Product) ([]entity.Product, error)
	UpdateProduct(id uint, updatedProduct entity.Product) error
	DeleteProduct(id uint) error
}

type ProductRepository interface {
	CreateProducts(products []model.Product) error
	GetProductByID(id uint) (model.Product, error)
	GetProducts(limit int, filter model.Product) ([]model.Product, error)
	UpdateProduct(id uint, updatedField model.Product) error
	DeleteProduct(id uint) error
}
