//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-20 14:40:45
 * @LastEditors: congz
 * @LastEditTime: 2020-08-09 20:03:55
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/pkg/util"
	"cmall/serializer"
	"time"
)

// VaildEmailService 绑定、解绑邮箱和修改密码的服务
type VaildEmailService struct {
	Token string `form:"token" json:"token"`
}

// Vaild 绑定邮箱
func (service *VaildEmailService) Vaild() serializer.Response {
	var userID uint
	var email string
	var password string
	var operationType uint
	code := e.SUCCESS
	//验证token
	if service.Token == "" {
		code = e.INVALID_PARAMS
	} else {
		claims, err := util.ParseEmailToken(service.Token)
		if err != nil {
			logging.Info(err)
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		} else {
			userID = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}

	if code != e.SUCCESS {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if operationType == 1 {
		//1:绑定邮箱
		if err := model.DB.Table("user").Where("id=?", userID).Update("email", email).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else if operationType == 2 {
		//2：解绑邮箱
		if err := model.DB.Table("user").Where("id=?", userID).Update("email", "").Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	//获取该用户信息
	var user model.User
	if err := model.DB.First(&user, userID).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//3：修改密码
	if operationType == 3 {
		// 加密密码
		if err := user.SetPassword(password); err != nil {
			logging.Info(err)
			code = e.ERROR_FAIL_ENCRYPTION
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		if err := model.DB.Save(&user).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		code = e.UPDATE_PASSWORD_SUCCESS
		//返回修改密码成功信息
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//返回用户信息
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}
