//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-17 14:45:17
 * @LastEditors: congz
 * @LastEditTime: 2020-07-18 14:31:47
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/pkg/util"
	"cmall/serializer"

	"github.com/jinzhu/gorm"
)

// AdminLoginService 管理员登录的服务
type AdminLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// Login 管理员登录函数
func (service *AdminLoginService) Login() serializer.Response {
	var admin model.Admin
	code := e.SUCCESS

	if err := model.DB.Where("user_name = ?", service.UserName).First(&admin).Error; err != nil {
		//如果查询不到，返回相应错误
		if gorm.IsRecordNotFoundError(err) {
			logging.Info(err)
			code = e.ERROR_NOT_EXIST_USER
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		code = e.ERROR_NOT_EXIST_USER
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if admin.CheckPassword(service.Password) == false {
		code = e.ERROR_NOT_COMPARE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	token, err := util.GenerateToken(service.UserName, service.Password, 1)
	if err != nil {
		logging.Info(err)
		code = e.ERROR_AUTH_TOKEN
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Data:   serializer.TokenData{User: serializer.BuildAdmin(admin), Token: token},
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
