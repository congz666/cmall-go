//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 11:10:13
 * @LastEditors: congz
 * @LastEditTime: 2020-07-22 11:01:23
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// CreateProductService 商品创建的服务
type CreateProductService struct {
	Name          string `form:"name" json:"name"`
	CategoryID    int    `form:"category_id" json:"category_id"`
	Title         string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info          string `form:"info" json:"info" binding:"max=1000"`
	ImgPath       string `form:"img_path" json:"img_path"`
	Price         string `form:"price" json:"price"`
	DiscountPrice string `form:"discount_price" json:"discount_price"`
}

// Create 创建商品
func (service *CreateProductService) Create() serializer.Response {
	product := model.Product{
		Name:          service.Name,
		CategoryID:    service.CategoryID,
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       service.ImgPath,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
	}
	code := e.SUCCESS

	err := model.DB.Create(&product).Error
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
		Data:   serializer.BuildProduct(product),
	}
}
