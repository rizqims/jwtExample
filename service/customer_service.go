package service

import (
	"apilaundry/model"
	"apilaundry/repository"
	"fmt"
)

type CustomerService interface {
	GetbyId(id string) (model.Customer, error)
	GetAll(page int, size int) ([]model.Customer, error)
}

type customerService struct {
	repo repository.CustomerRepository
}

// GetAll implmenets CustomerService
func (c *customerService) GetAll(page int, size int) ([]model.Customer, error) {
	panic("unimplemented")
}

// GetById implements CustomerService
func (c *customerService) GetbyId(id string) (model.Customer, error) {
	customer, err := c.repo.GetbyId(id)
	if err != nil {
		return model.Customer{}, fmt.Errorf("customer %v tidak ditemukan!", id)
	}
	return customer, nil
}

func NewCustomerService(repositori repository.CustomerRepository) CustomerService {
	return &customerService{repo: repositori}
}
