package main

import (
	"io"
	"net/http"
	"regexp"
)

func getComicUrl(comicUrl string) (comicUrlImage string) {

	resp, err := http.Get(comicUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`https:\/\/media.oglaf.com\/comic\/(.*).jpg`)
	comicUrlSlice := re.FindAllString(string(body), -1)
	comicUrlImage = comicUrlSlice[0]

	return comicUrlImage

}
