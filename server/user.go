package server

import (
	"entry-task/global"
	"entry-task/model"
)

type WebUserServer struct {
}

func (con *WebUserServer) Login(param model.UserLoginParam) uint64 {
	var user model.User
	global.Db.Where("username = ? and password = ?", param.UserName, param.Password).First(&user)
	return user.Id
}

func (con *WebUserServer) GetUserInfo(userId uint64) model.ViewUserInfo {
	var user model.User
	global.Db.Where("id = ?", userId).First(&user)
	showInfo := model.ViewUserInfo{
		UserName: user.UserName,
		Email:    user.Email,
		Avatar:   user.Avatar,
	}
	return showInfo
}
