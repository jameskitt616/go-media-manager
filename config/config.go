package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	ENV struct {
		WatchPath string `json:"watch_path"`
	} `json:"env"`
	TMDB struct {
		ApiKey string `json:"api_key"`
	} `json:"tmdb"`
}

func ParseConfigFromJson(fileName string) (conf *Config, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}

	conf = new(Config)
	err = json.NewDecoder(file).Decode(conf)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	return
}

func GetConfig() *Config {
	config, err := ParseConfigFromJson("./config.json")
	if err != nil {
		panic(err)
	}
	return config
}
