//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 13:00:37
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 17:55:28
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
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

	if err := model.DB.Model(model.Products{}).Where("category_id=?", CategoryID).Count(&total).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	if err := model.DB.Where("category_id=?", CategoryID).Limit(service.Limit).Offset(service.Start).Find(&products).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}
