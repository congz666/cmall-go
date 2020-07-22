//Package api ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 20:03:40
 * @LastEditors: congz
 * @LastEditTime: 2020-07-22 10:49:24
 */
package api

import (
	"cmall/pkg/logging"
	"cmall/service"

	"github.com/gin-gonic/gin"
)

// CreateProductImg 创建商品图片
func CreateProductImg(c *gin.Context) {
	service := service.CreateImgService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ShowProductImgs 商品详情接口
func ShowProductImgs(c *gin.Context) {
	service := service.ShowImgsService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}
