package config

import (
	"os"
	"path"
	"strings"

	// autoload .env file to env
	_ "github.com/joho/godotenv/autoload"
)

var (
	// Debug sets debug mode on or off
	Debug = false
	// Admins includes admins usernames
	Admins = []string{}
	// BotToken is our Telegram bot token
	BotToken string
	wd, _    = os.Getwd()
	// WifiCSVFileAddress is the path to save CSV file received from admins
	WifiCSVFileAddress = path.Join(wd, "credentials_data.csv")
)

// LoadConfig loads all configs
func LoadConfig() {
	if tempAdminUsers := os.Getenv("NULL_BOT_ADMINS"); tempAdminUsers != "" {
		Admins = strings.Split(tempAdminUsers, ",")
	}

	if tempBotToken := os.Getenv("NULL_BOT_TOKEN"); tempBotToken != "" {
		BotToken = tempBotToken
	}

	if tempWifiCSVFileAddress := os.Getenv("NULL_WIFI_CSV_FILE"); tempWifiCSVFileAddress != "" {
		WifiCSVFileAddress = tempWifiCSVFileAddress
	}

	if os.Getenv("BOT_DEBUG") != "" {
		Debug = true
	}
}
