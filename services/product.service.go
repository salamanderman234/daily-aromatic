package service

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/salamanderman234/daily-aromatic/domain"
	entity "github.com/salamanderman234/daily-aromatic/entities"
	model "github.com/salamanderman234/daily-aromatic/models"
)

type productService struct {
	repo domain.ProductRepository
	rdb  *redis.Client
}

func NewProductService(p domain.ProductRepository) domain.ProductService {
	return &productService{
		repo: p,
	}
}

func (p *productService) CreateProduct(c context.Context, product entity.Product) error {
	// creating model
	var productModel model.Product
	temp, _ := json.Marshal(product)
	json.Unmarshal(temp, &productModel)
	// calling repo
	err := p.repo.CreateProducts(c, []model.Product{
		productModel,
	})
	if err != nil {
		return err
	}

	return nil
}
func (p *productService) GetProduct(c context.Context, id uint) (entity.Product, error) {
	// calling repo
	var productEntity entity.Product
	product, err := p.repo.GetProductByID(c, id)

	if err != nil {
		return entity.Product{}, err
	}
	// creating entity
	temp, _ := json.Marshal(product)
	json.Unmarshal(temp, &productEntity)

	return productEntity, nil
}
func (p *productService) GetAllProducts(c context.Context, page int) ([]entity.Product, entity.Pagination, error) {
	var productsEntity []entity.Product
	limit := queryLimit
	offset := getOffset(page)
	// calling repo
	products, err := p.repo.GetProducts(c, limit, offset, model.Product{})
	if err != nil {
		return nil, entity.Pagination{}, err
	}
	// creating pagination obj
	maxPage := p.repo.GetProductTotal(c, model.Product{})
	pagination := entity.Pagination{
		CurrentPage:  page,
		NextPage:     page + 1,
		PreviousPage: page - 1,
		MaxPage:      int(maxPage),
	}
	// to entity
	temp, _ := json.Marshal(products)
	json.Unmarshal(temp, &productsEntity)

	return productsEntity, pagination, nil
}
func (p *productService) GetProductByFilter(c context.Context, page int, filter entity.Product) ([]entity.Product, entity.Pagination, error) {
	// creating model for repo
	var filterModel model.Product
	temp, _ := json.Marshal(filter)
	json.Unmarshal(temp, &filterModel)
	limit := queryLimit
	offset := getOffset(page)
	// calling repo
	products, err := p.repo.GetProducts(c, limit, offset, filterModel)
	if err != nil {
		return nil, entity.Pagination{}, err
	}

	// creating pagination obj
	maxPage := p.repo.GetProductTotal(c, filterModel)
	pagination := entity.Pagination{
		CurrentPage:  page,
		NextPage:     page + 1,
		PreviousPage: page - 1,
		MaxPage:      int(maxPage),
	}
	// convert model to entity
	var productsModel []entity.Product
	temp, _ = json.Marshal(products)
	json.Unmarshal(temp, &productsModel)

	return productsModel, pagination, nil
}
func (p *productService) UpdateProduct(c context.Context, id uint, updatedProduct entity.Product) error {
	var updateModel model.Product
	temp, _ := json.Marshal(updatedProduct)
	json.Unmarshal(temp, &updateModel)

	err := p.repo.UpdateProduct(c, id, updateModel)
	if err != nil {
		return err
	}
	return nil

}
func (p *productService) DeleteProduct(c context.Context, id uint) error {
	err := p.repo.DeleteProduct(c, id)
	if err != nil {
		return err
	}
	return nil
}
