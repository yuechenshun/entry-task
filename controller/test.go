package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Test struct {
}

// @BasePath /api

// TestFunc
// @Summary 测试
// @Schemes
// @Description 这是一个测试方法
// @Tags Swagger测试
// @Accept json
// @Produce json
// @Success 200 {string} TestFunc
// @Router /index [get]
func (con Test) TestFunc(c *gin.Context) {
	c.String(http.StatusOK, "entry-task by: yuechenshun")
}
