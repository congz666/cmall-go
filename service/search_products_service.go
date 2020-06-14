package service

import (
	"cmall/model"
	"cmall/serializer"
)

// SearchProductsService 搜索商品的服务
type SearchProductsService struct {
	Search string `form:"search" json:"search"`
}

// Show 搜索商品
func (service *SearchProductsService) Show() serializer.Response {
	products := []model.Products{}
	err := model.DB.Where("name LIKE ?", "%"+service.Search+"%").Find(&products).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "搜索失败",
			Error:  err.Error(),
		}
	}
	products1 := []model.Products{}
	err = model.DB.Where("info LIKE ?", "%"+service.Search+"%").Find(&products1).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "搜索失败",
			Error:  err.Error(),
		}
	}
	products = append(products, products1...)
	return serializer.Response{
		Data: serializer.BuildProducts(products),
	}
}
