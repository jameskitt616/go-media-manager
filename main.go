package main

import (
	"fmt"
	"go-media-manager/config"
)

var (
	conf = config.GetConfig()
)

func main() {
	fmt.Println(conf.TMDB.ApiKey)
}
