package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kilikili/util/e"
)

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
		fmt.Println(claims.Id, claims.Authority)
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
