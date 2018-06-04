package main

import (
	"database/sql"
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

func RepoCreateURLMap(t URLMap) URLMap {
	TinyURL := ""

	for {
		TinyURL = encode.TinyURL(6)
		if !IsExistsTinyURL(TinyURL) {
			break
		}
	}
	// Set a cache expiration period of one week
	redis.Set(TinyURL, t.OriginalURL, 604800)
	t.ID = redis.Incr("counter")
	t.ShortenURL = TinyURL
	t.Created = time.Now()
	InsertURLMap(t)
	return t
}

func RepoFindOriginalURL(tiny_url string) string {
	// todo: add redis find tiny_url here

	isInCache, _ := redis.Exists(tiny_url)
	if isInCache {
		return redis.Get(tiny_url)
	}

	if IsExistsTinyURL(tiny_url) {
		OriginalURL := SelectOriginalURL(tiny_url)
		// Set a cache expiration period of one week
		redis.Set(tiny_url, OriginalURL, 604800)
		return SelectOriginalURL(tiny_url)
	}

	return ""
}

func RepoFindTinyURL(original_url string) string {
	return SelectTinyURL(original_url)
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
	if err := db.QueryRow(sqlStatement, TinyURL).Scan(&original_url); err != nil {
		return ""
	}

	return original_url
}

func SelectTinyURL(OriginalURL string) string {
	dbinfo := postgres.Info()
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)

	tiny_url := ""
	sqlStatement := "SELECT tiny_url FROM public.url_map WHERE original_url = $1;"
	err = db.QueryRow(sqlStatement, OriginalURL).Scan(&tiny_url)
	if err := db.QueryRow(sqlStatement, OriginalURL).Scan(&tiny_url); err != nil {
		return ""
	}

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
