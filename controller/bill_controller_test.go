package controller

// go test ./... -v -coverprofile cover.out; go tool cover -html cover.out

import (
	"apilaundry/mocking"
	servicemock "apilaundry/mocking/service_mock"
	"apilaundry/model"
	"apilaundry/model/dto"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BillControllerTestSuite struct {
	suite.Suite
	serviceBillMock *servicemock.BillServiceMock
	rg              *gin.RouterGroup
	middlewareMock  *mocking.AuthMiddlewareMock
}

func (suite *BillControllerTestSuite) SetupTest() {
	suite.serviceBillMock = new(servicemock.BillServiceMock)
	rg := gin.Default()
	suite.rg = rg.Group("api/v1/transactions")
	suite.middlewareMock = new(mocking.AuthMiddlewareMock)
}

func TestBillControllerTestSuite(t *testing.T) {
	suite.Run(t, new(BillControllerTestSuite))
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
				Id:    "1",
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

func (suite *BillControllerTestSuite) TestCreatedHandler_Success() {
	record := httptest.NewRecorder()

	mockPayloadjson, err := json.Marshal(mockPayload)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(mockPayloadjson))
	assert.NoError(suite.T(), err)

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJFbmlnbWFDYW1wIiwiZXhwIjoxNzE5OTc1MTQ4LCJpYXQiOjE3MTk5NzE1NDgsInVzZXJJZCI6ImZmZDBhN2RkLWNlZDctNGM0MS1iYTFkLWZlYTI5OTAwYzQzMiIsInJvbGUiOiJhZG1pbiJ9.PXSDCw26EXmmFPJeXj1xXebKTlgZm6QD21C5ZzH4DL4"
	req.Header.Set("Authorization", "Bearer"+token)
	c, _ := gin.CreateTestContext(record)
	c.Request = req

	suite.serviceBillMock.On("CreateNewBill", mockPayload).Return(mockingBill, nil)

	billController := NewBillController(suite.serviceBillMock, suite.rg, suite.middlewareMock)
	billController.Route()
	billController.CreateHandler(c)

	mockBilljson, _ := json.Marshal(mockingBill)
	assert.Equal(suite.T(), http.StatusCreated, record.Code)
	assert.Equal(suite.T(), string(mockBilljson), record.Body.String())
}

func (suite *BillControllerTestSuite) TestCreatedHandler_FailedBinding(){
	billController := NewBillController(suite.serviceBillMock, suite.rg, suite.middlewareMock)
	billController.Route()
	record := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/",nil)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJFbmlnbWFDYW1wIiwiZXhwIjoxNzE5OTc1MTQ4LCJpYXQiOjE3MTk5NzE1NDgsInVzZXJJZCI6ImZmZDBhN2RkLWNlZDctNGM0MS1iYTFkLWZlYTI5OTAwYzQzMiIsInJvbGUiOiJhZG1pbiJ9.PXSDCw26EXmmFPJeXj1xXebKTlgZm6QD21C5ZzH4DL4"

	req.Header.Set("Authorization", "Bearer"+token)
	c, _ := gin.CreateTestContext(record)
	c.Request = req

	billController.CreateHandler(c)
	assert.Equal(suite.T(), http.StatusBadRequest, record.Code)

	// suite.repoBillMock.On("Create", mockBillayload).Return(model.Bill{}, errors.New("error"))
	// _, err := suite.bS.CreateNewBill(mockPayload)
	// assert.Error(suite.T(), err)
}

// func (suite *BillControllerTestSuite) TestCreatedHandler_FailedNganu(){
// 	billController := NewBillController(suite.serviceBillMock, suite.rg, suite.middlewareMock)
// 	billController.Route()
// 	record := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodPost, "/",nil)
// 	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJFbmlnbWFDYW1wIiwiZXhwIjoxNzE5OTc1MTQ4LCJpYXQiOjE3MTk5NzE1NDgsInVzZXJJZCI6ImZmZDBhN2RkLWNlZDctNGM0MS1iYTFkLWZlYTI5OTAwYzQzMiIsInJvbGUiOiJhZG1pbiJ9.PXSDCw26EXmmFPJeXj1xXebKTlgZm6QD21C5ZzH4DL4"

// 	req.Header.Set("Authorization", "Bearer"+token)
// 	c, _ := gin.CreateTestContext(record)
// 	c.Request = req

// 	billController.CreateHandler(c)
// 	assert.Equal(suite.T(), http.StatusCreated, record.Code)
// 	assert.Equal(suite.T(), http.StatusBadRequest, record.Code)
// }