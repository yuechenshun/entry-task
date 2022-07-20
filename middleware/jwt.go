package middleware

import (
	"entry-task/response"
	"entry-task/util"
	"github.com/gin-gonic/gin"
)

// JwtAuth 校验token
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("token")
		if tokenString == "" {
			response.Fail(c, "用户未登陆或非法访问")
			c.Abort()
			return
		}
		token, claims, err := util.VerifyToken(tokenString)
		if err != nil || !token.Valid {
			response.Fail(c, "登陆已过期，请重新登陆")
			c.Abort()
			return
		}
		// 从token中解析出来的数据挂载到上下文上,方便后面的控制器使用
		c.Set("username", claims.Username)
		c.Next()
	}
}
