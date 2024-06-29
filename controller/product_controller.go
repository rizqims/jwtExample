package controller

import (
	"apilaundry/service"
	"apilaundry/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service service.ProductService
	rg *gin.RouterGroup
}

func  (p *ProductController) GetAllHandler(c *gin.Context){
	page, err := strconv.Atoi(c.DefaultQuery("page","1"))
	size, err2 := strconv.Atoi(c.DefaultQuery("size","10"))
	if err != nil {
		util.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
	}
	if err2 != nil {
		util.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
	}

	data, paging, err := p.service.GetAll(page, size)
	if err != nil {
		util.SendErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}

	var listData []interface{}
	for _, pl := range data{
		listData = append(listData, pl)
	}

	util.SendPagingResponse(c,"success get data", listData, http.StatusOK, paging)
}

func (p *ProductController) Route(){
	rg := p.rg.Group("/products")
	rg.GET("/", p.GetAllHandler)
}

func NewProductController(service service.ProductService, rg *gin.RouterGroup) *ProductController{
	return &ProductController{service: service, rg: rg}
}