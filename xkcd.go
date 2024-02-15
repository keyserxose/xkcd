package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	rootPath := rootPath()
	configFile := readFlags()
	apiKey, chatId, database := readConfig(configFile)
	// all of the above are config variables
	downloadRss()
	title, comicUrl, lastBuildDateFormatted, comicUrlImage := readatom()
	fmt.Println(title)
	fmt.Println(comicUrl)
	fmt.Println(lastBuildDateFormatted)
	fmt.Println("XML Title:", title)
	fmt.Println("XML Last Build Date:", lastBuildDateFormatted)
	filename, executionDate := readDb(rootPath, database, title, lastBuildDateFormatted)
	fmt.Println("DB Title:", filename)
	fmt.Println("DB Execution:", executionDate)
	evaluate(rootPath, database, filename, title, executionDate, lastBuildDateFormatted, apiKey, chatId, comicUrl, comicUrlImage)

}
