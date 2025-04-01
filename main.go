package main

import (
	"fmt"
)

func main() {

	configFile := readFlags()
	apiKey, chatId := readConfig(configFile)
	body := getFeed("https://xkcd.com/atom.xml")
	//currentPath := currentPath()
	title, comicUrl, lastBuildDateFormatted, comicUrlImage, altText := readatom(body)
	fmt.Println("This is the title: " + title)
	fmt.Println("This is the comic url: " + comicUrl)
	fmt.Println("This is the date: " + lastBuildDateFormatted)
	sendToTelegram(apiKey, chatId, comicUrlImage, comicUrl, altText, title)

}
