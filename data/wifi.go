package data

import (
	"bufio"
	"encoding/csv"
	"os"
	"sync"

	"github.com/pkg/errors"

	"github.com/mamal72/cafenull-bot/config"
)

var (
	storage *WifiCredentialStorage
)

// WifiCredentialData is struct containing username and password for wifi credentials
type WifiCredentialData struct {
	Username string
	Password string
}

// WifiCredentialStorage is main struct with mux for credentials storage
type WifiCredentialStorage struct {
	credentials []WifiCredentialData
	sync.Mutex
}

// GetCredentialsStorage returns our CredentialStorage singleton
func GetCredentialsStorage() *WifiCredentialStorage {
	if storage == nil {
		storage = &WifiCredentialStorage{}
		storage.Lock()
		defer storage.Unlock()
		LoadFile()
	}
	return storage
}

// GetSlice returns a slice of string slices which can be saved in a CSV file
func GetSlice() [][]string {
	GetCredentialsStorage()

	var creds [][]string
	for _, item := range storage.credentials {
		currentCred := []string{item.Username, item.Password}
		creds = append(creds, currentCred)
	}
	return creds
}

// Pop returns a new credentialData
func Pop() (WifiCredentialData, error) {
	GetCredentialsStorage()

	// Lock storage to prevent multiple fs writing
	storage.Lock()
	defer storage.Unlock()

	if len(storage.credentials) == 0 {
		return WifiCredentialData{}, errors.New("no remaining usersnames")
	}

	// Pop last item and assign the rest to singleton credentials
	item, rest := storage.credentials[len(storage.credentials)-1], storage.credentials[:len(storage.credentials)-1]
	storage.credentials = rest
	err := os.Remove(config.WifiCSVFileAddress)
	file, err := os.Create(config.WifiCSVFileAddress)
	defer file.Close()
	if err != nil {
		return WifiCredentialData{}, err
	}

	// Write new credentials to CSV file
	writer := csv.NewWriter(file)
	err = writer.WriteAll(GetSlice())
	if err != nil {
		return WifiCredentialData{}, err
	}
	return item, nil
}

// LoadFile loads credentials CSV file in storage
func LoadFile() error {
	GetCredentialsStorage()
	// Return if CSV file does not exist
	if _, err := os.Stat(config.WifiCSVFileAddress); os.IsNotExist(err) {
		return nil
	}

	// Open CSV file
	file, err := os.Open(config.WifiCSVFileAddress)
	if err != nil {
		return err
	}

	// Read CSV file to a reader a remove
	reader := csv.NewReader(bufio.NewReader(file))
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// Remove header row if it's there
	if records[0][0] == "Login" {
		records = records[1:]
	}

	// Read credentials from CSV reader to storage
	var newCredentials []WifiCredentialData
	for _, item := range records {
		cred := WifiCredentialData{item[0], item[1]}
		newCredentials = append(newCredentials, cred)
	}
	storage.credentials = newCredentials

	return nil
}
