package main

import (
	"net/http"
	"net/url"
)

func sendToTelegram(apiKey string, chatId string, comicUrl string, comicArchiveUrl string) {

	botUrl := "https://api.telegram.org/bot" + apiKey + "/sendPhoto?"

	resp, err := http.PostForm(botUrl,
		url.Values{"chat_id": {chatId}, "photo": {comicUrl}, "caption": {comicArchiveUrl}})
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

}
