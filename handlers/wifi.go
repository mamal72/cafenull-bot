package handlers

import (
	"fmt"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/mamal72/cafenull-bot/config"
	"github.com/mamal72/cafenull-bot/data"
	"github.com/mamal72/cafenull-bot/helpers"
	"github.com/mamal72/cafenull-bot/messages"
)

func handleWifiButton(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) error {
	// Return if user is already in receiving queue
	if data.IsUserInQueue(msg.From) {
		return nil
	}

	// Send credentials message to user
	var credentialsMsg string
	credData, err := data.Pop()
	if err != nil {
		credentialsMsg = messages.NoWifiCredentials
	} else {
		credentialsMsg = fmt.Sprintf(messages.WifiCredentials, credData.Username, credData.Password)
	}
	helpers.SendMarkdownMessage(bot, msg.From.ID, credentialsMsg)

	// Add user to wait queue
	data.AddUserToQueue(msg.From)
	go func() {
		delay := config.WifiCredentialsSendDelay
		time.Sleep(delay)
		data.RemoveUserFromQueue(msg.From)
	}()

	return nil
}
