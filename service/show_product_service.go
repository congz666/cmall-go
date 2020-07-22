//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 13:35:13
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 17:57:14
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// ShowProductService 商品详情的服务
type ShowProductService struct {
}

// Show 视频
func (service *ShowProductService) Show(id string) serializer.Response {
	var product model.Product
	code := e.SUCCESS
	err := model.DB.First(&product, id).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//增加点击数
	product.AddView()

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}
}
