//Package api ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-12 23:08:23
 * @LastEditors: congz
 * @LastEditTime: 2020-08-04 11:22:48
 */
package api

import (
	"cmall/pkg/logging"
	"cmall/service"

	"github.com/gin-gonic/gin"
)

//ShowNotice 公告详情
func ShowNotice(c *gin.Context) {
	service := service.ShowNoticeService{}
	res := service.Show()
	c.JSON(200, res)
}

//CreateNotice 创建公告
func CreateNotice(c *gin.Context) {
	service := service.CreateNoticeService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//UpdateNotice 更新公告
func UpdateNotice(c *gin.Context) {
	service := service.UpdateNoticeService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
