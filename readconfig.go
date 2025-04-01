package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func readConfig(configFile *string) (apiKey, chatId string) {

	type Configuration struct {
		ApiKey string
		ChatId string
	}

	file, err := os.Open(*configFile)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration.ApiKey, configuration.ChatId

}
