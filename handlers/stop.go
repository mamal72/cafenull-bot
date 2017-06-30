package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/mamal72/cafenull-bot/helpers"
	"github.com/mamal72/cafenull-bot/messages"
)

func handleStopCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) error {
	err := helpers.SendTextMessage(bot, msg.From.ID, messages.Stop)
	return err
}
