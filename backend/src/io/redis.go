package io

import (
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	redisInstance *redis.Client
	redisOnce     sync.Once
)

func GetRedisInstance() *redis.Client {
	if redisInstance == nil {
		redisOnce.Do(func() {
			redisClient := redis.NewClient(&redis.Options{
				Addr:     "localhost:6379",
				Password: "", // no password set
				DB:       0,  // use default DB
			})

			_, err := redisClient.Ping(redisClient.Context()).Result()
			if err != nil {
				panic(err)
			}

			redisInstance = redisClient
		})
	}

	return redisInstance
}
