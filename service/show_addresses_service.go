//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-20 10:50:23
 * @LastEditors: congz
 * @LastEditTime: 2020-07-20 14:38:52
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// ShowAddressesService 展示收货地址的服务
type ShowAddressesService struct {
}

// Show 订单
func (service *ShowAddressesService) Show(id string) serializer.Response {
	var addresses []model.Address
	code := e.SUCCESS

	err := model.DB.Where("user_id=?", id).Order("created_at desc").Find(&addresses).Error
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
		Data:   serializer.BuildAddresses(addresses),
	}
}
