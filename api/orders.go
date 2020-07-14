package api

import (
	"cmall/service"

	"github.com/gin-gonic/gin"
)

// CreateOrder 创建订单
func CreateOrder(c *gin.Context) {
	service := service.CreateOrderService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowOrders 订单详情接口
func ShowOrders(c *gin.Context) {
	service := service.ShowOrdersService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
