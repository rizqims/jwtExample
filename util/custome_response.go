package util

import (
	"apilaundry/model/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSingleResponse(c *gin.Context, message string, data any, code int){
	c.JSON(http.StatusOK, dto.SingleResponse{
		Status: dto.Status{
			Code: code,
			Message: message,
		},
		Data: data,
	})
}

func SendPagingResponse(c *gin.Context, message string, data []any, code int, paging dto.Paging){
	c.JSON(http.StatusOK, dto.PagingResponse{
		Status: dto.Status{
			Code: code,
			Message: message,
		},
		Data: data,
		Paging: paging,
	})
}

func SendErrorResponse(c *gin.Context, message string, code int){
	c.JSON(code, dto.SingleResponse{
		Status: dto.Status{
			Code: code,
			Message: message,
		},
	})
}