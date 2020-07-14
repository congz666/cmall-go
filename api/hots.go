package api

import (
	"cmall/service"

	"github.com/gin-gonic/gin"
)

// ListEHots 家电列表接口
func ListEHots(c *gin.Context) {
	service := service.ListProductsService{}
	res := service.List()
	c.JSON(200, res)
} 
