package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/util"
	"cmall/serializer"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// Login 用户登录函数
func (service *UserLoginService) Login() serializer.Response {
	var user model.User
	var code int

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		code = e.ERROR_NOT_EXIST_USER
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if user.CheckPassword(service.Password) == false {
		code = e.ERROR_NOT_COMPARE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	token, err := util.GenerateToken(service.UserName, service.Password)
	if err != nil {
		code = e.ERROR_AUTH_TOKEN
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	code = e.SUCCESS
	return serializer.Response{
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Status: code,
		Msg:    e.GetMsg(code),
	}

}
