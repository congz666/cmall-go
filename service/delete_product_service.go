/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 14:30:52
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 11:36:57
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/serializer"
)

// DeleteProductService 删除商品的服务
type DeleteProductService struct {
}

// Delete 删除商品
func (service *DeleteProductService) Delete(id string) serializer.Response {
	var product model.Products
	code := e.SUCCESS

	err := model.DB.First(&product, id).Error
	if err != nil {
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err = model.DB.Delete(&product).Error
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
	}
}
