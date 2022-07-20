package _init

import (
	"entry-task/global"
	"entry-task/util"
)

func Init() {

	// 解析配置文件
	cfg, configErr := ParseConfig("./config/app.json")
	if configErr != nil {
		util.Log("配置文件解析失败")
		panic(configErr)
	}

	util.Log("配置文件解析成功")

	// 连接数据库
	db, dbErr := Mysql(cfg)
	if dbErr != nil {
		util.Log("数据库连接失败")
		panic(dbErr)
	}
	global.Db = db

	util.Log("数据库连接成功")

	// 连接redis
	rds := Redis(cfg)
	global.Rds = rds

	util.Log("redis连接成功")

	// 连接路由
	Router(cfg)
}
