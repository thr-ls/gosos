package cmd

import (
	"gitea.thrls.net/thr-ls/gosos/output"
)

// List function displays all URLs stored in gosos
func List() {
	urlList, err := loadURLs()
	if err != nil {
		output.PrintError("Error loading URLs: " + err.Error())
		return
	}

	if len(urlList.URLs) == 0 {
		output.PrintInfo("No URLs found. Use 'gosos add <url>' to add URLs.")
		return
	}

	output.PrintURLList(urlList.URLs)
}
