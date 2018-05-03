package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

func init() {
	New()
	fmt.Println("init")
}

func New() {
	client := redis.NewClient(&redis.Options{
		Addr:     "app_redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}
