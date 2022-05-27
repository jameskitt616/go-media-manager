package main

import (
	"go-media-manager/config"
	"go-media-manager/watcher"
)

var (
	conf = config.GetConfig()
)

func main() {
	//watcher.NewWatcher(conf.ENV.WatchPath)
	watcher.GetFolderEntriesByPath(conf.ENV.WatchPath)
}
