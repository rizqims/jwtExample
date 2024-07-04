package controller

import (
	"apilaundry/middleware"
	"apilaundry/model/dto"
	"apilaundry/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BillController struct {
	service service.BillService
	rg      *gin.RouterGroup
	middleware middleware.AuthMiddleware
}

func (b *BillController) CreateHandler(c *gin.Context) {
	var payload dto.BillRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response, errs := b.service.CreateNewBill(payload)
	if errs != nil {
		c.JSON(http.StatusBadRequest, errs.Error())
	}
	c.JSON(http.StatusCreated, response)
}

func (b *BillController) Route() {
	group := b.rg.Group("/transactions", b.middleware.CheckToken("admin"))
	group.POST("/", b.CreateHandler)
}

func NewBillController(service service.BillService, rg *gin.RouterGroup, aM middleware.AuthMiddleware) *BillController {
	return &BillController{service: service, rg: rg, middleware: aM}
}