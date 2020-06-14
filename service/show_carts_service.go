package service

import (
	"cmall/model"
	"cmall/serializer"
)

// ShowCartsService 订单详情的服务
type ShowCartsService struct {
}

// Show 订单
func (service *ShowCartsService) Show(id string) serializer.Response {
	var carts []model.Carts
	err := model.DB.Where("user_id=?", id).Find(&carts).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "查找购物车失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildCarts(carts),
	}
}
