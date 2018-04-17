package main

import "time"

type URLMap struct {
	ID          int       `json:"id"`
	OriginalURL string    `json:"original_url"`
	ShortenURL  string    `json:"shorten_url"`
	Created     time.Time `json:"created"`
}

type URLMaps []URLMap
