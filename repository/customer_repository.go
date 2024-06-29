package repository

import (
	"apilaundry/model"
	"database/sql"
	"fmt"
)

type CustomerRepository interface {
	GetbyId(id string) (model.Customer, error)
	GetAll(page int, size int) ([]model.Customer, error)
}

type customerRepository struct {
	db *sql.DB
}

// GetAll implmenets CustomerRepository
func (p *customerRepository) GetAll(page int, size int) ([]model.Customer, error) {
	panic("unimplemented")
}

// GetById implements CustomerRepository
func (p *customerRepository) GetbyId(id string) (model.Customer, error) {

	var customer model.Customer
	err := p.db.QueryRow(`SELECT id, name, phone_number, address, created_at, updated_at FROM customers WHERE id=$1`, id).
		Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		fmt.Println("customer not exists", customer.Id, err)
		return model.Customer{}, err
	}

	return customer, nil
}

func NewCustomerRepository(database *sql.DB) CustomerRepository {
	return &customerRepository{db: database}
}
