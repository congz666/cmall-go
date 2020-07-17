/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 14:21:05
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 11:37:25
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/serializer"
)

// ListCarouselsService 视频列表服务
type ListCarouselsService struct {
}

// List 视频列表
func (service *ListCarouselsService) List() serializer.Response {
	carousels := []model.Carousels{}
	code := e.SUCCESS

	if err := model.DB.Find(&carousels).Error; err != nil {
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCarousels(carousels),
	}
}
