package redislib

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var (
	redisClient *redis.Client
)

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:           viper.GetString("redis.addr"),
		Password:       viper.GetString("redis.password"),
		DB:             viper.GetInt("redis.db"),
		PoolSize:       viper.GetInt("redis.pool_size"),
		MinIdleConns:   viper.GetInt("redis.min_idle_conns"),
		MaxIdleConns:   viper.GetInt("redis.max_idle_conns"),
		MaxActiveConns: viper.GetInt("redis.max_active_conns"),
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}

func GetRedisClient() *redis.Client {
	return redisClient
}
