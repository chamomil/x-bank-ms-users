package config

import (
	"encoding/json"
	"os"
)

type (
	Config struct {
		Hs512SecretKey string `json:"hs512SecretKey"`
	}
)

func Read(filename string) (Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}
	defer func() { _ = f.Close() }()

	var config Config
	err = json.NewDecoder(f).Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
