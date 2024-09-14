package cmd

import (
	"gitea.thrls.net/thr-ls/gosos/output"
	"gitea.thrls.net/thr-ls/gosos/storage"
)

// loadURLs retrieves the list of URLs from storage
func loadURLs() (*storage.URLList, error) {
	urlList, err := storage.LoadURLs(storage.FileName)
	if err != nil {
		output.PrintError("Error loading URLs: " + err.Error())
		return &storage.URLList{}, err
	}
	return urlList, nil
}
