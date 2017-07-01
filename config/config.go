package config

import (
	"os"
	"path"
	"strings"
	"time"

	// autoload .env file to env
	"strconv"

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
	// WifiCredentialsSendDelay is delay before sending wifi credentials
	WifiCredentialsSendDelay = 10 * time.Second
	// ContactAddress is cafe address in string format
	ContactAddress = ""
	// ContactPhoneNumber is cafe phone number
	ContactPhoneNumber = ""
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

	if tempWifiCredentialsSendDelay := os.Getenv("NULL_WIFI_SEND_DELAY"); tempWifiCredentialsSendDelay != "" {
		duration, err := strconv.ParseInt(tempWifiCredentialsSendDelay, 10, 64)
		if err == nil {
			WifiCredentialsSendDelay = time.Duration(duration) * time.Second
		}
	}

	if tempContactAddress := os.Getenv("NULL_CONTACT_ADDRESS"); tempContactAddress != "" {
		ContactAddress = tempContactAddress
	}

	if tempContactPhoneNumber := os.Getenv("NULL_CONTACT_PHONE_NUMBER"); tempContactPhoneNumber != "" {
		ContactPhoneNumber = tempContactPhoneNumber
	}

	if os.Getenv("BOT_DEBUG") != "" {
		Debug = true
	}
}
