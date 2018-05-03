package redis

import (
	"time"

	"github.com/go-redis/redis"
)

var connect = New()

func New() redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "app_redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return *client
}

func Set(key string, value string, expiration_second int32) {
	err := connect.Set(key, value, time.Duration(expiration_second)*time.Second).Err()
	if err != nil {
		panic(err)
	}
}

func Get(key string) string {
	val, err := connect.Get(key).Result()
	if err != nil {
		panic(err)
	}
	return val
}

func Incr(key string) int64 {
	result, err := connect.Incr(key).Result()
	if err != nil {
		panic(err)
	}
	return result
}

func Decr(key string) int64 {
	result, err := connect.Decr(key).Result()
	if err != nil {
		panic(err)
	}
	return result
}
