//Package api ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-02 11:51:02
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 17:46:50
 */
package api

import (
	"cmall/pkg/logging"
	"cmall/service"

	"github.com/gin-gonic/gin"
)

// ShowRanking 排行
func ShowRanking(c *gin.Context) {
	service := service.ShowRankingService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
