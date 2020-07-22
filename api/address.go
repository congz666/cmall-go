//Package api ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-20 10:32:12
 * @LastEditors: congz
 * @LastEditTime: 2020-07-20 20:43:04
 */
package api

import (
	"cmall/pkg/logging"
	"cmall/service"

	"github.com/gin-gonic/gin"
)

// CreateAddress 新建收货地址
func CreateAddress(c *gin.Context) {
	service := service.CreateAddressService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ShowAddresses 展示收货地址
func ShowAddresses(c *gin.Context) {
	service := service.ShowAddressesService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// UpdateAddress 修改收货地址
func UpdateAddress(c *gin.Context) {
	service := service.UpdateAddressService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// DeleteAddress 删除收货地址
func DeleteAddress(c *gin.Context) {
	service := service.DeleteAddressService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
