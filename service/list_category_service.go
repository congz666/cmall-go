/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 13:00:37
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 11:41:48
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/serializer"
	"strconv"
)

// ShowCategoryService 视频列表服务
type ShowCategoryService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// Show 分类列表
func (service *ShowCategoryService) Show(CategoryID string) serializer.Response {
	products := []model.Products{}

	total := 0
	code := e.SUCCESS
	if service.Limit == 0 {
		service.Limit = 15
	}

	id, err := strconv.Atoi(CategoryID)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	if err := model.DB.Model(model.Products{}).Where("category_id=?", id).Count(&total).Error; err != nil {
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	if err := model.DB.Where("category_id=?", id).Limit(service.Limit).Offset(service.Start).Find(&products).Error; err != nil {
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}
