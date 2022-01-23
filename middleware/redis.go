package middleware

import (
	"time"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func getRedisClient() *redis.Client {
	if redisClient != nil {
		return redisClient
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return redisClient
}

func SetValue(key string, value string, expiry time.Duration) error {
	err := getRedisClient().Set(key, value, expiry).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetValue(key string) (string, error) {
	value, err := getRedisClient().Get(key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}
