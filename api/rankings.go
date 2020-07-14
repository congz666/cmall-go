package api

import (
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
	}
}
