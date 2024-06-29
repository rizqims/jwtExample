package repository

import (
	"apilaundry/model"
	"database/sql"
	"fmt"
	"time"
)

type BillRepository interface {
	Create(payload model.Bill) (model.Bill, error)
}

type billRepository struct {
	db *sql.DB
}

// GetById implements BillRepository
func (b *billRepository) Create(payload model.Bill) (model.Bill, error) {
	transaction, err := b.db.Begin()
	if err != nil {
		fmt.Println("begin err", err)
		return model.Bill{}, err
	}

	var bill model.Bill
	err = transaction.QueryRow(`INSERT INTO bills(bill_date, customer_id, user_id, created_at, updated_at) VALUES($1, $2, $3, $4, $5) RETURNING id, bill_date`,
		time.Now(), payload.Customer.Id, payload.User.Id, time.Now(), time.Now()).Scan(
		&bill.Id,
		&bill.BillDate,
	)
	bill.CreatedAt = time.Now()
	bill.UpdatedAt = time.Now()
	if err != nil {
		fmt.Println("queryrow bill err", err)
		return model.Bill{}, transaction.Rollback()
	}
	var billDetails []model.BillDetail
	// fmt.Println(payload.BillDetails)
	for _, bd := range payload.BillDetails{
		var billDetail model.BillDetail
		err = transaction.QueryRow(`INSERT INTO bill_details (bill_id, product_id, qty, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, qty, price`, bill.Id, bd.Product.Id, bd.Qty, bd.Price, time.Now(), time.Now()).Scan(
			&billDetail.Id,
			&billDetail.Qty,
			&billDetail.Price,
		)
		if err != nil {
			fmt.Println("billid",bd.BillId)
			return model.Bill{},transaction.Rollback()
		}
		billDetail.Product = bd.Product
		billDetail.CreatedAt = time.Now()
		billDetail.UpdatedAt = time.Now()
		billDetails = append(billDetails, billDetail)
	}
	bill.Customer = payload.Customer
	bill.User = payload.User
	bill.BillDetails = billDetails
	if err = transaction.Commit(); err != nil {
		fmt.Println("this get executed")
		return model.Bill{}, err
	}
	return bill, nil
}

func NewBillRepository(database *sql.DB) BillRepository {
	return &billRepository{db: database}
}
