package cmd

import (
	"git.thrls.net/thrls/gosos/output"
	"git.thrls.net/thrls/gosos/storage"
)

// loadURLs retrieves the list of URLs from storage
func loadURLs() (*storage.URLList, error) {
	urlList, err := storage.LoadURLs()
	if err != nil {
		output.PrintError("Error loading URLs: " + err.Error())
		return &storage.URLList{}, err
	}
	return urlList, nil
}
