package service

import (
	"cmall/model"
	"cmall/serializer"
)

// UpdateCartService 购物车修改的服务
type UpdateCartService struct {
	UserID    uint `form:"user_id" json:"user_id"`
	ProductID uint `form:"product_id" json:"product_id"`
	Num       uint `form:"num" json:"num"`
}

// Update 修改购物车信息
func (service *UpdateCartService) Update() serializer.Response {
	var cart model.Carts
	err := model.DB.Where("user_id=? AND product_id=?", service.UserID, service.ProductID).Find(&cart).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "查找购物车失败",
			Error:  err.Error(),
		}
	}
	cart.Num = service.Num
	err = model.DB.Save(&cart).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "更新购物车失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{}
}
