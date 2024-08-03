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

func getFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, fileName), nil
}
