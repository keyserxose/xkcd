package main

import (
	"log"
	"os"
	"path/filepath"
)

func currentPath() (currentPath string) {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	currentPath, err = os.Executable()
	if err != nil {
		log.Println(err)
	}

	currentPath, err = filepath.Rel(path, currentPath)
	if err != nil {
		panic(err)
	}

	currentPath = filepath.Dir(currentPath)

	return currentPath

}

func homePath() (dirname string) {

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return dirname
}
