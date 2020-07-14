package api

import (
	"cmall/service"

	"github.com/gin-gonic/gin"
)

// CreateFavorite 创建收藏
func CreateFavorite(c *gin.Context) {
	service := service.CreateFavoriteService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowFavorites 收藏夹详情接口
func ShowFavorites(c *gin.Context) {
	service := service.ShowFavoritesService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteFavorite 删除收藏夹的接口
func DeleteFavorite(c *gin.Context) {
	service := service.DeleteFavoriteService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
