//Package api ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 10:58:11
 * @LastEditors: congz
 * @LastEditTime: 2020-08-09 11:11:47
 */
package api

import (
	"cmall/pkg/logging"
	"cmall/pkg/util/sdk"
	"cmall/serializer"
	"cmall/service"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	session := sessions.Default(c)
	status := session.Get(sdk.GEETEST_SERVER_STATUS_SESSION_KEY)
	userID := session.Get("userId")
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register(userID, status)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	session := sessions.Default(c)
	status := session.Get(sdk.GEETEST_SERVER_STATUS_SESSION_KEY)
	userID := session.Get("userId")
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(userID, status)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// UserUpdate 用户修改信息
func UserUpdate(c *gin.Context) {
	var service service.UserUpdateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)

	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// CheckToken 用户详情
func CheckToken(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Status: 200,
		Msg:    "ok",
	})
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {

	c.JSON(200, serializer.Response{
		Status: 200,
		Msg:    "登出成功",
	})
}

// SendEmail 发送邮件接口
func SendEmail(c *gin.Context) {
	var service service.SendEmailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Send()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// VaildEmail 绑定和解绑邮箱接口
func VaildEmail(c *gin.Context) {
	var service service.VaildEmailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Vaild()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// InitGeetest 极验初始化
func InitGeetest(c *gin.Context) {
	gtLib := sdk.NewGeetestLib(os.Getenv("GEETEST_ID"), os.Getenv("GEETEST_KEY"))
	digestmod := "md5"
	userID := "test"
	params := map[string]string{
		"digestmod":   digestmod,
		"user_id":     userID,
		"client_type": "web",
		"ip_address":  "127.0.0.1",
	}
	result := gtLib.Register(digestmod, params)
	// 将结果状态写到session中，此处register接口存入session，后续validate接口会取出使用
	// 注意，此demo应用的session是单机模式，格外注意分布式环境下session的应用
	session := sessions.Default(c)
	session.Set(sdk.GEETEST_SERVER_STATUS_SESSION_KEY, result.Status)
	session.Set("userId", userID)
	session.Save()
	// 注意，不要更改返回的结构和值类型
	c.Header("Content-Type", "application/json;charset=UTF-8")
	c.String(http.StatusOK, result.Data)
}
