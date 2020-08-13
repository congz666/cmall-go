//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 11:52:23
 * @LastEditors: congz
 * @LastEditTime: 2020-08-12 20:32:43
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// ListProductsService 视频列表服务
type ListProductsService struct {
	Limit      int  `form:"limit" json:"limit"`
	Start      int  `form:"start" json:"start"`
	CategoryID uint `form:"category_id" json:"category_id"`
}

// List 视频列表
func (service *ListProductsService) List() serializer.Response {
	products := []model.Product{}

	total := 0
	code := e.SUCCESS

	if service.Limit == 0 {
		service.Limit = 15
	}
	if service.CategoryID == 0 {
		if err := model.DB.Model(model.Product{}).Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&products).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		if err := model.DB.Model(model.Product{}).Where("category_id=?", service.CategoryID).Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := model.DB.Where("category_id=?", service.CategoryID).Limit(service.Limit).Offset(service.Start).Find(&products).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}
