//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-14 17:04:28
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 17:57:31
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
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
	code := e.SUCCESS

	err := model.DB.Where("user_id=? AND product_id=?", service.UserID, service.ProductID).Find(&cart).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	cart.Num = service.Num
	err = model.DB.Save(&cart).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
