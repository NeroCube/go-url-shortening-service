package main

import (
	"log"
	"net/http"

	"./config"
)

func main() {

	// Read config
	config, err := config.FromDefault()
	if err != nil {
		log.Fatal(err)
	}

	router := NewRouter()

	log.Fatal(http.ListenAndServe(config.Server.Port, router))
}
