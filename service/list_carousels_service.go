package service

import (
	"cmall/model"
	"cmall/serializer"
)

// ListCarouselsService 视频列表服务
type ListCarouselsService struct {
}

// List 视频列表
func (service *ListCarouselsService) List() serializer.Response {
	carousels := []model.Carousels{}

	if err := model.DB.Find(&carousels).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildCarousels(carousels),
	}
}
