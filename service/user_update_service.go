/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-12 15:25:38
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 11:54:48
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
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
	code := e.SUCCESS
	//找到用户
	err := model.DB.First(&user, service.ID).Error
	if err != nil {
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	user.Nickname = service.Nickname
	if service.Avatar != "" {
		user.Avatar = service.Avatar
	}
	err = model.DB.Save(&user).Error
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
		Data:   serializer.BuildUser(user),
	}
}
