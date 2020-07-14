package service

import (
	"cmall/model"
	"cmall/serializer"
)

// ListProductsService 视频列表服务
type ListProductsService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// List 视频列表
func (service *ListProductsService) List() serializer.Response {
	products := []model.Products{}

	total := 0

	if service.Limit == 0 {
		service.Limit = 15
	}

	if err := model.DB.Model(model.Products{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&products).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}
