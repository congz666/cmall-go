//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 10:58:11
 * @LastEditors: congz
 * @LastEditTime: 2020-08-13 11:32:25
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/pkg/util"
	"cmall/pkg/util/sdk"
	"cmall/serializer"
	"os"

	"github.com/jinzhu/gorm"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName  string `form:"user_name" json:"user_name" binding:"required,min=5,max=15"`
	Password  string `form:"password" json:"password" binding:"required,min=8,max=16"`
	Challenge string `form:"challenge" json:"challenge"`
	Validate  string `form:"validate" json:"validate"`
	Seccode   string `form:"seccode" json:"seccode"`
}

// Login 用户登录函数
func (service *UserLoginService) Login(userID, status interface{}) serializer.Response {
	var user model.User
	code := e.SUCCESS
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
			"ip_address":  "120.25.207.25",
		}
		result = gtLib.SuccessValidate(service.Challenge, service.Validate, service.Seccode, params)
	} else {
		result = gtLib.FailValidate(service.Challenge, service.Validate, service.Seccode)
	}
	// 注意，不要更改返回的结构和值类型
	if result.Status != 1 {
		return serializer.Response{
			Status: 404,
			Msg:    result.Msg,
		}
	}

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		//如果查询不到，返回相应错误
		if gorm.IsRecordNotFoundError(err) {
			logging.Info(err)
			code = e.ERROR_NOT_EXIST_USER
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		logging.Info(err)
		code = e.ERROR_DATABASE
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

	token, err := util.GenerateToken(service.UserName, service.Password, 0)
	if err != nil {
		logging.Info(err)
		code = e.ERROR_AUTH_TOKEN
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Status: code,
		Msg:    e.GetMsg(code),
	}

}
