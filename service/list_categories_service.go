/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-12 22:28:31
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 11:38:01
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/serializer"
)

// ListCategoriesService 分类列表服务
type ListCategoriesService struct {
}

// List 视频列表
func (service *ListCategoriesService) List() serializer.Response {
	categories := []model.Categories{}
	code := e.SUCCESS

	if err := model.DB.Find(&categories).Error; err != nil {
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
		Data:   serializer.BuildCategories(categories),
	}
}
