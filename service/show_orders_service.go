package service

import (
	"cmall/model"
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

	if service.Limit == 0 {
		service.Limit = 5
	}

	if err := model.DB.Model(&orders).Where("user_id=?", id).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Where("user_id=?", id).Limit(service.Limit).Offset(service.Start).Find(&orders).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildOrders(orders), uint(total))
}
