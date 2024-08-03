package cmd

import (
	"flag"
	"git.thrls.net/thrls/gosos/output"
	"git.thrls.net/thrls/gosos/storage"
	"net/url"

	"golang.org/x/exp/slices"
)

func Add(args []string) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addCmd.Parse(args)

	if addCmd.NArg() < 1 {
		output.PrintError("Insufficient arguments")
		output.PrintInfo("Usage: gosos add <url>")
		return
	}

	urlStr := addCmd.Arg(0)

	// Validate URL
	parsedURL, err := url.Parse(urlStr)
	if err != nil || !isValidURL(parsedURL) {
		output.PrintError("Invalid URL: " + urlStr)
		return
	}

	urlList, err := storage.LoadURLs()
	if err != nil {
		output.PrintError("Error: " + err.Error())
		return
	}

	if slices.Contains(urlList.URLs, urlStr) {
		output.PrintWarning("URL already exists in gosos.")
		return
	}

	urlList.URLs = append(urlList.URLs, urlStr)
	err = storage.SaveURLs(urlList)
	if err != nil {
		output.PrintError("Error: " + err.Error())
		return
	}

	output.PrintSuccess("URL added successfully")
}

func isValidURL(u *url.URL) bool {
	return u.Scheme != "" && u.Host != ""
}
