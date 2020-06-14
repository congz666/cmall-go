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
