package servicemock

import (
	"apilaundry/model"
	"apilaundry/model/dto"

	"github.com/stretchr/testify/mock"
)

type BillServiceMock struct{
	mock.Mock
}

func (u *BillServiceMock) CreateNewBill(payload dto.BillRequest)(model.Bill, error){
	args := u.Called(payload)
	return args.Get(0).(model.Bill), args.Error(1)
	
}