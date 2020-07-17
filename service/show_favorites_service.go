/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-12 09:31:52
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 11:47:28
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
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
	code := e.SUCCESS

	if service.Limit == 0 {
		service.Limit = 12
	}
	if err := model.DB.Model(&favorites).Where("user_id=?", id).Count(&total).Error; err != nil {
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err := model.DB.Where("user_id=?", id).Limit(service.Limit).Offset(service.Start).Find(&favorites).Error
	if err != nil {
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildFavorites(favorites), uint(total))
}
