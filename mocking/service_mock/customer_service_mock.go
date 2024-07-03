package servicemock

import (
	"apilaundry/model"

	"github.com/stretchr/testify/mock"
)

type CustomerServiceMock struct{
	mock.Mock
}

func (u *CustomerServiceMock) GetbyId(id string)(model.Customer, error){
	args := u.Called(id)
	return args.Get(0).(model.Customer), args.Error(1)
}