/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-16 22:30:42
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 11:52:26
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/serializer"
)

// UpdateProductService 商品更新的服务
type UpdateProductService struct {
	ID            uint   `form:"id" json:"id"`
	Name          string `form:"name" json:"name"`
	CategoryID    int    `form:"category_id" json:"category_id"`
	Title         string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info          string `form:"info" json:"info" binding:"max=1000"`
	ImgPath       string `form:"img_path" json:"img_path"`
	Price         string `form:"price" json:"price"`
	DiscountPrice string `form:"discount_price" json:"discount_price"`
}

// Update 更新商品
func (service *UpdateProductService) Update() serializer.Response {
	product := model.Products{
		Name:          service.Name,
		CategoryID:    service.CategoryID,
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       service.ImgPath,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
	}
	product.ID = service.ID
	code := e.SUCCESS
	err := model.DB.Save(&product).Error
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
