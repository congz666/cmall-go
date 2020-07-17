package middleware

import (
	"cmall/pkg/e"
	"cmall/pkg/util"
	"time"

	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 200
		token := c.Query("token")
		if token == "" {
			code = 400
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
