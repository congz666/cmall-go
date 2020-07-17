/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-14 10:47:54
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 11:45:18
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/serializer"
)

// SearchProductsService 搜索商品的服务
type SearchProductsService struct {
	Search string `form:"search" json:"search"`
}

// Show 搜索商品
func (service *SearchProductsService) Show() serializer.Response {
	products := []model.Products{}
	code := e.SUCCESS

	err := model.DB.Where("name LIKE ?", "%"+service.Search+"%").Find(&products).Error
	if err != nil {
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	products1 := []model.Products{}
	err = model.DB.Where("info LIKE ?", "%"+service.Search+"%").Find(&products1).Error
	if err != nil {
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	products = append(products, products1...)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProducts(products),
	}
}
