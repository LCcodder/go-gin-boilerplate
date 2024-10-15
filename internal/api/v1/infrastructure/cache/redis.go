package cache

import (
	"example.com/m/internal/config"
	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func ConnectToRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     config.Config.RedisConnectionString,
		Password: config.Config.RedisPassword,
		DB:       0,
	})
}
