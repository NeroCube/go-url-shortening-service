package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nerocube/go-url-shortening-service/redis"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func URLIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(urlmaps); err != nil {
		panic(err)
	}
}

func URLShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var urlID int
	var err error
	if urlID, err = strconv.Atoi(vars["urlID"]); err != nil {
		panic(err)
	}
	urlmap := RepoFindURLMap(urlID)
	if urlmap.ID > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(urlmap); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"original_url":"https://github.com/NeroCube"}' http://localhost:8000/urls

*/
func URLCreate(w http.ResponseWriter, r *http.Request) {
	var urlmap URLMap
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &urlmap); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateURLMap(urlmap)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func URLRedirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tinyURL := vars["tinyURL"]
	isExists := redis.Exists(tinyURL)
	if isExists == 1 {
		http.Redirect(w, r, redis.Get(tinyURL), 301)
	} else {
		fmt.Fprint(w, "tinyURL not be used\n")
	}
}
