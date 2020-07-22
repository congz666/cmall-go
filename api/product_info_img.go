//Package api ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-21 23:27:34
 * @LastEditors: congz
 * @LastEditTime: 2020-07-21 23:50:50
 */
package api

import (
	"cmall/pkg/logging"
	"cmall/service"

	"github.com/gin-gonic/gin"
)

// CreateInfoImg 创建商品详情图片
func CreateInfoImg(c *gin.Context) {
	service := service.CreateInfoImgService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ShowInfoImgs 商品详情图片接口
func ShowInfoImgs(c *gin.Context) {
	service := service.ShowInfoImgsService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}
