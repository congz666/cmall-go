//Package api ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-14 13:07:14
 * @LastEditors: congz
 * @LastEditTime: 2020-08-12 16:01:44
 */
package api

import (
	"cmall/pkg/logging"
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
		logging.Info(err)
	}
}

// ListOrders 订单详情接口
func ListOrders(c *gin.Context) {
	service := service.ListOrdersService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ShowOrder 订单详情详情接口
func ShowOrder(c *gin.Context) {
	service := service.ShowOrderService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Param("num"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
