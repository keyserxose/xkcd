package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadRss() {

	url := "https://www.oglaf.com/feeds/rss/"

	out, err := os.Create("current.rss")
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
