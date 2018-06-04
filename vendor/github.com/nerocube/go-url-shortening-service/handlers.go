package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nerocube/go-url-shortening-service/redis"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

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
	TinyURL := RepoFindTinyURL(urlmap.OriginalURL)
	if TinyURL != "" {
		response := fmt.Sprintf("The Url has been established. The TinyURL is %s\n", TinyURL)
		fmt.Fprint(w, response)
	} else {
		t := RepoCreateURLMap(urlmap)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(t); err != nil {
			panic(err)
		}
	}
}

func URLRedirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tinyURL := vars["tinyURL"]

	if RepoFindOriginalURL(tinyURL) != "" {
		http.Redirect(w, r, redis.Get(tinyURL), 301)
	} else {
		fmt.Fprint(w, "tinyURL not be used\n")
	}
}
