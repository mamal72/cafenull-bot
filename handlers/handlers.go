package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/mamal72/cafenull-bot/helpers"
	"github.com/mamal72/cafenull-bot/messages"
)

func handleCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	switch msg.Text {
	case "/start":
		handleStartCommand(bot, msg)
	case "/stop":
		handleStopCommand(bot, msg)
	}
}

func handleKeyboardButton(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	switch msg.Text {
	case messages.WifiButton:
		handleWifiButton(bot, msg)
	case messages.ContactInfoButton:
		handleContactInfoButton(bot, msg)
	}
}

// Handle handles bot message
func Handle(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	if helpers.IsCommand(msg) {
		handleCommand(bot, msg)
		return
	}

	if helpers.IsFromAdmin(msg) && helpers.IsDocument(msg) {
		handleCSV(bot, msg)
		return
	}

	handleKeyboardButton(bot, msg)
}
