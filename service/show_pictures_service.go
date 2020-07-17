//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 20:08:41
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 11:48:52
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// ShowPicturesService 商品图片详情的服务
type ShowPicturesService struct {
}

// Show 商品图片
func (service *ShowPicturesService) Show(id string) serializer.Response {
	var pictures []model.Pictures
	code := e.SUCCESS

	err := model.DB.Where("product_id=?", id).Find(&pictures).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildPictures(pictures),
	}
}
