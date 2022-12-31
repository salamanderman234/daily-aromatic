package service

import (
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

func (p *productService) CreateProduct(product entity.Product) error {
	var productModel model.Product
	temp, _ := json.Marshal(product)
	json.Unmarshal(temp, &productModel)

	err := p.repo.CreateProducts([]model.Product{
		productModel,
	})
	if err != nil {
		return err
	}

	return nil
}
func (p *productService) CreateProductBatch(products []entity.Product) error {
	var productsModel []model.Product
	temp, _ := json.Marshal(products)
	json.Unmarshal(temp, &productsModel)

	err := p.repo.CreateProducts(productsModel)
	if err != nil {
		return err
	}

	return nil
}
func (p *productService) GetProduct(id uint) (entity.Product, error) {
	var productEntity entity.Product
	product, err := p.repo.GetProductByID(id)

	if err != nil {
		return entity.Product{}, err
	}

	temp, _ := json.Marshal(product)
	json.Unmarshal(temp, &productEntity)

	return productEntity, nil
}
func (p *productService) GetAllProducts(page int) ([]entity.Product, error) {
	var productsEntity []entity.Product
	products, err := p.repo.GetProducts(0, 0, model.Product{})
	if err != nil {
		return nil, err
	}

	temp, _ := json.Marshal(products)
	json.Unmarshal(temp, &productsEntity)

	return productsEntity, nil
}
func (p *productService) GetAllProductsWithReview(page int) ([]entity.Product, error) {
	return nil, nil
}
func (p *productService) GetProductByFilter(page int, filter entity.Product) ([]entity.Product, error) {
	var filterModel model.Product
	temp, _ := json.Marshal(filter)
	json.Unmarshal(temp, &filterModel)

	products, err := p.repo.GetProducts(0, 0, filterModel)
	if err != nil {
		return nil, err
	}

	var productsModel []entity.Product
	temp, _ = json.Marshal(products)
	json.Unmarshal(temp, &productsModel)

	return productsModel, nil
}
func (p *productService) UpdateProduct(id uint, updatedProduct entity.Product) error {
	var updateModel model.Product
	temp, _ := json.Marshal(updatedProduct)
	json.Unmarshal(temp, &updateModel)

	err := p.repo.UpdateProduct(id, updateModel)
	if err != nil {
		return err
	}
	return nil

}
func (p *productService) DeleteProduct(id uint) error {
	err := p.repo.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}
