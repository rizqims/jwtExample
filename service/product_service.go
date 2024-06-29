package service

import (
	"apilaundry/model"
	"apilaundry/model/dto"
	"apilaundry/repository"
	"fmt"
)

type ProductService interface {
	GetbyId(id string) (model.Product, error)
	GetAll(page int, size int) ([]model.Product, dto.Paging, error)
}

type productService struct {
	repo repository.ProductRepository
}

// GetAll implmenets ProductService
func (c *productService) GetAll(page int, size int) ([]model.Product, dto.Paging, error) {
	return c.repo.GetAll(page, size)
}

// GetById implements ProductService
func (c *productService) GetbyId(id string) (model.Product, error) {
	product, err := c.repo.GetbyId(id)
	if err != nil {
		return model.Product{}, fmt.Errorf("product %v tidak ditemukan!", id)
	}
	return product, nil
}

func NewProductService(repositori repository.ProductRepository) ProductService {
	return &productService{repo: repositori}
}
