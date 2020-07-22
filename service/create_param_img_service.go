//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 20:04:25
 * @LastEditors: congz
 * @LastEditTime: 2020-07-21 23:50:16
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// CreateParamImgService 商品参数图片创建的服务
type CreateParamImgService struct {
	ProductID uint   `form:"product_id" json:"product_id"`
	ImgPath   string `form:"img_path" json:"img_path"`
}

// Create 创建商品参数图片
func (service *CreateParamImgService) Create() serializer.Response {
	paramImg := model.ProductParamImg{
		ProductID: service.ProductID,
		ImgPath:   service.ImgPath,
	}
	code := e.SUCCESS
	err := model.DB.Create(&paramImg).Error
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
