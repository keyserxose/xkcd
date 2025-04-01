package main

import (
	"fmt"
	"io"
	"net/http"
)

var headerValue string

func getFeed(feedUrl string) (byteValue []byte) {

	//var headerTypeSend string

	client := &http.Client{}

	req, err := http.NewRequest("GET", feedUrl, nil)
	if err != nil {
		panic(err)
	}

	//req.Header.Set("User-Agent", "gotomic-rss/1.0")
	req.Header.Add("If-None-Match", readFile())
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// this prints the user agent
	//fmt.Println(req.UserAgent())

	byteValue, _ = io.ReadAll(resp.Body)

	if resp.StatusCode > 399 {
		fmt.Println("server returned an error code")
	}

	if resp.Status == "304 Not Modified" && resp.StatusCode < 399 {
		// if the Last-Modified Tag matches we do not do anything
		fmt.Println("ETag matches, process stops here, no need to get the feed")

	}
	if resp.Status != "304 Not Modified" && resp.StatusCode < 399 {
		if resp.Header.Get("ETag") != "" {
			headerValue = resp.Header.Get("ETag")
			writeFile(headerValue)
			fmt.Println("found ETag")
			fmt.Println(headerValue)
		}
		/* if validateFeedType(byteValue) == "atom" {
			postTitle, postLink, postLastBuildDate, feedTitle = readAtom(byteValue)
			if headerValue != "" {
				writeLastModified(headerValue, headerType, id)
			}

		}*/

	}

	return byteValue

}
