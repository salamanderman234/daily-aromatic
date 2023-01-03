package service

import (
	"context"
	"encoding/json"

	"github.com/salamanderman234/daily-aromatic/domain"
	entity "github.com/salamanderman234/daily-aromatic/entities"
	model "github.com/salamanderman234/daily-aromatic/models"
)

type productService struct {
	repo domain.ProductRepository
}

func NewProductService(p domain.ProductRepository) domain.ProductService {
	return &productService{
		repo: p,
	}
}

func (p *productService) CreateProduct(c context.Context, product entity.Product) error {
	var productModel model.Product
	temp, _ := json.Marshal(product)
	json.Unmarshal(temp, &productModel)

	err := p.repo.CreateProducts(c, []model.Product{
		productModel,
	})
	if err != nil {
		return err
	}

	return nil
}
func (p *productService) GetProduct(c context.Context, id uint) (entity.Product, error) {
	var productEntity entity.Product
	product, err := p.repo.GetProductByID(c, id)

	if err != nil {
		return entity.Product{}, err
	}

	temp, _ := json.Marshal(product)
	json.Unmarshal(temp, &productEntity)

	return productEntity, nil
}
func (p *productService) GetAllProducts(c context.Context, page int) ([]entity.Product, error) {
	var productsEntity []entity.Product
	products, err := p.repo.GetProducts(c, 0, 0, model.Product{})
	if err != nil {
		return nil, err
	}

	temp, _ := json.Marshal(products)
	json.Unmarshal(temp, &productsEntity)

	return productsEntity, nil
}
func (p *productService) GetProductByFilter(c context.Context, page int, filter entity.Product) ([]entity.Product, error) {
	var filterModel model.Product
	temp, _ := json.Marshal(filter)
	json.Unmarshal(temp, &filterModel)

	products, err := p.repo.GetProducts(c, 0, 0, filterModel)
	if err != nil {
		return nil, err
	}

	var productsModel []entity.Product
	temp, _ = json.Marshal(products)
	json.Unmarshal(temp, &productsModel)

	return productsModel, nil
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
