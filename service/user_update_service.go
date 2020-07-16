package service

import (
	"cmall/model"
	"cmall/serializer"
)

// UserUpdateService 用户修改信息的服务
type UserUpdateService struct {
	ID       uint   `form:"id" json:"id"`
	Nickname string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	Avatar   string `form:"avatar" json:"avatar"`
}

// Update 用户修改信息
func (service *UserUpdateService) Update() serializer.Response {
	var user model.User
	//找到用户
	err := model.DB.First(&user, service.ID).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "查询用户失败",
			Error:  err.Error(),
		}
	}

	user.Nickname = service.Nickname
	if service.Avatar != "" {
		user.Avatar = service.Avatar
	}
	err = model.DB.Save(&user).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "用户信息保存失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildUser(user),
	}
}
