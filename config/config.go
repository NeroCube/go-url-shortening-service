package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

// Config contains the configuration of the url shortener.
type Config struct {
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"server"`
	Redis struct {
		Host     string `json:"host"`
		Password string `json:"password"`
		DB       string `json:"db"`
	} `json:"redis"`
	Postgres struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DB       string `json:"db"`
	} `json:"postgres"`
	Options struct {
		Prefix string `json:"prefix"`
	} `json:"options"`
}

// FromFile returns a configuration parsed from the given file.
func FromFile(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// FromDefault returns a configuration parsed from the default json file.
func FromDefault() (*Config, error) {
	configPath := flag.String("config", "./config/config.json", "path of the config file")

	flag.Parse()

	b, err := ioutil.ReadFile(*configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
