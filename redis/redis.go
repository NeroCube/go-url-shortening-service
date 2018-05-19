package redis

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

var connect = New()

func New() redis.Client {
	address := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	password := os.Getenv("REDIS_PASSWORD")
	database, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	checkErr(err)

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       database,
	})

	return *client
}

func Set(key string, value string, expiration_second int32) {
	err := connect.Set(key, value, time.Duration(expiration_second)*time.Second).Err()
	checkErr(err)
}

func Get(key string) string {
	val, err := connect.Get(key).Result()
	checkErr(err)
	return val
}

func Incr(key string) int64 {
	result, err := connect.Incr(key).Result()
	checkErr(err)
	return result
}

func Decr(key string) int64 {
	result, err := connect.Decr(key).Result()
	checkErr(err)
	return result
}

func Exists(key string) (bool, error) {
	result, err := connect.Exists(key).Result()
	if result == 1 {
		return true, err

	} else {
		return false, err
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
