package service

import (
	"cmall/model"
	"cmall/serializer"
)

// CreatePictureService 商品图片创建的服务
type CreatePictureService struct {
	ProductID uint   `form:"product_id" json:"product_id"`
	ImgPath   string `form:"img_path" json:"img_path"`
}

// Create 创建商品图片
func (service *CreatePictureService) Create() serializer.Response {
	picture := model.Pictures{
		ProductID: service.ProductID,
		ImgPath:   service.ImgPath,
	}

	err := model.DB.Create(&picture).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "商品图片保存失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildPicture(picture),
	}
}
