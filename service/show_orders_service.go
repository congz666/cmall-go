//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-14 14:14:08
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 17:57:01
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// ShowOrdersService 订单详情的服务
type ShowOrdersService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// Show 订单
func (service *ShowOrdersService) Show(id string) serializer.Response {
	var orders []model.Orders

	total := 0
	code := e.SUCCESS
	if service.Limit == 0 {
		service.Limit = 5
	}

	if err := model.DB.Model(&orders).Where("user_id=?", id).Count(&total).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	if err := model.DB.Where("user_id=?", id).Limit(service.Limit).Offset(service.Start).Find(&orders).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildOrders(orders), uint(total))
}
