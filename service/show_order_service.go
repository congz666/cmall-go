//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-14 14:14:08
 * @LastEditors: congz
 * @LastEditTime: 2020-07-22 14:36:08
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"

	"github.com/jinzhu/gorm"
)

// ShowOrderService 订单详情的服务
type ShowOrderService struct {
}

// Show 订单
func (service *ShowOrderService) Show(id string) serializer.Response {
	var order model.Order
	var product model.Product
	var address model.Address
	code := e.SUCCESS
	//根据id查找order
	if err := model.DB.Where("id=?", id).First(&order).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//根据order查找product
	if err := model.DB.Where("id=?", order.ProductID).First(&product).Error; err != nil {
		//如果查询不到，返回相应错误
		if gorm.IsRecordNotFoundError(err) {
			logging.Info(err)
			code = e.ERROR_NOT_EXIST_PRODUCT
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//根据order查找address
	if err := model.DB.Where("id=?", order.AddressID).First(&address).Error; err != nil {
		//如果查询不到，返回相应错误
		if gorm.IsRecordNotFoundError(err) {
			logging.Info(err)
			code = e.ERROR_NOT_EXIST_ADDRESS
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data: serializer.OrderDetails{
			Order:   serializer.BuildOrder(order, product),
			Address: serializer.BuildAddress(address),
		},
	}
}
