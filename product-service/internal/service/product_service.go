package service

import (
	"github.com/omarelweshy/EcomMaster-product-service/internal/model"
	"github.com/omarelweshy/EcomMaster-product-service/internal/repository"
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
