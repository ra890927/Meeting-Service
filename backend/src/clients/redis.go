package clients

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	redisInstance *redis.Client
	redisOnce     sync.Once
)

func GetRedisInstance() *redis.Client {
	if redisInstance == nil {
		redisOnce.Do(func() {
			redisClient := redis.NewClient(getRedisOptionsFromConfig())

			_, err := redisClient.Ping(redisClient.Context()).Result()
			if err != nil {
				panic(err)
			}

			redisInstance = redisClient
		})
	}

	return redisInstance
}

func getRedisOptionsFromConfig() *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.database"),
	}
}
