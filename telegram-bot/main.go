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

	data := bot.GetUpdatesChan(tgbotapi.UpdateConfig{
		Timeout:        2,
		AllowedUpdates: []string{"ClarksonsFarmSeries"},
	})
	println(data)

}

func Init(botToken string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalln(err)
	}
	bot.Debug = true
	return bot
}
