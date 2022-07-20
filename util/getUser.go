package util

import (
	"entry-task/global"
	"entry-task/model"
)

// GetUserExample token中存储的username获取当前用户实例
func GetUserExample(username string) model.User {
	var user model.User
	global.Db.Where("username = ?", username).Find(&user)
	return user
}
