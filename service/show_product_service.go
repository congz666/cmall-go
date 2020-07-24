//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 13:35:13
 * @LastEditors: congz
 * @LastEditTime: 2020-07-23 14:47:22
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

// Show 商品
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
	if product.CategoryID == 2 || product.CategoryID == 3 {
		product.AddElecRank()
	}
	if product.CategoryID == 5 || product.CategoryID == 6 || product.CategoryID == 7 || product.CategoryID == 8 {
		product.AddAcceRank()
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}
}
