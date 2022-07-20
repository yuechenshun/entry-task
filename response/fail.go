package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Fail 请求失败返回
func Fail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"Code":    400,
		"Message": msg,
	})
}
