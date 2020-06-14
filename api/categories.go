package api

import (
	"cmall/service"

	"github.com/gin-gonic/gin"
)

// CreateCategory 创建分类
func CreateCategory(c *gin.Context) {
	service := service.CreateCategoryService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListCategories 分类列表接口
func ListCategories(c *gin.Context) {
	service := service.ListCategoriesService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
