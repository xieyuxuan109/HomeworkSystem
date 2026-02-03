package pkg

import (
	"github.com/gin-gonic/gin"
)

// 成功响应
func GoodResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(200, gin.H{
		"code":    0,
		"message": message,
		"data":    data,
	})
}
func BadResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(400, gin.H{
		"code":    10001,
		"message": message,
		"data":    data,
	})
}
