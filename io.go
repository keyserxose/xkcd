package main

import (
	"fmt"
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
	dat, _ := os.ReadFile("etag.txt")
	defer func() {
		if r := recover(); r != nil {

			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	etag = string(dat)

	return etag
}
