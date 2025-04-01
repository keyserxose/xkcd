package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	configFile := readFlags()
	apiKey, chatId := readConfig(configFile)
	body := getFeed("https://xkcd.com/atom.xml")
	//rootPath := rootPath()
	title, comicUrl, lastBuildDateFormatted, comicUrlImage, altText := readatom(body)
	fmt.Println("This is the title: " + title)
	fmt.Println("This is the comic url: " + comicUrl)
	fmt.Println("This is the date: " + lastBuildDateFormatted)
	fmt.Println("XML Title:", title)
	fmt.Println("XML Last Build Date:", lastBuildDateFormatted)
	//evaluate(rootPath, database, filename, title, executionDate, lastBuildDateFormatted, apiKey, chatId, comicUrl, comicUrlImage, altText)
	sendToTelegram(apiKey, chatId, comicUrlImage, comicUrl, altText, title)

}
