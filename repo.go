package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
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
	t.ShortenURL = TinyURL(6)
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

func TinyURL(random_length int) string {
	var b bytes.Buffer
	var charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for i := 0; i < random_length; i++ {
		rand.Seed(time.Now().UnixNano())
		// the length of charSet is 62
		b.WriteString(fmt.Sprintf("%v", string(charSet[rand.Intn(62)])))
	}

	return b.String()
}
