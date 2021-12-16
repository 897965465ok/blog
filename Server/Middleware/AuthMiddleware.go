package middleware

import (
	util "main/Util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取Authorization
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusInternalServerError,
				"mes":  "权限不足",
			})
			ctx.Abort()
			return
		} else {
			// 取出bearer空格 后面的token字符串
			tokenString = tokenString[7:]
			token, _, err := util.ParseToken(tokenString)
			// token.Valid 如果令牌有效就是ture
			if err != nil || !token.Valid {
				ctx.JSON(http.StatusOK, gin.H{
					"code": http.StatusUnauthorized,
					"meg":  "权限不足",
				})
				ctx.Abort()
				return
			}

		}
		ctx.Next()

	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
