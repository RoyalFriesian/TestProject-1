package utils

import (
	"github.com/go-redis/redis"
)

func CheckUserInRedis(username string) bool {
	client := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPassword,
		DB:       RedisDB,
	})

	_, err := client.Get(username).Result()
	return err == nil
}
