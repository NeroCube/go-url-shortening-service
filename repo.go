package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/nerocube/go-url-shortening-service/encode"
	"github.com/nerocube/go-url-shortening-service/postgres"
	"github.com/nerocube/go-url-shortening-service/redis"
)

var urlmaps URLMaps

// Give us some seed data
func init() {
	redis.Set("counter", "0", 0)
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

func RepoCreateURLMap(t URLMap) URLMap {
	TinyURL := SelectTinyURL(t.OriginalURL)

	for {
		TinyURL = encode.TinyURL(6)
		if !IsExistsTinyURL(TinyURL) {
			break
		}
	}

	redis.Set(TinyURL, t.OriginalURL, 0)
	t.ShortenURL = TinyURL
	t.Created = time.Now()
	InsertURLMap(t)
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

func RepoFindOriginalURL(tiny_url string) string {
	if redis.Exists(tiny_url) {
		return redis.Get(tiny_url)
	}

	if IsExistsTinyURL(tiny_url) {
		return SelectOriginalURL(tiny_url)
	}

	return ""
}

func InsertURLMap(t URLMap) {
	dbinfo := postgres.Info()
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO public.url_map( original_url, tiny_url, created_at) VALUES ($1, $2, $3);")
	checkErr(err)

	_, err = stmt.Exec(t.OriginalURL, t.ShortenURL, t.Created)
	checkErr(err)

	defer db.Close()
}

func SelectOriginalURL(TinyURL string) string {
	dbinfo := postgres.Info()
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)

	original_url := ""
	sqlStatement := "SELECT original_url FROM public.url_map WHERE tiny_url = $1;"
	err = db.QueryRow(sqlStatement, TinyURL).Scan(&original_url)
	checkErr(err)

	return original_url
}

func SelectTinyURL(OriginalURL string) string {
	dbinfo := postgres.Info()
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)

	tiny_url := ""
	sqlStatement := "SELECT tiny_url FROM public.url_map WHERE original_url = $1;"
	err = db.QueryRow(sqlStatement, OriginalURL).Scan(&tiny_url)
	checkErr(err)

	return tiny_url
}

func IsExistsTinyURL(TinyURL string) bool {
	if SelectOriginalURL(TinyURL) == "" {
		return false
	}

	return true
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
