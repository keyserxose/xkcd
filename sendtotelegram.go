package main

import (
	"net/http"
	"net/url"
)

func sendToTelegram(apiKey, chatId, comicUrl, comicArchiveUrl, altText string) {

	botUrl := "https://api.telegram.org/bot" + apiKey + "/sendPhoto?"

	resp, err := http.PostForm(botUrl,
		url.Values{"chat_id": {chatId}, "photo": {comicUrl}, "caption": {comicArchiveUrl}})
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	botUrlText := "https://api.telegram.org/bot" + apiKey + "/sendMessage?"

	resp2, err := http.PostForm(botUrlText,
		url.Values{"chat_id": {chatId}, "text": {altText}})
	if err != nil {
		panic(err)
	}

	defer resp2.Body.Close()

}
