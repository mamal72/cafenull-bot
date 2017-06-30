package handlers

import (
	"fmt"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/mamal72/cafenull-bot/config"
	"github.com/mamal72/cafenull-bot/data"
	"github.com/mamal72/cafenull-bot/helpers"
	"github.com/mamal72/cafenull-bot/messages"
)

func handleCSV(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	// Get CSV file url
	url, err := bot.GetFileDirectURL(msg.Document.FileID)
	helpers.CheckErr(err)

	// Download new CSV file
	err = helpers.DownloadFile(url, config.WifiCSVFileAddress)
	helpers.CheckErr(err)

	// Load new CSV file into storage
	err = data.LoadFile()
	helpers.CheckErr(err)

	// Send success message to admin
	helpers.SendTextMessage(bot, msg.From.ID, fmt.Sprintf(messages.NewUserNamesReceived, len(data.GetSlice())))
}
