//Package api ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 10:58:11
 * @LastEditors: congz
 * @LastEditTime: 2020-08-12 22:06:11
 */
package api

import (
	"cmall/pkg/logging"
	"cmall/service"

	"github.com/gin-gonic/gin"
)

// InitPay 初始化支付
func InitPay(c *gin.Context) {
	service := service.InitPayService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Init()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ConfirmPay 接收FM支付回调接口
func ConfirmPay(c *gin.Context) {
	service := service.ConfirmPayService{}
	if err := c.ShouldBind(&service); err == nil {
		service.Confirm()
		c.String(200, "success")
	} else {
		c.String(200, "success")
		logging.Info(err)
	}
}
