package api

import (
	"cmall/service"

	"github.com/gin-gonic/gin"
)

// CreatePicture 创建商品图片
func CreatePicture(c *gin.Context) {
	service := service.CreatePictureService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowPictures 商品详情接口
func ShowPictures(c *gin.Context) {
	service := service.ShowPicturesService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}
