package main

import (
	"fmt"
	"os"
)

func writeFile() {

	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("etag.txt", d1, 0644)
	if err != nil {
		panic(1)
	}

}

func readFile() {
	dat, err := os.ReadFile("etag.txt")
	if err != nil {
		panic(1)
	}
	fmt.Print(string(dat))
}
