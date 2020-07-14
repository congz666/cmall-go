package service

import (
	"cmall/model"
	"cmall/serializer"
)

// ShowFavoritesService 展示收藏夹详情的服务
type ShowFavoritesService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// Show 商品图片
func (service *ShowFavoritesService) Show(id string) serializer.Response {
	var favorites []model.Favorites
	total := 0

	if service.Limit == 0 {
		service.Limit = 12
	}
	if err := model.DB.Model(&favorites).Where("user_id=?", id).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	err := model.DB.Where("user_id=?", id).Limit(service.Limit).Offset(service.Start).Find(&favorites).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "收藏夹获取失败",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildFavorites(favorites), uint(total))
}
