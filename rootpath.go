package main

import (
	"log"
	"os"
	"path/filepath"
)

func currentPath() (rootPath string) {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	rootPath, err = os.Executable()
	if err != nil {
		log.Println(err)
	}

	rootPath, err = filepath.Rel(path, rootPath)
	if err != nil {
		panic(err)
	}

	rootPath = filepath.Dir(rootPath)

	return rootPath

}
