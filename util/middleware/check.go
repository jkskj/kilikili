package middleware

import (
	"github.com/gin-gonic/gin"
	"kilikili/util/e"
)

// Check 检查是否是管理员
func Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = 200
		token := c.GetHeader("Authorization")
		claims, _ := ParseToken(token)
		if claims.Authority != 1 {
			code = e.ErrorNotAdmin
		}
		if code != e.SUCCESS {
			c.JSON(400, gin.H{
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
