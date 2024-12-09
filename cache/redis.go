package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func SetCache(key string, value string) error {
	return RDB.Set(context.Background(), key, value, 0).Err()
}

func GetCache(key string) (string, error) {
	return RDB.Get(context.Background(), key).Result()
}
