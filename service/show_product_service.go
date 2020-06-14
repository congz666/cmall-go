package service

import (
	"cmall/model"
	"cmall/serializer"
)

// ShowProductService 商品详情的服务
type ShowProductService struct {
}

// Show 视频
func (service *ShowProductService) Show(id string) serializer.Response {
	var product model.Products
	err := model.DB.First(&product, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "商品不存在",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildProduct(product),
	}
}
