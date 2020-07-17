package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/serializer"
)

// ShowProductService 商品详情的服务
type ShowProductService struct {
}

// Show 视频
func (service *ShowProductService) Show(id string) serializer.Response {
	var product model.Products
	code := e.SUCCESS
	err := model.DB.First(&product, id).Error
	if err != nil {
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//增加点击数
	product.AddView()

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}
}
