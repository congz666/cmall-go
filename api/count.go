/*Package api ...
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-08-13 09:44:41
 * @LastEditors: congz
 * @LastEditTime: 2020-08-13 09:45:02
 */
package api

import (
	"cmall/service"

	"github.com/gin-gonic/gin"
)

// ShowCount 获取数量服务
func ShowCount(c *gin.Context) {
	service := service.ShowCountService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}
