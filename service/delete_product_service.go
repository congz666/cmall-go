package service

import (
	"cmall/model"
	"cmall/serializer"
)

// DeleteProductService 删除商品的服务
type DeleteProductService struct {
}

// Delete 删除商品
func (service *DeleteProductService) Delete(id string) serializer.Response {
	var product model.Products
	err := model.DB.First(&product, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "商品不存在",
			Error:  err.Error(),
		}
	}

	err = model.DB.Delete(&product).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "商品删除失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{}
}
