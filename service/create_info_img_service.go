//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 20:04:25
 * @LastEditors: congz
 * @LastEditTime: 2020-07-21 23:49:45
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// CreateInfoImgService 商品详情图片创建的服务
type CreateInfoImgService struct {
	ProductID uint   `form:"product_id" json:"product_id"`
	ImgPath   string `form:"img_path" json:"img_path"`
}

// Create 创建商品详情图片
func (service *CreateInfoImgService) Create() serializer.Response {
	infoImg := model.ProductInfoImg{
		ProductID: service.ProductID,
		ImgPath:   service.ImgPath,
	}
	code := e.SUCCESS
	err := model.DB.Create(&infoImg).Error
	if err != nil {
		logging.Info(err)
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
	}
}
