//Package api ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-12 23:08:23
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 17:46:42
 */
package api

import (
	"cmall/service"

	"github.com/gin-gonic/gin"
)

// ReadMe 发送ReadMe到前端
func ReadMe(c *gin.Context) {
	service := service.ReadMeService{}
	res := service.Read()
	c.JSON(200, res)
}
