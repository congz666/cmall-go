/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 20:04:25
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 11:34:14
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
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
	code := e.SUCCESS
	err := model.DB.Create(&picture).Error
	if err != nil {
		code := e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildPicture(picture),
	}
}
