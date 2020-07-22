//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-12 09:03:27
 * @LastEditors: congz
 * @LastEditTime: 2020-07-22 11:01:11
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// CreateFavoriteService 收藏创建的服务
type CreateFavoriteService struct {
	UserID    uint `form:"user_id" json:"user_id"`
	ProductID uint `form:"product_id" json:"product_id"`
}

// Create 创建收藏夹图片
func (service *CreateFavoriteService) Create() serializer.Response {
	favorite := model.Favorite{
		UserID:    service.UserID,
		ProductID: service.ProductID,
	}
	product := model.Product{}
	code := e.SUCCESS

	err := model.DB.First(&product, service.ProductID).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err = model.DB.Create(&favorite).Error
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
		Data:   serializer.BuildFavorite(favorite, product),
	}
}
