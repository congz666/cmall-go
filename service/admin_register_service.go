/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-15 15:47:27
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 11:30:17
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/serializer"
)

// AdminRegisterService 管理用户注册服务
type AdminRegisterService struct {
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// Valid 验证表单
func (service *AdminRegisterService) Valid() *serializer.Response {
	var code int
	if service.PasswordConfirm != service.Password {
		code = e.ERROR_NOT_COMPARE_PASSWORD
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	count := 0
	model.DB.Model(&model.Admin{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		code = e.ERROR_EXIST_USER
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return nil
}

// Register 用户注册
func (service *AdminRegisterService) Register() *serializer.Response {
	admin := model.Admin{
		UserName: service.UserName,
	}
	code := e.SUCCESS
	// 表单验证
	if res := service.Valid(); res != nil {
		return res
	}

	// 加密密码
	if err := admin.SetPassword(service.Password); err != nil {
		code = e.ERROR_FAIL_ENCRYPTION
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 创建用户
	if err := model.DB.Create(&admin).Error; err != nil {
		code = e.ERROR_DATABASE
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
