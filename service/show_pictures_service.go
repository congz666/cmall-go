package service

import (
	"cmall/model"
	"cmall/serializer"
)

// ShowPicturesService 商品图片详情的服务
type ShowPicturesService struct {
}

// Show 商品图片
func (service *ShowPicturesService) Show(id string) serializer.Response {
	var pictures []model.Pictures
	err := model.DB.Where("product_id=?", id).Find(&pictures).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "商品不存在",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildPictures(pictures),
	}
}
