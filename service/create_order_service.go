/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-14 13:22:30
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 11:33:43
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/serializer"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// CreateOrderService 订单创建的服务
type CreateOrderService struct {
	UserID    uint `form:"user_id" json:"user_id"`
	ProductID uint `form:"product_id" json:"product_id"`
	Num       uint `form:"num" json:"num"`
}

// Create 创建订单
func (service *CreateOrderService) Create() serializer.Response {
	order := model.Orders{
		UserID:    service.UserID,
		ProductID: service.ProductID,
		Num:       service.Num,
	}
	product := model.Products{}
	code := e.SUCCESS

	err := model.DB.First(&product, service.ProductID).Error
	if err != nil {
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//生成随机订单号
	orderNum := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000000))
	productNum := strconv.Itoa(int(service.ProductID))
	userNum := strconv.Itoa(int(service.UserID))
	orderNum = orderNum + productNum + userNum
	orderID, err := strconv.ParseUint(orderNum, 10, 64)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	order.OrderID = orderID
	err = model.DB.Create(&order).Error
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
		Data:   serializer.BuildOrder(order, product),
	}
}
