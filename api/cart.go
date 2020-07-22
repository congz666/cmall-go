//Package api ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-14 15:29:54
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 17:47:50
 */
package api

import (
	"cmall/pkg/logging"
	"cmall/service"

	"github.com/gin-gonic/gin"
)

// CreateCart 加入购物车
func CreateCart(c *gin.Context) {
	service := service.CreateCartService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ShowCarts 购物车详情接口
func ShowCarts(c *gin.Context) {
	service := service.ShowCartsService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// UpdateCart 修改购物车信息
func UpdateCart(c *gin.Context) {
	service := service.UpdateCartService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// DeleteCart 删除购物车
func DeleteCart(c *gin.Context) {
	service := service.DeleteCartService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
