package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var service_port string

	flag.StringVar(&service_port, "port", "8000", "The port you want to run your service")
	flag.Parse()

	router := NewRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", service_port), router))
}
