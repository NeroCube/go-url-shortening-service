package main

import (
	"fmt"
	"time"

	"github.com/nerocube/go-url-shortening-service/encode"
	"github.com/nerocube/go-url-shortening-service/redis"
)

var urlmaps URLMaps

// Give us some seed data
func init() {
	redis.Set("counter", "0", 0)
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
	TinyURL := ""
	for {
		TinyURL = encode.TinyURL(6)
		isExists := redis.Exists(TinyURL)
		if !isExists {
			break
		}
	}
	redis.Set(TinyURL, t.OriginalURL, 0)
	t.ShortenURL = TinyURL
	t.ID = int(redis.Incr("counter"))
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
