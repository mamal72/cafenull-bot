package helpers

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/mamal72/cafenull-bot/config"
)

// CheckErr checks for error and panics if it exists
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetMainKeyboard returns main bot keyboard
func GetMainKeyboard() tgbotapi.ReplyKeyboardMarkup {
	wifiButton := tgbotapi.NewKeyboardButton("اطلاعات وای‌فای 📶")
	keyboardButtons := tgbotapi.NewKeyboardButtonRow(wifiButton)
	return tgbotapi.NewReplyKeyboard(keyboardButtons)
}

// IsPrivate returns true if message is a private message
func IsPrivate(msg *tgbotapi.Message) bool {
	if msg.Chat.IsGroup() || msg.Chat.IsSuperGroup() {
		return false
	}
	return true
}

// IsFromAdmin returns true if message is from an admin
func IsFromAdmin(msg *tgbotapi.Message) bool {
	for _, username := range config.Admins {
		if msg.From.UserName == username {
			return true
		}
	}
	return false
}

// IsDocument returns true if message is a file
func IsDocument(msg *tgbotapi.Message) bool {
	return msg.Document != nil
}

// IsCommand returns true if message text starts with /
func IsCommand(msg *tgbotapi.Message) bool {
	return strings.HasPrefix(msg.Text, "/")
}

// DownloadFile downloads a file and saves it
func DownloadFile(url string, filepath string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// SendMessage sends a message to a chat
func SendMessage(bot *tgbotapi.BotAPI, chatID int, text string, markdown bool, keyboard bool) error {
	msg := tgbotapi.NewMessage(int64(chatID), text)
	if markdown {
		msg.ParseMode = tgbotapi.ModeMarkdown
	}
	if keyboard {
		msg.ReplyMarkup = GetMainKeyboard()
	}
	_, err := bot.Send(msg)
	return err
}

// SendTextMessage sends a normal text message to a chat
func SendTextMessage(bot *tgbotapi.BotAPI, chatID int, text string) error {
	return SendMessage(bot, chatID, text, false, true)
}

// SendMarkdownMessage sends a markdown message to a chat
func SendMarkdownMessage(bot *tgbotapi.BotAPI, chatID int, text string) error {
	return SendMessage(bot, chatID, text, true, true)
}
