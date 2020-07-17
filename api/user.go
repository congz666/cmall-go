package api

import (
	"cmall/serializer"
	"cmall/service"

	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
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
