package service

import (
	repomock "apilaundry/mocking/repo_mock"
	servicemock "apilaundry/mocking/service_mock"
	"apilaundry/model"
	"apilaundry/model/dto"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)


type BillServiceTestSuite struct {
	suite.Suite
	repoBillMock        *repomock.BillRepoMock
	userServiceMock     *servicemock.UserServiceMock
	productServiceMock  *servicemock.ProductServiceMock
	customerServiceMock *servicemock.CustomerServiceMock
	bS                  BillService
}

func (suite *BillServiceTestSuite) SetupTest() {
	suite.repoBillMock = new(repomock.BillRepoMock)
	suite.userServiceMock = new(servicemock.UserServiceMock)
	suite.productServiceMock = new(servicemock.ProductServiceMock)
	suite.customerServiceMock = new(servicemock.CustomerServiceMock)
	suite.bS = NewBillService(suite.repoBillMock, suite.userServiceMock, suite.productServiceMock, suite.customerServiceMock)
}

func TestBillServiceTestSuite(t *testing.T) {
	suite.Run(t, new(BillServiceTestSuite))
}

var mockingBill = model.Bill{
	Id:       "1",
	BillDate: time.Now(),
	Customer: model.Customer{
		Id:   "1",
		Name: "Arfian",
	},
	User: model.User{
		Id:   "2",
		Name: "Dimas",
	},
	BillDetails: []model.BillDetail{
		{
			Id:     "1",
			BillId: "1",
			Product: model.Product{
				Id: "1",
				Price: 1,
			},
			Qty:   1,
			Price: 1,
		},
	},
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var mockPayload = dto.BillRequest{
	Customer: "1",
	User:     "1",
	BillDetails: []model.BillDetail{
		{
			Product: model.Product{
				Id: "1",
			},
			Qty: 1,
		},
	},
}

func (suite *BillServiceTestSuite) TestCreateNewBill_Success() {
	suite.customerServiceMock.On("GetbyId", mockPayload.Customer).Return(mockingBill.Customer, nil)
	suite.userServiceMock.On("GetbyId", mockPayload.User).Return(mockingBill.User, nil)

	var billDetails []model.BillDetail
	for _, v := range mockPayload.BillDetails {
		suite.productServiceMock.On("GetbyId", v.Product.Id).Return(mockingBill.BillDetails[0].Product, nil)
		billDetails = append(billDetails, model.BillDetail{
			Product: mockingBill.BillDetails[0].Product,
			Qty:     v.Qty,
			Price:   mockingBill.BillDetails[0].Price,
		})
	}

	mockBillayload := model.Bill{
		Customer:    mockingBill.Customer,
		User:        mockingBill.User,
		BillDetails: billDetails,
	}

	suite.repoBillMock.On("Create", mockBillayload).Return(mockingBill, nil)

	_, err := suite.bS.CreateNewBill(mockPayload)
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)

}
