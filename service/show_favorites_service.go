package service

import (
	"cmall/model"
	"cmall/serializer"
)

// ShowFavoritesService 商品图片详情的服务
type ShowFavoritesService struct {
}

// Show 商品图片
func (service *ShowFavoritesService) Show(id string) serializer.Response {
	var favorites []model.Favorites
	err := model.DB.Where("user_id=?", id).Find(&favorites).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "收藏夹获取失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildFavorites(favorites),
	}
}
