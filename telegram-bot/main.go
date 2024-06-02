package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Get Bot token from environment variables
	botToken := os.Getenv("TOKEN")

	bot := Init(botToken)

	update := tgbotapi.NewUpdate(0)
	//update.AllowedUpdates = []string{"ClarksonsFarmSeries"}
	update.Timeout = 2
	data := bot.GetUpdatesChan(update)
	for i := range data {
		if i.Message == nil {
			continue
		} else {
			print(i.Message)
		}
	}
}

func Init(botToken string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalln(err)
	}
	bot.Debug = true
	return bot
}
