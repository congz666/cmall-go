package service

import (
	"cmall/model"
	"cmall/serializer"
)

// CreateCarouselService 轮播图创建的服务
type CreateCarouselService struct {
	ImgPath string `form:"img_path" json:"img_path"`
}

// Create 创建商品
func (service *CreateCarouselService) Create() serializer.Response {
	carousel := model.Carousels{
		ImgPath: service.ImgPath,
	}

	err := model.DB.Create(&carousel).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "商品保存失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildCarousel(carousel),
	}
}
