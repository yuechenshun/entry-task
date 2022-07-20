package controller

import (
	"entry-task/model"
	"entry-task/response"
	"entry-task/server"
	"entry-task/util"
	"github.com/gin-gonic/gin"
)

var user server.WebUserServer

// @BasePath /api

// Login
// @Summary 登陆
// @Schemes
// @Description 用于用户登陆
// @Tags 无需登陆
// @Accept json
// @Produce json
// @Param object body model.UserLoginParam true "登陆参数"
// @Success 200 {string} json "{"msg":"用户登陆成功"}"
// @Failure 400 {string} json "{"msg":"用户登陆失败"}"
// @Router /login [post]
func Login(c *gin.Context) {
	// 1.解析用户传入参数
	var param model.UserLoginParam
	if err := c.ShouldBind(&param); err != nil {
		response.Fail(c, "用户参数解析错误")
		return
	}
	// 2.生成token
	uid := user.Login(param)
	if uid > 0 {
		token, _ := util.GenerateToke(param.UserName)
		userInfo := model.UserLogin{
			Uid:   uid,
			Token: token,
		}
		response.Success(c, "用户登陆成功", userInfo)
		return
	}
	response.Fail(c, "用户名或密码不存在")
}

// @BasePath /api

// ViewUserInfo
// @Summary 查看用户信息
// @Schemes
// @Description 用于已登陆的用户查看个人信息，包括用户名、邮箱、头像
// @Tags 需要登陆
// @Accept json
// @Produce json
// @Success 200 {string} json "{"msg":"用户信息查询成功"}"
// @Failure 400 {string} json "{"msg":"用户信息查询失败"}"
// @Router /user_info [get]
func ViewUserInfo(c *gin.Context) {
	// 获取当前登陆用户实例
	username := c.MustGet("username").(string)
	u := util.GetUserExample(username)

	userInfo := user.GetUserInfo(u.Id)
	if userInfo.UserName != "" {
		response.Success(c, "用户信息查询成功", userInfo)
	} else {
		response.Fail(c, "查询失败，用户信息无效")
	}
}
