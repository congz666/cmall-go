package service

import (
	"cmall/model"
	"cmall/serializer"
)

// DeleteFavoriteService 删除收藏的服务
type DeleteFavoriteService struct {
	UserID    uint `form:"user_id" json:"user_id"`
	ProductID uint `form:"product_id" json:"product_id"`
}

// Delete 删除收藏
func (service *DeleteFavoriteService) Delete() serializer.Response {
	var favorite model.Favorites
	err := model.DB.Where("user_id=? AND product_id=?", service.UserID, service.ProductID).Find(&favorite).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "该收藏不存在",
			Error:  err.Error(),
		}
	}

	err = model.DB.Delete(&favorite).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "收藏删除失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{}
}
