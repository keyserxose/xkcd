package main

import (
	"fmt"
	"os"
)

func evaluate(rootPath, database, filename, title, executionDate, lastBuildDateFormatted, apiKey, chatId, comicUrl, comicUrlImage, altText string) {

	if filename == title && executionDate == lastBuildDateFormatted {
		fmt.Println("No new comics available")
		os.Exit(1)
	}

	if filename != title && executionDate == lastBuildDateFormatted {
		fmt.Println("found date but not title")
		os.Exit(1)
	}

	if filename == title && executionDate != lastBuildDateFormatted {
		fmt.Println("found title but not date")
		os.Exit(1)
	}

	if filename != title && executionDate != lastBuildDateFormatted {
		fmt.Println("It looks like there is a new comic")
		fmt.Println("Getting today's comic URL: " + comicUrl)
		fmt.Println("Getting Image url: " + comicUrlImage)
		sendToTelegram(apiKey, chatId, comicUrlImage, comicUrl, altText, title)
	}

}
