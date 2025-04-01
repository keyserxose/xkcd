package main

import (
	"fmt"
)

func main() {

	configFile := readFlags()
	apiKey, chatId := readConfig(configFile)
	title, comicUrl, lastBuildDateFormatted, comicUrlImage, altText := readatom(getFeed("https://xkcd.com/atom.xml"))
	fmt.Println("this is the title: " + title)
	fmt.Println("this is the comic url: " + comicUrl)
	fmt.Println("this is the date: " + lastBuildDateFormatted)
	sendToTelegram(apiKey, chatId, comicUrlImage, comicUrl, altText, title)

}
