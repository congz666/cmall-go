//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-14 14:14:08
 * @LastEditors: congz
 * @LastEditTime: 2020-08-12 19:39:57
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// ListOrdersService 订单详情的服务
type ListOrdersService struct {
	Limit int  `form:"limit" json:"limit"`
	Start int  `form:"start" json:"start"`
	Type  uint `form:"type" json:"type" `
}

// List 订单
func (service *ListOrdersService) List(id string) serializer.Response {
	var orders []model.Order

	total := 0
	code := e.SUCCESS
	if service.Limit == 0 {
		service.Limit = 5
	}

	if service.Type == 0 {
		if err := model.DB.Model(&orders).Where("user_id=?", id).Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := model.DB.Where("user_id=?", id).Limit(service.Limit).Offset(service.Start).Order("created_at desc").Find(&orders).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		if err := model.DB.Model(&orders).Where("user_id=? AND type=?", id, service.Type).Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := model.DB.Where("user_id=? AND type=?", id, service.Type).Limit(service.Limit).Offset(service.Start).Order("created_at desc").Find(&orders).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	return serializer.BuildListResponse(serializer.BuildOrders(orders), uint(total))
}
