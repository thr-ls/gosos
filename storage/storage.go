package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type URLList struct {
	URLs []string `json:"urls"`
}

const fileName = ".gosos-urls.json"

// LoadURLs reads the list of URLs from a file and returns it as a URLList struct
func LoadURLs() (*URLList, error) {
	filePath, err := getFilePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(filePath)
	if os.IsNotExist(err) {
		return &URLList{}, nil
	} else if err != nil {
		return nil, err
	}

	var urls URLList
	err = json.Unmarshal(data, &urls)
	return &urls, err
}

// SaveURLs writes the provided URLList struct to a file in JSON format
func SaveURLs(urls *URLList) error {
	data, err := json.MarshalIndent(urls, "", "  ")
	if err != nil {
		return err
	}

	filePath, err := getFilePath()
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0600)
}

// getFilePath returns the full path to the file where URLs are stored
func getFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, fileName), nil
}
