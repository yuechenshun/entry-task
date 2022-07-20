package _init

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func Mysql(cfg *Config) (*gorm.DB, error) {
	username := cfg.Database.Username
	password := cfg.Database.Password
	url := cfg.Database.Url
	var dsn = fmt.Sprintf("%s:%s@%s", username, password, url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 设置连接池
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
	return db, err
}
