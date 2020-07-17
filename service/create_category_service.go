/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-12 22:16:18
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 11:29:26
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/serializer"
)

// CreateCategoryService 收藏创建的服务
type CreateCategoryService struct {
	CategoryID   uint   `form:"category_id" json:"category_id"`
	CategoryName string `form:"category_name" json:"category_name"`
}

// Create 创建分类
func (service *CreateCategoryService) Create() serializer.Response {
	category := model.Categories{
		CategoryID:   service.CategoryID,
		CategoryName: service.CategoryName,
	}
	code := e.SUCCESS
	err := model.DB.Create(&category).Error
	if err != nil {
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
		Data:   serializer.BuildCategory(category),
	}
}
