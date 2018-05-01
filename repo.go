package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/nerocube/go-url-shortening-service/encode"
)

var currentId int

var urlmaps URLMaps

// Give us some seed data
func init() {
	RepoCreateURLMap(URLMap{OriginalURL: "https://github.com"})
	RepoCreateURLMap(URLMap{OriginalURL: "https://google.com/"})
}

func RepoFindURLMap(id int) URLMap {
	for _, t := range urlmaps {
		if t.ID == id {
			return t
		}
	}
	// return empty URLMap if not found
	return URLMap{}
}

//this is bad, I don't think it passes race conditions
func RepoCreateURLMap(t URLMap) URLMap {
	currentId += 1
	t.ID = currentId
	t.ShortenURL = encode.TinyURL(6)
	ExampleNewClient()
	t.Created = time.Now()
	urlmaps = append(urlmaps, t)
	return t
}

func RepoDestroyURLMap(id int) error {
	for i, t := range urlmaps {
		if t.ID == id {
			urlmaps = append(urlmaps[:i], urlmaps[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find URLMap with id of %d to delete", id)
}

func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "app_redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}
