package repomock

import (
	"apilaundry/model"
	"apilaundry/model/dto"

	"github.com/stretchr/testify/mock"
)

type ProductRepoMock struct{
	mock.Mock
}

func (u *ProductRepoMock) GetById(id string)(model.User, error){
	args := u.Called(id)
	return args.Get(0).(model.User), args.Error(1)
	
}

func (u *ProductRepoMock) GetAll(page int, size int)([]model.Product, dto.Paging, error){
	args := u.Called(page, size)
	return args.Get(0).([]model.Product), args.Get(1).(dto.Paging), args.Error(2)
}