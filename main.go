package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	//writeFile("etag")
	body := getFeed("https://xkcd.com/atom.xml")

	//rootPath := rootPath()
	configFile := readFlags()
	apiKey, chatId := readConfig(configFile)
	// all of the above are config variables
	//downloadRss()
	title, comicUrl, lastBuildDateFormatted, comicUrlImage, altText := readatom(body)
	_ = comicUrlImage
	_ = altText
	fmt.Println("This is the title: " + title)
	fmt.Println("This is the comic url: " + comicUrl)
	fmt.Println("This is the date: " + lastBuildDateFormatted)
	fmt.Println("XML Title:", title)
	fmt.Println("XML Last Build Date:", lastBuildDateFormatted)
	//filename, executionDate := readDb(rootPath, database, title, lastBuildDateFormatted)
	//fmt.Println("DB Title:", filename)
	//fmt.Println("DB Execution:", executionDate)
	//evaluate(rootPath, database, filename, title, executionDate, lastBuildDateFormatted, apiKey, chatId, comicUrl, comicUrlImage, altText)
	sendToTelegram(apiKey, chatId, comicUrlImage, comicUrl, altText, title)

}
