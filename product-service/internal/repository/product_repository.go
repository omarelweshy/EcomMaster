package repository

import (
	"errors"

	"product-service/internal/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func (r *ProductRepository) CreateProduct(product *model.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepository) GetProductById(id int) (*model.Product, error) {
	var product model.Product
	result := r.DB.Where("id = ?", id).First(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *ProductRepository) GetProducts() ([]model.Product, error) {
	var products []model.Product
	result := r.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
func (r *ProductRepository) UpdateProduct(id int, updatedProduct *model.Product) error {
	product, err := r.GetProductById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}
	product.Name = updatedProduct.Name
	product.Description = updatedProduct.Description
	product.Price = updatedProduct.Price
	product.Stock = updatedProduct.Stock
	product.Categories = updatedProduct.Categories
	return r.DB.Save(product).Error
}

func (r *ProductRepository) DeleteProduct(id int) error {
	product, err := r.GetProductById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}
	return r.DB.Delete(product).Error
}
