package servicemock

import (
	"apilaundry/model"
	"apilaundry/model/dto"

	"github.com/stretchr/testify/mock"
)

type ProductServiceMock struct{
	mock.Mock
}

func (u *ProductServiceMock) GetbyId(id string)(model.Product, error){
	args := u.Called(id)
	return args.Get(0).(model.Product), args.Error(1)
}

func (u *ProductServiceMock) GetAll(page int, size int)([]model.Product, dto.Paging, error){
	args := u.Called(page, size)
	return args.Get(0).([]model.Product), args.Get(1).(dto.Paging), args.Error(2)
}