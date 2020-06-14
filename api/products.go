package api

import (
	"cmall/service"

	"github.com/gin-gonic/gin"
)

// CreateProduct 创建商品
func CreateProduct(c *gin.Context) {
	service := service.CreateProductService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListProducts 商品列表接口
func ListProducts(c *gin.Context) {
	service := service.ListProductsService{}
	res := service.List()
	c.JSON(200, res)
}

// ShowProduct 商品详情接口
func ShowProduct(c *gin.Context) {
	service := service.ShowProductService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// ShowCategory 展示分类商品列表接口
func ShowCategory(c *gin.Context) {
	service := service.ShowCategoryService{}
	res := service.Show(c.Param("category_id"))
	c.JSON(200, res)

}

// DeleteProduct 删除商品的接口
func DeleteProduct(c *gin.Context) {
	service := service.DeleteProductService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}

// SearchProducts 搜索商品的接口
func SearchProducts(c *gin.Context) {
	service := service.SearchProductsService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
