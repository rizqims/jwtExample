package service

import (
	"apilaundry/model"
	"apilaundry/repository"
	"fmt"
)

type CustomerService interface {
	GetbyId(id string) (model.Customer, error)
}

type customerService struct {
	repo repository.CustomerRepository
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
