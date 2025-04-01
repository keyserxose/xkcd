package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func getFeed(feedUrl string) (byteValue []byte) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", feedUrl, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("If-None-Match", readFile())
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	byteValue, _ = io.ReadAll(resp.Body)

	if resp.StatusCode > 399 {
		fmt.Println("server returned an error code")
		os.Exit(0)
	}

	if resp.Status == "304 Not Modified" && resp.StatusCode < 399 {
		fmt.Println("ETag matches, no need to get the feed")
		os.Exit(0)

	}
	if resp.Status != "304 Not Modified" && resp.StatusCode < 399 {
		if resp.Header.Get("ETag") != "" {
			headerValue := resp.Header.Get("ETag")
			writeFile(headerValue)
			fmt.Println("this is the ETag: " + headerValue)
		}

	}

	return byteValue

}
