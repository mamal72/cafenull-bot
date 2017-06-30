package handlers

import (
	"fmt"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/mamal72/cafenull-bot/data"
	"github.com/mamal72/cafenull-bot/helpers"
	"github.com/mamal72/cafenull-bot/messages"
)

func handleWifiButton(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) error {
	var credentialsMsg string
	helpers.SendTextMessage(bot, msg.From.ID, fmt.Sprintf(messages.SendingWifiCredentials, 10))
	go func() {
		time.Sleep(10 * time.Second)
		credData, err := data.Pop()
		if err != nil {
			credentialsMsg = messages.NoWifiCredentials
		} else {
			credentialsMsg = fmt.Sprintf(messages.WifiCredentials, credData.Username, credData.Password)
		}
		helpers.SendMarkdownMessage(bot, msg.From.ID, credentialsMsg)
	}()
	return nil
}
