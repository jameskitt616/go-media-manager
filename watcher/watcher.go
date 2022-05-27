package watcher

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

func GetFolderEntriesByPath(directory string) {
	items, err := ioutil.ReadDir(directory)
	if err != nil {
		return
	}

	//useDirectory := false
	var largestFile fs.FileInfo
	//TODO: make recursive for sub-folder
	for _, item := range items {
		if item.Size() != 0 {
			if largestFile == nil {
				largestFile = item
				continue
			}

			if largestFile.Size() < item.Size() {
				//fmt.Println(largestFile)
				largestFile = item
			}
		}
		//if strings.HasSuffix(item.Name(), "docs") {
		//	continue
		//}
		//if strings.HasSuffix(item.Name(), "swagger") {
		//	continue
		//}
		//
		//if item.IsDir() && item.Name()[0] != '.' {
		//	readAppDirectories(directory+"/"+item.Name(), paths)
		//	continue
		//}
		//
		//if useDirectory {
		//	continue
		//}
		//
		//*paths = append(*paths, directory)
		//useDirectory = true
	}
	fmt.Println(largestFile.Name())
}
