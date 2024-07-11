package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func Which() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("please provide an argument!")
		return
	}

	files := arguments[1:]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	for _, file := range files {
	for _, directory := range pathSplit{
		fullPath := filepath.Join(directory, file)
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			if mode.IsRegular(){ // проверка является ли файл исполняемым
				if mode&0111 != 0 {
					fmt.Println(fullPath)
					continue
				}
			}
		}
	}
}
}