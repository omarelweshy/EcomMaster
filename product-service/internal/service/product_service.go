package service

import (
	"product-service/internal/model"
	"product-service/internal/repository"
)

type ProductService struct {
	Repo repository.ProductRepository
}

func (s *ProductService) CreateProduct(name string, description string, price float64, stock int) error {
	product := model.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
		// Categories:  categories,
	}
	return s.Repo.CreateProduct(&product)
}

func (s *ProductService) GetProducts() ([]model.Product, error) {
	products, err := s.Repo.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}
