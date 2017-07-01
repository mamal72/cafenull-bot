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

	// Add user to receiving queue
	data.AddUserToQueue(msg.From)

	// Ask user to wait for credentials
	var credentialsMsg string
	delay := config.WifiCredentialsSendDelay
	helpers.SendTextMessage(
		bot,
		msg.From.ID,
		fmt.Sprintf(messages.SendingWifiCredentials, int(delay.Seconds())),
	)

	// Send credentials message with a short delay
	go func() {
		time.Sleep(delay)
		credData, err := data.Pop()
		if err != nil {
			credentialsMsg = messages.NoWifiCredentials
		} else {
			credentialsMsg = fmt.Sprintf(messages.WifiCredentials, credData.Username, credData.Password)
		}
		helpers.SendMarkdownMessage(bot, msg.From.ID, credentialsMsg)
		data.RemoveUserFromQueue(msg.From)
	}()

	return nil
}
