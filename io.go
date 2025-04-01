package main

import (
	"os"
)

func writeFile(etag string) {

	d1 := []byte(etag)
	err := os.WriteFile("etag.txt", d1, 0644)
	if err != nil {
		panic(1)
	}

}

func readFile() (etag string) {
	dat, err := os.ReadFile("etag.txt")
	if err != nil {
		panic(1)
	}

	etag = string(dat)

	return etag
}
