package handlers

import (
	"fmt"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/mamal72/cafenull-bot/config"
	"github.com/mamal72/cafenull-bot/helpers"
	"github.com/mamal72/cafenull-bot/messages"
)

func handleContactInfoButton(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) error {
	return helpers.SendTextMessage(bot, msg.From.ID, fmt.Sprintf(messages.ContactInfo, config.ContactAddress, config.ContactPhoneNumber))
}
