package service

import (
	"cmall/model"
	"cmall/serializer"
)

// ListEHotsService 热门家电列表服务
type ListEHotsService struct {
}

// List 视频列表
func (service *ListEHotsService) List() serializer.Response {
	products := []model.Products{}

	if err := model.DB.Find(&products).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildProducts(products),
	}
}
