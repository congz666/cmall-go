//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-12 15:25:38
 * @LastEditors: congz
 * @LastEditTime: 2020-08-13 11:33:17
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// UserUpdateService 用户修改信息的服务
type UserUpdateService struct {
	ID       uint   `form:"id" json:"id"`
	Nickname string `form:"nickname" json:"nickname" binding:"required,min=2,max=10"`
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=15"`
	Avatar   string `form:"avatar" json:"avatar"`
}

// Update 用户修改信息
func (service *UserUpdateService) Update() serializer.Response {
	var user model.User
	code := e.SUCCESS
	//找到用户
	err := model.DB.First(&user, service.ID).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	user.Nickname = service.Nickname
	user.UserName = service.UserName
	if service.Avatar != "" {
		user.Avatar = service.Avatar
	}
	err = model.DB.Save(&user).Error
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
		Data:   serializer.BuildUser(user),
	}
}
