package _init

import (
	"github.com/go-redis/redis/v8"
)

func Redis(cfg *Config) *redis.Client {
	local := cfg.Redis.Local
	password := cfg.Redis.Password
	db := cfg.Redis.Db
	rds := redis.NewClient(&redis.Options{
		Addr:     local,
		Password: password,
		DB:       db,
	})
	return rds
}
