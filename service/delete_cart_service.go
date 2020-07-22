//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-14 17:28:09
 * @LastEditors: congz
 * @LastEditTime: 2020-07-22 11:01:30
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// DeleteCartService 购物车删除的服务
type DeleteCartService struct {
	UserID    uint `form:"user_id" json:"user_id"`
	ProductID uint `form:"product_id" json:"product_id"`
}

// Delete 删除购物车
func (service *DeleteCartService) Delete() serializer.Response {
	var cart model.Cart
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

	err = model.DB.Delete(&cart).Error
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
