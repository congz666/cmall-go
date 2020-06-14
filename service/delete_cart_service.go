package service

import (
	"cmall/model"
	"cmall/serializer"
)

// DeleteCartService 购物车删除的服务
type DeleteCartService struct {
	UserID    uint `form:"user_id" json:"user_id"`
	ProductID uint `form:"product_id" json:"product_id"`
}

// Delete 删除购物车
func (service *DeleteCartService) Delete() serializer.Response {
	var cart model.Carts
	err := model.DB.Where("user_id=? AND product_id=?", service.UserID, service.ProductID).Find(&cart).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "查找购物车失败",
			Error:  err.Error(),
		}
	}

	err = model.DB.Delete(&cart).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "删除购物车失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{}
}
