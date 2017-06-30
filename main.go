package main

import (
	"github.com/mamal72/cafenull-bot/config"
	"github.com/mamal72/cafenull-bot/handlers"
	"github.com/mamal72/cafenull-bot/helpers"
)

func main() {
	// Load config
	config.LoadConfig()

	// Create bot
	bot, err := helpers.GetNewBot(config.BotToken, config.Debug)
	helpers.CheckErr(err)

	// Create updates chan
	updates, err := helpers.GetBotUpdatesChan(bot)
	helpers.CheckErr(err)

	for update := range updates {
		handlers.Handle(bot, update.Message)
	}
}
