//Package api ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-21 23:27:24
 * @LastEditors: congz
 * @LastEditTime: 2020-07-21 23:50:37
 */
package api

import (
	"cmall/pkg/logging"
	"cmall/service"

	"github.com/gin-gonic/gin"
)

// CreateParamImg 创建商品参数图片
func CreateParamImg(c *gin.Context) {
	service := service.CreateParamImgService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ShowParamImgs 商品参数图片接口
func ShowParamImgs(c *gin.Context) {
	service := service.ShowParamImgsService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}
