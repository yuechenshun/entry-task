package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Success 请求成功返回
func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"Code":    200,
		"Message": msg,
		"Data":    data,
	})
}
