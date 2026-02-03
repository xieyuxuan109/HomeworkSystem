package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xieyuxuan109/homeworksystem/pkg"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			pkg.BadResponse(c, "需要登录",nil)
			c.Abort()
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := pkg.VerifyToken(tokenStr, "access")
		if err != nil {
			pkg.BadResponse(c, "token无效或已过期",err)
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Set("department", claims.Department)
		c.Next()
	}
}
