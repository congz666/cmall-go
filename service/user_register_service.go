//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 10:58:11
 * @LastEditors: congz
 * @LastEditTime: 2020-08-05 10:43:09
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/pkg/util/sdk"
	"cmall/serializer"
	"os"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
	Challenge       string `form:"challenge" json:"challenge"`
	Validate        string `form:"validate" json:"validate"`
	Seccode         string `form:"seccode" json:"seccode"`
}

// Valid 验证表单
func (service *UserRegisterService) Valid(userID, status interface{}) *serializer.Response {
	var code int

	//极验SDK验证
	gtLib := sdk.NewGeetestLib(os.Getenv("GEETEST_ID"), os.Getenv("GEETEST_KEY"))
	var result *sdk.GeetestLibResult
	if status.(int) == 1 {
		/*
			   自定义参数,可选择添加
			       user_id 客户端用户的唯一标识，确定用户的唯一性；作用于提供进阶数据分析服务，可在register和validate接口传入，不传入也不影响验证服务的使用；若担心用户信息风险，可作预处理(如哈希处理)再提供到极验
				   client_type 客户端类型，web：电脑上的浏览器；h5：手机上的浏览器，包括移动应用内完全内置的web_view；native：通过原生sdk植入app应用的方式；unknown：未知
				   ip_address 客户端请求sdk服务器的ip地址
		*/
		params := map[string]string{
			"user_id":     userID.(string),
			"client_type": "web",
			"ip_address":  "127.0.0.1",
		}
		result = gtLib.SuccessValidate(service.Challenge, service.Validate, service.Seccode, params)
	} else {
		result = gtLib.FailValidate(service.Challenge, service.Validate, service.Seccode)
	}
	// 注意，不要更改返回的结构和值类型
	if result.Status != 1 {
		return &serializer.Response{
			Status: 404,
			Msg:    result.Msg,
		}
	}

	//检验密码
	if service.PasswordConfirm != service.Password {
		code = e.ERROR_NOT_COMPARE_PASSWORD
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	count := 0
	err := model.DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count).Error
	if err != nil {
		code = e.ERROR_DATABASE
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if count > 0 {
		code = e.ERROR_EXIST_NICK
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	count = 0
	err = model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count).Error
	if err != nil {
		code = e.ERROR_DATABASE
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
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
func (service *UserRegisterService) Register(userID, status interface{}) *serializer.Response {
	user := model.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
		Status:   model.Active,
	}
	code := e.SUCCESS
	// 表单验证
	if res := service.Valid(userID, status); res != nil {
		return res
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = e.ERROR_FAIL_ENCRYPTION
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	user.Avatar = "img/avatar/avatar1.jpg"

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		logging.Info(err)
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
