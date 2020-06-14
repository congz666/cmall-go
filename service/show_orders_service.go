package service

import (
	"cmall/model"
	"cmall/serializer"
)

// ShowOrdersService 订单详情的服务
type ShowOrdersService struct {
}

// Show 订单
func (service *ShowOrdersService) Show(id string) serializer.Response {
	var orders []model.Orders
	err := model.DB.Where("user_id=?", id).Find(&orders).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "查找订单失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildOrders(orders),
	}
}
