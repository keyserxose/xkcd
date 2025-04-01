package main

import (
	"net/http"
	"net/url"
)

func sendToTelegram(apiKey, chatId, comicUrlImage, comicUrl, altText, title string) {

	botUrl := "https://api.telegram.org/bot" + apiKey + "/sendPhoto?"

	titleUrl := "[" + title + "]" + "(" + comicUrl + ")"

	resp, err := http.PostForm(botUrl,
		url.Values{"chat_id": {chatId}, "photo": {comicUrlImage}, "parse_mode": {"markdown"}, "caption": {titleUrl}})
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	botUrlText := "https://api.telegram.org/bot" + apiKey + "/sendMessage?"

	resp2, err := http.PostForm(botUrlText,
		url.Values{"chat_id": {chatId}, "parse_mode": {"markdown"}, "text": {altText}})
	if err != nil {
		panic(err)
	}

	defer resp2.Body.Close()

}
