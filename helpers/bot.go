package helpers

import "github.com/go-telegram-bot-api/telegram-bot-api"

// GetNewBot returns a new Telegram bot
func GetNewBot(token string, debug bool) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	if debug {
		bot.Debug = true
	}
	return bot, err
}

// GetBotUpdatesChan returns updates chan of bot
func GetBotUpdatesChan(bot *tgbotapi.BotAPI) (<-chan tgbotapi.Update, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	return updates, err
}
