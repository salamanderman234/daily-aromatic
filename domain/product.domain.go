package domain

import (
	"context"

	entity "github.com/salamanderman234/daily-aromatic/entities"
	model "github.com/salamanderman234/daily-aromatic/models"
)

type ProductService interface {
	CreateProduct(c context.Context, product entity.Product) error
	GetProduct(c context.Context, id uint) (entity.Product, error)
	GetAllProducts(c context.Context, page int) ([]entity.Product, entity.Pagination, error)
	GetProductByFilter(c context.Context, page int, filter entity.Product) ([]entity.Product, entity.Pagination, error)
	GetProductNotInUserReview(c context.Context, page int, filter entity.Product, idUser uint) ([]entity.Product, entity.Pagination, error)
	UpdateProduct(c context.Context, id uint, updatedProduct entity.Product) error
	DeleteProduct(c context.Context, id uint) error
}

type ProductRepository interface {
	GetProductTotal(c context.Context, filter model.Product) int64
	CreateProducts(c context.Context, products []model.Product) error
	GetProductByID(c context.Context, id uint) (model.Product, error)
	GetProducts(c context.Context, limit int, skip int, filter model.Product) ([]model.Product, error)
	GetProductsForUserReview(c context.Context, limit int, skip int, filter model.Product, idUser uint) ([]model.Product, error)
	UpdateProduct(c context.Context, id uint, updatedField model.Product) error
	DeleteProduct(c context.Context, id uint) error
}
