package main

import (
	"fmt"
	"go-media-manager/config"
)

var (
	conf = config.GetConfig()
)

func GetConfig() *config.Config {
	return conf
}

func main() {
	fmt.Println(conf.TMDB.ApiKey)
}
