//Package api ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-02 11:51:02
 * @LastEditors: congz
 * @LastEditTime: 2020-07-23 14:42:59
 */
package api

import (
	"cmall/service"

	"github.com/gin-gonic/gin"
)

// ListRanking 排行
func ListRanking(c *gin.Context) {
	service := service.ListRankingService{}
	res := service.List()
	c.JSON(200, res)

}

// ListElecRanking 家电排行
func ListElecRanking(c *gin.Context) {
	service := service.ListElecRankingService{}
	res := service.List()
	c.JSON(200, res)
}

// ListAcceRanking 配件排行
func ListAcceRanking(c *gin.Context) {
	service := service.ListAcceRankingService{}
	res := service.List()
	c.JSON(200, res)
}
