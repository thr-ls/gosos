package cmd

import (
	"git.thrls.net/thrls/gosos/output"
	"git.thrls.net/thrls/gosos/storage"
)

func List() {
	urlList, err := storage.LoadURLs()
	if err != nil {
		output.PrintError("Error: " + err.Error())
		return
	}

	if len(urlList.URLs) == 0 {
		output.PrintInfo("The URL list is currently empty. Use 'gosos add <url>' to add URLs.")
		return
	}

	output.PrintURLList(urlList.URLs)
}
