package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadRss() {

	url := "https://xkcd.com/atom.xml"

	out, err := os.Create("atom.xml")
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println(err)
	}

}
