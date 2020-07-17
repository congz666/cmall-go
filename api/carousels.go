//Package api ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 14:09:57
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 17:48:01
 */
package api

import (
	"cmall/pkg/logging"
	"cmall/service"

	"github.com/gin-gonic/gin"
)

// CreateCarousel 创建轮播图
func CreateCarousel(c *gin.Context) {
	service := service.CreateCarouselService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ListCarousels 轮播图列表接口
func ListCarousels(c *gin.Context) {
	service := service.ListCarouselsService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
