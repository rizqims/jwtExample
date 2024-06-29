package service

import (
	"apilaundry/model"
	"apilaundry/model/dto"
	"apilaundry/repository"
)

type BillService interface {
	CreateNewBill(payload dto.BillRequest) (model.Bill, error)
}

type billService struct {
	repo            repository.BillRepository
	userService     UserService
	productService  ProductService
	customerService CustomerService
}

func NewBillService(repo repository.BillRepository, uS UserService, pS ProductService, cS CustomerService) BillService {
	return &billService{
		repo:            repo,
		userService:     uS,
		productService:  pS,
		customerService: cS,
	}
}

func (b *billService) CreateNewBill(payload dto.BillRequest) (model.Bill, error) {
	customer, err := b.customerService.GetbyId(payload.Customer)
	if err != nil {
		return model.Bill{}, err
	}
	user, err := b.userService.GetbyId(payload.User)
	if err != nil {
		return model.Bill{}, err
	}

	var billDetails []model.BillDetail
	for _, bd := range payload.BillDetails{
		product, err := b.productService.GetbyId(bd.Product.Id)
		if err != nil {
			return model.Bill{}, err
		}
		billDetails = append(billDetails, model.BillDetail{Product: product, Qty: bd.Qty, Price: product.Price})
	}
	newPayload := model.Bill{
		Customer: customer,
		User: user,
		BillDetails: billDetails,
	}
	bill, err := b.repo.Create(newPayload)
	if err != nil {
		return model.Bill{}, err
	}
	return bill, nil
}