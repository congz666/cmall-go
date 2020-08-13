//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-20 10:50:23
 * @LastEditors: congz
 * @LastEditTime: 2020-08-13 10:12:08
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// ShowCountService 展示收货地址的服务
type ShowCountService struct {
}

// Show 订单
func (service *ShowCountService) Show(id string) serializer.Response {
	code := e.SUCCESS
	var favoriteTotal int
	var notPayTotal int
	var payTotal int
	if err := model.DB.Model(model.Favorite{}).Where("user_id=?", id).Count(&favoriteTotal).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if err := model.DB.Model(model.Order{}).Where("user_id=? AND type=?", id, 1).Count(&notPayTotal).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if err := model.DB.Model(model.Order{}).Where("user_id=? AND type=?", id, 2).Count(&payTotal).Error; err != nil {
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
		Data:   serializer.BuildCount(favoriteTotal, notPayTotal, payTotal),
	}
}
