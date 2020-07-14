package service

import (
	"cmall/model"
	"cmall/serializer"
)

// AdminLoginService 管理员登录的服务
type AdminLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// Login 管理员登录函数
func (service *AdminLoginService) Login() (model.User, *serializer.Response) {
	var admin model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&admin).Error; err != nil {
		return admin, &serializer.Response{
			Status: 40001,
			Msg:    "不存在该账号",
		}
	}

	if admin.CheckPassword(service.Password) == false {
		return admin, &serializer.Response{
			Status: 40001,
			Msg:    "账号或密码错误",
		}
	}

	if admin.Limit != 1 {
		return admin, &serializer.Response{
			Status: 40001,
			Msg:    "权限不足",
		}
	}
	return admin, nil
}
