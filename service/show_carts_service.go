/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-14 15:48:10
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 11:46:37
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/serializer"
)

// ShowCartsService 订单详情的服务
type ShowCartsService struct {
}

// Show 订单
func (service *ShowCartsService) Show(id string) serializer.Response {
	var carts []model.Carts
	code := e.SUCCESS

	err := model.DB.Where("user_id=?", id).Find(&carts).Error
	if err != nil {
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
		Data:   serializer.BuildCarts(carts),
	}
}
