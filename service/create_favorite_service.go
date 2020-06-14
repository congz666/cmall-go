package service

import (
	"cmall/model"
	"cmall/serializer"
)

// CreateFavoriteService 收藏创建的服务
type CreateFavoriteService struct {
	UserID    uint `form:"user_id" json:"user_id"`
	ProductID uint `form:"product_id" json:"product_id"`
}

// Create 创建收藏夹图片
func (service *CreateFavoriteService) Create() serializer.Response {
	favorite := model.Favorites{
		UserID:    service.UserID,
		ProductID: service.ProductID,
	}
	product := model.Products{}
	err := model.DB.First(&product, service.ProductID).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "找不到该商品",
			Error:  err.Error(),
		}
	}

	err = model.DB.Create(&favorite).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "收藏保存失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildFavorite(favorite, product),
	}
}
