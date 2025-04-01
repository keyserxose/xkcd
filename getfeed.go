package main

import (
	"fmt"
	"io"
	"net/http"
)

func getFeed(feedUrl string, id int) (postTitle, postLink, postLastBuildDate, feedTitle string, nextStep bool) {

	headerTypeFromDb := readLastModifiedType(id)

	var headerTypeSend string

	client := &http.Client{}

	req, err := http.NewRequest("GET", feedUrl, nil)
	if err != nil {
		panic(err)
	}

	if headerTypeFromDb == "Etag" {
		headerTypeSend = "If-None-Match"
		req.Header.Add(headerTypeSend, readLastModifiedValue(id))

	}
	if headerTypeFromDb == "Last-Modified" {
		headerTypeSend = "If-Modified-Since"
		req.Header.Add(headerTypeSend, readLastModifiedValue(id))

	}

	req.Header.Set("User-Agent", "gotomic-rss/1.0")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// this prints the user agent
	//fmt.Println(req.UserAgent())

	byteValue, _ := io.ReadAll(resp.Body)

	if resp.StatusCode > 399 {
		nextStep = false
		fmt.Println("server returned an error code")
	}

	if resp.Status == "304 Not Modified" && resp.StatusCode < 399 {
		nextStep = false
		// if the Last-Modified Tag matches we do not do anything
		fmt.Println(headerTypeFromDb + " matches, process stops here, no need to get the feed")

	}
	if resp.Status != "304 Not Modified" && resp.StatusCode < 399 {
		nextStep = true
		// if the Last-Modified Tag does not match, we update the tag on the db
		fmt.Println(headerTypeFromDb + " does not match, getting the feed")
		if resp.Header.Get("Last-Modified") != "" {
			headerValue = resp.Header.Get("Last-Modified")
			headerType = "Last-Modified"
			fmt.Println("found Last-Modified")
			fmt.Println(headerValue)
		}
		if resp.Header.Get("Etag") != "" {
			headerType = "Etag"
			headerValue = resp.Header.Get("Etag")
			fmt.Println("found Etag")
			fmt.Println(headerValue)
		}
		if validateFeedType(byteValue) == "atom" {
			postTitle, postLink, postLastBuildDate, feedTitle = readAtom(byteValue)
			if headerValue != "" {
				writeLastModified(headerValue, headerType, id)
			}

		}
		if validateFeedType(byteValue) == "rss" {
			postTitle, postLink, postLastBuildDate, feedTitle = readRss(byteValue)
			if headerValue != "" {
				writeLastModified(headerValue, headerType, id)
			}
		}

	}

	return postTitle, postLink, postLastBuildDate, feedTitle, nextStep

}
