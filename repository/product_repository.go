package repository

import (
	"apilaundry/model"
	"apilaundry/model/dto"
	"database/sql"
	"math"
	"fmt"
)

type ProductRepository interface {
	GetbyId(id string) (model.Product, error)
	GetAll(page int, size int) ([]model.Product, dto.Paging, error)
}

type productRepository struct {
	db *sql.DB
}

// GetAll implmenets ProductRepository
func (p *productRepository) GetAll(page int, size int) ([]model.Product, dto.Paging, error) {
	var listData []model.Product
	skip := (page-1) * size

	rows, err := p.db.Query("SELECT * FROM products LIMIT $1 OFFSET $2", size, skip)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	totalRows := 0
	err = p.db.QueryRow("SELECT COUNT(*) FROM products").Scan(&totalRows)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	for rows.Next(){
		var product model.Product

		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Type,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, dto.Paging{}, err
		}

		listData = append(listData, product)
	}

	paging := dto.Paging{
		Page: page,
		Size: size,
		TotalRows: totalRows,
		TotalPage: int(math.Ceil(float64(totalRows)/float64(size))),
	}
	return listData, paging, nil
}

// GetById implements ProductRepository
func (p *productRepository) GetbyId(id string) (model.Product, error) {

	var product model.Product
	err := p.db.QueryRow(`SELECT id, name, price, type, created_at, updated_at FROM products WHERE id=$1`, id).
		Scan(&product.Id, &product.Name, &product.Price, &product.Type, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		fmt.Println("customer not exists", product.Id, err)
		return model.Product{}, err
	}

	return product, nil
}

func NewProductRepository(database *sql.DB) ProductRepository {
	return &productRepository{db: database}
}
